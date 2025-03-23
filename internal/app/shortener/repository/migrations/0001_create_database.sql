create table if not exists users
(
    id            varchar(50) primary key,
    login         varchar(50) unique not null,
    password      text               not null,
    cookie        text,
    cookie_finish TIMESTAMP
);

create table if not exists urls
(
    id           serial
        primary key,
    short_url    varchar(50) not null
        unique,
    original_url text        not null,
    user_id      varchar(50) REFERENCES users (id) ON DELETE CASCADE
);