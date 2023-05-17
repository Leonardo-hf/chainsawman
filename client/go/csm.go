package _go

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/jsonx"
	"net/http"
	"strings"
)

type CSMClient struct {
	rdb     *redis.Client
	service string
	batch   int
}

type Graph struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (c *CSMClient) InitGraph(ctx context.Context, name string) (*Graph, error) {
	return c.InitGraphWithDesc(ctx, name, "")
}

func (c *CSMClient) InitGraphWithDesc(_ context.Context, name string, desc string) (*Graph, error) {
	payload := strings.NewReader(fmt.Sprintf("{\"graph\":%v, \"desc\":\"%v\"}", name, desc))
	req, _ := http.NewRequest("POST", c.service, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	r := make([]byte, 0)
	_, err = response.Body.Read(r)
	if err != nil {
		return nil, err
	}
	graph := &Graph{}
	err = jsonx.Unmarshal(r, graph)
	if err != nil {
		return nil, err
	}
	return graph, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *CSMClient) Creates(ctx context.Context, graphID int64, nodes NodesBody) error {
	for i := 0; i < len(nodes.Nodes); i += c.batch {
		msg := NodesBody{Nodes: nodes.Nodes[i:min(len(nodes.Nodes), i+c.batch)]}
		body, _ := jsonx.MarshalToString(msg)
		err := c.send(ctx, map[string]interface{}{
			"opt":      Creates,
			"graph_id": graphID,
			"body":     body,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CSMClient) Updates(ctx context.Context, graphID int64, edges EdgesBody) error {
	for i := 0; i < len(edges.Edges); i += c.batch {
		msg := EdgesBody{Edges: edges.Edges[i:min(len(edges.Edges), i+c.batch)]}
		body, _ := jsonx.MarshalToString(msg)
		err := c.send(ctx, map[string]interface{}{
			"opt":      Updates,
			"graph_id": graphID,
			"body":     body,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CSMClient) Deletes(ctx context.Context, graphID int64, nodes NodesBody) error {
	body, err := jsonx.MarshalToString(nodes)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"opt":      Deletes,
		"graph_id": graphID,
		"body":     body,
	})
}

func (c *CSMClient) InsertEdges(ctx context.Context, graphID int64, edges SplitEdgesBody) error {
	for i := 0; i < len(edges.Edges); i += c.batch {
		msg := SplitEdgesBody{Edges: edges.Edges[i:min(len(edges.Edges), i+c.batch)]}
		body, _ := jsonx.MarshalToString(msg)
		err := c.send(ctx, map[string]interface{}{
			"opt":      InsertEdges,
			"graph_id": graphID,
			"body":     body,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CSMClient) DeleteEdges(ctx context.Context, graphID int64, edges SplitEdgesBody) error {
	body, err := jsonx.MarshalToString(edges)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"opt":      DeleteEdges,
		"graph_id": graphID,
		"body":     body,
	})
}

func (c *CSMClient) send(ctx context.Context, values map[string]interface{}) error {
	cmd := c.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: topic,
		Values: values,
	})
	return cmd.Err()
}

type CSMBuilder struct {
	mq      string
	service string
	batch   int
}

func (b *CSMBuilder) Mq(addr string) *CSMBuilder {
	b.mq = addr
	return b
}

func (b *CSMBuilder) Service(addr string) *CSMBuilder {
	b.service = addr
	return b
}

func (b *CSMBuilder) Batch(batch int) *CSMBuilder {
	b.batch = batch
	return b
}

func (b *CSMBuilder) Build() *CSMClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: b.mq,
	})
	ctx := context.Background()
	err := rdb.XGroupCreateMkStream(ctx, topic, group, "0").Err()
	// TODO: 重复创建会报错，怎么避免？
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}
	return &CSMClient{
		rdb:     rdb,
		service: b.service,
		batch:   b.batch,
	}
}

const (
	topic = "import"
	group = "import_consumers"
)
