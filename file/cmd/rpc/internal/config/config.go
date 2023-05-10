package config

import (
	"chainsawman/file/db"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Minio db.MinioConfig
}
