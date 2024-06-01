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

 Date: 23/05/2024 16:52:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
                              PRIMARY KEY (`id`) USING BTREE
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
                               PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `access_level` int NOT NULL,
                          `user_menus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `created_at` datetime NULL DEFAULT NULL,
                          `updated_at` datetime NULL DEFAULT NULL,
                          PRIMARY KEY (`id`, `role_name`) USING BTREE,
                          UNIQUE INDEX `role_name`(`role_name` ASC) USING BTREE,
                          INDEX `id`(`id` ASC) USING BTREE
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
                               PRIMARY KEY (`id`) USING BTREE
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
                          PRIMARY KEY (`id`, `username`) USING BTREE,
                          UNIQUE INDEX `username_index`(`username` ASC) USING BTREE
);

DROP TABLE IF EXISTS `tokens`;
CREATE TABLE `tokens`  (
                           `token_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `token_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `expires_at` datetime NULL DEFAULT NULL,
                           `created_at` datetime NULL DEFAULT NULL,
                           `updated_at` datetime NULL DEFAULT NULL,
                           UNIQUE INDEX `token_key_type_index`(`token_key` ASC, `token_type` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
