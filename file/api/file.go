package main

import (
	"flag"
	"fmt"

	"chainsawman/file/api/internal/config"
	"chainsawman/file/api/internal/handler"
	"chainsawman/file/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file.yaml", "the config api")

func main() {
	flag.Parse()

	var c config.Config
	_ = conf.Load(*configFile, &c)
	c.Init()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
