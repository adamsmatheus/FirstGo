CREATE
DATABASE IF NOT EXISTS appgo;
USE
appgo;
GO;
DROP TABLE IF EXISTS usuarios;
GO;
CREATE TABLE usuarios
(
    id       int auto_increment primary key,
    nome  varchar(50) not null,
    nick  varchar(50) not null unique,
    email varchar(50) not null unique,
    senha    varchar(100) not null,
    createAt timestamp default current_timestamp()
) ENGINE=INNODB
GO;
ALTER TABLE appgo.usuarios MODIFY
    COLUMN senha varchar (100) CHARACTER SET utf8mb4
    COLLATE utf8mb4_0900_ai_ci NOT NULL;
GO;
create table seguidores
(
    usuario_id int not null,
    foreign key (usuario_id)
        references usuarios (id)
        on delete cascade,

    seguidor_id int not null,
    foreign key (usuario_id)
        references usuarios (id)
        on delete cascade,

    primary key (usuario_id,seguidor_id)
) ENGINE = INNODB;
