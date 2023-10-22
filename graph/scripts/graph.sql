create database if not exists graph;

drop table if exists graph.tasks;
drop table if exists graph.nodes_attr;
drop table if exists graph.nodes;
drop table if exists graph.edges_attr;
drop table if exists graph.edges;
drop table if exists graph.graphs;
drop table if exists graph.groups;
drop table if exists graph.algos;
drop table if exists graph.algos_param;

create table if not exists graph.groups
(
    id     int auto_increment
        primary key,
    name   varchar(255) not null,
    `desc` text         null
);

INSERT INTO graph.groups(name, `desc`)
VALUES ("default", "默认图谱");

create table if not exists graph.graphs
(
    id         int auto_increment
        primary key,
    name       varchar(255)       not null,
    status     int      default 0 not null,
    numNode    int      default 0 not null,
    numEdge    int      default 0 not null,
    groupID    int      default 0 not null,
    createTime datetime DEFAULT CURRENT_TIMESTAMP,
    updateTime datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    constraint graphs_groups_id_fk
        foreign key (groupID) references `groups` (id)
            on update cascade on delete cascade
);

create table if not exists graph.nodes
(
    id      int auto_increment
        primary key,
    groupID int          not null,
    name    varchar(255) not null,
    `desc`  text         null,
    display varchar(255) null,
    constraint nodes_groups_id_fk
        foreign key (groupID) references `groups` (id)
            on update cascade on delete cascade
);

INSERT INTO graph.nodes(groupID, name, `desc`)
VALUES (1, "default", "标准节点");

create table if not exists graph.nodes_attr
(
    id        int auto_increment
        primary key,
    nodeID    int                  not null,
    name      varchar(255)         not null,
    `desc`    text                 null,
    type      int        default 0 not null,
    `primary` tinyint(1) default 0 not null,
    constraint nodes_attr_nodes_id_fk
        foreign key (nodeID) references nodes (id)
            on update cascade on delete cascade
);

INSERT INTO graph.nodes_attr(nodeID, name, `desc`, type, `primary`)
VALUES (1, "name", "名称", 0, 1);
INSERT INTO graph.nodes_attr(nodeID, name, `desc`, type, `primary`)
VALUES (1, "desc", "描述", 0, 0);

create table if not exists graph.edges
(
    id      int auto_increment
        primary key,
    groupID int                  not null,
    name    varchar(255)         not null,
    `desc`  text                 null,
    direct  tinyint(1) default 1 not null,
    display varchar(255)         null,
    constraint edges_groups_id_fk
        foreign key (groupID) references `groups` (id)
            on update cascade on delete cascade
);

INSERT INTO graph.edges(groupID, name, `desc`)
VALUES (1, "default", "标准边");
create table if not exists graph.edges_attr
(
    id        int auto_increment
        primary key,
    edgeID    int                  not null,
    name      varchar(255)         not null,
    `desc`    text                 null,
    type      int        default 0 not null,
    `primary` tinyint(1) default 0 not null,
    constraint edges_attr_edges_id_fk
        foreign key (edgeID) references edges (id)
            on update cascade on delete cascade
);

create table if not exists graph.tasks
(
    id          int auto_increment
        primary key,
    params      varchar(1024) null,
    status      int default 0 not null,
    result      mediumtext    null,
    graphID     int           not null,
    visible     tinyint(1)    null,
    tid         varchar(255)  null,
    idf         varchar(255)  not null,
    createTime datetime DEFAULT CURRENT_TIMESTAMP,
    updateTime datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    constraint task___graph_fk
        foreign key (graphID) references graph.graphs (id)
            on update cascade on delete cascade
);

create table if not exists graph.algos
(
    id        int auto_increment
        primary key,
    name      varchar(255)      not null,
    `desc`    text              null,
    type      int     default 0 not null,
    jarPath   varchar(255)      null,
    mainClass varchar(255)      null,
    isCustom  tinyint default 0 null
);

INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (1, "degree", "度中心度算法，测量网络中一个节点与所有其它节点相联系的程度。", 0, "s3a://lib/degree-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (2, "pagerank", "PageRank是Google使用的对其搜索引擎搜索结果中的网页进行排名的一种算法。能够衡量集合范围内某一元素的相关重要性。", 0, "s3a://lib/pagerank-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (3, "betweenness", "中介中心性用于衡量一个顶点出现在其他任意两个顶点对之间的最短路径的次数。", 0, "s3a://lib/betweenness-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (4, "closeness", "接近中心性反映在网络中某一节点与其他节点之间的接近程度。", 0, "s3a://lib/closeness-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (5, "average clustering coefficient", "平均聚类系数。描述图中的节点与其相连节点之间的聚集程度。", 2, "s3a://lib/clusteringCoefficient-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (6, "louvain", "一种基于模块度的社区发现算法。其基本思想是网络中节点尝试遍历所有邻居的社区标签，并选择最大化模块度增量的社区标签。", 1, "s3a://lib/louvain-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (7, "quantity", "广度排序算法，基于假设：节点入度越大越重要。使用邻居意见的Voterank算法衡量节点的相对入度", 0, "s3a://lib/voterank-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (8, "depth", "深度排序算法，基于假设：在更多路径中处于头部的节点更重要。使用改进的closeness算法衡量节点在头部的程度", 0, "s3a://lib/depth-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (9, "integration", "中介度排序算法，基于假设：在更多路径中处于中部的节点更重要。使用改进的betweenness算法衡量节点中介的程度", 0, "s3a://lib/betweenness-latest.jar", "applerodite.Main");
INSERT INTO graph.algos(id, name, `desc`, type, jarPath, mainClass)
VALUES (10, "ecology", "子图稳定性排序算法，基于假设：具有高稳定性的衍生子图的节点更重要。使用基于最小渗流的collective influence算法计算子图稳定性", 0, "s3a://lib/ecology-latest.jar", "applerodite.Main");

create table if not exists graph.algos_param
(
    id        int auto_increment
        primary key,
    algoID    int           not null,
    fieldName varchar(255)  not null,
    fieldDesc varchar(255)  null,
    fieldType int default 0 not null,
    initValue double        null,
    `max`       double        null,
    `min`       double        null,
    constraint algoParam___id
        foreign key (algoID) references algos (id)
            on update cascade on delete cascade
);

INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`, `max`)
VALUES (2, "iter", "迭代次数", 2, 3, 1, 100);
INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`, `max`)
VALUES (2, "prob", "阻尼系数", 0, 0.85, 0.1, 1);
INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`, `max`)
VALUES (6, "maxIter", "外部迭代次数", 2, 10, 1, 100);
INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`, `max`)
VALUES (6, "internalIter", "内部迭代次数", 2, 5, 1, 50);
INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`, `max`)
VALUES (6, "tol", "最小增加量", 0, 0.3, 0.1, 1);
INSERT INTO graph.algos_param(algoID, fieldName, fieldDesc, fieldType, initValue, `min`)
VALUES (7, "iter", "迭代次数", 2, 1, 100);