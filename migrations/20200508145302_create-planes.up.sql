CREATE TABLE planes (
    id bigserial not null primary key,
    name varchar not null unique,
    manufacture_year varchar not null
);