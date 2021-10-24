-- coordinates
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:12:35.000000', 35.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:45:35.000000', 36.1234, 139.1234);

-- albums
INSERT INTO albums (user_id, title, started_at, ended_at, is_public, thumbnail_image_id) VALUES (1, 'うんこモグモグおいしいにょ', '2021-10-23 13:00:00.000000', '2021-10-23 14:00:00.000000', 1, 1);

-- images
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:25:35.000000', 1);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:30:35.000000', 2);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:35:35.000000', 3);

-- users
INSERT INTO users (name, password) VALUES ('Bob', 'pass');
INSERT INTO users (name, password) VALUES ('Tom', 'pass');

-- friends
INSERT INTO friends (user_id, follow_user_id) VALUES (1, 2);