CREATE DATABASE IF NOT EXISTS `project_management`;

USE `project_management`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` INT(10) AUTO_INCREMENT,
  `username` VARCHAR(20) NOT NULL,
  `password` VARCHAR(20) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) DEFAULT 0,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS  `user_role`;
CREATE TABLE `user_role` (
  `id` INT(2) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` INT(10) AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  `status` INT(2) NOT NULL,
  `creator` INT(10) NOT NULL,
  `created_time` INT(13) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `group_member`;
CREATE TABLE `group_member` (
  `group_id` INT(10) NOT NULL,
  `uid` INT(10) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) NOT NULL,
  `created_date` INT(13) NOT NULL,
  PRIMARY KEY(`group_id`, `uid`)
);

DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` INT(10) AUTO_INCREMENT,
  `creator` INT(10) NOT NULL,
  `group_id` INT(10) NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `status` INT(2) NOT NULL,
  `created_date` INT(13) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `project_member`;
CREATE TABLE `project_member` (
  `project_id` INT(10) NOT NULL,
  `uid` INT(10) NOT NULL,
  `role_id` INT(2) NOT NULL,
  `status` INT(2) NOT NULL,
  `created_date` INT(13) NOT NULL,
  PRIMARY KEY(`project_id`, `uid`)
);

DROP TABLE IF EXISTS  `member_role`;
CREATE TABLE `member_role` (
  `id` INT(2) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `task_group`;
CREATE TABLE `task_group` (
  `id` INT(10) AUTO_INCREMENT,
  `title` VARCHAR(20) NOT NULL,
  `desc` VARCHAR(20) NOT NULL,
  `creator` INT(10) NOT NULL,
  `status` INT(2) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` INT(10) AUTO_INCREMENT,
  `title` VARCHAR(20) NOT NULL,
  `desc` VARCHAR(200),
  `creator` INT(10) NOT NULL,
  `executor` INT(10) NOT NULL,
  `status` INT(2) NOT NULL,
  `created_date` INT(13) NOT NULL,
  `expire_date` INT(13) NOT NULL,
  `task_project_id` INT(10),
  `task_group_id` INT(10),
  `parent_task_id` INT(10),
  `type` INT(2) NOT NULL,
  `priority` INT(2) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS `task_comment`;
CREATE TABLE `task_comment` (
  `id` INT(10) AUTO_INCREMENT,
  `task_id` INT(10) NOT NULL,
  `creator` INT(10) NOT NULL,
  `text` TEXT NOT NULL,
  `status` INT(2) NOT NULL,
  `created_date` INT(13) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS  `task_type`;
CREATE TABLE `task_type` (
  `id` INT(2) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS  `priority_type`;
CREATE TABLE `priority_type` (
  `id` INT(2) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `desc` TEXT,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS  `task_participant`;
CREATE TABLE `task_participant` (
  `task_id` INT(10) NOT NULL,
  `uid` INT(10) NOT NULL,
  `add_date` INT(13) NOT NULL,
  PRIMARY KEY(`task_id`, `uid`)
);

DROP TABLE IF EXISTS  `tag`;
CREATE TABLE `tag` (
  `id` INT(10) AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `add_date` INT(13) NOT NULL,
  PRIMARY KEY(`id`)
);

DROP TABLE IF EXISTS  `task_tag`;
CREATE TABLE `task_tag` (
  `task_id` INT(10) NOT NULL,
  `tag_id` INT(10) NOT NULL,
  PRIMARY KEY(`task_id`, `tag_id`)
);

INSERT INTO `user_role` (`text`, `desc`)
VALUES ('管理员', '管理员身份'), ('用户', '普通用户');

INSERT INTO `user` (`username`, `password`, `role_id`, `status`)
VALUES ('admin1', '123456', 1, 1), ('admin2', '123456', 1, 1), ('user1', '123456', 2, 1), ('user2', '123456', 2, 1);

INSERT INTO `group` (`name`, `status`, `creator`, `created_time`)
VALUES ('测试小组1', 1, 3, 1616478832), ('测试小组2', 1, 4, 1616478832);

INSERT INTO `member_role` (`text`, `desc`)
VALUES ('管理员', '群组/项目管理员'), ('普通用户', '普通参与者');

INSERT INTO `group_member` (`group_id`, `uid`, `role_id`, `status`, `created_date`)
VALUES (1, 1, 1, 1, 1616478832), (2, 2, 2, 1, 1616478832), (2, 3, 1, 1, 1616478832);

INSERT INTO `project` (`creator`, `group_id`, `name`, `status`)
VALUES (3, 1, '测试项目 1', 1), (4, 2, '测试项目 2', 1);

INSERT INTO `project_member` (`project_id`, `uid`, `role_id`, `status`, `created_date`)
VALUES (1, 1, 1, 1, 1616478832), (2, 2, 2, 1, 1616478832), (2, 3, 1, 1, 1616478832);

INSERT INTO `priority_type` (`text`)
VALUES ('紧急'), ('重要'), ('普通');

INSERT INTO `task_type` (`text`)
VALUES ('需求'), ('缺陷');

INSERT INTO `task` (`title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_group_id`, `task_project_id`, `type`, `priority`)
VALUES ('测试任务 1', '任务描述', 3, 3, 1, 1616478832, 1614419750, 1, 1, 1, 1), ('测试任务 2', '任务描述', 2, 3, 1, 1616478832, 1614439750, 2, 2, 1, 2), ('测试任务 3', '任务描述', 2, 1, 0, 1616478832, 1614439800, 1, 3, 1, 1), ('测试任务 4', '任务描述', 3, 1, 1, 1616478832, 1614449750, 1, 4, 1, 1);

INSERT INTO `task_participant` (`task_id`, `uid`, `add_date`)
VALUES (1, 2, 1614419750), (2, 3, 1614419750);

SELECT * FROM tag;
SELECT * FROM task;
SELECT * FROM task_group;

SELECT `id`, `title`, `desc`, `creator`, `executor`, `status`, `created_date`, `expire_date`, `task_project_id`, `task_group_id`, `parent_task_id`, `type`, `priority` FROM task WHERE task_project_id = 1 ORDER BY task_project_id DESC, task_group_id DESC, parent_task_id DESC, priority DESC, created_date DESC LIMIT 0 , 10;

SELECT id, title, `desc`, creator, status FROM task_group WHERE id IN (1, 2);

UPDATE task SET type = 1 WHERE id = 8;

UPDATE task SET executor = 2 WHERE executor = 6;

SELECT * FROM tag

SELECT * FROM task WHERE parent_task_id = 1