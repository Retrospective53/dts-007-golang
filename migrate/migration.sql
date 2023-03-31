CREATE DATABASE BOOKS_MANAGEMENT;

CREATE TABLE books(
  id serial NOT NULL PRIMARY KEY,
  title VARCHAR(255),
  author VARCHAR(255),
  description text,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz
);



