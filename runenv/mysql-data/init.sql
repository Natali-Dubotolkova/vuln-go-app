SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;
create database vulnapp;
create table vulnapp.user (id int not null auto_increment primary key, name varchar(255) not null,mail varchar(255),age int not null,passwd varchar(255) not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into vulnapp.user (name,mail,age,passwd) values ("Amuro Ray","RX-78-2@EFSF.com",15,"Mieru,Mieruzo!"),("Char Aznable","MS-06-S@Zeon.com",20,"BouyaDakarasa..."),("Kamille Bidan","MSZ-006@AEUG.com",17,"Kikoeru...Koega..."),("Judau Ashta","MSZ-010@AEUG.com",14,"Hamaaaan!!"),("Banagher Links","RX-0@londo.bell",16,"HitoNoMiraiHa...HitoGaTsukuruMonoDa!!!");
create table vulnapp.sessions (uid int,sessionid varchar(128));
create table vulnapp.userdetails (uid int not null primary key, userimage varchar(64), address varchar(64), animal varchar(32), word varchar(64));
insert vulnapp.userdetails(uid,userimage,address,animal,word) values (1,"amuro.png","SIDE-7","GANDOM","アムロ、行きまーす！"),(2,"char.png","SIDE-3","ZAKU","Misetemoraouka...RenpouNoMStoYarawo!!"),(3,"kamiyu.png","SIDE-7","Z-GANDOM","遊びでやってんじゃないんだよ!!!"),(4,"judou.png","SIDE-1","ZZ-GANDOM","貴様のようなやつは････許せないんだよ!!"),(5,"unicorn.png","INDUSTRIAL-7","UNICORN GANDOM","人の未来は...人が作るものだ!!");
create table vulnapp.posts (postid int not null primary key auto_increment, uid int not null, post varchar(256) not null, created_at timestamp not null default current_timestamp);
