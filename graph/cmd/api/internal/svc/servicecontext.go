package svc

import (
	"chainsawman/common"
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
	svc := &ServiceContext{
		Config:       c,
		NebulaClient: db.InitNebulaClient(&c.Nebula),
		RedisClient:  db.InitRedisClient(&c.Redis),
		MysqlClient:  db.InitMysqlClient(&c.Mysql),
		//AlgoRPC: algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn()),
	}
	if c.TaskMqEd == common.TaskMqEd {
		svc.TaskMq = mq.InitTaskMq(&c.TaskMq)
	} else {
		svc.TaskMq = mq.InitTaskMqV2(&c.TaskMqV2)
	}
	return svc
}
