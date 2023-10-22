package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"strconv"
)

type UpdateGraph struct {
}

// 期望为：一次最多在内存中存入 100M 数据
const maxInsertNum = 100000

func (h *UpdateGraph) Handle(task *model.KVTask) (string, error) {
	params, taskID := task.Params, task.Id
	req := &types.UpdateGraphRequest{}
	if err := jsonx.UnmarshalFromString(params, req); err != nil {
		return "", err
	}
	ctx := context.Background()
	group, err := config.MysqlClient.GetGroupByGraphId(ctx, req.GraphID)
	if err != nil {
		return "", err
	}
	// XLS/XLSX 最大行数有限，尽管需要全部读入内存，但不会造成问题，CSV 则可能过大，但可以流式读取，分批插入
	numNode, numEdge := 0, 0
	// 读所有边文件
	degMap := make(map[int64]int64)
	for _, edgeType := range req.EdgeFileList {
		// 确认边类型
		t, edgeFileId := edgeType.Key, edgeType.Value
		tid, _ := strconv.ParseInt(t, 10, 64)
		var edge *model.Edge
		for _, e := range group.Edges {
			if e.ID == tid {
				edge = e
				break
			}
		}
		if edge == nil {
			return "", fmt.Errorf("no such edge type")
		}
		// 获取文件
		file, err := config.OSSClient.FetchSource(ctx, edgeFileId)
		if err != nil {
			return "", err
		}
		// 解析文件
		parser, err := common.NewExcelParser(file)
		if err != nil {
			return "", err
		}
		if !(parser.HasColumn(common.KeySrc) && parser.HasColumn(common.KeyTgt)) {
			return "", fmt.Errorf("update graph fail, source files of nodes needs `source` & `target`")
		}
		edgeRecords := make([]*common.Record, 0)
		for {
			r, err := parser.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				logx.Errorf("[Consumer] parse csv fail, err: %v", err)
				continue
			}
			src, err := r.GetAsInt(common.KeySrc)
			if err != nil {
				return "", err
			}
			tgt, err := r.GetAsInt(common.KeyTgt)
			if err != nil {
				return "", err
			}
			if d, ok := degMap[src]; ok {
				degMap[src] = d + 1
			} else {
				degMap[src] = 0
			}
			if d, ok := degMap[tgt]; ok {
				degMap[tgt] = d + 1
			} else {
				degMap[tgt] = 0
			}
			edgeRecords = append(edgeRecords, r)
			if len(edgeRecords) > maxInsertNum {
				// 直接插入边，nebula先插入边，再插入顶点
				_, err = config.NebulaClient.MultiInsertEdges(req.GraphID, edge, edgeRecords)
				if err != nil {
					return "", err
				}
				// 合计边数目
				numEdge += len(edgeRecords)
				// 清零
				edgeRecords = make([]*common.Record, 0)
			}
		}
		// 插入剩余边
		_, err = config.NebulaClient.MultiInsertEdges(req.GraphID, edge, edgeRecords)
		if err != nil {
			return "", err
		}
		// 合计剩余边数目
		numEdge += len(edgeRecords)
	}
	// 读所有节点文件
	for _, nodeType := range req.NodeFileList {
		// 确认节点类型
		t, nodeFileId := nodeType.Key, nodeType.Value
		tid, _ := strconv.ParseInt(t, 10, 64)
		var node *model.Node
		for _, n := range group.Nodes {
			if n.ID == tid {
				node = n
				break
			}
		}
		if node == nil {
			return "", fmt.Errorf("no such node type")
		}
		// 获得文件流
		file, err := config.OSSClient.FetchSource(ctx, nodeFileId)
		if err != nil {
			return "", err
		}
		// 解析文件
		parser, err := common.NewExcelParser(file)
		if err != nil {
			return "", err
		}
		if !parser.HasColumn(common.KeyID) {
			return "", fmt.Errorf("update graph fail, source files of nodes needs `id`")
		}
		nodeRecords := make([]*common.Record, 0)
		for {
			r, err := parser.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				logx.Errorf("[Consumer] parse csv fail, err: %v", err)
				continue
			}
			nodeRecords = append(nodeRecords, r)
			if len(nodeRecords) > maxInsertNum {
				// 插入节点
				_, err = config.NebulaClient.MultiInsertNodes(req.GraphID, node, nodeRecords)
				if err != nil {
					return "", err
				}
				// 合计节点数目
				numNode += len(nodeRecords)
				// 清零
				nodeRecords = make([]*common.Record, 0)
			}
		}
		// 插入剩余节点
		_, err = config.NebulaClient.MultiInsertNodes(req.GraphID, node, nodeRecords)
		if err != nil {
			return "", err
		}
		// 合计剩余节点数目
		numNode += len(nodeRecords)
	}
	// 更新节点度数
	_, err = config.NebulaClient.MultiIncNodesDeg(req.GraphID, degMap)
	// 查询图原始大小，更新图大小
	size, err := config.MysqlClient.GetGraphSizeByID(ctx, req.GraphID)
	size.NumNode += int64(numNode)
	size.NumEdge += int64(numEdge)

	if err != nil {
		return "", err
	}
	_, err = config.MysqlClient.UpdateGraphByID(ctx, &model.Graph{
		ID:      req.GraphID,
		Status:  common.GraphStatusOK,
		NumNode: size.NumNode,
		NumEdge: size.NumEdge,
	})
	if err != nil {
		return "", err
	}

	resp := &types.GraphInfoReply{
		Base: &types.BaseReply{
			TaskID:     taskID,
			TaskStatus: int64(model.KVTask_Finished),
		},
		Graph: &types.Graph{
			Id:      req.GraphID,
			GroupID: group.ID,
			NumNode: size.NumNode,
			NumEdge: size.NumEdge,
		},
	}
	return jsonx.MarshalToString(resp)
}
