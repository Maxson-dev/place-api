-- +goose Up

create table if not exists place
(
    id              bigserial          primary key,
    name            text               not null,
    lat             double precision   not null,
    lng             double precision   not null,
    created_at      timestamptz        not null,
    updated_at      timestamptz        not null
);

-- +goose Down

drop table if exists place;