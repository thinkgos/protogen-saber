-- Code generated by protoc-gen-saber-seaql. DO NOT EDIT.
-- versions:
--   - protoc-gen-saber-seaql v0.0.2
--   - protoc             v4.22.2
-- source: seaql.proto
-- 字典表
CREATE TABLE 
	`dict` (
		`id` bigint NOT NULL AUTO_INCREMENT COMMENT '系统序号',
		`key` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
		`name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
		`is_pin` tinyint(1) NOT NULL COMMENT '是否锁定',
		`created_at` datetime NOT NULL COMMENT '创建时间',
		`updated_at` datetime NOT NULL COMMENT '更新时间',
		PRIMARY KEY (`id`)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表';
