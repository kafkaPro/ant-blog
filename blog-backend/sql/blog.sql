/*
 Navicat Premium Data Transfer

 Source Server         : gfblog
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : database:23156
 Source Schema         : ant-blog

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : 65001

 Date: 02/09/2021 16:55:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户唯一标识UUID',
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'QMPlusUser' COMMENT '用户昵称',
  `header_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg' COMMENT '用户头像',
  `role_id` int(10) UNSIGNED NULL DEFAULT 888 COMMENT '用户角色ID',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admins_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 'b4c54e5a-d015-4f8c-9f01-624c527a63ae', '超级管理员', 'http://qmplusimg.henrongyi.top/1571627762timg.jpg', 888, 'admin', '$2a$10$zF5PNCWobve/0RBk.3k03eAGwyvDevFBFd3AZUwETjMhYpZwNooba');
INSERT INTO `admins` VALUES (2, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 'fd6ef79b-944c-4888-8377-abe2d2608858', 'QMPlusUser', 'http://qmplusimg.henrongyi.top/1572075907logo.png', 9528, 'a303176530', '$2a$10$zF5PNCWobve/0RBk.3k03eAGwyvDevFBFd3AZUwETjMhYpZwNooba');

-- ----------------------------
-- Table structure for apis
-- ----------------------------
DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_apis_delete_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 148 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of apis
-- ----------------------------
INSERT INTO `apis` VALUES (1, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO `apis` VALUES (2, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/base/register', '用户注册', 'base', 'POST');
INSERT INTO `apis` VALUES (3, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `apis` VALUES (4, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `apis` VALUES (5, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `apis` VALUES (6, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `apis` VALUES (7, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `apis` VALUES (8, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `apis` VALUES (9, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `apis` VALUES (10, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `apis` VALUES (11, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `apis` VALUES (12, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO `apis` VALUES (13, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `apis` VALUES (14, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `apis` VALUES (15, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `apis` VALUES (16, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `apis` VALUES (17, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `apis` VALUES (18, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `apis` VALUES (19, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `apis` VALUES (20, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `apis` VALUES (21, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/changePassword', '修改密码', 'user', 'POST');
INSERT INTO `apis` VALUES (22, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/uploadHeaderImg', '上传头像', 'user', 'POST');
INSERT INTO `apis` VALUES (23, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/getInfoList', '分页获取用户列表', 'user', 'POST');
INSERT INTO `apis` VALUES (24, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `apis` VALUES (25, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/setUserAuthority', '修改用户角色', 'user', 'POST');
INSERT INTO `apis` VALUES (26, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO `apis` VALUES (27, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO `apis` VALUES (28, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `apis` VALUES (29, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `apis` VALUES (30, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO `apis` VALUES (31, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单', 'jwt', 'POST');
INSERT INTO `apis` VALUES (32, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `apis` VALUES (33, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO `apis` VALUES (34, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO `apis` VALUES (35, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/customer/customer', '创建客户', 'customer', 'POST');
INSERT INTO `apis` VALUES (36, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/customer/customer', '更新客户', 'customer', 'PUT');
INSERT INTO `apis` VALUES (37, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/customer/customer', '删除客户', 'customer', 'DELETE');
INSERT INTO `apis` VALUES (38, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/customer/customer', '获取单一客户', 'customer', 'GET');
INSERT INTO `apis` VALUES (39, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/customer/customerList', '获取客户列表', 'customer', 'GET');
INSERT INTO `apis` VALUES (40, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `apis` VALUES (41, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/autoCode/createTemp', '自动化代码', 'autoCode', 'POST');
INSERT INTO `apis` VALUES (42, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO `apis` VALUES (43, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
INSERT INTO `apis` VALUES (44, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO `apis` VALUES (45, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', 'sysDictionaryDetail', 'POST');
INSERT INTO `apis` VALUES (46, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', 'sysDictionaryDetail', 'DELETE');
INSERT INTO `apis` VALUES (47, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', 'sysDictionaryDetail', 'PUT');
INSERT INTO `apis` VALUES (48, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', 'sysDictionaryDetail', 'GET');
INSERT INTO `apis` VALUES (49, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', 'sysDictionaryDetail', 'GET');
INSERT INTO `apis` VALUES (50, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionary/createSysDictionary', '新增字典', 'sysDictionary', 'POST');
INSERT INTO `apis` VALUES (51, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', 'sysDictionary', 'DELETE');
INSERT INTO `apis` VALUES (52, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionary/updateSysDictionary', '更新字典', 'sysDictionary', 'PUT');
INSERT INTO `apis` VALUES (53, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', 'sysDictionary', 'GET');
INSERT INTO `apis` VALUES (54, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', 'sysDictionary', 'GET');
INSERT INTO `apis` VALUES (55, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/autoCode/getTables', '获取数据库表', 'autoCode', 'GET');
INSERT INTO `apis` VALUES (56, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/autoCode/getDB', '获取所有数据库', 'autoCode', 'GET');
INSERT INTO `apis` VALUES (57, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/autoCode/getColume', '获取所选table的所有字段', 'autoCode', 'GET');
INSERT INTO `apis` VALUES (58, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/user/setUserInfo', '设置用户信息', 'user', 'PUT');
INSERT INTO `apis` VALUES (59, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/simpleUploader/upload', '插件版分片上传', 'simpleUploader', 'POST');
INSERT INTO `apis` VALUES (60, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', 'simpleUploader', 'GET');
INSERT INTO `apis` VALUES (61, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', 'simpleUploader', 'GET');

-- ----------------------------
-- Table structure for breakpoint_chucks
-- ----------------------------
DROP TABLE IF EXISTS `breakpoint_chucks`;
CREATE TABLE `breakpoint_chucks`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `exa_file_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '文件id',
  `file_chunk_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '切片路径',
  `file_chunk_number` int(11) NULL DEFAULT NULL COMMENT '切片标号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of breakpoint_chucks
-- ----------------------------

-- ----------------------------
-- Table structure for breakpoint_files
-- ----------------------------
DROP TABLE IF EXISTS `breakpoint_files`;
CREATE TABLE `breakpoint_files`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `file_md5` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件md5',
  `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  `chunk_id` int(11) NULL DEFAULT NULL COMMENT '关联id',
  `chunk_total` int(11) NULL DEFAULT NULL COMMENT '切片总数',
  `is_finish` tinyint(1) NULL DEFAULT NULL COMMENT '是否完整',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_breakpoint_files_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of breakpoint_files
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/register', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/uploadHeaderImg', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getInfoList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserInfo', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customerList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/createTemp', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getTables', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getDB', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getColume', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/createSysDictionary', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/findSysDictionary', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/upload', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/checkFileMd5', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/mergeFileMd5', 'GET', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for customers
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `customer_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '客户电话',
  `sys_user_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '负责员工id',
  `sys_user_role_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责员工角色',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of customers
-- ----------------------------

-- ----------------------------
-- Table structure for dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `dictionaries`;
CREATE TABLE `dictionaries`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of dictionaries
-- ----------------------------
INSERT INTO `dictionaries` VALUES (1, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '性别', 'sex', 1, '性别字典');
INSERT INTO `dictionaries` VALUES (2, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '数据库int类型', 'int', 1, '数据库int类型');
INSERT INTO `dictionaries` VALUES (3, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `dictionaries` VALUES (4, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `dictionaries` VALUES (5, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `dictionaries` VALUES (6, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `dictionary_details`;
CREATE TABLE `dictionary_details`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` int(11) NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序标记',
  `dictionary_id` int(11) NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of dictionary_details
-- ----------------------------
INSERT INTO `dictionary_details` VALUES (1, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `dictionary_details` VALUES (2, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `dictionary_details` VALUES (3, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `dictionary_details` VALUES (4, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `dictionary_details` VALUES (5, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'data', 0, 1, 0, 3);
INSERT INTO `dictionary_details` VALUES (6, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `dictionary_details` VALUES (7, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `dictionary_details` VALUES (8, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `dictionary_details` VALUES (9, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'timestamp', 4, 1, 4, 3);
INSERT INTO `dictionary_details` VALUES (10, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `dictionary_details` VALUES (11, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'double', 0, 1, 0, 4);
INSERT INTO `dictionary_details` VALUES (12, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'decimal', 0, 1, 0, 4);
INSERT INTO `dictionary_details` VALUES (13, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (14, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'varchar', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (15, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'tinyblob', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (16, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'tinytext', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (17, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'text', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (18, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'blob', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (19, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'mediumblob', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (20, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'mediumtext', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (21, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'longblob', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (22, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'longtext', 0, 1, 0, 5);
INSERT INTO `dictionary_details` VALUES (23, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL, 'tinyint', 4, 1, 4, 6);

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_files_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of files
-- ----------------------------
INSERT INTO `files` VALUES (1, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, '10.png', 'http://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `files` VALUES (2, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for jwts
-- ----------------------------
DROP TABLE IF EXISTS `jwts`;
CREATE TABLE `jwts`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `jwt` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwts_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 145 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of jwts
-- ----------------------------

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录用户名',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录ip地址',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录的浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Windows' COMMENT '登录的操作系统类型，默认Windows',
  `status` int(11) NULL DEFAULT 0 COMMENT '登录结果(0->失败,1->成功)',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录结果信息(例如:登录成功)',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `login_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登录日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_log
-- ----------------------------
INSERT INTO `login_log` VALUES (1, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:29:19');
INSERT INTO `login_log` VALUES (2, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:29:55');
INSERT INTO `login_log` VALUES (3, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:30:25');
INSERT INTO `login_log` VALUES (4, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:30:26');
INSERT INTO `login_log` VALUES (5, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:30:27');
INSERT INTO `login_log` VALUES (6, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:30:27');
INSERT INTO `login_log` VALUES (7, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:30:28');
INSERT INTO `login_log` VALUES (8, 'admin', '[::1]', 'PostmanRuntime', 'Windows10', 1, '登录成功', '内网IP', '2021-09-02 15:39:47');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `menu_level` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '菜单等级(预留字段)',
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端文件路径',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '是否缓存',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '是否是基础路由(开发中)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menus_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 513 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', '仪表盘', 'setting', 1, 0, 0);
INSERT INTO `menus` VALUES (2, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'about', 'about', 0, 'view/about/index.vue', '关于我们', 'info', 7, 0, 0);
INSERT INTO `menus` VALUES (3, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'superAdmin', 'admin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 3, 0, 0);
INSERT INTO `menus` VALUES (4, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 1, 0, 0);
INSERT INTO `menus` VALUES (5, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 2, 1, 0);
INSERT INTO `menus` VALUES (6, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, 1, 0);
INSERT INTO `menus` VALUES (7, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 4, 0, 0);
INSERT INTO `menus` VALUES (8, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'message-solid', 4, 0, 0);
INSERT INTO `menus` VALUES (9, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'example', 'example', 0, 'view/example/index.vue', '示例文件', 's-management', 6, 0, 0);
INSERT INTO `menus` VALUES (10, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', 1, 0, 0);
INSERT INTO `menus` VALUES (11, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', 2, 0, 0);
INSERT INTO `menus` VALUES (12, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', 3, 0, 0);
INSERT INTO `menus` VALUES (13, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 4, 0, 0);
INSERT INTO `menus` VALUES (14, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', 5, 0, 0);
INSERT INTO `menus` VALUES (15, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', 6, 0, 0);
INSERT INTO `menus` VALUES (16, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', 7, 0, 0);
INSERT INTO `menus` VALUES (17, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', '系统工具', 's-cooperation', 5, 0, 0);
INSERT INTO `menus` VALUES (18, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 17, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', '代码生成器', 'cpu', 1, 1, 0);
INSERT INTO `menus` VALUES (19, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 17, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', '表单生成器', 'magic-stick', 2, 1, 0);
INSERT INTO `menus` VALUES (20, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 17, 'system', 'system', 0, 'view/systemTools/system/system.vue', '系统配置', 's-operation', 3, 0, 0);
INSERT INTO `menus` VALUES (21, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'iconList', 'iconList', 0, 'view/iconList/index.vue', '图标集合', 'star-on', 2, 0, 0);
INSERT INTO `menus` VALUES (22, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', '字典管理', 'notebook-2', 5, 0, 0);
INSERT INTO `menus` VALUES (23, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', '字典详情', 's-order', 1, 0, 0);
INSERT INTO `menus` VALUES (24, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', '操作历史', 'time', 6, 0, 0);
INSERT INTO `menus` VALUES (25, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 0, 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', '官方网站', 's-home', 0, 0, 0);
INSERT INTO `menus` VALUES (26, '2021-08-31 19:08:04', '2021-08-31 19:08:04', NULL, 0, 9, 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', '断点续传（插件版）', 'upload', 6, 0, 0);

-- ----------------------------
-- Table structure for online_user
-- ----------------------------
DROP TABLE IF EXISTS `online_user`;
CREATE TABLE `online_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户的uuid',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '在线用户名',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '在线用户token',
  `explorer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '在线用户浏览器类型',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '在线用户的操作系统类型',
  `login_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `is_admin` int(11) NULL DEFAULT NULL COMMENT '是否是管理员',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '在线用户列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of online_user
-- ----------------------------
INSERT INTO `online_user` VALUES (7, '4ea40fac7acff8dc3aa303b73758a6d1', 'admin', 'pQ3dDeZGzCIZRz21PaBfYVorKf5ZDTC77y7KVFgSfafTNqlWDtTA5pHmtYFR09kX', 'PostmanRuntime', '', '2021-09-02 15:40:10', 1);

-- ----------------------------
-- Table structure for operate_log
-- ----------------------------
DROP TABLE IF EXISTS `operate_log`;
CREATE TABLE `operate_log`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `module` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作的模块的名称',
  `opt_type` int(11) NULL DEFAULT NULL COMMENT '操作类型(0->其他, 1->新增, 2->修改, 3->删除)',
  `opt_person` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作员姓名',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `request_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求的URL',
  `request_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法名',
  `request_param` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求的参数',
  `status` int(11) NULL DEFAULT 0 COMMENT '操作状态(0->正常, 1->异常)',
  `opt_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operate_log
-- ----------------------------

-- ----------------------------
-- Table structure for operations
-- ----------------------------
DROP TABLE IF EXISTS `operations`;
CREATE TABLE `operations`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求路由',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '报错信息',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '请求Body',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '响应Body',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_operations_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 392 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of operations
-- ----------------------------

-- ----------------------------
-- Table structure for parameters
-- ----------------------------
DROP TABLE IF EXISTS `parameters`;
CREATE TABLE `parameters`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `base_menu_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT 'BaseMenu的ID',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of parameters
-- ----------------------------

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `role_id` int(10) UNSIGNED NOT NULL COMMENT '角色ID',
  `role_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`role_id`) USING BTREE,
  INDEX `idx_sys_authorities_deleted_at`(`delete_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (888, '超级管理员', 0, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL);
INSERT INTO `role` VALUES (8881, '超级管理员子账号', 888, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL);
INSERT INTO `role` VALUES (9528, '测试角色', 0, '2021-08-31 19:08:05', '2021-08-31 19:08:05', NULL);

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `role_id` int(10) UNSIGNED NOT NULL COMMENT '权限id',
  `menu_id` int(10) UNSIGNED NOT NULL COMMENT '菜单id',
  INDEX `menu_id`(`menu_id`) USING BTREE,
  INDEX `authority_id`(`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (888, 1);
INSERT INTO `role_menu` VALUES (888, 2);
INSERT INTO `role_menu` VALUES (888, 3);
INSERT INTO `role_menu` VALUES (888, 4);
INSERT INTO `role_menu` VALUES (888, 5);
INSERT INTO `role_menu` VALUES (888, 6);
INSERT INTO `role_menu` VALUES (888, 7);
INSERT INTO `role_menu` VALUES (888, 8);
INSERT INTO `role_menu` VALUES (888, 9);
INSERT INTO `role_menu` VALUES (888, 10);
INSERT INTO `role_menu` VALUES (888, 11);
INSERT INTO `role_menu` VALUES (888, 12);
INSERT INTO `role_menu` VALUES (888, 13);
INSERT INTO `role_menu` VALUES (888, 14);
INSERT INTO `role_menu` VALUES (888, 15);
INSERT INTO `role_menu` VALUES (888, 16);
INSERT INTO `role_menu` VALUES (888, 17);
INSERT INTO `role_menu` VALUES (888, 18);
INSERT INTO `role_menu` VALUES (888, 19);
INSERT INTO `role_menu` VALUES (888, 20);
INSERT INTO `role_menu` VALUES (888, 21);
INSERT INTO `role_menu` VALUES (888, 22);
INSERT INTO `role_menu` VALUES (888, 23);
INSERT INTO `role_menu` VALUES (888, 24);
INSERT INTO `role_menu` VALUES (888, 25);
INSERT INTO `role_menu` VALUES (888, 26);

-- ----------------------------
-- Table structure for role_resources
-- ----------------------------
DROP TABLE IF EXISTS `role_resources`;
CREATE TABLE `role_resources`  (
  `role_id` int(10) UNSIGNED NOT NULL COMMENT '权限id',
  `resources_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '资源id',
  INDEX `resources_id`(`resources_id`) USING BTREE,
  INDEX `authority_id`(`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_resources
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
