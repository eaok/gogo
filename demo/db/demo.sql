SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo` (
	`uid` INT (10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR (64) NULL DEFAULT NULL,
	`departname` VARCHAR (64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

DROP TABLE IF EXISTS `userdetail`;
CREATE TABLE `userdetail` (
	`uid` INT (10) NOT NULL DEFAULT ''0'',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
	`id` INT (4) NOT NULL,
	`name` VARCHAR (30) DEFAULT NULL,
	`money` FLOAT (8, 2) DEFAULT NULL,
	PRIMARY KEY (`id`)
)
INSERT INTO `account` VALUES ('1', '杨超越', '3000.00');
INSERT INTO `account` VALUES ('2', '王一博', '1000.00');