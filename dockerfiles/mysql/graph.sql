create database if not exists graph;

use graph;

create table if not exists `group`
(
    id       int auto_increment
        primary key,
    name     varchar(255)  not null,
    `desc`   text          null,
    parentID int default 1 null comment '标识父策略组，子策略组继承父策略组的全部节点与边缘'
);

create table if not exists algo
(
    id        int auto_increment
        primary key,
    name      varchar(255)  not null,
    `desc`    varchar(255)  null,
    detail    text          null,
    jarPath   varchar(255)  null,
    mainClass varchar(255)  null,
    tag       varchar(255)  null,
    groupId   int default 1 null comment '约束算法应用于某个策略组的图谱，此外：
1......应用于全部策略组',
    constraint algos_groups_id_fk
        foreign key (groupId) references `group` (id)
            on update cascade on delete set default
);

create table if not exists algoParam
(
    id        int auto_increment
        primary key,
    algoID    int           not null,
    name      varchar(255)  not null,
    `desc`    varchar(255)  null,
    type      int default 0 not null,
    `default` varchar(255)        null,
    `max`       varchar(255)        null,
    `min`       varchar(255)        null,
    constraint algoParam___id
        foreign key (algoID) references algo (id)
            on update cascade on delete cascade
);

create table if not exists edge
(
    id        int auto_increment
        primary key,
    groupID   int                         not null,
    name      varchar(255)                not null,
    `desc`    text                        null,
    direct    tinyint(1)   default 1      not null,
    display   varchar(255) default 'real' null,
    `primary` varchar(255)                null,
    constraint edges_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists edgeAttr
(
    id     int auto_increment
        primary key,
    edgeID int           not null,
    name   varchar(255)  not null,
    `desc` text          null,
    type   int default 0 not null,
    constraint edges_attr_edges_id_fk
        foreign key (edgeID) references edge (id)
            on update cascade on delete cascade
);

create table if not exists graph
(
    id         int auto_increment
        primary key,
    name       varchar(255)                        not null,
    status     int       default 0                 not null,
    numNode    int       default 0                 not null,
    numEdge    int       default 0                 not null,
    groupID    int       default 0                 not null,
    createTime timestamp default CURRENT_TIMESTAMP null,
    updateTime timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint graphs_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists `exec`
(
    id         int auto_increment
        primary key,
    status     int       default 0                 null,
    params     varchar(1024)                       null,
    algoID     int                                 not null,
    graphID    int                                 not null,
    output     varchar(255)                        null,
    appID        varchar(255)                      null,
    updateTime timestamp default CURRENT_TIMESTAMP not null,
    createTime timestamp default CURRENT_TIMESTAMP not null,
    constraint exec_algos_id_fk
        foreign key (algoID) references algo (id),
    constraint exec_graphs_id_fk
        foreign key (graphID) references graph (id)
)
    comment '算法执行结果';

create table if not exists node
(
    id        int auto_increment
        primary key,
    groupID   int                          not null,
    name      varchar(255)                 not null,
    `desc`    text                         null,
    display   varchar(255) default 'color' null,
    `primary` varchar(255)                 null,
    constraint nodes_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists nodeAttr
(
    id     int auto_increment
        primary key,
    nodeID int           not null,
    name   varchar(255)  not null,
    `desc` text          null,
    type   int default 0 not null,
    constraint nodes_attr_nodes_id_fk
        foreign key (nodeID) references node (id)
            on update cascade on delete cascade
);