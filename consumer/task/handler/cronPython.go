package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"

	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/package-url/packageurl-go"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type CronPython struct {
}

func (h *CronPython) Handle(task *model.KVTask) (string, error) {
	// 准备数据
	t := common.TaskIdf(task.Idf)
	//maxID, err := config.NebulaClient.GetMaxID(t.FixedGraph.GraphID)
	//if err != nil {
	//	return "", nil
	//}
	// TODO: 更新软件，版本，依赖，属于 合计四种文件
	edgeCSV := ""
	//nodeCSV := ""
	// 获得更新
	software := getUpdate()
	// 从 sca 查询依赖
	for _, s := range software {
		purls := make([]string, 0)
		idMap := make(map[string]int64)
		resp, _ := config.ScaClient.GetDeps(s.String(), "python")
		modules := resp.Packages.Modules
		if len(modules) > 0 {
			for _, dep := range modules[0].Dependencies {
				purls = append(purls, dep.Purl)
				idMap[dep.Purl] = 0
			}
		}
		// 将软件转化为图谱ID
		names := make([]string, 0, len(idMap))
		for k := range idMap {
			names = append(names, k)
		}
		idMap, err := config.NebulaClient.GetNodeIDsByNames(t.FixedGraph.GraphID, names)
		if err != nil {
			return "", err
		}
		for _, p := range purls {
			edgeCSV = edgeCSV + fmt.Sprintf("%v,%v\n", idMap[s.String()], idMap[p])
		}
	}
	// 包装为文件，完成上传
	filename := fmt.Sprintf("%v_%v", task.Idf, time.Now().String())
	config.OSSClient.AddSource(context.Background(), filename, []byte(edgeCSV))
	// 更新依赖
	updateReq := &types.UpdateGraphRequest{}
	updateReqStr, _ := jsonx.MarshalToString(updateReq)
	updateHandler := UpdateGraph{}
	_, err := updateHandler.Handle(&model.KVTask{
		Params:  updateReqStr,
		GraphID: t.FixedGraph.GraphID,
	})
	return "", err
}

func getUpdate() []*packageurl.PackageURL {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://pypi.org/rss/updates.xml")
	pList := make([]*packageurl.PackageURL, len(feed.Items))
	re, _ := regexp.Compile("https://pypi.org/project/([^/]+?)/([^/]+?)/")
	for i, item := range feed.Items {
		res := re.FindStringSubmatch(item.Link)
		pList[i] = &packageurl.PackageURL{
			Name:    res[1],
			Version: res[2],
		}
	}
	return pList
}
