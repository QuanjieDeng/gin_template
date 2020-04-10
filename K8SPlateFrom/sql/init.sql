-- 创建表 编码为utf8_general_ci
CREATE  DATABASE  gvp;


SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for gvp_auth
-- ----------------------------
DROP TABLE IF EXISTS `gvp`.`gvp_auth`;
CREATE TABLE `gvp`.`gvp_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='认证表';

-- ----------------------------
-- Table structure for gvp_game
-- ----------------------------
DROP TABLE IF EXISTS `gvp`.`gvp_game`;
CREATE TABLE `gvp`.`gvp_game` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `game` varchar(100) NOT NULL COMMENT '游戏名称',
  `game_type` int(10) DEFAULT '0' COMMENT '游戏类型',
  `app_id` varchar(100) NOT NULL COMMENT 'appid',
  `app_key` varchar(100) NOT NULL COMMENT 'appkey',
  `tel_num` varchar(20) NOT NULL COMMENT '手机号码',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `state` int(10) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `game` (`game`),
  UNIQUE KEY `app_id` (`app_id`),
  UNIQUE KEY `app_key` (`app_key`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8 COMMENT='游戏列表';

-- ----------------------------
-- Table structure for gvp_game_service
-- ----------------------------
DROP TABLE IF EXISTS `gvp`.`gvp_game_service`;
CREATE TABLE `gvp`.`gvp_game_service` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `game` varchar(100) NOT NULL COMMENT '游戏名称',
  `service_id` int(10) NOT NULL COMMENT '服务商ID',
  `game_id` int(10) NOT NULL COMMENT '游戏ID',
  `deleted_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='游戏服务商绑定关系';

-- ----------------------------
-- Table structure for gvp_voice_service
-- ----------------------------
DROP TABLE IF EXISTS `gvp`.`gvp_voice_service`;
CREATE TABLE `gvp`.`gvp_voice_service` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '服务名称',
  `app_id` varchar(100) NOT NULL COMMENT 'appid',
  `app_key` varchar(100) NOT NULL COMMENT 'appkey',
  `user_id` varchar(100) DEFAULT '' COMMENT '用户ID',
  `url` varchar(100) DEFAULT '' COMMENT '连接地址',
  `description` varchar(300) DEFAULT '' COMMENT '服务描述',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `state` int(10) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `app_id` (`app_id`),
  UNIQUE KEY `app_key` (`app_key`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8 COMMENT='服务商信息表';

-- 认证用户数据初始化
INSERT INTO `gvp`.`gvp_auth` (`id`, `username`, `password`) VALUES (null, 'admin', 'admin');