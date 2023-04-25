create table graphs
(
    id     int auto_increment
        primary key,
    name   varchar(255)  not null,
    `desc` text          null,
    status int default 0 not null,
    nodes  int           null,
    edges  int           null
);

create table task
(
    id          int auto_increment
        primary key,
    params      varchar(1024) null,
    update_time int           not null,
    name        varchar(255)  not null,
    status      int           not null,
    result      text          null,
    create_time int           not null
);

