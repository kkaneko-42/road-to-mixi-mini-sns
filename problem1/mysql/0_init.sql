CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT '' NOT NULL,
  PRIMARY KEY (`user_id`)
);

CREATE TABLE `friend_link` (
  `user_id` int(11) NOT NULL,
  `friend_id` int(11) NOT NULL
);
-- subject blocks object
CREATE TABLE `block_list` (
  `subject_id` int(11) NOT NULL,
  `object_id` int(11) NOT NULL
);

-- data for debug

INSERT INTO users (name) VALUES ("kkaneko");
INSERT INTO users (name) VALUES ("hogehoge");
INSERT INTO users (name) VALUES ("fuga");
INSERT INTO users (name) VALUES ("piyopiyp");
INSERT INTO users (name) VALUES ("asdf");
INSERT INTO users (name) VALUES ("webfo");
INSERT INTO users (name) VALUES ("qwpeo");
INSERT INTO users (name) VALUES ("qwevboi");
INSERT INTO users (name) VALUES ("asdofu");
INSERT INTO users (name) VALUES ("asfvo");

INSERT INTO friend_link VALUES (1, 2);
INSERT INTO friend_link VALUES (2, 3);
INSERT INTO friend_link VALUES (1, 3);
INSERT INTO friend_link VALUES (3, 4);
INSERT INTO friend_link VALUES (2, 5);
INSERT INTO friend_link VALUES (2, 6);
INSERT INTO friend_link VALUES (2, 7);
INSERT INTO friend_link VALUES (2, 8);
INSERT INTO friend_link VALUES (2, 9);
INSERT INTO friend_link VALUES (2, 10);

INSERT INTO block_list values (1, 4);
