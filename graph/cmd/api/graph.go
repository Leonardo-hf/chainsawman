package main

import (
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/handler"
	"chainsawman/graph/cmd/api/internal/svc"
	"os"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	flag.Parse()
	defaultCfg := "graph/cmd/api/etc/graph.yaml"
	switch os.Getenv("CHS_ENV") {
	case "pre":
		defaultCfg = "graph/cmd/api/etc/graph-pre.yaml"
	}
	var configFile = flag.String("f", defaultCfg, "the config api")
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
