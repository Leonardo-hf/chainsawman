package config

import (
	"chainsawman/common"
	"chainsawman/consumer/task/db"
	"chainsawman/consumer/task/mq"
	"chainsawman/consumer/task/types/rpc/algo"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig
	Minio  db.MinioConfig

	TaskMq   mq.TaskMqConfig
	TaskMqV2 mq.AsynqConfig

	TaskMqEd string

	AlgoRPC zrpc.RpcClientConf
}

var (
	NebulaClient db.NebulaClient
	MysqlClient  db.MysqlClient
	RedisClient  db.RedisClient
	OSSClient    db.OSSClient
)

var (
	AlgoRPC algo.AlgoClient
)

var TaskMq mq.TaskMq

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)
	MysqlClient = db.InitMysqlClient(&c.Mysql)
	RedisClient = db.InitRedisClient(&c.Redis)
	OSSClient = db.InitMinioClient(&c.Minio)

	AlgoRPC = algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn())

	if c.TaskMqEd == common.TaskMqEd {
		TaskMq = mq.InitTaskMq(&c.TaskMq)
	}
}

func (c *Config) IsTaskV2Enabled() bool {
	return c.TaskMqEd == common.TaskMqEd2
}
