package config

import (
	"chainsawman/consumer/db"
	"chainsawman/consumer/types/rpc/algo"
	"chainsawman/consumer/types/rpc/file"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig

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

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)
	MysqlClient = db.InitMysqlClient(&c.Mysql)
	RedisClient = db.InitRedisClient(&c.Redis)

	FileRPC = file.NewFileClient(zrpc.MustNewClient(c.FileRPC).Conn())
	//AlgoRPC = algo.NewAlgoClient(zrpc.MustNewClient(c.AlgoRPC).Conn())
}
