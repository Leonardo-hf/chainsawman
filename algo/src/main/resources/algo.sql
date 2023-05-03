drop table if exists algoParam;
drop table if exists algo;

create table if not exists graph.algo
(
    id   int auto_increment
        primary key,
    name varchar(255)  not null,
    note text          null,
    type int default 0 not null
);

create table if not exists graph.algoParam
(
    id        int auto_increment
        primary key,
    algoID    int           not null,
    fieldName varchar(255)  not null,
    fieldType int default 0 not null,
    constraint algoParam___id
        foreign key (algoID) references graph.algo (id)
            on update cascade on delete cascade
);

