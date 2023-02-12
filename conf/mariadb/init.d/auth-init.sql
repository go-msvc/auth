CREATE DATABASE IF NOT EXISTS `auth`;
GRANT ALL PRIVILEGES ON `auth`.* to 'auth'@'%' IDENTIFIED BY 'auth';

DROP TABLE IF EXISTS `user_roles`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `role_permissions`;
DROP TABLE IF EXISTS `roles`;
DROP TABLE IF EXISTS `permissions`;
DROP TABLE IF EXISTS `accounts`;

CREATE TABLE `accounts` (
  `id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `active` boolean DEFAULT true,
  `name` VARCHAR(100) NOT NULL,
  UNIQUE KEY `account_id` (`id`),
  UNIQUE KEY `account_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `permissions` (
  `account_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  UNIQUE KEY `permission_id` (`id`),
  UNIQUE KEY `permission_name` (`account_id`,`name`),
  FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `roles` (
  `account_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  UNIQUE KEY `role_id` (`id`),
  UNIQUE KEY `role_name` (`account_id`,`name`),
  FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `role_permissions` (
  `role_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `permission_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`),
  FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `users` (
  `account_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  UNIQUE KEY `user_id` (`id`),
  UNIQUE KEY `user_name` (`account_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `user_roles` (
  `user_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  `role_id` VARCHAR(40) DEFAULT (uuid()) NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

