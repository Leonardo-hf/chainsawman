package svc

import (
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/types/rpc/algo"
	"chainsawman/graph/db"
	"chainsawman/graph/mq"
)

type ServiceContext struct {
	Config config.Config

	NebulaClient db.NebulaClient
	RedisClient  db.RedisClient
	MysqlClient  db.MysqlClient

	TaskMq  mq.TaskMq
	AlgoRPC algo.AlgoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		NebulaClient: db.InitNebulaClient(&c.Nebula),
		RedisClient:  db.InitRedisClient(&c.Redis),
		MysqlClient:  db.InitMysqlClient(&c.Mysql),
		TaskMq:       mq.InitTaskMq(&c.TaskMq),
		AlgoRPC:      nil, /*algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn())*/
	}
}
