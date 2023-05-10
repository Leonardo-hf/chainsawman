package config

import (
	"chainsawman/file/db"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Minio db.MinioConfig
}
