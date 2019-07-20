-- db.sql
use test;

CREATE TABLE `my_posts` (
	`id` int UNSIGNED AUTO_INCREMENT,
	`uid` int NOT NULL,
	`title` varchar(255) NOT NULL,
	`content` text NOT NULL,
	`type` tinyint NOT NULL,
	`created_at` timestamp NULL,
	`updated_at` timestamp NULL,
	`deleted_at` timestamp NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `my_user` (
	`id` int UNSIGNED AUTO_INCREMENT,
	`username` varchar(255),
	`password` varchar(255),
	`email` varchar(255),
	`phone` varchar(255),
	PRIMARY KEY (`id`)
);