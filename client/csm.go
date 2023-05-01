package client

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

func (c *CSMClient) CreateNode(ctx context.Context, graphID int64, node NodeBody) error {
	body, err := jsonx.MarshalToString(node)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"optFlag":    Node,
		"entityFlag": Create,
		"graphID":    graphID,
		"body":       body,
	})
}

func (c *CSMClient) UpdateNode(ctx context.Context, graphID int64, node NodeBody) error {
	body, err := jsonx.MarshalToString(node)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"optFlag":    Node,
		"entityFlag": Update,
		"graphID":    graphID,
		"body":       body,
	})
}

func (c *CSMClient) DeleteNode(ctx context.Context, graphID int64, nodeID int64) error {
	body, err := jsonx.MarshalToString(&NodeBody{ID: nodeID})
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"optFlag":    Delete,
		"entityFlag": Create,
		"graphID":    graphID,
		"body":       body,
	})
}

func (c *CSMClient) CreateEdge(ctx context.Context, graphID int64, edge EdgeBody) error {
	body, err := jsonx.MarshalToString(edge)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"optFlag":    Edge,
		"entityFlag": Create,
		"graphID":    graphID,
		"body":       body,
	})
}

func (c *CSMClient) DeleteEdge(ctx context.Context, graphID int64, edge EdgeBody) error {
	body, err := jsonx.MarshalToString(edge)
	if err != nil {
		return err
	}
	return c.send(ctx, map[string]interface{}{
		"optFlag":    Edge,
		"entityFlag": Delete,
		"graphID":    graphID,
		"body":       body,
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
}

func (b *CSMBuilder) Mq(addr string) *CSMBuilder {
	b.mq = addr
	return b
}

func (b *CSMBuilder) Service(addr string) *CSMBuilder {
	b.service = addr
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
	}
}

const (
	topic = "import"
	group = "import_consumers"
)
