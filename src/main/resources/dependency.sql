/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : dependency

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 12/07/2022 09:47:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for edgeJava
-- ----------------------------
DROP TABLE IF EXISTS `edgeExample`;
create table edgeJava
(
    id         int auto_increment
        primary key,
    source     int          not null,
    target     int          not null,
    attributes varchar(255) null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

create index source__index
    on edgeJava (source);

create index target__index
    on edgeJava (target);


-- ----------------------------
-- Table structure for nodeJava
-- ----------------------------
DROP TABLE IF EXISTS `nodeExample`;
CREATE TABLE `nodeJava`
(
    `id`         int NOT NULL,
    `name`       varchar(255)                                                  DEFAULT NULL,
    `attributes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for graph
-- ----------------------------
DROP TABLE IF EXISTS `graph`;
CREATE TABLE `graph`
(
    `id`   int NOT NULL,
    `name` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
