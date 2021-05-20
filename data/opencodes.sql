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
