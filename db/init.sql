CREATE DATABASE `survey`;

USE `survey`;

DROP TABLE IF EXISTS `Admin`;
CREATE TABLE `Admin` (
  id INT(11) AUTO_INCREMENT,
  username VARCHAR(20) NOT NULL,
  password VARCHAR(64) NOT NULL,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS  `QuestionInfo`;
CREATE TABLE `QuestionInfo` (
  id INT(11) AUTO_INCREMENT,
  title VARCHAR(30) NOT NULL,
  creater INT(11) NOT NULL,
  datetime DATETIME,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `Answer`;
CREATE TABLE `Answer` (
  id INT(11) AUTO_INCREMENT,
  qid INT(11) NOT NULL,
  answer TEXT NOT NULL,
  option_a TEXT NOT NULL,
  option_b TEXT,
  option_c TEXT,
  option_d TEXT,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `User`;
CREATE TABLE `User` (
  id INT(11) AUTO_INCREMENT,
  phone VARCHAR(11) NOT NULL,
  password VARCHAR(64) NOT NULL,
  name VARCHAR(32) NOT NULL,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
);

DROP TABLE IF EXISTS `QuestionList`;
CREATE TABLE `QuestionList` (
  qid INT(11) NOT NULL,
  aid INT(11) NOT NULL,
  uid INT(11) NOT NULL,
  result TEXT NOT NULL,
  PRIMARY KEY(qid, aid, uid)
);

INSERT INTO Admin (username, password, status)
VALUES ('admin', '1234', 1);

INSERT INTO QuestionInfo (title, creater, datetime, status)
VALUES ('问卷调查1', 1, '2020-04-01 16:00', 1), ('问卷调查2', 1, '2020-04-01 10:00', 1), ('问卷调查3', 1, '2020-02-01 16:00', 1), ('问卷调查4', 1, '2020-03-20 16:00', 0);