INSERT INTO appgo.usuarios (nome, nick, email, senha, createAt)
VALUES ('Bernardo', 'Be', 'bernado123@gmail.com', '$2a$10$rZxMR7LO3u2120l7fAXyMOsl2Eo1NcRZ/BnnHVt/AggcOUsAIXQeO',
        CURRENT_TIMESTAMP),
       ('Usuario_1', '1', 'usu1@gmail.com', '$2a$10$rZxMR7LO3u2120l7fAXyMOsl2Eo1NcRZ/BnnHVt/AggcOUsAIXQeO',
        CURRENT_TIMESTAMP),
       ('Usuario_2', '2', 'usu2@gmail.com', '$2a$10$rZxMR7LO3u2120l7fAXyMOsl2Eo1NcRZ/BnnHVt/AggcOUsAIXQeO',
        CURRENT_TIMESTAMP);


INSERT INTO appgo.seguidores
    (usuario_id, seguidor_id)
VALUES (1, 2),
       (3, 1),
       (1, 3)