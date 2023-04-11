package svc

import (
	"chainsawman/graph/api/internal/config"
	"chainsawman/graph/db"
)

type ServiceContext struct {
	Config       config.Config
	NebulaClient db.NebulaClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		NebulaClient: db.InitNebulaClient(&c.NebulaConfig),
	}
}
