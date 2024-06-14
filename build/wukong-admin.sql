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

 Date: 15/06/2024 01:36:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mfa_app_bind_sessions
-- ----------------------------
DROP TABLE IF EXISTS `mfa_app_bind_sessions`;
CREATE TABLE `mfa_app_bind_sessions`  (
                                          `id` int NOT NULL AUTO_INCREMENT,
                                          `user_id` int NOT NULL,
                                          `totp_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                          `expire_at` datetime NOT NULL,
                                          `updated_at` datetime NULL DEFAULT NULL,
                                          `created_at` datetime NULL DEFAULT NULL,
                                          PRIMARY KEY (`id`) USING BTREE,
                                          UNIQUE INDEX `uq_index_totp_key`(`totp_secret` ASC) USING BTREE
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
-- Table structure for mfa_login_sessions
-- ----------------------------
DROP TABLE IF EXISTS `mfa_login_sessions`;
CREATE TABLE `mfa_login_sessions`  (
                                       `id` int NOT NULL AUTO_INCREMENT,
                                       `user_id` int NOT NULL,
                                       `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                       `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                       `mfa_provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                       `expire_at` datetime NOT NULL,
                                       `updated_at` datetime NULL DEFAULT NULL,
                                       `created_at` datetime NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`) USING BTREE
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
-- Table structure for oauth_codes
-- ----------------------------
DROP TABLE IF EXISTS `oauth_codes`;
CREATE TABLE `oauth_codes`  (
                                `id` int NOT NULL,
                                `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                `client_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                `client_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                `redirect_uri` varchar(1020) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                `scope` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                                `expires_at` datetime NOT NULL,
                                `created_at` datetime NULL DEFAULT NULL,
                                `updated_at` datetime NULL DEFAULT NULL,
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `uq_index_code_client_id`(`code` ASC, `client_id` ASC) USING BTREE
);

-- ----------------------------
-- Table structure for pass_key_login_sessions
-- ----------------------------
DROP TABLE IF EXISTS `pass_key_login_sessions`;
CREATE TABLE `pass_key_login_sessions`  (
                                            `id` int NOT NULL AUTO_INCREMENT,
                                            `user_id` int NOT NULL,
                                            `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                            `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                            `session_raw` varchar(1020) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                            `expire_at` datetime NOT NULL,
                                            `updated_at` datetime NULL DEFAULT NULL,
                                            `created_at` datetime NULL DEFAULT NULL,
                                            PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for pass_key_register_sessions
-- ----------------------------
DROP TABLE IF EXISTS `pass_key_register_sessions`;
CREATE TABLE `pass_key_register_sessions`  (
                                               `id` int NOT NULL AUTO_INCREMENT,
                                               `user_id` int NOT NULL,
                                               `session_raw` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                               `expire_at` datetime NOT NULL,
                                               `updated_at` datetime NULL DEFAULT NULL,
                                               `created_at` datetime NULL DEFAULT NULL,
                                               PRIMARY KEY (`id`) USING BTREE
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
                              `credential_raw` varchar(1020) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
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
-- Table structure for store_tokens
-- ----------------------------
DROP TABLE IF EXISTS `store_tokens`;
CREATE TABLE `store_tokens`  (
                                 `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `token_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `token_str` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `expires_at` datetime NULL DEFAULT NULL,
                                 `created_at` datetime NULL DEFAULT NULL,
                                 `updated_at` datetime NULL DEFAULT NULL,
                                 PRIMARY KEY (`uuid`) USING BTREE
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
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `display_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `is_active` int NOT NULL,
                          `created_at` datetime NULL DEFAULT NULL,
                          `updated_at` datetime NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `uq_index_username`(`username` ASC) USING BTREE
);

SET FOREIGN_KEY_CHECKS = 1;
