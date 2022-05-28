CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT '' NOT NULL,
  PRIMARY KEY (`id`)
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
