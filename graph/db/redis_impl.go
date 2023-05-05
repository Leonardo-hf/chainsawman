package db

import (
	"chainsawman/graph/model"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
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

func (r *RedisClientImpl) GetTaskById(ctx context.Context, id int64) (*model.KVTask, error) {
	cmd := r.rdb.Get(ctx, strconv.FormatInt(id, 10))
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	res := &model.KVTask{}
	task, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal([]byte(task), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RedisClientImpl) UpsertTask(ctx context.Context, task *model.KVTask) error {
	v, err := proto.Marshal(task)
	if err != nil {
		return err
	}
	cmd := r.rdb.Set(ctx, strconv.FormatInt(task.Id, 10), string(v), r.expiration)
	return cmd.Err()
}

func (r *RedisClientImpl) DropTask(ctx context.Context, id int64) (int64, error) {
	cmd := r.rdb.Del(ctx, strconv.FormatInt(id, 10))
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	return cmd.Result()
}
