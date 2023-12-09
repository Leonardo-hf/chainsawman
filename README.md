使用 go-zero, graphin, nebula 和 spark graphx 构建一个通用、可扩展的在线“图可视化+分析”系统。特别地，基于软件依赖、元数据构建图谱并提供一系列分析方法

> use go-zero, graphin, nebula and spark graphx to build an common and extended visual system for graph analysis. Specially, we support for the graph based on dependencies and metadata of software and provide a few analysis method.

### pre-prep

* 语言（language）：java(maven), go, python, js(npm)
* 数据库（database）：[nebula](https://www.nebula-graph.com.cn/database), mysql, redis, minio
* 批处理（batch calculate）：livy, spark

### services list:

* graph，HTTP 服务，处理图谱结构、图算法 CRUD 的前端请求，将更新图谱节点与边、执行图算法等高耗时操作交付 consumer 处理

  > http service, handle CRUD about structures and algorithms of graphs, passing time cost tasks like inserting nodes and edges, executing algorithms to consumers

  * `run graph/scripts/graph.sql`
  * `go run graph/cmd/api/graph.go`

* consumer，消费者，通过基于 redis 的消息队列获取高耗时任务，并写入任务结果

  > consumer, handle time cost tasks passed by a task queue based on redis and then write result

  * `go run consumer/connector/main.go`
  * `go run consumer/task/main.go`

* sca，HTTP 服务，提供对 java、python、go 语言的代码的依赖树解析功能

  > http service, resolve requests about software composition analysis

  * `python sca/main.py`

* front，前端，提供图谱可视化和执行图算法的功能

  > front-end, provide an inspect for graph

  * `cd front`
  * `npm run install`
  * `npm run dev`

### others:

* k8s, 支持通过 K8S 进行简单的部署

  > support deploying simply by k8s

* dockerfiles，支持通过 docker-compose 进行部署

  > support deploying simply by docker-compose

* algo，图算法以 jar 文件的形式发送到 spark 执行

  > spark algorithms in jar

  * common，为编写图算法提供一些输入输出的方法

    > manage common interface with nebula(input) & minio(output)

  * others，提供一些已编写好的图算法，部分来自 nebula-algorithm

    > a series of algorithms provided, with a part from nebula-algorithm

### TODO LIST:

* 基于 PYPI RSS 提供一个每日更新的 python 依赖图谱

  > support a dependency graph of python software updated daily base on PYPI RSS

* client，提供 SDK 对线上图谱进行增删改查

  > packaged client for interacting with graph service

* 基于 openfaas 部署消费者

  > use openfass to deploy consumers

* 验证 graph 服务中的请求

  > validate data interface of graph

* Mysql 存在慢查询问题，可能是外键设置过多的原因

  > solve problems of slow query in mysql

* nebula 可能需要全文索引

  > use full text index to make query nodes quicker