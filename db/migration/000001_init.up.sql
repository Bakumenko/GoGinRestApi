CREATE TABLE users
(
    uuid        uuid not null unique primary key,
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