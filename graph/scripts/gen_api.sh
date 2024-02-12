#!/bin/bash
# api
API="graph/cmd/api"
SCRIPT="graph/scripts"
goctl api go -api ${API}/graph.api -dir ${API}
cp ${API}/internal/types/types.go consumer/task/types/
# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename graph.json" -api ${API}/graph.api -dir ${API}
sed -i '/requestBody/d' ${API}/graph.json
cp ${API}/graph.json front/
# gorm gen
# TODO: mysql -uroot -p${MYSQL_ROOT_PASSWORD}
GSQL="${SCRIPT}/gsql"
INIT_SQL="dockerfiles/mysql/graph.sql"
cp ${SCRIPT}/graph.sql ${INIT_SQL}
for file in `ls ${GSQL}`
do
if test -f ${GSQL}/${file}
then
  cat ${GSQL}/${file} >> ${INIT_SQL}
fi
done
go run ${SCRIPT}/gen.go
go run consumer/task/scripts/gen.go