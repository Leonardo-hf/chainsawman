package svc

import (
	"chainsawman/file/cmd/api/internal/config"
	"chainsawman/file/db"
	"github.com/google/uuid"
)

type ServiceContext struct {
	Config    config.Config
	IDGen     IDGenerator
	OSSClient db.OSSClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		IDGen:     &IDGeneratorImpl{},
		OSSClient: db.InitMinioClient(&c.Minio),
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
