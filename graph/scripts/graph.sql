create table if not exists graph.graphs
(
    id     int auto_increment
        primary key,
    name   varchar(255)  not null,
    `desc` text          null,
    status int default 0 not null,
    nodes  int           null,
    edges  int           null
);

create table if not exists graph.task
(
    id          int auto_increment
        primary key,
    params      varchar(1024) null,
    status      int default 0 not null,
    result      mediumtext    null,
    graphID     int           not null,
    visible     tinyint(1)    null,
    tid         varchar(255)  null,
    idf         int           not null,
    create_time int           not null,
    update_time int           not null,
    constraint task___graph_fk
        foreign key (graphID) references graph.graphs (id)
            on update cascade on delete cascade
);

