package mq

import (
	"chainsawman/graph/model"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
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
	_ = rdb.Del(ctx, cfg.Topic)
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

func (r *TaskMqImpl) ProduceTaskMsg(ctx context.Context, task *model.KVTask) (string, error) {
	cmd := r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: r.topic,
		Values: map[string]interface{}{
			"id":     task.Id,
			"idf":    task.Idf,
			"params": task.Params,
		},
		ID: strconv.FormatInt(task.Id, 10),
	})
	return cmd.Result()
}

func (r *TaskMqImpl) DelTaskMsg(ctx context.Context, id string) error {
	cmd := r.rdb.XDel(ctx, r.topic, id)
	return cmd.Err()
}
