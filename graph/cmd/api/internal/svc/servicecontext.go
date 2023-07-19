package svc

import (
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/types/rpc/algo"
	"chainsawman/graph/db"
	"chainsawman/graph/mq"
	"github.com/google/uuid"
)

type ServiceContext struct {
	Config config.Config

	NebulaClient db.NebulaClient
	RedisClient  db.RedisClient
	MysqlClient  db.MysqlClient
	OSSClient    db.OSSClient

	TaskMq  mq.TaskMq
	AlgoRPC algo.AlgoClient

	IDGen IDGenerator
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:       c,
		NebulaClient: db.InitNebulaClient(&c.Nebula),
		RedisClient:  db.InitRedisClient(&c.Redis),
		MysqlClient:  db.InitMysqlClient(&c.Mysql),
		OSSClient:    db.InitMinioClient(&c.Minio),
		IDGen:        &IDGeneratorImpl{},
		//AlgoRPC: algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn()),
	}
	if c.IsTaskV2Enabled() {
		svc.TaskMq = mq.InitTaskMqV2(&c.TaskMqV2)
	} else {
		svc.TaskMq = mq.InitTaskMq(&c.TaskMq)
	}
	return svc
}

// TODO 放到其他包里去
type IDGenerator interface {
	New() string
}

type IDGeneratorImpl struct{}

func (g *IDGeneratorImpl) New() string {
	return uuid.New().String()
}
