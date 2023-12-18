#!/bin/bash
# api
goctl api go -api graph/cmd/api/graph.api -dir graph/cmd/api/
cp graph/cmd/api/internal/types/types.go consumer/task/types/
# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename graph.json" -api graph/cmd/api/graph.api -dir graph/cmd/api/
sed -i '/requestBody/d' graph/cmd/api/graph.json
cp graph/cmd/api/graph.json front/
# gorm gen
# TODO: mysql -uroot -p${MYSQL_ROOT_PASSWORD}
cp graph/scripts/graph.sql dockerfiles/mysql/
cp -r graph/scripts/gsql/ dockerfiles/mysql/
go run graph/scripts/gen.go
go run consumer/task/scripts/gen.go