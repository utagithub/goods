/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50743 (5.7.43)
 Source Host           : localhost:3306
 Source Schema         : good

 Target Server Type    : MySQL
 Target Server Version : 50743 (5.7.43)
 File Encoding         : 65001

 Date: 27/09/2023 16:10:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` (`id`, `cat_name`) VALUES (1, '配件');
INSERT INTO `category` (`id`, `cat_name`) VALUES (2, '小吃');
INSERT INTO `category` (`id`, `cat_name`) VALUES (3, '零食');
COMMIT;

-- ----------------------------
-- Table structure for good
-- ----------------------------
DROP TABLE IF EXISTS `good`;
CREATE TABLE `good` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `goods_name` varchar(32) DEFAULT NULL COMMENT '商品名称',
  `cat_id` tinyint(4) DEFAULT NULL COMMENT '商品分类',
  `goods_price` varchar(128) DEFAULT NULL COMMENT '商品价格',
  `goods_thumb` varchar(255) DEFAULT NULL COMMENT '商品缩略图',
  `goods_desc` varchar(255) DEFAULT NULL COMMENT '商品描述',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_good_deleted_at` (`deleted_at`),
  KEY `idx_good_create_by` (`create_by`),
  KEY `idx_good_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of good
-- ----------------------------
BEGIN;
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (1, '888', 1, '666', '666', '666', '2023-09-24 22:37:25.393', '2023-09-24 22:37:25.393', '2023-09-24 23:07:35.770', NULL, 1);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (2, '888', 1, '666', '666', '666', '2023-09-24 22:37:26.273', '2023-09-24 23:14:52.578', '2023-09-24 23:14:52.580', NULL, 1);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (3, '777', 1, '555', '555', '555', '2023-09-24 22:37:26.940', '2023-09-24 23:16:15.719', '2023-09-24 23:16:15.720', 0, 1);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (4, '888', 1, '666', '666', '666', '2023-09-24 22:49:23.435', '2023-09-24 22:49:23.435', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (5, '888', 1, '666', '666', '666', '2023-09-24 23:30:30.619', '2023-09-24 23:30:30.619', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (6, '888', 1, '666', '666', '666', '2023-09-24 23:30:31.131', '2023-09-24 23:30:31.131', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (7, '888', 1, '666', '666', '666', '2023-09-24 23:30:31.602', '2023-09-24 23:30:31.602', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (8, '888', 1, '666', '666', '666', '2023-09-24 23:30:32.057', '2023-09-24 23:30:32.057', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (9, '888', 1, '666', '666', '666', '2023-09-24 23:30:32.454', '2023-09-24 23:30:32.454', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (10, '888', 1, '666', '666', '666', '2023-09-24 23:30:32.925', '2023-09-24 23:30:32.925', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (11, '888', 1, '666', '666', '666', '2023-09-24 23:30:33.532', '2023-09-24 23:30:33.532', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (12, '888', 1, '666', '666', '666', '2023-09-24 23:30:34.358', '2023-09-24 23:30:34.358', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (13, '888', 1, '666', '666', '666', '2023-09-24 23:30:34.911', '2023-09-24 23:30:34.911', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (14, '888', 1, '666', '666', '666', '2023-09-24 23:30:35.303', '2023-09-24 23:30:35.303', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (15, '888', 1, '666', '666', '666', '2023-09-24 23:30:35.624', '2023-09-24 23:30:35.624', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (16, '888', 1, '666', '666', '666', '2023-09-24 23:30:35.974', '2023-09-24 23:30:35.974', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (17, '888', 1, '666', '666', '666', '2023-09-24 23:30:36.292', '2023-09-24 23:30:36.292', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (18, '888', 1, '666', '666', '666', '2023-09-24 23:30:37.360', '2023-09-24 23:30:37.360', NULL, 1, NULL);
INSERT INTO `good` (`id`, `goods_name`, `cat_id`, `goods_price`, `goods_thumb`, `goods_desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (19, '888', 3, '666', '666', '666', '2023-09-24 23:30:56.844', '2023-09-24 23:30:56.844', NULL, 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `path` varchar(128) DEFAULT NULL COMMENT '地址',
  `type` varchar(16) DEFAULT NULL COMMENT '接口类型',
  `action` varchar(16) DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`),
  KEY `idx_sys_api_create_by` (`create_by`),
  KEY `idx_sys_api_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`),
  KEY `idx_sys_dept_create_by` (`create_by`),
  KEY `idx_sys_dept_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_name` varchar(255) DEFAULT NULL,
  `job_group` varchar(255) DEFAULT NULL,
  `job_type` tinyint(4) DEFAULT NULL,
  `cron_expression` varchar(255) DEFAULT NULL,
  `invoke_target` varchar(255) DEFAULT NULL,
  `args` varchar(255) DEFAULT NULL,
  `misfire_policy` bigint(20) DEFAULT NULL,
  `concurrent` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `entry_id` smallint(6) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`job_id`),
  KEY `idx_sys_job_create_by` (`create_by`),
  KEY `idx_sys_job_update_by` (`update_by`),
  KEY `idx_sys_job_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (1, 'testJob', 'DEFAULT', 2, '0/5 * * * *', 'ExamplesOne', 'args', 1, 2, 1, 0, '2023-09-25 00:40:30.459', '2023-09-25 00:55:52.853', '2023-09-25 00:55:52.855', 1, 1);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (2, 'testJob', 'DEFAULT', 2, '0/5 * * * *', 'ExamplesOne', 'args', 1, 2, 1, 0, '2023-09-25 00:40:33.939', '2023-09-25 00:55:52.853', '2023-09-25 00:55:52.855', 1, 1);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (3, 'testJob11', 'DEFAULT', 2, '0/5 * * * *', 'ExamplesOne', 'args-ss', 1, 2, 1, 0, '2023-09-25 00:40:34.571', '2023-09-25 00:55:52.853', '2023-09-25 00:55:52.855', 1, 1);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (8, 'testJob11', 'DEFAULT', 2, '0/5 * * * *', 'Examples1', 'args-ss', 1, 2, 1, 0, '2023-09-25 00:57:06.461', '2023-09-26 00:19:54.724', NULL, 1, 1);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (9, 'testJob', 'DEFAULT', 2, '0/5 * * * *', 'Examples2', 'args', 1, 2, 1, 0, '2023-09-25 00:57:06.809', '2023-09-25 00:57:06.809', NULL, 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` smallint(6) DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` tinyint(4) DEFAULT NULL,
  `visible` varchar(1) DEFAULT NULL,
  `is_frame` varchar(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`),
  KEY `idx_sys_menu_create_by` (`create_by`),
  KEY `idx_sys_menu_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_menu_id` bigint(20) NOT NULL,
  `sys_api_id` bigint(20) NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`),
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_migration
-- ----------------------------
DROP TABLE IF EXISTS `sys_migration`;
CREATE TABLE `sys_migration` (
  `version` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `apply_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sys_migration
-- ----------------------------
BEGIN;
INSERT INTO `sys_migration` (`version`, `apply_time`) VALUES ('1695532856746', '2023-09-24 14:00:21.309');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` bigint(20) DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_update_by` (`update_by`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_sys_role_create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL,
  `menu_id` bigint(20) NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `role_id` bigint(20) DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门',
  `post_id` bigint(20) DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` tinyint(4) DEFAULT NULL COMMENT '创建者',
  `update_by` tinyint(4) DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_create_by` (`create_by`),
  KEY `idx_sys_user_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'zhangwj', '13818888888', 1, '', '', '1', '1@qq.com', 1, 1, '', '2', '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '密码',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$09$kMo41Zct9ZxXaF7r.GfSAu1ZtbvEJQW237dFbppU8PqSmWoWPOYN.', NULL, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
