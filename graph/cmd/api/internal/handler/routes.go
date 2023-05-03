// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"chainsawman/graph/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getAll",
				Handler: getAllGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: getGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/get",
				Handler: getNeighborsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/drop",
				Handler: dropGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: createGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/degree",
				Handler: algoDegreeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/pr",
				Handler: algoPageRankHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/vr",
				Handler: algoVoteRankHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/betweenness",
				Handler: algoBetweennessHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/closeness",
				Handler: algoClosenessHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/avgCC",
				Handler: algoAvgCCHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/louvain",
				Handler: algoLouvainHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/getAll",
				Handler: getGraphTasksHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/createEmpty",
				Handler: createEmptyGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/algo/getAll",
				Handler: algoGetAllHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/graph"),
	)
}
