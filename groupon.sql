/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50517
Source Host           : localhost:3306
Source Database       : groupon

Target Server Type    : MYSQL
Target Server Version : 50517
File Encoding         : 65001

Date: 2018-12-14 09:11:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `admins`
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `role_id` int(11) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES ('1', 'janet', '$2a$10$2UFj1SxfSAKd90b74GHiguCqOoGtkY8xOHhvgTQ1S6kS2iwtYDp3m', '3', '2018-12-07 10:32:37', '2018-12-07 11:14:24');
INSERT INTO `admins` VALUES ('2', 'jianlinker', '$2a$10$P39piGYJiZxJFLUjNYFi9OJpHvUn56EJHYLZGknPSD73u0t0eeZ5C', '3', '2018-12-07 10:49:48', '2018-12-07 10:49:48');
INSERT INTO `admins` VALUES ('3', 'wertsafc', '$2a$10$XTFXDJCVgDfljLxHrQIuUeOuolirNH4I5X7Z1qomUO2jU9e7FDSQy', '3', '2018-12-11 11:06:51', '2018-12-11 11:06:51');

-- ----------------------------
-- Table structure for `banners`
-- ----------------------------
DROP TABLE IF EXISTS `banners`;
CREATE TABLE `banners` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `order` int(11) DEFAULT NULL,
  `cli_num` int(11) DEFAULT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order` (`order`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of banners
-- ----------------------------
INSERT INTO `banners` VALUES ('1', '可口可乐', 'www.baidu.com', 'static/upload/banner/2018/12/10/0b6a78eaace3f7e5134b9ad4939a9b47.jpg', '1', '0', '2018-12-10 12:50:58', '2018-12-10 12:50:58');
INSERT INTO `banners` VALUES ('7', '可口可乐3', 'www.baidu.com', 'static/upload/banner/2018/12/10/a533828b7433a4efce503c6726e4819e.jpg', '3', '0', '2018-12-10 13:49:06', '2018-12-10 13:49:06');
INSERT INTO `banners` VALUES ('8', '消消乐', 'www.taobao.com', 'static/upload/banner/2018/12/10/d6daf57c153bf7631881859d2f11475e.jpg', '9', '50', '2018-12-10 13:49:12', '2018-12-10 16:18:15');
INSERT INTO `banners` VALUES ('9', '可口可乐5', 'www.baidu.com', 'static/upload/banner/2018/12/10/5bfb6757b5543ac4294ddd8c509b2f6e.jpg', '5', '0', '2018-12-10 13:49:16', '2018-12-10 13:49:16');
INSERT INTO `banners` VALUES ('10', '可口可乐2', 'www.baidu.com', 'static/upload/banner/2018/12/10/f715a2e033697386161fdcbca822eadf.jpg', '2', '0', '2018-12-10 15:53:39', '2018-12-10 15:53:39');

-- ----------------------------
-- Table structure for `categorys`
-- ----------------------------
DROP TABLE IF EXISTS `categorys`;
CREATE TABLE `categorys` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `pid` bigint(11) NOT NULL,
  `sort_name` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of categorys
-- ----------------------------
INSERT INTO `categorys` VALUES ('2', '0', '时尚');
INSERT INTO `categorys` VALUES ('4', '2', '面膜');
INSERT INTO `categorys` VALUES ('5', '0', '食品');
INSERT INTO `categorys` VALUES ('6', '2', '口红');
INSERT INTO `categorys` VALUES ('7', '5', '面包');

-- ----------------------------
-- Table structure for `clients`
-- ----------------------------
DROP TABLE IF EXISTS `clients`;
CREATE TABLE `clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` int(11) DEFAULT NULL,
  `client_id` int(11) DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of clients
-- ----------------------------

-- ----------------------------
-- Table structure for `coupons`
-- ----------------------------
DROP TABLE IF EXISTS `coupons`;
CREATE TABLE `coupons` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `kind` varchar(64) DEFAULT NULL,
  `goods_id` int(11) DEFAULT NULL,
  `coupons_rule` float DEFAULT NULL,
  `coupons_norm` float DEFAULT NULL,
  `coupons_status` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of coupons
-- ----------------------------

-- ----------------------------
-- Table structure for `goods`
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_name` varchar(64) DEFAULT NULL,
  `goods_price` varchar(64) DEFAULT NULL,
  `goods_sort` varchar(64) DEFAULT NULL,
  `shop_id` int(11) DEFAULT NULL,
  `goods_stock` int(11) DEFAULT NULL,
  `person_num` int(11) DEFAULT NULL,
  `group_aging` int(11) DEFAULT NULL,
  `discount_price` float DEFAULT NULL,
  `created_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goods
-- ----------------------------

-- ----------------------------
-- Table structure for `groups`
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groups_status` varchar(64) DEFAULT NULL,
  `header_id` int(11) DEFAULT NULL,
  `goods_id` int(11) DEFAULT NULL,
  `groups_aging` datetime DEFAULT NULL,
  `groups_price` float DEFAULT NULL,
  `groups_discount` float DEFAULT NULL,
  `goods_num` int(11) DEFAULT NULL,
  `actually_paid` float DEFAULT NULL,
  `total_price` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of groups
