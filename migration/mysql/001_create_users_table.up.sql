CREATE TABLE users (
  id int PRIMARY KEY AUTO_INCREMENT,
  username varchar(255),
  password varchar(255),
  created_at datetime DEFAULT CURRENT_TIMESTAMP
);