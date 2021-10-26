CREATE TABLE IF NOT EXISTS coordinates (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    album_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    is_show BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS albums (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(128) NOT NULL,
    started_at TIMESTAMP NOT NULL,
    ended_at TIMESTAMP NOT NULL,
    is_public BOOLEAN NOT NULL,
    thumbnail_image_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS images (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    url VARCHAR(128) NOT NULL,
    album_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    coordinate_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    device_id VARCHAR(256),
    name VARCHAR(128),
    profile_image_url VARCHAR(128),
    introduction VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS friends (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    follow_user_id INT NOT NULL
);