CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
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

INSERT INTO users VALUES (1, "kkaneko");
INSERT INTO users VALUES (2, "hogehoge");
INSERT INTO users VALUES (3, "fuga");
INSERT INTO users VALUES (4, "piyopiyp");

INSERT INTO friend_link VALUES (1, 2);
INSERT INTO friend_link VALUES (2, 1);
INSERT INTO friend_link VALUES (1, 3);
INSERT INTO friend_link VALUES (3, 4);
