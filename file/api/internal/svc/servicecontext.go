package svc

import (
	"chainsawman/file/api/internal/config"
	"github.com/google/uuid"
	"os"
)

type ServiceContext struct {
	Config config.Config
	IDGen  IDGenerator
}

func initStorage(c *config.Config) {
	if stat, err := os.Stat(c.Path); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(c.Path, os.ModePerm); err != nil {
				panic(err)
			}
		}
	} else {
		if !stat.IsDir() {
			panic("Existing file with the same name with path.")
		}
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	initStorage(&c)
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
