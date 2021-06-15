CREATE TABLE users
(
    id        uuid not null unique primary key,
    firstname varchar(20),
    lastname  varchar(30),
    email     varchar(40) unique,
    age       int,
    role_id   int not null,
    created_at timestamp default now()
);

CREATE TABLE roles
(
    id      serial not null unique primary key,
    name    varchar(20)
);

INSERT INTO roles VALUES
('0', 'ROLE_ADMIN');

INSERT INTO roles VALUES
('1', 'ROLE_USER');