package handler

import (
	"chainsawman/common"
	"chainsawman/consumer/task/config"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"strconv"
)

type UpdateGraph struct {
}

// TODO: 更新
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
	// TODO: 是否存在占用过多内存的问题？
	numNode, numEdge := 0, 0
	nodeMap := make(map[int64]*common.Record)
	nodeTypeList := make([]int64, 0)
	nodeRecordList := make([][]*common.Record, 0)
	// 读所有节点文件, 暂存节点数据
	for _, nodeType := range req.NodeFileList {
		t, nodeFileId := nodeType.Key, nodeType.Value
		tid, _ := strconv.ParseInt(t, 10, 64)
		file, err := config.OSSClient.Fetch(ctx, nodeFileId)
		if err != nil {
			return "", err
		}
		records, err := handle(file)
		if err != nil {
			return "", err
		}
		for _, r := range records {
			nodeID, err := r.GetAsInt(common.KeyID)
			if err != nil {
				return "", fmt.Errorf("update graph fail, source files of nodes needs `id`")
			}
			r.Put(common.KeyDeg, common.DefaultDeg)
			nodeMap[nodeID] = r
		}
		// 合计节点数目
		numNode += len(records)
		nodeTypeList = append(nodeTypeList, tid)
		nodeRecordList = append(nodeRecordList, records)
	}
	// 读所有边文件
	for _, edgeType := range req.EdgeFileList {
		t, edgeFileId := edgeType.Key, edgeType.Value
		tid, _ := strconv.ParseInt(t, 10, 64)
		file, err := config.OSSClient.Fetch(ctx, edgeFileId)
		if err != nil {
			return "", err
		}
		records, err := handle(file)
		edgeRecords := make([]*common.Record, 0)
		if err != nil {
			return "", err
		}
		// 如果边的两端节点不存在，则舍弃该边
		// 两端节点度数+1
		for _, r := range records {
			src, err := r.GetAsInt(common.KeySrc)
			if err != nil {
				return "", fmt.Errorf("update graph fail, source files of edges needs `source`")
			}
			tgt, err := r.GetAsInt(common.KeyTgt)
			if err != nil {
				return "", fmt.Errorf("update graph fail, source files of edges needs `target`")
			}
			var ok bool
			var srcRecord, tgtRecord *common.Record
			if srcRecord, ok = nodeMap[src]; !ok {
				continue
			}
			if tgtRecord, ok = nodeMap[tgt]; !ok {
				continue
			}
			edgeRecords = append(edgeRecords, r)
			v, _ := srcRecord.GetAsInt(common.KeyDeg)
			srcRecord.Put(common.KeyDeg, strconv.FormatInt(v+1, 10))
			v, _ = tgtRecord.GetAsInt(common.KeyDeg)
			tgtRecord.Put(common.KeyDeg, strconv.FormatInt(v+1, 10))
		}
		// 合计边数目
		numEdge += len(edgeRecords)
		// 直接插入边，nebula先插入边，再插入顶点
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
		_, err = config.NebulaClient.MultiInsertEdges(req.GraphID, edge, edgeRecords)
		if err != nil {
			return "", err
		}
	}

	// 插入节点
	for i, tid := range nodeTypeList {
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
		_, err = config.NebulaClient.MultiInsertNodes(req.GraphID, node, nodeRecordList[i])
		if err != nil {
			return "", err
		}
	}

	// 更新图状态
	_, err = config.MysqlClient.UpdateGraphByID(ctx, &model.Graph{
		ID:      req.GraphID,
		Status:  common.GraphStatusOK,
		NumNode: int64(numNode),
		NumEdge: int64(numEdge),
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
			NumNode: int64(numNode),
			NumEdge: int64(numEdge),
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
