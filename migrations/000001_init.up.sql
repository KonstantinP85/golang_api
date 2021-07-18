CREATE TABLE users
(
    id         serial not null unique,
    email      varchar(255) not null,
    password   varchar(255) not null unique,
    is_active  boolean
);

CREATE TABLE author
(
    id         serial not null unique,
    name       varchar(255) not null
);

CREATE TABLE book
(
    id         serial not null unique,
    author_id  int references author (id) on delete cascade not null,
    title      varchar(255) not null,
    years      int not null
);