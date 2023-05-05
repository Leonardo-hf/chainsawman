package mq

import (
	"chainsawman/consumer/connector/msg"
	"fmt"

	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type ImportMqImpl struct {
	rdb   *redis.Client
	topic string
	group string
}

type ImportMqConfig struct {
	Addr  string
	Topic string
	Group string
}

func InitImportMq(cfg *ImportMqConfig) ImportMq {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	rdb.XGroupCreate(context.Background(), cfg.Topic, cfg.Group, "0")
	return &ImportMqImpl{
		rdb:   rdb,
		topic: cfg.Topic,
		group: cfg.Group,
	}
}

// ConsumeImportMsg TODO: 消费失败了消息会丢失，应该解决
func (r *ImportMqImpl) ConsumeImportMsg(ctx context.Context, consumer string, handle func(ctx context.Context, m *msg.Msg) error) error {
	result, err := r.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    r.group,
		Streams:  []string{r.topic, ">"},
		Consumer: consumer,
		Count:    1,
	}).Result()
	if err != nil {
		return err
	}
	for _, mi := range result[0].Messages {
		fmt.Println(mi)
		opt, _ := strconv.Atoi(mi.Values["Opt"].(string))
		entity, _ := strconv.Atoi(mi.Values["Entity"].(string))
		m := &msg.Msg{
			Opt:    msg.OptFlag(opt),
			Entity: msg.EntityFlag(entity),
			Body:   mi.Values["body"].(string),
		}
		if err = handle(ctx, m); err != nil {
			return err
		}
		cmd := r.rdb.XAck(ctx, r.topic, r.group, mi.ID)
		return cmd.Err()
	}
	return nil
}
