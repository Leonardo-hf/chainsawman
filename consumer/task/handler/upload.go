package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"time"
)

type Upload struct {
}

func (h *Upload) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.UploadRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	nodeFile, err := config.OSSClient.Fetch(ctx, req.NodeID)
	if err != nil {
		return "", err
	}
	edgeFile, err := config.OSSClient.Fetch(ctx, req.EdgeID)
	if err != nil {
		return "", err
	}
	// 处理文件
	var nodes []*model.Node
	var edges []*model.Edge
	nodeMap := make(map[int64]*model.Node)
	records, err := handle(nodeFile)
	if err != nil {
		return "", err
	}
	for _, record := range records {
		id, idErr := record.GetAsInt("id")
		name, nameErr := record.Get("name")
		desc, descErr := record.Get("desc")
		if idErr == nil && nameErr == nil && descErr == nil {
			node := &model.Node{ID: id, Name: name, Desc: desc}
			nodes = append(nodes, node)
			// 把点放到nodeMap, 后面赋度数
			nodeMap[id] = node
		}
	}
	records, err = handle(edgeFile)
	for _, record := range records {
		source, sourceErr := record.GetAsInt("source")
		target, targetErr := record.GetAsInt("target")
		if sourceErr != nil || targetErr != nil {
			continue
		}
		if _, sourceExist := nodeMap[source]; !sourceExist {
			logx.Error("source doesn't exist", source)
			continue
		}
		if _, targetExist := nodeMap[target]; !targetExist {
			logx.Error("target doesn't exist", target)
			continue
		}
		edge := &model.Edge{Source: source, Target: target}
		edges = append(edges, edge)
		// 给点加度数
		nodeMap[source].Deg++
		nodeMap[target].Deg++
	}

	// 插入数据
	err = config.NebulaClient.CreateGraph(req.GraphID)
	if err != nil {
		return "", err
	}
	// TODO: az，时间贼长！
	time.Sleep(20 * time.Second)
	_, err = config.NebulaClient.MultiInsertNodes(req.GraphID, nodes)
	if err != nil {
		return "", err
	}
	_, err = config.NebulaClient.MultiInsertEdges(req.GraphID, edges)
	if err != nil {
		return "", err
	}

	_, err = config.MysqlClient.UpdateGraphByID(ctx, &model.Graph{
		ID:    req.GraphID,
		Nodes: int64(len(nodes)),
		Edges: int64(len(edges)),
	})
	if err != nil {
		return "", err
	}

	resp := &types.SearchGraphReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Graph: &types.Graph{
			Name:  req.Graph,
			Desc:  req.Desc,
			Id:    req.GraphID,
			Nodes: int64(len(nodes)),
			Edges: int64(len(edges)),
		},
	}
	return jsonx.MarshalToString(resp)
}

func handle(content io.Reader) ([]*common.Record, error) {
	parser, err := common.NewExcelParser(content)
	if err != nil {
		return nil, err
	}
	var records []*common.Record
	for {
		record, err := parser.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
			//return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