-- ----------------------------

-- ----------------------------
-- Table structure for `group_users`
-- ----------------------------
DROP TABLE IF EXISTS `group_users`;
CREATE TABLE `group_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `header_id` int(11) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group_users
-- ----------------------------

-- ----------------------------
-- Table structure for `orders`
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(64) DEFAULT NULL,
  `goods_id` int(11) DEFAULT NULL,
  `client_id` int(11) DEFAULT NULL,
  `order_price` float DEFAULT NULL,
  `order_status` varchar(64) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  `order_addr` varchar(64) DEFAULT NULL,
  `created_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders
-- ----------------------------

-- ----------------------------
-- Table structure for `orders_return`
-- ----------------------------
DROP TABLE IF EXISTS `orders_return`;
CREATE TABLE `orders_return` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) DEFAULT NULL,
  `client_id` int(11) DEFAULT NULL,
  `shoper_id` int(11) DEFAULT NULL,
  `reason` varchar(164) DEFAULT NULL,
  `status` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders_return
-- ----------------------------

-- ----------------------------
-- Table structure for `person_msgs`
-- ----------------------------
DROP TABLE IF EXISTS `person_msgs`;
CREATE TABLE `person_msgs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_id` int(11) DEFAULT NULL,
  `msg_id` int(11) DEFAULT NULL,
  `msg_status` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of person_msgs
-- ----------------------------

-- ----------------------------
-- Table structure for `roles`
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(64) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES ('1', '普通用户', '2018-12-07 15:17:35', '2018-12-07 15:17:35');
INSERT INTO `roles` VALUES ('2', '商家', '2018-12-07 15:24:33', '2018-12-07 15:24:33');
INSERT INTO `roles` VALUES ('3', '管理员', '2018-12-07 15:25:05', '2018-12-07 15:25:05');
INSERT INTO `roles` VALUES ('4', '房地产商', '2018-12-07 15:17:35', '2018-12-07 16:25:12');
INSERT INTO `roles` VALUES ('5', '审核管理员', '2018-12-07 15:24:33', '2018-12-07 15:24:33');
INSERT INTO `roles` VALUES ('6', '商品管理员', '2018-12-07 15:25:05', '2018-12-07 15:25:05');

-- ----------------------------
-- Table structure for `shops`
-- ----------------------------
DROP TABLE IF EXISTS `shops`;
CREATE TABLE `shops` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `shop_name` varchar(64) NOT NULL,
  `owner_cert` varchar(64) NOT NULL,
  `owner_id` bigint(11) NOT NULL,
  `shop_addr` varchar(64) NOT NULL,
  `shop_phone` varchar(64) NOT NULL,
  `shop_cert` varchar(64) NOT NULL,
  `shop_qq` varchar(64) NOT NULL,
  `shop_logo` varchar(255) NOT NULL,
  `shop_intro` text NOT NULL,
  `is_review` varchar(64) NOT NULL,
  `mark` varchar(64) DEFAULT NULL,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shops
