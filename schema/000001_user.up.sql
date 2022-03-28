CREATE TABLE IF NOT EXISTS users(
   id serial       not null unique,
   name VARCHAR (50) not null unique,
   age VARCHAR (50) not null,
   friends text[]
);

