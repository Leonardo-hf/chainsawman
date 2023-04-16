package config

import (
	"chainsawman/consumer/db"
	file "chainsawman/consumer/types/rpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig

	FileRPC zrpc.RpcClientConf
}

var NebulaClient db.NebulaClient

var MysqlClient db.MysqlClient

var RedisClient db.RedisClient

var FileRPC file.FileClient

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)
	MysqlClient = db.InitMysqlClient(&c.Mysql)
	RedisClient = db.InitRedisClient(&c.Redis)

	FileRPC = file.NewFileClient(zrpc.MustNewClient(c.FileRPC).Conn())
}
