package config

import (
	"chainsawman/common"
	"chainsawman/consumer/task/db"
	"chainsawman/consumer/task/mq"
	"chainsawman/consumer/task/types/rpc/algo"
	"chainsawman/consumer/task/types/rpc/file"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig

	TaskMq   mq.TaskMqConfig
	TaskMqV2 mq.AsynqConfig

	TaskMqEd string

	FileRPC zrpc.RpcClientConf
	AlgoRPC zrpc.RpcClientConf
}

var (
	NebulaClient db.NebulaClient
	MysqlClient  db.MysqlClient
	RedisClient  db.RedisClient
)

var (
	FileRPC file.FileClient
	AlgoRPC algo.AlgoClient
)

var TaskMq mq.TaskMq

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)

	MysqlClient = db.InitMysqlClient(&c.Mysql)
	RedisClient = db.InitRedisClient(&c.Redis)

	FileRPC = file.NewFileClient(zrpc.MustNewClient(c.FileRPC).Conn())
	//AlgoRPC = algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn())

	if c.TaskMqEd == common.TaskMqEd {
		TaskMq = mq.InitTaskMq(&c.TaskMq)
	}
}

func (c *Config) IsTaskV2Enabled() bool {
	return c.TaskMqEd == common.TaskMqEd2
}
