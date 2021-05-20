CREATE TABLE `o_order` (
   `id` INT ( 10 ) DEFAULT '0' COMMENT '主键ID',
   `user_id` INT ( 10 ) DEFAULT '0' COMMENT '用户ID',
   `good_name` VARCHAR ( 255 ) DEFAULT '' COMMENT '商品名称',
   `good_price` DECIMAL ( 10, 4 ) DEFAULT '0.0' COMMENT '价格',
   `good_counts` INT ( 10 ) DEFAULT '0' COMMENT '商品数量',
   `created_at` INT ( 10 ) DEFAULT '0' COMMENT '创建时间',
   `updated_at` INT ( 10 ) DEFAULT '0' COMMENT '更新时间',
   PRIMARY KEY `pk_id` ( `id` ),
   KEY `idx_user_id` ( `user_id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '用户订单表';

DROP TABLE `tb_sys_settings`;
CREATE  TABLE `tb_sys_settings` (
    `id` INT (10) AUTO_INCREMENT COMMENT '主键ID',
    `group` VARCHAR (32) NOT NULL COMMENT '配置分组',
    `key` VARCHAR (32) NOT NULL  COMMENT '设置键',
    `type` varchar (32) NOT NULL DEFAULT 'string' COMMENT 'string 类型 object',
    `content` text NOT NULL  COMMENT '设置内容',
    `created_at` int(10) NOT NULL COMMENT '创建时间',
    `updated_at` int(10) NOT NULL  COMMENT '更新时间',
    `deleted_at` int(10) NOT NULL  DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY `pk_id` (`id`),
    UNIQUE KEY  `uk_group_key` (`key`,`group`)
) ENGINE=INNODB DEFAULT CHARSET= utf8 COMMENT '系统配置表';
