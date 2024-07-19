/*
 Navicat Premium Data Transfer

 Source Server         : mysql-8
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : wukong

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 18/06/2024 18:00:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs`  (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `updated_at` datetime NULL DEFAULT NULL,
                            `created_at` datetime NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for certs
-- ----------------------------
DROP TABLE IF EXISTS `certs`;
CREATE TABLE `certs`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `private_key` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `public_key` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `purpose` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `created_at` datetime NULL DEFAULT NULL,
                          `updated_at` datetime NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for login_securities
-- ----------------------------
DROP TABLE IF EXISTS `login_securities`;
CREATE TABLE `login_securities`  (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `user_id` int NOT NULL,
                                     `last_login_at` datetime NOT NULL,
                                     `login_fail_count` int NULL DEFAULT NULL,
                                     `lock_at` datetime NULL DEFAULT NULL,
                                     `updated_at` datetime NULL DEFAULT NULL,
                                     `created_at` datetime NULL DEFAULT NULL,
                                     PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for mfa_apps
-- ----------------------------
DROP TABLE IF EXISTS `mfa_apps`;
CREATE TABLE `mfa_apps`  (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int NOT NULL,
                             `totp_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `updated_at` datetime NULL DEFAULT NULL,
                             `created_at` datetime NULL DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `uq_index_totp_key`(`totp_secret` ASC) USING BTREE,
                             UNIQUE INDEX `uq_index_user_id`(`user_id` ASC) USING BTREE
);

-- ----------------------------
-- Table structure for oauth_apps
-- ----------------------------
DROP TABLE IF EXISTS `oauth_apps`;
CREATE TABLE `oauth_apps`  (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `app_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                               `client_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                               `client_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                               `scope` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                               `redirect_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                               `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                               `created_at` datetime NULL DEFAULT NULL,
                               `updated_at` datetime NULL DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `index_app_name`(`app_name` ASC) USING BTREE,
                               UNIQUE INDEX `index_client_id`(`client_id` ASC) USING BTREE
);

-- ----------------------------
-- Table structure for pass_keys
-- ----------------------------
DROP TABLE IF EXISTS `pass_keys`;
CREATE TABLE `pass_keys`  (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `display_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                              `user_id` int NOT NULL,
                              `credential_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                              `credential_raw` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                              `last_used_at` datetime NULL DEFAULT NULL,
                              `updated_at` datetime NULL DEFAULT NULL,
                              `created_at` datetime NULL DEFAULT NULL,
                              PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for resources
-- ----------------------------
DROP TABLE IF EXISTS `resources`;
CREATE TABLE `resources`  (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `resource_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `is_auth` int NOT NULL,
                              `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `created_at` datetime NULL DEFAULT NULL,
                              `updated_at` datetime NULL DEFAULT NULL,
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE INDEX `index_resource_path`(`resource_path` ASC) USING BTREE
);

-- ----------------------------
-- Table structure for role_binds
-- ----------------------------
DROP TABLE IF EXISTS `role_binds`;
CREATE TABLE `role_binds`  (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `role_id` int NOT NULL,
                               `resource_id` int NOT NULL,
                               `created_at` datetime NULL DEFAULT NULL,
                               `updated_at` datetime NULL DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               INDEX `ct_role_bind_resource_id`(`resource_id` ASC) USING BTREE,
                               INDEX `ct_role_bind_role_id`(`role_id` ASC) USING BTREE,
                               CONSTRAINT `role_bind_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                               CONSTRAINT `xrole_bind_resource_id` FOREIGN KEY (`resource_id`) REFERENCES `resources` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `access_level` int NOT NULL,
                          `user_menus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `created_at` datetime NULL DEFAULT NULL,
                          `updated_at` datetime NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `index_role_name`(`role_name` ASC) USING BTREE
);

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
                             `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `session_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `user_id` int NOT NULL,
                             `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `session_raw` varchar(1020) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                             `expires_at` datetime NULL DEFAULT NULL,
                             `created_at` datetime NULL DEFAULT NULL,
                             PRIMARY KEY (`session_id`) USING BTREE
);

-- ----------------------------
-- Table structure for user_binds
-- ----------------------------
DROP TABLE IF EXISTS `user_binds`;
CREATE TABLE `user_binds`  (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `role_id` int NOT NULL,
                               `user_id` int NOT NULL,
                               `created_at` datetime NULL DEFAULT NULL,
                               `updated_at` datetime NULL DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               INDEX `ct_user_bind_role_id`(`role_id` ASC) USING BTREE,
                               INDEX `ct_user_bind_user_id`(`user_id` ASC) USING BTREE,
                               CONSTRAINT `constraint_user_bind_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                               CONSTRAINT `constraint_user_bind_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

-- ----------------------------
-- Table structure for users
-- ----------------------------
CREATE TABLE `users`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `display_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `is_active` int NOT NULL,
                          `gitlab_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `dingtalk_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `wecom_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `created_at` datetime NULL DEFAULT NULL,
                          `updated_at` datetime NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `uq_index_username`(`username` ASC) USING BTREE,
                          UNIQUE INDEX `uq_index_gitlab_id`(`gitlab_id` ASC) USING BTREE,
                          UNIQUE INDEX `uq_index_dingtalk_id`(`dingtalk_id` ASC) USING BTREE,
                          UNIQUE INDEX `uq_index_wecom_id`(`wecom_id` ASC) USING BTREE

SET FOREIGN_KEY_CHECKS = 1;
