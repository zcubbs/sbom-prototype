create table if not exists scan
(
    id         serial,
    uuid       uuid      default gen_random_uuid() not null,
    created_at timestamp default now()             not null,
    updated_at timestamp default now()             not null,
    image     text     not null,
    status    text    not null default 'pending',
    sbom      text,
    report    text,
    artifact_id uuid,
    artifact_name text,
    artifact_version text
);

create index if not exists scan_pk_index
    on scan (id);

create index if not exists scan_uuid_index
    on scan (uuid);


