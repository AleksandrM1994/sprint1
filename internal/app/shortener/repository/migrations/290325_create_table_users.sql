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
    ADD COLUMN user_id VARCHAR(50),
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
-- +goose StatementEnd