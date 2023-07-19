package config

import (
	"chainsawman/common"
	"chainsawman/graph/db"
	"chainsawman/graph/mq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig
	Minio  db.MinioConfig

	TaskMqEd string
	TaskMq   mq.TaskMqConfig
	TaskMqV2 mq.AsynqConfig

	//AlgoRPC zrpc.RpcClientConf
}

func (c *Config) IsTaskV2Enabled() bool {
	return c.TaskMqEd == common.TaskMqEd2
}
