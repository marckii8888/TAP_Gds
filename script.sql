DROP DATABASE IF EXISTS gds_assessment;

CREATE DATABASE gds_assessment;

USE gds_assessment;

DROP TABLE IF EXISTS urls;

CREATE TABLE urls(
original_url VARCHAR(100) NOT NULL,
short_url VARCHAR(30) NOT NULL,
code VARCHAR(30) NOT NULL,
PRIMARY KEY (code));