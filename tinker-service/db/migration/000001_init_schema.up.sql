create table
    if not exists app_service (
        id serial4 primary key,
        container_name varchar not null,
        database_vendor varchar check (database_vendor in ('postgres', 'mysql')),
        default_database varchar not null,
        db_username varchar not null,
        db_password varchar not null,
        created_at timestamp not null default now()
    );