-- ----------------------------
INSERT INTO `shops` VALUES ('1', '消消乐', '4526852385456451', '2', '顺德', '18320399457', '54687132164', '931242644', 'static/upload/merchants/2018/12/13/f58c0463dfef000ff7c4ca9bc1985de8.jpg', 'sdsadasd13212465', '审核通过', '', '2018-12-13 16:07:45', '2018-12-14 08:47:55');
INSERT INTO `shops` VALUES ('3', '消消乐2', '4526852385456451', '5', '顺德2', '18320399457', '54687132164', '931242644', 'static/upload/merchants/2018/12/13/39db0c9e7f0e97a67bab8a3cb8c072a6.jpg', 'sdsadasd13212465', '审核通过', '', '2018-12-13 16:24:32', '2018-12-14 08:54:21');
INSERT INTO `shops` VALUES ('4', '消消乐2', '4526852385456451', '6', '顺德2', '18320399457', '54687132164', '931242644', 'static/upload/merchants/2018/12/13/24cfede81c555d245ec9cc9d2cc6a0eb.jpg', 'sdsadasd13212465', '正在审核', null, '2018-12-13 16:24:36', '2018-12-13 16:24:36');

-- ----------------------------
-- Table structure for `shop_msgs`
-- ----------------------------
DROP TABLE IF EXISTS `shop_msgs`;
CREATE TABLE `shop_msgs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shoper_id` int(11) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  `content` varchar(64) DEFAULT NULL,
  `created_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shop_msgs
-- ----------------------------

-- ----------------------------
-- Table structure for `tb_users`
-- ----------------------------
DROP TABLE IF EXISTS `tb_users`;
CREATE TABLE `tb_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_users
-- ----------------------------
INSERT INTO `tb_users` VALUES ('0', 'admin', '$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG', '2018-05-27 16:25:33', '2018-05-27 16:25:33', null);
INSERT INTO `tb_users` VALUES ('3', 'kong', '$2a$10$nxzKu7HAl4PlZZ90LuXNqO5wc.9xY4Hp5Z47YupkNHwnexQZvRew.', '2018-11-27 13:14:30', '2018-11-27 13:14:30', null);
INSERT INTO `tb_users` VALUES ('4', 'kong2', '$2a$10$EQz6dmNaHG1xWlddnO6m7.F6cqVvXhxQhcKUZANLqlbuCLH5DL19i', '2018-11-27 13:36:38', '2018-11-27 13:36:38', null);

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `phonenum` varchar(64) NOT NULL,
  `nick_name` varchar(64) DEFAULT NULL,
  `password` varchar(64) NOT NULL,
  `address` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `head_image` varchar(255) DEFAULT NULL,
  `sex` varchar(64) DEFAULT '男',
  `account` float unsigned DEFAULT '0',
  `role_id` int(11) NOT NULL DEFAULT '1',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phoneindex` (`phonenum`(25)) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('2', '18320399457', '小白', '$2a$10$PI4zH/PTaMPV4vc4mRM1n.Bdt222bdgUIER1yKOjRlZHP1yMrL6Uu', '顺德', '伟林', 'static/upload/user/2018/11/29/25c73b25cc67595a78fb387eea0c4972.jpg', '男', '0', '2', '2018-11-29 16:47:08', '2018-12-14 08:47:53');
INSERT INTO `users` VALUES ('4', '13659763182', '', '$2a$10$hsxR4ofpweH1inhlsNo2OeClufMDwEvn3xj9Q6rndPN5JyqUpyyx6', '陆丰', '富华', '', '男', '0', '1', '2018-11-30 09:55:26', '2018-11-30 09:55:26');
INSERT INTO `users` VALUES ('5', '13420120750', '', '$2a$10$7SyJNdSzEGd521IKFJR73eTnl8chkzDXupto8tWWIbZljdq27koHi', '廉江', '景润', '', '男', '0', '2', '2018-11-30 10:13:02', '2018-12-14 08:54:21');
INSERT INTO `users` VALUES ('6', '13420121717', '小茂', '$2a$10$LtwwPtQ8/R4Rv7RcxgRXPug2UAP1C6TGnf3UlREjwPUAogYPSON02', '深圳', '林佳茂', '', '男', '0', '1', '2018-11-30 18:54:46', '2018-11-30 21:27:25');

-- ----------------------------
-- Table structure for `user_coupons`
-- ----------------------------
DROP TABLE IF EXISTS `user_coupons`;
CREATE TABLE `user_coupons` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `coupons_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_coupons
-- ----------------------------
