package main

import (
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/handler"
	"chainsawman/graph/cmd/api/internal/svc"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "graph/cmd/api/etc/graph.yaml", "the config api")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}