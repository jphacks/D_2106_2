CREATE DATABASE IF NOT EXISTS sample_db;

use sample_db;

CREATE TABLE users
(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(128),
    password VARCHAR(128)
);

INSERT INTO users (name, password) VALUES ('Bob', 'pass');
INSERT INTO users (name, password) VALUES ('Tom', 'pass');
INSERT INTO users (name, password) VALUES ('Nancy', 'pass');