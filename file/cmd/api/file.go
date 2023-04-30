package main

import (
	"chainsawman/file/cmd/api/internal/config"
	"chainsawman/file/cmd/api/internal/handler"
	"chainsawman/file/cmd/api/internal/svc"

	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "file/cmd/api/etc/file.yaml", "the config api")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/file/get/:1",
			Handler: dirhandler("/api/file/get", c.Path),
		})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func dirhandler(pattern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		h := http.StripPrefix(pattern, http.FileServer(http.Dir(filedir)))
		h.ServeHTTP(w, req)
	}
}
