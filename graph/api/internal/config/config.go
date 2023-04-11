package config

import (
	"chainsawman/graph/db"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	NebulaConfig db.NebulaConfig
}
