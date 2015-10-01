
-- MYSQL
DROP TABLE IF EXISTS USERS;
CREATE TABLE USERS (
  ID VARCHAR(50) NOT NULL,
  CODE VARCHAR(50),
  NAME VARCHAR(255),
  PASSWORD VARCHAR(255),
  MEMO VARCHAR(1000),
  GENRE VARCHAR(50),
  STATUS VARCHAR(50),
  ROLES VARCHAR(1000),
  POSTS VARCHAR(1000),
  ORG VARCHAR(50),
  AREA VARCHAR(50),
  CREATER VARCHAR(50),
  CREATED BIGINT,
  MODIFIER VARCHAR(50),
  MODIFIED BIGINT,
  VERSION INT(11),
  DELETION INT(11),
  PRIMARY KEY (ID)
);

-- ORACLE
CREATE TABLE USERS (
  ID VARCHAR2(50) NOT NULL,
  CODE VARCHAR2(50),
  NAME VARCHAR2(255),
  PASSWORD VARCHAR2(255),
  MEMO VARCHAR2(1000),
  GENRE VARCHAR2(50),
  STATUS VARCHAR2(50),
  ROLES VARCHAR2(1000),
  POSTS VARCHAR2(1000),
  ORG VARCHAR2(50),
  AREA VARCHAR2(50),
  CREATER VARCHAR2(50),
  CREATED NUMBER(20),
  MODIFIER VARCHAR2(50),
  MODIFIED NUMBER(20),
  VERSION NUMBER(11),
  DELETION NUMBER(11),
  PRIMARY KEY (ID)
);

INSERT INTO USERS (ID, CODE, NAME, PASSWORD, MEMO, GENRE, STATUS, ROLES, POSTS, ORG, AREA, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION) VALUES ('demo-i-01', '01', '01', '01', '01', '01', '01', '01', '01', '01', '01', '01', 1443253329, '01', 1443253329, 0, 0);
INSERT INTO USERS (ID, CODE, NAME, PASSWORD, MEMO, GENRE, STATUS, ROLES, POSTS, ORG, AREA, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION) VALUES ('demo-i-02', '02', 'user02', '02', '02', '02', '02', '02', '02', '02', '02', '02', 1443253329, '02', 1443253329, 0, 1);
INSERT INTO USERS (ID, CODE, NAME, PASSWORD, MEMO, GENRE, STATUS, ROLES, POSTS, ORG, AREA, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION) VALUES ('demo-i-03', '03', '03', '03', '03', '03', '03', '03', '03', '03', '03', '03', 1443253329, '03', 1443253329, 0, 0);
INSERT INTO USERS (ID, CODE, NAME, PASSWORD, MEMO, GENRE, STATUS, ROLES, POSTS, ORG, AREA, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION) VALUES ('demo-i-04', '04', '04', '04', '04', '04', '04', '04', '04', '04', '04', '04', 1443253329, '04', 1443253329, 0, 0);