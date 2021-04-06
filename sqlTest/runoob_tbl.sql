/*
 Navicat Premium Data Transfer

 Source Server         : local_mysql_ubuntu
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : 192.168.3.8:3306
 Source Schema         : godb

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 06/04/2021 23:30:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for runoob_tbl
-- ----------------------------
DROP TABLE IF EXISTS `runoob_tbl`;
CREATE TABLE `runoob_tbl` (
  `runoob_id` int unsigned NOT NULL AUTO_INCREMENT,
  `runoob_title` varchar(100) NOT NULL,
  `runoob_author` varchar(40) NOT NULL,
  `submission_date` date DEFAULT NULL,
  PRIMARY KEY (`runoob_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of runoob_tbl
-- ----------------------------
BEGIN;
INSERT INTO `runoob_tbl` VALUES (1, 'go', 'tp', '2021-04-06');
INSERT INTO `runoob_tbl` VALUES (2, ' MySQL', '', '2021-04-06');
INSERT INTO `runoob_tbl` VALUES (3, 'goa', 'tp2', NULL);
INSERT INTO `runoob_tbl` VALUES (4, 'hyh01', 'huo', '2021-04-06');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
