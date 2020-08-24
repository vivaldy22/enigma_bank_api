CREATE DATABASE enigma_bank;
USE enigma_bank;

CREATE TABLE login (
                       login_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
                       username VARCHAR(100) NOT NULL,
                       password VARCHAR(100) NOT NULL,
                       status_del INT DEFAULT 0
);

CREATE TABLE user (
                      user_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
                      login_owner_id INT NOT NULL,
                      balance INT DEFAULT 0,
                      status_del INT DEFAULT 0
);

CREATE TABLE transaction (
                             trans_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
                             user_owner_id INT NOT NULL,
                             trans_date DATETIME NOT NULL,
                             destination VARCHAR(100) NOT NULL,
                             amount INT DEFAULT 0,
                             description VARCHAR(255) NOT NULL,
                             status_del INT DEFAULT 0
);

drop database enigma_bank;
drop table transaction;

insert into login values (null, 'admin', 'admin', 0);
insert into user values (null, 1, 200, 0);
insert into transaction values (null, 1, NOW(), 'Vivaldy', 50, 'Pay for lunch', 0);
select * from login;
select * from user where login_owner_id = 1;
select * from transaction where user_owner_id = 1;