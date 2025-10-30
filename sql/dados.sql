insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com",  "$2a$10$B3OOBT/N2tviWYzTGDMb.OPax3SzR3F4mIV5A6pz47dUauOfy.HxW" ),
("Usuario 2", "usuario_2", "usuario2@gmail.com",  "$2a$10$B3OOBT/N2tviWYzTGDMb.OPax3SzR3F4mIV5A6pz47dUauOfy.HxW" ),
("Usuario 3", "usuario_3", "usuario3@gmail.com",  "$2a$10$B3OOBT/N2tviWYzTGDMb.OPax3SzR3F4mIV5A6pz47dUauOfy.HxW" );

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do Usuario 1", "Essa e a publicacao do usuario 1! Oba!", 1),
("Publicacao do Usuario 2", "Essa e a publicacao do usuario 2! Oba!", 2),
("Publicacao do Usuario 3", "Esa e a publicao do usuario 3! Oba!", 3);