package config

import (
	"os"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Path        string
	Expired     int
	IDGenerator uuid.UUID
}

func (c *Config) initUUID() {
	c.IDGenerator = uuid.New()
}

func (c *Config) initStorage() {
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

func (c *Config) Init() {
	c.initUUID()
	c.initStorage()
}
