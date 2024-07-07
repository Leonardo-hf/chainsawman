package main

import (
	"chainsawman/common"
	"chainsawman/graph/cmd/api/internal/config"
	"chainsawman/graph/cmd/api/internal/handler"
	"chainsawman/graph/cmd/api/internal/logic"
	"chainsawman/graph/cmd/api/internal/svc"
	"chainsawman/graph/cmd/api/internal/types"
	"chainsawman/graph/cmd/api/internal/util"
	"github.com/zeromicro/go-zero/core/logx"

	"context"
	"flag"
	"fmt"
	"os"
	"time"

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
	// 启动定时任务
	go schedule(ctx)

	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func schedule(ctx *svc.ServiceContext) {
	graphExisted := true
	// 准备图谱
	for graphID, info := range common.PreparedGraph {
		key := fmt.Sprintf("prepare_%d", graphID)
		// 获得锁
		if ok, _ := ctx.RedisClient.CheckIdempotent(context.Background(), key, time.Hour); ok {
			// 判断图谱是否存在
			if graph, _ := ctx.MysqlClient.GetGraphByID(context.Background(), graphID); graph == nil {
				graphExisted = false
				// 不存在则创建图谱
				_, _ = logic.NewCreateGraphLogic(context.Background(), ctx).CreateGraph(&types.CreateGraphRequest{
					GraphID: graphID,
					Graph:   info.Name,
					GroupID: info.GroupID,
				})
			}
		}
	}
	// 保障图谱创建完成
	if !graphExisted {
		time.Sleep(time.Minute)
	}
	// Debug 立刻执行一次任务
	//ctx.TaskMq.ProduceTaskMsg(context.Background(), &model.KVTask{Idf: common.CronPython})

	// 开始调度任务
	for _, idf := range []string{common.CronPython} {
		_, _ = util.CreateCronTask(context.Background(), ctx, idf)
	}
}
