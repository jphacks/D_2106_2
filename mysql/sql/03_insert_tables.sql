-- coordinates
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:12:35.000000', 35.1234, 135.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:25:35.000000', 37.1234, 137.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:45:35.000000', 36.1234, 139.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:12:35.000000', 34.1234, 135.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:25:35.000000', 37.1234, 137.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:45:35.000000', 39.1234, 139.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:12:35.000000', 35.1234, 132.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:25:35.000000', 37.1234, 135.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 13:45:35.000000', 36.1234, 133.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:12:35.000000', 34.1234, 136.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:25:35.000000', 37.1234, 132.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (1, '2021-10-23 14:45:35.000000', 39.1234, 135.1234, false);

INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:12:35.000000', 26.2455, 128.2355, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:25:35.000000', 37.1234, 137.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:45:35.000000', 36.1234, 139.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:12:35.000000', 34.1234, 135.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:25:35.000000', 37.1234, 137.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:45:35.000000', 39.1234, 139.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:12:35.000000', 30.1234, 130.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:25:35.000000', 37.1234, 137.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 13:45:35.000000', 36.1234, 139.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:12:35.000000', 34.1234, 135.1234, false);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:25:35.000000', 37.1234, 137.1234, true);
INSERT INTO coordinates (album_id, timestamp, latitude, longitude, is_show) VALUES (2, '2021-10-24 14:45:35.000000', 39.1234, 139.1234, false);


-- albums
INSERT INTO albums (user_id, title, started_at, ended_at, is_public, spot ,thumbnail_image_id) VALUES ('dummy_device_id_1', 'うんこモグモグおいしいにょ', '2021-10-23 13:00:00.000000', '2021-10-23 14:00:00.000000', 1, "愛知県名古屋市", 0);
INSERT INTO albums (user_id, title, started_at, ended_at, is_public, spot ,thumbnail_image_id) VALUES ('dummy_device_id_2', 'Hey Bob How are you I`m fine thank you', '2021-10-24 13:00:00.000000', '2021-10-23 14:00:00.000000', 1, "愛知県名古屋市",0);

-- images
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:30:35.000000', 2);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:25:35.000000', 1);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-thumbnail.png", 1, '2021-10-23 13:35:35.000000', 3);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:25:35.000000', 4);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:30:35.000000', 5);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:35:35.000000', 6);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:25:35.000000', 7);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:30:35.000000', 8);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:35:35.000000', 9);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:25:35.000000', 10);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:30:35.000000', 11);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg","1-image.png", 1, '2021-10-23 13:35:35.000000', 12);

INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-thumbnail.png", 2, '2021-10-24 13:25:35.000000', 13);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:30:35.000000', 14);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:35:35.000000', 15);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:25:35.000000', 16);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:30:35.000000', 17);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:35:35.000000', 18);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:25:35.000000', 19);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:30:35.000000', 20);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:35:35.000000', 21);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:25:35.000000', 22);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:30:35.000000', 23);
INSERT INTO images (url, name, album_id, created_at, coordinate_id) VALUES ("https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg","2-image.png", 2, '2021-10-24 13:35:35.000000', 24);

-- users
INSERT INTO users (id, name, profile_image_url, introduction) VALUES ('dummy_device_id_1', 'Bob', 'https://pbs.twimg.com/media/Cw5hdTBUUAAWuIo.jpg', 'I am fine very very much');
INSERT INTO users (id, name, profile_image_url, introduction) VALUES ('dummy_device_id_2', 'Tom', 'https://i.ytimg.com/vi/gF4m7sCQ-4c/maxresdefault.jpg', 'I want to go to zoo');

-- friends
INSERT INTO friends (user_id, follow_user_id) VALUES ('dummy_device_id_1', 'dummy_device_id_2');