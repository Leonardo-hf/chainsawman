package svc

import (
	"chainsawman/file/cmd/rpc/internal/config"
	"github.com/google/uuid"
)

type ServiceContext struct {
	Config config.Config
	IDGen  IDGenerator
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		IDGen:  &IDGeneratorImpl{},
	}
}

// TODO 放到其他包里去
type IDGenerator interface {
	New() string
}

type IDGeneratorImpl struct{}

func (g *IDGeneratorImpl) New() string {
	return uuid.New().String()
}
