create database shows;
create user show_admin@'%' identified by 'show_admin';
GRANT ALL PRIVILEGES ON shows.* TO 'show_admin'@'%';

use shows;

create table cinema_halls
(
    id        int            not null primary key AUTO_INCREMENT,
    cinema_id int   unsigned not null,
    name      varchar(20)    not null,
    created_at    datetime    not null,
    updated_at    datetime    null,
    deleted_at    datetime    null
);

create table cinema_seats
(
    id             int          not null primary key AUTO_INCREMENT,
    cinema_hall_id int unsigned not null,
    number         varchar(10)  not null,
    created_at    datetime    not null,
    updated_at    datetime    null,
    deleted_at    datetime    null
);

create table cinemas
(
    id   int         not null primary key AUTO_INCREMENT,
    name varchar(20) not null,
    city tinyint(1)     not null,
    created_at    datetime    not null,
    updated_at    datetime    null,
    deleted_at    datetime    null
);

create table movies
(
    id            int         not null primary key AUTO_INCREMENT,
    title         varchar(30) not null,
    duration      int         not null,
    `desc`        longtext    not null,
    come_out_date datetime    null,
    created_at    datetime    not null,
    updated_at    datetime    null,
    deleted_at    datetime    null
);

create table shows
(
    id             int          not null primary key AUTO_INCREMENT,
    movie_id       int unsigned not null,
    cinema_hall_id int unsigned not null,
    date           datetime not null,
    start_time     datetime not null,
    end_time       datetime not null,
    created_at    datetime    not null,
    updated_at    datetime    null,
    deleted_at    datetime    null
);

