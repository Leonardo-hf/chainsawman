package config

import (
	"chainsawman/graph/db"
	"chainsawman/graph/mq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig

	TaskMq mq.TaskMqConfig

	AlgoRPC zrpc.RpcClientConf
}
