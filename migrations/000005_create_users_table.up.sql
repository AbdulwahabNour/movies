CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    create_at timestamp(0) with time zone not null default now(),
    name text not null,
    email CITEXT UNIQUE not null,
    password_hash bytea not null,
    activated boolean not null default false,
    version uuid not null default uuid_generate_v4()
);