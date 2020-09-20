CREATE DATABASE `survey`;

USE `survey`;

CREATE TABLE `Admin` {
  id INT(11) AUTO_INCREMENT,
  username VARCHAR(20) NOT NULL,
  password VARCHAR(64) NOT NULL,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
}

CREATE TABLE `QuestionInfo` {
  id INT(11) AUTO_INCREMENT,
  title VARCHAR(30) NOT NULL,
  creater INT(11) NOT NULL,
  datetime DATETIME,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
}

CREATE TABLE `QuestionList` {
  id INT(11) AUTO_INCREMENT,
  qid INT(11) NOT NULL,
  answer TEXT NOT NULL,
  option_a TEXT NOT NULL,
  option_b TEXT,
  option_c TEXT,
  option_d TEXT,
  PRIMARY KEY(id)
}

CREATE TABLE `User` {
  id INT(11) AUTO_INCREMENT,
  phone VARCHAR(11) NOT NULL,
  password VARCHAR(64) NOT NULL,
  name VARCHAR(32) NOT NULL,
  status INT(1) DEFAULT 0,
  PRIMARY KEY(id)
}

CREATE TABLE `QuestionList` {
  qid INT(11) NOT NULL,
  aid INT(11) NOT NULL,
  uid INT(11) NOT NULL,
  result TEXT NOT NULL,
  PRIMARY KEY(qid, aid, uid)
}
