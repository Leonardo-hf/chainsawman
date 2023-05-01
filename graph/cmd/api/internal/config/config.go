package config

import (
	"chainsawman/graph/db"
	"chainsawman/graph/mq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig

	TaskMq mq.TaskMqConfig
}
