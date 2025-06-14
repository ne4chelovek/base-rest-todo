-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users (id) on DELETE cascade not null,
    list_id int references todo_lists (id) on DELETE cascade not null
);

CREATE TABLE todo_items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

CREATE TABLE lists_items
(
    id serial not null unique,
    item_id int references todo_items (id) on DELETE cascade not null,
    list_id int references todo_lists (id) on DELETE cascade not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lists_items;
DROP TABLE IF EXISTS users_lists;
DROP TABLE IF EXISTS todo_items;
DROP TABLE IF EXISTS todo_lists;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd