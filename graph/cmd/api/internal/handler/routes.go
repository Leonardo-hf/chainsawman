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
				Method:  http.MethodPost,
				Path:    "/group/create",
				Handler: createGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/drop",
				Handler: dropGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: createGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: updateGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/drop",
				Handler: dropGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAll",
				Handler: getAllGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: getGraphInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: getGraphHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/getAll",
				Handler: getNodesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/getMatch",
				Handler: getMatchNodesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/getMatchByTag",
				Handler: getMatchNodesByTagHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/nbr",
				Handler: getNeighborsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/graph"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getAll",
				Handler: algoGetAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: algoCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/drop",
				Handler: algoDropHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exec",
				Handler: algoExecHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/get",
				Handler: getAlgoTaskByIDHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/getAll",
				Handler: getAlgoTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/task/drop",
				Handler: dropAlgoTaskHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/graph/algo"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/put/source",
				Handler: fileSourcePutPresignedHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/put/lib",
				Handler: fileLibPutPresignedHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/algo",
				Handler: fileAlgoGetPresignedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/graph/file"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hot",
				Handler: getHotHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/hhi",
				Handler: getHHIHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/graph/se"),
	)
}
