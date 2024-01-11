CREATE DATABASE IF NOT EXISTS appgo;
USE appgo;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios
(
    id       int auto_increment primary key,
    nome     varchar(50) not null,
    nick     varchar(50) not null unique,
    email    varchar(50) not null unique,
    senha    varchar(100) not null,
    createAt timestamp default current_timestamp()
) ENGINE=INNODB

ALTER TABLE appgo.usuarios MODIFY COLUMN senha varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL;
