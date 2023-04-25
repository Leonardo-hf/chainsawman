use go-zero, graphin, nebula and spark graphx to build a extended visual system for graph analysis

#### ONLY SUPPORT LOCAL DEPLOY CURRENTLY !!

### pre-prep
* [nebula](https://www.nebula-graph.com.cn/database), mysql, redis
* npm, maven, goctl, protobuf, etcd
* spark


### services list:
* graph api, resolve requests from front-end
  * `run graph/scripts/graph.sql`
  * `go run graph/cmd/api/graph.go`
* file api, allow front-end to upload files
  * `run file/cmd/api/file.go`
* file rpc, allow back-end to fetch files
  * `run file/cmd/rpc/file/go`
* consumer, exec requests which cost time
  * `run consumer/main.go`
* algo rpc, interact with spark to resolve graph algorithm
  * `cd algo`
  * `mvn clean compile dependency:properties exec:exec@server`
* front, front-end, allow CRUD and inspect for graph

### TODO LIST:
* support for graph algorithm
* multiple data source, connector
  * a default datasource for python requirements updated daily
* hdfs or cloud rds for file service
* register algo rpc on ETCD, or use kubernetes