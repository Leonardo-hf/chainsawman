package db

import (
	"chainsawman/consumer/task/model"
	"time"

	"context"
	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisClientImpl struct {
	rdb        *redis.Client
	expiration time.Duration
}

type RedisConfig struct {
	Addr    string
	Expired int64
}

func InitRedisClient(cfg *RedisConfig) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	return &RedisClientImpl{
		rdb:        rdb,
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
