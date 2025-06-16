-- 创建用户表
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL, 
  `nick_name` varchar(255) NOT NULL,
  `gender` int unsigned DEFAULT '1',
  `avatar_url` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- 创建群组表
CREATE TABLE `group` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` smallint NOT NULL comment "1表示单聊, 2表示群聊",
  `status` smallint default 0 comment "1表示有效, 2表示无效(未同意), 3表示无效(拉黑)",
  `config` json comment "群聊配置", 
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_type` (`type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建群组成员表
CREATE TABLE `group_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `group_id` varchar(255) NOT NULL,
  `user_id` bigint NOT NULL,
  `alias_name` varchar(255) comment "用户对该群的备注名",
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_group_id` (`group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- 创建聊天消息表
CREATE TABLE `chat_msg` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `group_id` varchar(255) NOT NULL,
  `sender_id` bigint NOT NULL,
  `type` int DEFAULT 1 COMMENT "1文本, 2图片, 3视频, 4音频", 
  `content` varchar(2048) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `uuid` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid` (`uuid`),
  INDEX `idx_group_id` (`group_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 