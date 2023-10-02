-- +goose Up
create table scheduled_event
(
    id          uuid         not null primary key,
    status      integer      not null default 0,  -- 0 - new, 1 - done, 3 - failed
    type        varchar(255) not null,
    payload     jsonb        not null,
    datetime    timestamptz  not null,
    attempt     integer not  null default 0,
    created_at  timestamptz  not null,
    updated_at  timestamptz  not null
);

create index event__status__idx on scheduled_event (status);

-- +goose Down

drop table if exists scheduled_event;