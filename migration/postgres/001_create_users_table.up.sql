CREATE TABLE users (
  id serial PRIMARY KEY,
  username varchar(255),
  password varchar(255),
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);