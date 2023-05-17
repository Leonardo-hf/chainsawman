package mq

import (
	"chainsawman/consumer/task/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	//rdb.XGroupCreate(context.Background(), cfg.Topic, cfg.Group, "0")
	return &TaskMqImpl{
		rdb:   rdb,
		topic: cfg.Topic,
		group: cfg.Group,
	}
}

// ConsumeTaskMsg TODO: 消费失败了消息会丢失，应该解决
func (r *TaskMqImpl) ConsumeTaskMsg(ctx context.Context, consumer string, handle func(ctx context.Context, task *model.KVTask) error) error {
	result, err := r.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    r.group,
		Streams:  []string{r.topic, ">"},
		Consumer: consumer,
		Count:    1,
	}).Result()
	if err != nil {
		return err
	}
	for _, msg := range result[0].Messages {
		logx.Info(msg)
		createTime, _ := strconv.ParseInt(msg.Values["create_time"].(string), 10, 64)
		task := &model.KVTask{
			Id:         msg.Values["id"].(string),
			Idf:        msg.Values["idf"].(string),
			Params:     msg.Values["params"].(string),
			CreateTime: createTime,
		}
		if err = handle(ctx, task); err != nil {
			return err
		}
		cmd := r.rdb.XAck(ctx, r.topic, r.group, msg.ID)
		return cmd.Err()
	}
	return nil
}
