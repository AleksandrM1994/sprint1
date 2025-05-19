-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id            varchar(50) primary key,
    login         varchar(50) unique not null,
    password      text               not null,
    cookie        text,
    cookie_finish TIMESTAMP
);

ALTER TABLE urls
    ADD COLUMN user_id VARCHAR(50);

ALTER TABLE urls
    ADD COLUMN is_deleted BOOL default false;
-- +goose StatementEnd