use go-zero, graphin, nebula and spark graphx to build a extended visual system for graph analysis

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
* algo rpc, interact with spark to resolve graph algorithm
  * `cd algo`
  * `mvn clean compile dependency:properties exec:exec@server`
* sca api, resolve requests about software composition analysis
* front, front-end, allow CRUD and inspect for graph
  * `npm run install`
  * `npm run dev`
### others:
* k8s, support deploying simply by k8s
* client, packaged client for interacting with graph service
### TODO LIST:
* multiple data source, connector
  * a default datasource for python requirements updated daily
* deploy by docker compose & kubernetes
* use openfass to deploy consumers
* data interface to format
* replace redis with abase