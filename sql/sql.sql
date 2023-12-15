CREATE DATABASE IF NOT EXISTS appgo;
USE appgo;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null unique,
    createAt timestamp default current_timestamp()
)ENGINE=INNODB