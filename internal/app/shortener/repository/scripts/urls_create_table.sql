create table if not exists urls
(
    id           serial primary key,
    short_url    varchar(50) unique not null,
    original_url text               not null
)