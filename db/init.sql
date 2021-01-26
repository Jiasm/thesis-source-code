CREATE DATABASE IF NOT EXISTS `project_management`;

USE `project_management`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` INT(11) AUTO_INCREMENT,
  `username` VARCHAR(20) NOT NULL,
  `password` VARCHAR(20) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) DEFAULT 0,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS  `user_role`;
CREATE TABLE `user_role` (
  `id` INT(11) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` INT(11) AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `status` INT(2) NOT NULL,
  `creator` INT(10) NOT NULL,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `group_member`;
CREATE TABLE `group_member` (
  `group_id` INT(10) NOT NULL,
  `uid` INT(10) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) NOT NULL,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` INT(11) AUTO_INCREMENT,
  `creator` INT(10) NOT NULL,
  `group_id` INT(10) NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `status` INT(2) NOT NULL,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `project_member`;
CREATE TABLE `project_member` (
  `project_id` INT(10) NOT NULL,
  `uid` INT(10) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) NOT NULL,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS  `member_role`;
CREATE TABLE `member_role` (
  `id` INT(11) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `task_group`;
CREATE TABLE `task_group` (
  id INT(10) AUTO_INCREMENT,
  title VARCHAR(20) NOT NULL,
  desc VARCHAR(20) NOT NULL,
  creator INT(10) NOT NULL,
  status INT(2) NOT NULL,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  id INT(10) AUTO_INCREMENT,
  title VARCHAR(20) NOT NULL,
  desc VARCHAR(20) NOT NULL,
  creator INT(10) NOT NULL,
  status INT(2) NOT NULL,
  expire_date INT(13) NOT NULL,
  task_group_id INT(10),
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `task_comment`;
CREATE TABLE `task_comment` (
  id INT(10) AUTO_INCREMENT,
  task_id INT(10) NOT NULL,
  creator INT(10) NOT NULL,
  text TEXT NOT NULL,
  status INT(2) NOT NULL,
  datetime INT(13) NOT NULL,
  PRIMARY KEY(id)
);
