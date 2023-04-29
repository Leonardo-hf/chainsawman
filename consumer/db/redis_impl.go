package db

import (
	"chainsawman/consumer/model"
	"time"

	"context"
	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisClientImpl struct {
	rdb        *redis.Client
	topic      string
	group      string
	expiration time.Duration
}

type RedisConfig struct {
	Addr    string
	Topic   string
	Group   string
	Expired int64
}

func InitRedisClient(cfg *RedisConfig) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	rdb.XGroupCreate(context.Background(), cfg.Topic, cfg.Group, "0")
	return &RedisClientImpl{
		rdb:        rdb,
		topic:      cfg.Topic,
		group:      cfg.Group,
		expiration: time.Duration(cfg.Expired) * time.Second,
	}
}

func (r *RedisClientImpl) UpsertTask(ctx context.Context, task *model.KVTask) error {
	v, err := proto.Marshal(task)
	if err != nil {
		return err
	}
	cmd := r.rdb.Set(ctx, strconv.FormatInt(task.Id, 10), v, r.expiration)
	return cmd.Err()
}

// ConsumeTaskMsg TODO: 消费失败了消息会丢失，应该解决
func (r *RedisClientImpl) ConsumeTaskMsg(ctx context.Context, consumer string, handle func(ctx context.Context, task *model.KVTask) error) error {
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
		id, _ := strconv.ParseInt(msg.Values["id"].(string), 10, 64)
		task := &model.KVTask{
			Id: id,
			//Name:   msg.Values["name"].(string),
			Params: msg.Values["params"].(string),
		}
		if err = handle(ctx, task); err != nil {
			return err
		}
		cmd := r.rdb.XAck(ctx, r.topic, r.group, msg.ID)
		return cmd.Err()
	}
	return nil
}
