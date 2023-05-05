package config

import (
	"chainsawman/consumer/connector/db"
	"chainsawman/consumer/connector/mq"
)

type Config struct {
	Nebula db.NebulaConfig

	ImportMq mq.ImportMqConfig
}

var (
	NebulaClient db.NebulaClient
	ImportMq     mq.ImportMq
)

func Init(c *Config) {
	NebulaClient = db.InitNebulaClient(&c.Nebula)
	ImportMq = mq.InitImportMq(&c.ImportMq)
}
