package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/rpc"
	"chainsawman/consumer/task/types"
	"github.com/zeromicro/go-zero/core/logx"

	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/package-url/packageurl-go"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type CronPython struct {
	libraryIDCache *collection.Cache
	releaseIDCache *collection.Cache
}

func (h *CronPython) Handle(task *model.KVTask) (string, error) {
	// 初始化缓存
	if h.libraryIDCache == nil {
		h.libraryIDCache, _ = collection.NewCache(24 * time.Hour)
	}
	if h.releaseIDCache == nil {
		h.releaseIDCache, _ = collection.NewCache(24 * time.Hour)
	}
	// 准备数据
	// 获得图谱信息
	t := common.TaskIdf(task.Idf)
	// 获得最大ID
	maxLibraryID, err := config.NebulaClient.GetMaxLibraryID(t.FixedGraph.GraphID)
	if err != nil {
		return "", err
	}
	libraryGen := &IDGen{id: maxLibraryID + 1}
	maxReleaseID, err := config.NebulaClient.GetMaxReleaseID(t.FixedGraph.GraphID)
	if err != nil {
		return "", err
	}
	releaseGen := &IDGen{id: maxReleaseID + 1}
	// 更新软件，版本，依赖，属于 合计四种文件
	libraryCSV := "id,artifact,topic,desc,home\n"
	releaseCSV := "id,idf,artifact,version,createTime\n"
	dependsCSV := "source,target\n"
	belong2CSV := "source,target\n"
	// 获得更新的软件
	software := getUpdate()
	// 从 sca 查询依赖
	for _, s := range software {
		logx.Infof("[Cron] handle Python software: %v", s.String())
		// 获得待更新的库的ID
		sourceID, err := h.handleReleaseIfAbsent(t.FixedGraph.GraphID, s, libraryGen, releaseGen, &libraryCSV, &releaseCSV, &belong2CSV)
		if err != nil {
			return "", err
		}
		resp, err := config.ScaClient.GetDeps(s.String(), "python")
		// TODO：查询依赖失败，则跳过
		if err != nil {
			logx.Error(err)
			continue
		}
		for _, dep := range resp.Deps.Dependencies {
			purl, _ := packageurl.FromString(dep.Purl)
			targetID, err := h.handleReleaseIfAbsent(t.FixedGraph.GraphID, &purl, libraryGen, releaseGen, &libraryCSV, &releaseCSV, &belong2CSV)
			logx.Debugf("[Cron] handle Python dependency: %v -> %v, err: %v", sourceID, targetID, err)
			if err != nil {
				return "", err
			}
			// 写入发行版本依赖, TODO: 没写入？
			dependsCSV = dependsCSV + fmt.Sprintf("%v,%v\n", sourceID, targetID)
		}
	}
	// 包装为文件，完成上传
	libraryKey, err := config.OSSClient.AddSource(context.Background(), fmt.Sprintf("%v_library_%v", task.Idf, time.Now().String()), []byte(libraryCSV))
	if err != nil {
		return "", err
	}
	releaseKey, err := config.OSSClient.AddSource(context.Background(), fmt.Sprintf("%v_release_%v", task.Idf, time.Now().String()), []byte(releaseCSV))
	if err != nil {
		return "", err
	}
	dependsKey, err := config.OSSClient.AddSource(context.Background(), fmt.Sprintf("%v_depend_%v", task.Idf, time.Now().String()), []byte(dependsCSV))
	if err != nil {
		return "", err
	}
	belong2Key, err := config.OSSClient.AddSource(context.Background(), fmt.Sprintf("%v_belong2_%v", task.Idf, time.Now().String()), []byte(belong2CSV))
	if err != nil {
		return "", err
	}
	// 更新图谱
	updateReq := &types.UpdateGraphRequest{
		TaskID:  common.GraphUpdate,
		GraphID: t.FixedGraph.GraphID,
		NodeFileList: []*types.Pair{
			{
				Key:   "library",
				Value: libraryKey,
			}, {
				Key:   "release",
				Value: releaseKey,
			},
		},
		EdgeFileList: []*types.Pair{
			{
				Key:   "depend",
				Value: dependsKey,
			}, {
				Key:   "belong2",
				Value: belong2Key,
			},
		}}
	updateReqStr, _ := jsonx.MarshalToString(updateReq)
	updateHandler := UpdateGraph{}
	_, err = updateHandler.Handle(&model.KVTask{
		Params:  updateReqStr,
		GraphID: t.FixedGraph.GraphID,
	})
	return "", err
}

