-- coordinates
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:12:35.000000', 35.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:45:35.000000', 36.1234, 139.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:12:35.000000', 34.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:45:35.000000', 39.1234, 139.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:12:35.000000', 35.1234, 132.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:25:35.000000', 37.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 13:45:35.000000', 36.1234, 133.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:12:35.000000', 34.1234, 136.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:25:35.000000', 37.1234, 132.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (1, '2021-10-23 14:45:35.000000', 39.1234, 135.1234);

INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:12:35.000000', 26.2455, 128.2355);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:45:35.000000', 36.1234, 139.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:12:35.000000', 34.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:45:35.000000', 39.1234, 139.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:12:35.000000', 30.1234, 130.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 13:45:35.000000', 36.1234, 139.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:12:35.000000', 34.1234, 135.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:25:35.000000', 37.1234, 137.1234);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude) VALUES (2, '2021-10-24 14:45:35.000000', 39.1234, 139.1234);


-- albums
INSERT INTO albums (user_id, title, started_at, ended_at, is_public, thumbnail_image_id) VALUES (1, 'うんこモグモグおいしいにょ', '2021-10-23 13:00:00.000000', '2021-10-23 14:00:00.000000', 1, 1);
INSERT INTO albums (user_id, title, started_at, ended_at, is_public, thumbnail_image_id) VALUES (2, 'Hey Bob How are you I`m fine thank you', '2021-10-24 13:00:00.000000', '2021-10-23 14:00:00.000000', 1, 13);

-- images
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:25:35.000000', 1);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:30:35.000000', 2);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:35:35.000000', 3);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:25:35.000000', 4);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:30:35.000000', 5);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:35:35.000000', 6);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:25:35.000000', 7);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:30:35.000000', 8);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:35:35.000000', 9);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:25:35.000000', 10);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:30:35.000000', 11);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://wired.jp/app/uploads/2018/01/GettyImages-522585140_w3200.webp", 1, '2021-10-23 13:35:35.000000', 12);

INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:25:35.000000', 13);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:30:35.000000', 14);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:35:35.000000', 15);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:25:35.000000', 16);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:30:35.000000', 17);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:35:35.000000', 18);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:25:35.000000', 19);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:30:35.000000', 20);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:35:35.000000', 21);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:25:35.000000', 22);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:30:35.000000', 23);
INSERT INTO images (url, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg", 2, '2021-10-24 13:35:35.000000', 24);

-- users
INSERT INTO users (name, password) VALUES ('Bob', 'pass');
INSERT INTO users (name, password) VALUES ('Tom', 'pass');

-- friends
INSERT INTO friends (user_id, follow_user_id) VALUES (1, 2);