CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null unique
);

CREATE TABLE tasks
(
    id serial not null unique,
    tag varchar(255),
    title varchar(255) not null,
    description varchar(255),
    deadline date,
    done boolean not null default false
);

CREATE TABLE tasks_users
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    task_id int references tasks(id) on delete cascade not null
);

