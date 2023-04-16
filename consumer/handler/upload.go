package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/config"
	"chainsawman/consumer/model"
	"chainsawman/consumer/types"
	file "chainsawman/consumer/types/rpc"
	"context"
	"io"
	"os"

	"github.com/zeromicro/go-zero/core/jsonx"
)

type Upload struct {
}

func (h *Upload) Handle(params string, taskID int64) (string, error) {
	req := &types.UploadRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	nodeFile, err := config.FileRPC.FetchFile(ctx, &file.IDReq{Id: req.NodeID})
	if err != nil {
		return "", err
	}
	edgeFile, err := config.FileRPC.FetchFile(ctx, &file.IDReq{Id: req.EdgeID})
	if err != nil {
		return "", err
	}
	// 处理文件
	var nodes []*model.Node
	var edges []*model.Edge
	nodeMap := make(map[string]*model.Node)
	records, err := handle(nodeFile.Data)
	for _, record := range records {
		name, nameOK := record["name"]
		desc, descOK := record["desc"]
		if nameOK && descOK {
			node := &model.Node{Name: name, Desc: desc}
			nodes = append(nodes, node)
			// 把点放到nodeMap, 后面赋度数
			nodeMap[name] = node
		}
	}
	records, err = handle(edgeFile.Data)
	for _, record := range records {
		source, sourceOK := record["source"]
		target, targetOK := record["target"]
		if _, sourceExist := nodeMap[source]; !sourceExist {
			continue
		}
		if _, targetExist := nodeMap[target]; !targetExist {
			continue
		}
		if sourceOK && targetOK {
			edge := &model.Edge{Source: source, Target: target}
			edges = append(edges, edge)
			// 给点加度数
			nodeMap[source].Deg++
			nodeMap[target].Deg++
		}
	}
	// 插入数据
	err = config.NebulaClient.CreateGraph(req.Graph)
	if err != nil {
		return "", err
	}
	_, err = config.NebulaClient.MultiInsertNodes(req.Graph, nodes)
	if err != nil {
		return "", err
	}
	_, err = config.NebulaClient.MultiInsertEdges(req.Graph, edges)
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
			Nodes: int64(len(nodes)),
			Edges: int64(len(edges)),
		},
	}
	return jsonx.MarshalToString(resp)
}

func handle(content []byte) ([]common.Record, error) {
	path, err := tempSave(content)
	if err != nil {
		return nil, err
	}
	parser, err := common.NewExcelParser(path)
	if err != nil {
		return nil, err
	}
	var records []common.Record
	for {
		record, err := parser.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func tempSave(content []byte) (string, error) {
	tempFile, err := os.CreateTemp("", "*.csv")
	defer tempFile.Close()
	if err != nil {
		return "", err
	}
	tempFile.Write(content)
	return tempFile.Name(), nil
}
