create table algo
(
    id       int auto_increment
        primary key,
    name     varchar(255)         not null,
    note     text                 null,
    isCustom tinyint(1) default 0 not null,
    type     int        default 0 not null
);

create table algoParam
(
    id        int auto_increment
        primary key,
    algoID    int           not null,
    fieldName varchar(255)  not null,
    fieldType int default 0 not null,
    constraint algoParam___id
        foreign key (algoID) references algo (id)
            on update cascade on delete cascade
);
