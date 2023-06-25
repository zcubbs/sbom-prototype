create table if not exists scan
(
    id         serial,
    uuid       uuid      default gen_random_uuid() not null,
    created_at timestamp default now()             not null
);

alter table scan
    owner to postgres;

create index if not exists scan_pk_index
    on scan (id);

create index if not exists scan_uuid_index
    on scan (uuid);