func (h *CronPython) handleReleaseIfAbsent(graph int64, release *packageurl.PackageURL,
	libraryIDGen *IDGen, releaseIDGen *IDGen, libraryCSV *string, releaseCSV *string, belong2CSV *string) (int64, error) {
	// 1. 查询发行版本是否存在
	if release.Version == "" {
		// 无版本则默认为 latest
		release.Version = "latest"
	}
	releaseIDMap, err := h.getIDsFromCache(graph, []string{release.String()}, h.releaseIDCache, config.NebulaClient.GetReleaseIDsByNames)
	if err != nil {
		return 0, err
	}
	// 1.1 不存在则创建
	if releaseIDMap[release.String()] == 0 {
		// 获得新ID
		releaseID := releaseIDGen.get()
		releaseIDMap[release.String()] = releaseID
		// 写入缓存中
		h.releaseIDCache.Set(release.String(), releaseID)
		meta := &rpc.MetaResponse{Meta: rpc.ModuleMeta{}}
		// 若非latest, 查询软件信息
		if release.Version != "latest" {
			res, err := config.ScaClient.GetMeta(release.ToString(), "python")
			if err == nil {
				meta = res
			}
		}
		// 新发行版本写入CSV
		*releaseCSV = *releaseCSV + fmt.Sprintf("%v,%v,%v,%v,%v\n", releaseID, release.String(), release.Name, release.Version, meta.Meta.UploadTime)
		// 2. 查询库是否存在
		libraryIDMap, err := h.getIDsFromCache(graph, []string{release.Name}, h.libraryIDCache, config.NebulaClient.GetLibraryIDsByNames)
		if err != nil {
			return 0, err
		}
		// 2.1 不存在则创建
		if libraryIDMap[release.Name] == 0 {
			// 获得新ID
			libraryID := libraryIDGen.get()
			libraryIDMap[release.Name] = libraryID
			// 写入缓存中
			h.libraryIDCache.Set(release.Name, libraryID)
			// 新库写入CSV
			*libraryCSV = *libraryCSV + fmt.Sprintf("%v,%v,,%v,%v\n", libraryID, release.Name, meta.Meta.Desc, meta.Meta.Homepage)
		}
		// 3. 归属关系写入CSV
		*belong2CSV = *belong2CSV + fmt.Sprintf("%v,%v\n", releaseIDMap[release.String()], libraryIDMap[release.Name])
	}
	// 1.2 存在则返回
	return releaseIDMap[release.String()], nil
}

func (h *CronPython) getIDsFromCache(graph int64, names []string, cache *collection.Cache,
	callback func(graph int64, names []string) (map[string]int64, error)) (map[string]int64, error) {
	res := make(map[string]int64)
	filterNames := make([]string, 0)
	// 优先从缓存中读取
	for _, name := range names {
		if id, ok := cache.Get(name); ok {
			res[name] = id.(int64)
			continue
		}
		filterNames = append(filterNames, name)
	}
	// 未读取到，则从数据库查询
	if len(filterNames) != 0 {
		filterIdMap, err := callback(graph, filterNames)
		if err != nil {
			return nil, err
		}
		for k, v := range filterIdMap {
			res[k] = v
			// 写入缓存
			cache.Set(k, v)
		}
	}
	return res, nil
}

func getUpdate() []*packageurl.PackageURL {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://pypi.org/rss/updates.xml")
	pList := make([]*packageurl.PackageURL, len(feed.Items))
	re, _ := regexp.Compile("https://pypi.org/project/([^/]+?)/([^/]+?)/")
	for i, item := range feed.Items {
		res := re.FindStringSubmatch(item.Link)
		pList[i] = &packageurl.PackageURL{
			Name:      res[1],
			Version:   res[2],
			Namespace: "pypi",
		}
	}
	return pList
}

type IDGen struct {
	id int64
}

func (g *IDGen) get() int64 {
	currentID := g.id
	g.id += 1
	return currentID
}
