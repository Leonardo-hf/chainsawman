#!/bin/bash
# api
goctl api go -api graph/cmd/api/graph.api -dir graph/cmd/api/
cp graph/cmd/api/internal/types/types.go consumer/task/types/
# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename graph.json" -api graph/cmd/api/graph.api -dir graph/cmd/api/
sed -i '/requestBody/d' graph/cmd/api/graph.json
cp graph/cmd/api/graph.json front/
# algo grpc
protoc  algo/src/main/protobuf/algo.proto --go_out=consumer/task/types --go-grpc_out=consumer/task/types
# gorm gen
# TODO: mysql -uroot -p${MYSQL_ROOT_PASSWORD}
cp graph/scripts/graph.sql dockerfiles/mysql/graph.sql
go run graph/scripts/gen.go
go run consumer/task/scripts/gen.go