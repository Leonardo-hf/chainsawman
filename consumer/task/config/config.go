package config

import (
	"chainsawman/common"
	"chainsawman/consumer/task/db"
	"chainsawman/consumer/task/mq"
)

type Config struct {
	Nebula db.NebulaConfig
	Redis  db.RedisConfig
	Mysql  db.MysqlConfig
	Minio  db.MinioConfig

	TaskMq   mq.TaskMqConfig
	TaskMqV2 mq.AsynqConfig

	TaskMqEd string
}

var (
	NebulaClient db.NebulaClient
	MysqlClient  db.MysqlClient
	RedisClient  db.RedisClient
	OSSClient    db.OSSClient
)

var TaskMq mq.TaskMq

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)
	MysqlClient = db.InitMysqlClient(&c.Mysql)
	RedisClient = db.InitRedisClient(&c.Redis)
	OSSClient = db.InitMinioClient(&c.Minio)

	if c.TaskMqEd == common.TaskMqEd {
		TaskMq = mq.InitTaskMq(&c.TaskMq)
	}
}

func (c *Config) IsTaskV2Enabled() bool {
	return c.TaskMqEd == common.TaskMqEd2
}
