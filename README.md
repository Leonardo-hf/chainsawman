use go-zero, graphin, nebula and spark graphx to build an extended visual system for graph analysis

### pre-prep
* java(maven), go, python, js(npm)
* [nebula](https://www.nebula-graph.com.cn/database), mysql, redis, minio
* spark

### services list:
* graph api, resolve requests from front-end
  * `run graph/scripts/graph.sql`
  * `go run graph/cmd/api/graph.go`
* consumer, exec requests which cost time
  * `go run consumer/connector/main.go`
  * `go run consumer/task/main.go`
* sca api, resolve requests about software composition analysis
  * `python sca/main.py`
* front, front-end, allow CRUD and inspect for graph
  * `cd front`
  * `npm run install`
  * `npm run dev`
### others:
* k8s, support deploying simply by k8s[TODO]
* dockerfiles, support deploying simply by docker-compose
* client, packaged client for interacting with graph service
* algo, spark algo in jar
  * common, manage interface with nebula(input) & minio(output)
  * others, algo logic
### TODO LIST:
* multiple data source, connector
  * a default datasource for python requirements updated daily
* use openfass to deploy consumers
* format data interface
* may replace redis with abase in future
* Mysql 存在慢查询问题
* nebula 需要全文索引
* 父策略组