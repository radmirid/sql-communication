create table users (
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    registered_at timestamp not null default now()
);

insert into users (name, email, password) values ('Alice', 'alice@mail.ru', 'qwerty');

create table logs (
    id serial not null unique,
    entity varchar(255) not null,
    action varchar(255) not null,
    time timestamp not null default now()
);