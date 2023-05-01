package mq

import (
	"chainsawman/graph/model"

	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type TaskMqImpl struct {
	rdb   *redis.Client
	topic string
	group string
}

type TaskMqConfig struct {
	Addr  string
	Topic string
	Group string
}

func InitTaskMq(cfg *TaskMqConfig) TaskMq {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	ctx := context.Background()
	err := rdb.XGroupCreateMkStream(ctx, cfg.Topic, cfg.Group, "0").Err()
	// TODO: 重复创建会报错，怎么避免？
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}
	return &TaskMqImpl{
		rdb:   rdb,
		topic: cfg.Topic,
		group: cfg.Group,
	}
}

func (r *TaskMqImpl) ProduceTaskMsg(ctx context.Context, task *model.KVTask) error {
	cmd := r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: r.topic,
		Values: map[string]interface{}{
			"id":     task.Id,
			"idf":    task.Idf,
			"params": task.Params,
		},
	})
	return cmd.Err()
}

func (r *TaskMqImpl) DelTaskMsg(ctx context.Context, id int64) error {
	cmd := r.rdb.XDel(ctx, r.topic, strconv.FormatInt(id, 10))
	return cmd.Err()
}
