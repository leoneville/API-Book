CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS books;

CREATE TABLE books(
    id          int auto_increment primary key,
    titulo      varchar(70) not null,
    autor       varchar(100) not null,
    qtd_paginas int not null,
    editora     varchar(50) not null,
    criadoEm    timestamp default current_timestamp()
) ENGINE=INNODB;