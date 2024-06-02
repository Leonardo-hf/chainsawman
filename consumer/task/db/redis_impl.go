package db

import (
	"chainsawman/consumer/task/model"
	"fmt"
	"time"

	"context"
	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
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
	cmd := r.rdb.Set(ctx, task.Id, v, r.expiration)
	return cmd.Err()
}

func (r *RedisClientImpl) DropDuplicate(ctx context.Context, set string, keys []string, expired time.Duration) ([]string, error) {
	formatKeys := make([]string, len(keys))
	msetReq := make(map[string]interface{})
	for i, k := range keys {
		formatKeys[i] = fmt.Sprintf("%v_%v", set, k)
		msetReq[formatKeys[i]] = 1
	}
	// 查询哪些Key存在
	cmd := r.rdb.MGet(ctx, formatKeys...)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	// 筛选不存在的key
	newKeys := make([]string, 0)
	for i, r := range res {
		if r == nil {
			newKeys = append(newKeys, keys[i])
		}
	}
	// 更新过期时间
	_ = r.rdb.MSet(ctx, msetReq, expired)
	return newKeys, err
}
