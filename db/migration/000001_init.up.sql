CREATE TABLE users
(
    uuid        uuid not null unique primary key,
    firstname varchar(20),
    lastname  varchar(30),
    email     varchar(40) unique,
    age       int,
    created_at timestamp default now()
);