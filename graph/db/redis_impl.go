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
		expiration: time.Duration(cfg.Expired) * time.Second,
		group:      cfg.Group,
	}
}

func (r *RedisClientImpl) GetTaskById(ctx context.Context, id int64) (*model.KVTask, error) {
	cmd := r.rdb.Get(ctx, strconv.FormatInt(id, 10))
	if cmd.Err() == redis.Nil {
		return nil, nil
	} else if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	res := &model.KVTask{}
	err := proto.Unmarshal([]byte(cmd.String()), res)
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
	cmd := r.rdb.Set(ctx, strconv.FormatInt(task.Id, 10), v, r.expiration)
	return cmd.Err()
}

func (r *RedisClientImpl) ProduceTaskMsg(ctx context.Context, task *model.KVTask) error {
	cmd := r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: r.topic,
		Values: map[string]interface{}{
			"id":     task.Id,
			"name":   task.Name,
			"params": task.Params,
		},
	})
	return cmd.Err()
}

func (r *RedisClientImpl) DelTaskMsg(ctx context.Context, id int64) error {
	cmd := r.rdb.XDel(ctx, r.topic, strconv.FormatInt(id, 10))
	return cmd.Err()
}

func (r *RedisClientImpl) ConsumeTaskMsg(ctx context.Context, consumer string, handle func(task *model.KVTask) error) error {
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
		task := &model.KVTask{
			Id:     msg.Values["id"].(int64),
			Name:   msg.Values["name"].(string),
			Params: msg.Values["params"].(string),
		}
		if err = handle(task); err == nil {
			cmd := r.rdb.XAck(ctx, r.topic, r.group, msg.ID)
			return cmd.Err()
		}
	}
	return nil
}
