package logic

import (
	"chainsawman/graph/model"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"chainsawman/common"
	"chainsawman/graph/api/internal/svc"
	"chainsawman/graph/api/internal/types"
	"chainsawman/graph/config"
	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 1 << 31

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Upload TODO: 改变为提交task
func (l *UploadLogic) Upload(r *http.Request) (resp *types.SearchGraphReply, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	graph := r.FormValue("graph")
	err = config.NebulaClient.CreateGraph(graph)
	if err != nil {
		return nil, err
	}
	file, _, err := r.FormFile("nodes")
	if err != nil {
		return nil, err
	}
	var nodes []*model.Node
	records, err := handle(file)
	for _, record := range records {
		name, nameOK := record["name"]
		desc, descOK := record["desc"]
		if nameOK && descOK {
			node := &model.Node{Name: name, Desc: desc}
			nodes = append(nodes, node)
		}
	}
	_, err = config.NebulaClient.MultiInsertNodes(graph, nodes)
	if err != nil {
		return nil, err
	}
	file, _, err = r.FormFile("edges")
	if err != nil {
		return nil, err
	}
	var edges []*model.Edge
	records, err = handle(file)
	for _, record := range records {
		source, sourceOK := record["source"]
		target, targetOK := record["target"]
		if sourceOK && targetOK {
			edge := &model.Edge{Source: source, Target: target}
			edges = append(edges, edge)
		}
	}
	_, err = config.NebulaClient.MultiInsertEdges(graph, edges)
	if err != nil {
		return nil, err
	}

	return &types.SearchGraphReply{
		Graph: &types.Graph{
			Name:  graph,
			Nodes: int64(len(nodes)),
			Edges: int64(len(edges)),
		},
	}, nil
}

func handle(file multipart.File) ([]common.Record, error) {
	path, err := tempSave(file)
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

func tempSave(file multipart.File) (string, error) {
	tempFile, err := os.CreateTemp("", "*.csv")
	defer tempFile.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	io.Copy(tempFile, file)
	return tempFile.Name(), nil
}
