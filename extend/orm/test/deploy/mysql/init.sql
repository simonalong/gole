CREATE TABLE `iot_product_model` (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                     `model_name` varchar(64) NOT NULL DEFAULT '' COMMENT '应用名',
                                     `model_key` varchar(32) NOT NULL DEFAULT '' COMMENT '模型key',
                                     `node_type` varchar(32) NOT NULL DEFAULT '' COMMENT '节点类型',
                                     `create_user` varchar(12) NOT NULL DEFAULT '' COMMENT '创建者',
                                     `update_user` varchar(12) NOT NULL DEFAULT '' COMMENT '更新者',
                                     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     KEY `k_model` (`model_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品模型表';


CREATE TABLE `iot_device` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `device_name` varchar(64) NOT NULL DEFAULT '' COMMENT '设备名',
                              `model_id` varchar(32) NOT NULL DEFAULT '' COMMENT '所属的产品模型id',
                              `model_key` varchar(32) NOT NULL DEFAULT '' COMMENT '所属的产品模型key',
                              `node_type` varchar(32) NOT NULL DEFAULT '' COMMENT '节点类型',
                              `online_status` tinyint NOT NULL DEFAULT 0 COMMENT '在线状态：0-离线；1-在线',
                              `create_user` varchar(12) NOT NULL DEFAULT '' COMMENT '创建者',
                              `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              KEY `k_model_id` (`model_id`),
                              KEY `k_model_key` (`model_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备表';
