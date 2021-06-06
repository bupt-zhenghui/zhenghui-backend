CREATE TABLE `user_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
  `age` int DEFAULT NULL COMMENT '年龄',
  `gender` tinyint DEFAULT NULL COMMENT '性别',
  `tel-number` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '联系电话',
  `email` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '电子邮箱',
  `company` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所在单位',
  `university` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '毕业院校',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;