-- +goose Up

create table if not exists file
(
    id              uuid         primary key,
    size            bigint       not null default 0,
    bucket          text         not null,
    key             text         not null,
    name            text         not null,
    s3_url          text         not null,
    created_at      timestamptz  not null,
    updated_at      timestamptz  not null,
    is_deleted      boolean      not null default false
);

-- +goose Down

drop table if exists file;