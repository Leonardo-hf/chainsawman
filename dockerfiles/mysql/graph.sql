create database if not exists graph;

drop table if exists graph.tasks;
drop table if exists graph.nodes_attr;
drop table if exists graph.nodes;
drop table if exists graph.edges_attr;
drop table if exists graph.edges;
drop table if exists graph.graphs;
drop table if exists graph.groups;

create table if not exists graph.groups
(
    id     int auto_increment
        primary key,
    name   varchar(255) not null,
    `desc` text         null
);

INSERT INTO graph.groups(name, `desc`)
VALUES ("group_default", "默认分组");

create table if not exists graph.graphs
(
    id         int auto_increment
        primary key,
    name       varchar(255)       not null,
    `desc`     text               null,
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
VALUES (1, "node_default", "标准节点");

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
    direct  tinyint(1) default 0 not null,
    display varchar(255)         null,
    constraint edges_groups_id_fk
        foreign key (groupID) references `groups` (id)
            on update cascade on delete cascade
);

INSERT INTO graph.edges(groupID, name, `desc`)
VALUES (1, "edge_default", "标准边");
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

