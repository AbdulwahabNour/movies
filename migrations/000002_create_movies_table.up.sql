Create table if not exists movies(
    id bigserial primary key,
    create_at timestamp(0) with time zone not null default now(),
    title text not null,
    year integer not null,
    runtime integer not null,
    genres text[] not null,
    version uuid not null default uuid_generate_v4()
);