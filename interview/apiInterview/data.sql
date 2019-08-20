DROP TABLE IF EXISTS `student` ;
CREATE TABLE `student` (
  `student_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'student_id',
  `firstname` VARCHAR(10) NOT NULL COMMENT 'firstname',
  `lastname` VARCHAR(10) NOT NULL COMMENT 'lastname',
  `age` INT(11) UNSIGNED NOT NULL COMMENT 'age',
  PRIMARY KEY (`student_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='@学生数据表';

INSERT INTO `student` VALUES ('1', '理工', '吴', '22');
INSERT INTO `student` VALUES ('2', '莎拉', '将', '21');
INSERT INTO `student` VALUES ('3', '工', '欧阳', '20');
INSERT INTO `student` VALUES ('4', '算法', '王', '22');

DROP TABLE IF EXISTS `user` ;
CREATE TABLE `user` (
  `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'user_id',
  `name` VARCHAR(10) NOT NULL COMMENT 'name',
  `age` INT(11) UNSIGNED NOT NULL COMMENT 'age',
  PRIMARY KEY (`user_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='@用户数据表';

INSERT INTO `user` VALUES ('100029', 'peter', '22');
INSERT INTO `user` VALUES ('100030', 'tom', '50');
INSERT INTO `user` VALUES ('100031', 'karl', '34');
INSERT INTO `user` VALUES ('100032', 'mary', '18');

DROP TABLE IF EXISTS `int_auth_token_cache` ;
CREATE TABLE `int_auth_token_cache` (
  `int_auth_token` VARCHAR(30) NOT NULL COMMENT 'int_auth_token',
  `user_id` INT(11) UNSIGNED NOT NULL COMMENT 'user_id',
  `device` VARCHAR(30) NOT NULL COMMENT 'device',
  `ip` VARCHAR(30) NOT NULL COMMENT 'ip',
  PRIMARY KEY (`int_auth_token`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='@token数据表';

INSERT INTO `int_auth_token_cache` VALUES ('xxxyyyzzz', '100029', 'iPhone6', '192.168.1.88');
INSERT INTO `int_auth_token_cache` VALUES ('aaabbbccc', '100029', 'Samsung Galaxy S3', '177.15.33.8');
INSERT INTO `int_auth_token_cache` VALUES ('pppqqqsss', '100030', 'Samsung Note Bomb', '192.168.7.55');
INSERT INTO `int_auth_token_cache` VALUES ('dddeeefff', '100031', 'Xiaomi 5X', '192.168.7.58');
INSERT INTO `int_auth_token_cache` VALUES ('eeefffggg', '100031', 'Xiaomi 4', '111.20.3.7');
INSERT INTO `int_auth_token_cache` VALUES ('yuqbajnnr', '100032', 'iPhone SE', '121.2.88.137');
