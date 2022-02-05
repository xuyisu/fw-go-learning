
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `activity_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '活动id',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动名称',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '活动状态',
  `main_image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动图片地址',
  `start_time` datetime(0) NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime(0) NULL DEFAULT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of activity
-- ----------------------------
INSERT INTO `activity` VALUES (1, '2020-11-29 08:47:54', '2020-11-29 08:47:54', 0, 0, 0, 1, '新手活动', 1, '', '2020-11-29 16:47:42', '2029-11-29 16:47:45');

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `user_id` bigint(11) NULL DEFAULT 0 COMMENT '用户id',
  `activity_id` bigint(11) NULL DEFAULT 0 COMMENT '活动id',
  `activity_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动名称',
  `product_id` bigint(11) NOT NULL DEFAULT 0 COMMENT '商品id',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_subtitle` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '商品简要描述',
  `product_main_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '商品图片地址',
  `quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '数量',
  `product_unit_price` decimal(20, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '单价',
  `selected` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否已选择 1是 0 否',
  `product_total_price` decimal(20, 2) NOT NULL DEFAULT 0.00 COMMENT '总价格',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '购物车' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES (4, '2020-12-06 11:01:02', '2021-11-27 12:37:26', 0, 0, 1, 0, 0, '新手活动', 1332985086563454976, '小米10 至尊纪念版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 8, 9999.00, 1, 79992.00);
INSERT INTO `cart` VALUES (5, '2020-12-09 07:58:33', '2021-11-27 12:37:26', 0, 0, 1, 0, 0, '新手活动', 1332985087444258816, '小米10', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 7, 9999.00, 1, 69993.00);
INSERT INTO `cart` VALUES (6, '2021-11-27 03:24:04', '2021-11-27 12:37:26', 0, 0, 1, 0, 0, '新手活动', 1332985089260392448, 'Redmi 10X 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (7, '2021-11-27 12:38:59', '2021-11-27 12:39:11', 0, 0, 1, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (8, '2021-11-27 12:39:02', '2021-11-27 12:39:11', 0, 0, 1, 0, 0, '新手活动', 1332985088656412672, 'Redmi K30 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (9, '2021-11-27 12:46:01', '2021-11-27 12:46:07', 0, 0, 1, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (10, '2021-11-27 12:47:22', '2021-11-27 12:47:33', 0, 0, 1, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (11, '2021-11-27 12:48:42', '2021-11-27 12:48:53', 0, 0, 1, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (12, '2021-11-27 14:44:09', '2021-11-27 17:56:11', 0, 0, 1, 0, 0, '新手活动', 1332985089461719040, 'Redmi 10X 4G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (13, '2021-11-29 21:33:15', '2021-11-29 21:39:29', 0, 0, 1, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (14, '2021-11-29 21:34:46', '2021-11-29 21:39:29', 0, 0, 1, 0, 0, '新手活动', 1332985089260392448, 'Redmi 10X 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (15, '2021-11-29 21:37:33', '2021-11-29 21:39:29', 0, 0, 1, 0, 0, '新手活动', 1332985089876955136, 'Redmi Note 9 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/970c6b287eb89620f5ee8e2b347f6f3d.png', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (16, '2021-11-29 21:38:56', '2021-11-29 21:39:29', 0, 0, 1, 0, 0, '新手活动', 1332985089461719040, 'Redmi 10X 4G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (17, '2021-11-29 21:39:12', '2021-11-29 21:39:29', 0, 0, 1, 0, 0, '新手活动', 1332985089067454464, 'Redmi 10X Pro', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/22419dee5ddfa9f4a91782231d07e189.png', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (18, '2021-11-29 22:03:21', '2021-11-29 22:05:05', 0, 0, 1, 0, 0, '新手活动', 1332985088656412672, 'Redmi K30 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (19, '2021-11-29 22:04:16', '2021-11-29 22:05:05', 0, 0, 1, 0, 0, '新手活动', 1332985089671434240, 'Redmi Note 9 Pro', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/76378ce289a36fcfa29f704ba90b4155.png', 2, 9999.00, 1, 19998.00);
INSERT INTO `cart` VALUES (20, '2021-11-29 22:04:39', '2021-11-29 22:05:05', 0, 0, 1, 0, 0, '新手活动', 1332985089260392448, 'Redmi 10X 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (21, '2021-11-29 22:09:31', '2021-11-29 22:09:31', 0, 0, 0, 0, 0, '新手活动', 1332985088450891776, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (22, '2021-11-29 22:09:35', '2021-11-29 22:09:35', 0, 0, 0, 0, 0, '新手活动', 1332985088656412672, 'Redmi K30 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 1, 9999.00, 1, 9999.00);
INSERT INTO `cart` VALUES (23, '2021-11-29 22:09:44', '2021-11-29 22:09:58', 0, 0, 0, 0, 1, '新手活动', 1332985087444258816, '小米10', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 12, 9999.00, 1, 119988.00);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父id',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '启用禁用状态 1启用 0禁用',
  `sort_order` int(11) NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '类目' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '手机 电话卡', 1, 100);
INSERT INTO `category` VALUES (2, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '电视 盒子', 1, 90);
INSERT INTO `category` VALUES (3, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '笔记本 显示器', 1, 80);
INSERT INTO `category` VALUES (4, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '家电 插线板', 1, 70);
INSERT INTO `category` VALUES (5, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '出行 穿戴', 1, 60);
INSERT INTO `category` VALUES (6, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '智能 路由器', 1, 50);
INSERT INTO `category` VALUES (7, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '电源 配件', 1, 40);
INSERT INTO `category` VALUES (8, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '健康 儿童', 1, 30);
INSERT INTO `category` VALUES (9, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '耳机 音响', 1, 20);
INSERT INTO `category` VALUES (10, '2020-11-29 08:43:45', '2020-11-29 08:43:45', 0, 0, 0, 0, '生活 箱包', 1, 10);

-- ----------------------------
-- Table structure for order_detail
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_detail_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单明细编号',
  `activity_id` bigint(20) NULL DEFAULT 0 COMMENT '活动id',
  `activity_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动名称',
  `activity_main_image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '活动图片地址',
  `product_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '商品id',
  `product_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_main_image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品图片地址',
  `current_unit_price` decimal(20, 2) NULL DEFAULT 0.00 COMMENT '单价',
  `quantity` int(11) NULL DEFAULT 0 COMMENT '数量',
  `total_price` decimal(20, 2) NULL DEFAULT 0.00 COMMENT '总价',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '购买人id',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '订单状态',
  `status_desc` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '状态描述',
  `cancel_time` datetime(0) NULL DEFAULT NULL COMMENT '取消时间',
  `cancel_reason` int(4) NULL DEFAULT 0 COMMENT '取消原因',
  `send_time` datetime(0) NULL DEFAULT NULL COMMENT '发货时间',
  `receive_time` datetime(0) NULL DEFAULT NULL COMMENT '签收时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单明细' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_detail
-- ----------------------------
INSERT INTO `order_detail` VALUES (1, '2020-12-09 08:36:10', '2020-12-09 09:59:33', 0, 0, 0, '1607502969', '1607502970050', 1, '', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 1, 9999.00, 0, 0, '已取消', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (2, '2020-12-09 08:37:23', '2020-12-09 08:37:23', 0, 0, 0, '1607503019', '1607503043156', 1, '', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (3, '2020-12-09 08:38:07', '2020-12-09 08:38:07', 0, 0, 0, '1607503086', '1607503086944', 1, '新手活动', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (4, '2020-12-09 08:38:40', '2020-12-09 08:38:40', 0, 0, 0, '1607503119', '1607503119753', 1, '新手活动', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (5, '2020-12-09 08:38:40', '2020-12-09 08:38:40', 0, 0, 0, '1607503119', '1607503119806', 1, '新手活动', '', 1332985087444258816, '小米10', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (6, '2020-12-09 08:41:31', '2020-12-09 08:41:31', 0, 0, 0, '1607503290', '1607503291059', 1, '新手活动', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (7, '2020-12-09 08:41:31', '2020-12-09 08:41:31', 0, 0, 0, '1607503290', '1607503291187', 1, '新手活动', '', 1332985087444258816, '小米10', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (8, '2021-11-27 12:37:26', '2021-11-27 14:40:12', 0, 0, 0, '1637987847', '1637987847247', 1, '新手活动', '', 1332985086563454976, '小米10 至尊纪念版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 9999.00, 8, 79992.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (9, '2021-11-27 12:37:26', '2021-11-27 14:40:12', 0, 0, 0, '1637987847', '1637987847270', 1, '新手活动', '', 1332985087444258816, '小米10', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 9999.00, 7, 69993.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (10, '2021-11-27 12:37:26', '2021-11-27 14:40:12', 0, 0, 0, '1637987847', '1637987847287', 1, '新手活动', '', 1332985089260392448, 'Redmi 10X 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (11, '2021-11-27 12:39:11', '2021-11-27 13:24:14', 0, 0, 0, '1637987952', '1637987952102', 1, '新手活动', '', 1332985088450891776, 'Redmi K30 Pro 变焦版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (12, '2021-11-27 12:39:11', '2021-11-27 13:24:14', 0, 0, 0, '1637987952', '1637987952124', 1, '新手活动', '', 1332985088656412672, 'Redmi K30 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (13, '2021-11-27 12:46:07', '2021-11-27 12:46:07', 0, 0, 0, '1637988368', '1637988368641', 1, '新手活动', '', 1332985088450891776, 'Redmi K30 Pro 变焦版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 9999.00, 1, 9999.00, 0, 10, '未付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (14, '2021-11-27 12:47:33', '2021-11-27 13:30:17', 0, 0, 0, '1637988454', '1637988454045', 1, '新手活动', '', 1332985088450891776, 'Redmi K30 Pro 变焦版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (15, '2021-11-27 12:48:53', '2021-11-27 13:13:13', 0, 0, 0, '1637988534', '1637988534186', 1, '新手活动', '', 1332985088450891776, 'Redmi K30 Pro 变焦版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (16, '2021-11-27 17:56:11', '2021-11-27 17:56:14', 0, 0, 0, '1638006971', '1638006971074', 1, '新手活动', '', 1332985089461719040, 'Redmi 10X 4G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (17, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', '1638193169836', 1, '新手活动', '', 1332985088450891776, 'Redmi K30 Pro 变焦版', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (18, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', '1638193169857', 1, '新手活动', '', 1332985089260392448, 'Redmi 10X 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (19, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', '1638193169870', 1, '新手活动', '', 1332985089876955136, 'Redmi Note 9 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/970c6b287eb89620f5ee8e2b347f6f3d.png', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (20, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', '1638193169886', 1, '新手活动', '', 1332985089461719040, 'Redmi 10X 4G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (21, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', '1638193169899', 1, '新手活动', '', 1332985089067454464, 'Redmi 10X Pro', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/22419dee5ddfa9f4a91782231d07e189.png', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (22, '2021-11-29 22:05:05', '2021-11-29 22:08:58', 0, 0, 0, '1638194705', '1638194705408', 1, '新手活动', '', 1332985088656412672, 'Redmi K30 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (23, '2021-11-29 22:05:05', '2021-11-29 22:08:58', 0, 0, 0, '1638194705', '1638194705420', 1, '新手活动', '', 1332985089671434240, 'Redmi Note 9 Pro', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/76378ce289a36fcfa29f704ba90b4155.png', 9999.00, 2, 19998.00, 0, 20, '已付款', NULL, 0, NULL, NULL);
INSERT INTO `order_detail` VALUES (24, '2021-11-29 22:05:05', '2021-11-29 22:08:58', 0, 0, 0, '1638194705', '1638194705432', 1, '新手活动', '', 1332985089260392448, 'Redmi 10X 5G', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 9999.00, 1, 9999.00, 0, 20, '已付款', NULL, 0, NULL, NULL);

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `payment` decimal(20, 2) NULL DEFAULT 0.00 COMMENT '支付金额',
  `payment_type` tinyint(4) NULL DEFAULT 0 COMMENT '支付类型',
  `payment_type_desc` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '支付类型描述',
  `postage` decimal(20, 2) NULL DEFAULT 0.00 COMMENT '邮费',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '订单状态',
  `status_desc` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '状态描述',
  `payment_time` datetime(0) NULL DEFAULT NULL COMMENT '支付时间',
  `address_id` bigint(20) NULL DEFAULT 0 COMMENT '地址id',
  `receive_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '收货人',
  `receive_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '联系号码',
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '省份',
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '城市',
  `area` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '区',
  `street` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '详细地址',
  `postal_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '邮编',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '购买人id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_info
-- ----------------------------
INSERT INTO `order_info` VALUES (2, '2020-12-09 08:36:10', '2020-12-09 09:59:33', 0, 0, 0, '1607502969', 9999.00, 1, '在线支付', 0.00, 0, '已取消', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (3, '2020-12-09 08:37:23', '2020-12-09 09:57:48', 0, 0, 0, '1607503019', 9999.00, 1, '在线支付', 0.00, 0, '已取消', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (4, '2020-12-09 08:38:07', '2020-12-09 09:57:49', 0, 0, 0, '1607503086', 9999.00, 1, '在线支付', 0.00, 0, '已取消', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (5, '2020-12-09 08:38:40', '2020-12-09 09:57:50', 0, 0, 0, '1607503119', 19998.00, 1, '在线支付', 0.00, 0, '已取消', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (6, '2020-12-09 08:41:31', '2020-12-09 09:57:51', 0, 0, 0, '1607503290', 19998.00, 1, '在线支付', 0.00, 0, '已取消', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (7, '2021-11-27 12:37:26', '2021-11-27 14:40:12', 0, 0, 0, '1637987847', 159984.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-27 14:40:13', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (8, '2021-11-27 12:39:11', '2021-11-27 13:24:13', 0, 0, 0, '1637987952', 19998.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-27 13:24:15', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (9, '2021-11-27 12:46:07', '2021-11-27 12:46:07', 0, 0, 0, '1637988368', 9999.00, 1, '在线支付', 0.00, 10, '未付款', NULL, 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (10, '2021-11-27 12:47:33', '2021-11-27 13:30:17', 0, 0, 0, '1637988454', 9999.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-27 13:30:18', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (11, '2021-11-27 12:48:53', '2021-11-27 13:13:13', 0, 0, 0, '1637988534', 9999.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-27 13:13:14', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (12, '2021-11-27 17:56:11', '2021-11-27 17:56:14', 0, 0, 0, '1638006971', 19998.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-27 17:56:14', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (13, '2021-11-29 21:39:29', '2021-11-29 21:39:33', 0, 0, 0, '1638193169', 79992.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-29 21:39:34', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);
INSERT INTO `order_info` VALUES (14, '2021-11-29 22:05:05', '2021-11-29 22:08:58', 0, 0, 0, '1638194705', 49995.00, 1, '在线支付', 0.00, 20, '已付款', '2021-11-29 22:08:58', 1332985086563454976, '1', '11111111111', '上海市', '上海市', '闵行区', '测试街道', '', 0);

-- ----------------------------
-- Table structure for order_status_record
-- ----------------------------
DROP TABLE IF EXISTS `order_status_record`;
CREATE TABLE `order_status_record`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_detail_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单明细编号',
  `product_id` int(11) NOT NULL DEFAULT 0 COMMENT '商品id',
  `product_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '商品名称',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '订单状态',
  `status_desc` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '状态描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `product_id` bigint(20) NULL DEFAULT 0 COMMENT '商品id',
  `category_id` bigint(20) NULL DEFAULT 0 COMMENT '品类id',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '商品名称',
  `sub_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '简要描述',
  `main_image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '商品图片地址',
  `sub_images` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '子图片列表',
  `activity_id` bigint(20) NULL DEFAULT NULL COMMENT '活动id',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '商品状态',
  `price` decimal(20, 2) NOT NULL DEFAULT 0.00 COMMENT '商品单价',
  `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (2, '2021-11-27 01:52:34', '2021-11-27 01:52:34', 0, 0, 0, 1332985086563454976, 100012, '小米10 至尊纪念版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/5ca871528d3420622f21f25be7aba58c.png', 1, 1, 9999.00, 100);
INSERT INTO `product` VALUES (3, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985087444258816, 100012, '小米10', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4a7a4f24061860901f724b6ba6d75fd9.png', 1, 1, 9999.00, 100);
INSERT INTO `product` VALUES (4, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985087637196800, 100012, '小米10 青春版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/cc563c76a7383d8030d1c23f31c60cb9.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/cc563c76a7383d8030d1c23f31c60cb9.png', 1, 1, 9999.00, 100);
INSERT INTO `product` VALUES (5, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985087838523392, 100012, 'Redmi K30S 至尊纪念版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/69a0df49cff71107977bbe3b48af15fa.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/69a0df49cff71107977bbe3b48af15fa.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (6, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985088039849984, 100012, 'Redmi K30 至尊纪念版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/0803dc7125ffa3447b04b1ae60f57a2b.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/0803dc7125ffa3447b04b1ae60f57a2b.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (7, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985088241176576, 100012, '小米云服务', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/8ade0163e9fdbd92a4b7d7ee3aeab905.jpg', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/8ade0163e9fdbd92a4b7d7ee3aeab905.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (8, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985088450891776, 100012, 'Redmi K30 Pro 变焦版', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/461bf013d08a7a91423cafcbc5ff9339.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (9, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985088656412672, 100012, 'Redmi K30 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7a98170e97ca5df8f5ad2470a8a2d01e.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (10, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985088870322176, 100012, 'Redmi K30 4G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b2b8b609aab05d9ad184bbe5c8e8be28.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b2b8b609aab05d9ad184bbe5c8e8be28.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (11, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985089067454464, 100012, 'Redmi 10X Pro', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/22419dee5ddfa9f4a91782231d07e189.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/22419dee5ddfa9f4a91782231d07e189.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (12, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985089260392448, 100012, 'Redmi 10X 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f774ef8332f2204a7c8fbf1b92f29e8a.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (13, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985089461719040, 100012, 'Redmi 10X 4G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/afa28e7477639537f556eb46e3ca5f43.jpeg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (14, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985089671434240, 100012, 'Redmi Note 9 Pro', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/76378ce289a36fcfa29f704ba90b4155.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/76378ce289a36fcfa29f704ba90b4155.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (15, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985089876955136, 100012, 'Redmi Note 9 5G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/970c6b287eb89620f5ee8e2b347f6f3d.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/970c6b287eb89620f5ee8e2b347f6f3d.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (16, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985090078281728, 100012, 'Redmi Note 9  4G', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/947721c8bc4a4ecc3bca17237ee64dea.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/947721c8bc4a4ecc3bca17237ee64dea.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (17, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985090267025408, 100012, 'Redmi Note 8 Pro', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/cba27e53367b74271a38a4515a692816.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/cba27e53367b74271a38a4515a692816.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (18, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985090476740608, 100012, 'Redmi Note 8', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b4534b9fd22f478314f3cbbf286b7338.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b4534b9fd22f478314f3cbbf286b7338.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (19, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985090665484288, 100012, '手机上门维修', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/84e19ece0125992449a6276e2ba333e9.png', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/84e19ece0125992449a6276e2ba333e9.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (20, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985090858422272, 100012, 'Redmi 9', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/0cadc8b00dbe3b5615bd6ab657715baf.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/0cadc8b00dbe3b5615bd6ab657715baf.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (21, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985091063943168, 100012, 'Redmi 9A', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1db88cd66ff1a6a3e75116988b7f3e12.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1db88cd66ff1a6a3e75116988b7f3e12.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (22, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985091265269760, 100012, 'Redmi 8', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/47057d897ee2c05c9215e059e1308dc6.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/47057d897ee2c05c9215e059e1308dc6.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (23, '2021-11-27 01:52:35', '2021-11-27 01:52:35', 0, 0, 0, 1332985091483373568, 100012, 'Redmi 8A', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/77bfd346ad97807237beca297cfe2fba.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/77bfd346ad97807237beca297cfe2fba.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (24, '2021-11-27 01:52:36', '2021-11-27 01:52:36', 0, 0, 0, 1332985091680505856, 100012, '腾讯黑鲨游戏手机', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f382e29367ad5852fc1adfdae1332d5c.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f382e29367ad5852fc1adfdae1332d5c.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (25, '2021-11-27 01:52:36', '2021-11-27 01:52:36', 0, 0, 0, 1332985091898609664, 100012, '中国电信', '3200万+4800万 前后双旗舰相机', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4d8753c35974699754af66388fc03a16.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4d8753c35974699754af66388fc03a16.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (26, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985385919320064, 1, '小米透明电视', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/37576cef378ef2805c20b9ab2a9d6d3d.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/37576cef378ef2805c20b9ab2a9d6d3d.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (27, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985386661711872, 1, '小米电视5 Pro 55英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/942d4f28d406f3946f0ce551c125c641.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/942d4f28d406f3946f0ce551c125c641.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (28, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985386858844160, 1, '小米电视5 65英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d8264eb9b900ead63eea69f6fc1e413a.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d8264eb9b900ead63eea69f6fc1e413a.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (29, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985387051782144, 1, 'Redmi 智能电视 A系列', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/93b55d27cd6e2100c25611eb828396ae.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/93b55d27cd6e2100c25611eb828396ae.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (30, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985387253108736, 1, '量子点电视', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da1680fb62238cb5b01998d80fc34e25.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da1680fb62238cb5b01998d80fc34e25.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (31, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985387433463808, 1, 'Redmi 智能电视 X系列', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/070b329c18c1ab167df2a42564a67813.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/070b329c18c1ab167df2a42564a67813.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (32, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985387626401792, 1, '小米电视大师 82英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b8f12c8d20878e88b66047ef32566e48.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/b8f12c8d20878e88b66047ef32566e48.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (33, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985387819339776, 1, '小米电视5 Pro 65英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d096389e19ba6b79b324b7d0acfab7ca.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d096389e19ba6b79b324b7d0acfab7ca.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (34, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985388033249280, 1, '小米电视5 55英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/23316088a9182a08c624958319924d8d.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/23316088a9182a08c624958319924d8d.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (35, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985388234575872, 1, '小米全面屏电视E43K', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/319700f48d268f522664e1fc876b0044.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/319700f48d268f522664e1fc876b0044.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (36, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985388440096768, 1, '金属全面屏电视', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d219d8b9c4a74eda4ae6c0222d194cd2.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d219d8b9c4a74eda4ae6c0222d194cd2.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (37, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985388620451840, 1, '小米盒子', '', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/554a4f8ea0c2c3ed19d843bea7321203.jpg', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/554a4f8ea0c2c3ed19d843bea7321203.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (38, '2020-12-03 11:51:07', '2020-12-03 11:51:07', 0, 0, 0, 1332985388825972736, 1, '大师电视 65英寸 OLED', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7da49e7ada2229c963c9c7353a58d49b.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7da49e7ada2229c963c9c7353a58d49b.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (39, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389018910720, 1, '小米电视5 Pro 75英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da1680fb62238cb5b01998d80fc34e25.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da1680fb62238cb5b01998d80fc34e25.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (40, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389207654400, 1, '小米电视4A 60英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1df560d16cc983b0f34d24fc484bd697.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1df560d16cc983b0f34d24fc484bd697.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (41, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389392203776, 1, 'Redmi 智能电视 X55', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/070b329c18c1ab167df2a42564a67813.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/070b329c18c1ab167df2a42564a67813.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (42, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389610307584, 1, '小米全面屏电视', '', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/8aeae36f5d088f22c84b8953bd0d08b4.jpg', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/8aeae36f5d088f22c84b8953bd0d08b4.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (43, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389811634176, 1, '电视音箱', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/381acfeddeebcce4720ea7b8e481252f.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/381acfeddeebcce4720ea7b8e481252f.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (44, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985389996183552, 1, 'Redmi 智能电视 MAX 98\"', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9d649ec3d7d14da7090e396d56b7cc92.jpg', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9d649ec3d7d14da7090e396d56b7cc92.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (45, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985390176538624, 1, '小米电视5 75英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d707cdac2a19703b57e65212f32fd25b.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d707cdac2a19703b57e65212f32fd25b.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (46, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985390373670912, 1, '小米电视4A 70英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/6797917392e912577135d9eb8ad92f1f.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/6797917392e912577135d9eb8ad92f1f.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (47, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985390562414592, 1, '小米电视4S 75英寸', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/8cb5791161d8b25897f60385f0ec7ab2.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/8cb5791161d8b25897f60385f0ec7ab2.png', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (48, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985390784712704, 1, '小米壁画电视', '', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/5732e8f317eb86bfcb174fbb49d00c4a.jpg', 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-mimall-media/5732e8f317eb86bfcb174fbb49d00c4a.jpg', NULL, 1, 9999.00, 100);
INSERT INTO `product` VALUES (49, '2020-12-03 11:51:08', '2020-12-03 11:51:08', 0, 0, 0, 1332985391027982336, 1, '激光投影电视', '', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4fa264c15d3829e6de4283a552651e22.png', 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4fa264c15d3829e6de4283a552651e22.png', NULL, 1, 9999.00, 100);

-- ----------------------------
-- Table structure for product_detail
-- ----------------------------
DROP TABLE IF EXISTS `product_detail`;
CREATE TABLE `product_detail`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `product_id` bigint(20) NOT NULL COMMENT '商品id',
  `detail` json NULL COMMENT '商品详情',
  `param` json NULL COMMENT '商品参数',
  `rotation` json NULL COMMENT '轮播图片',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品明细' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '启用标志',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `union_phone_idx`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2021-11-27 17:52:14', '2021-11-27 17:52:14', 0, 0, 0, 1, 'xuyisu', '123@qq.com', '11111111111', 'e10adc3949ba59abbe56e057f20f883e');

-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS `user_address`;
CREATE TABLE `user_address`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `address_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '地址id',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新人',
  `delete_flag` tinyint(4) NOT NULL DEFAULT 0 COMMENT '删除标志',
  `default_flag` tinyint(4) NULL DEFAULT 0 COMMENT '默认标志',
  `receive_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人',
  `receive_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系号码',
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `province_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份编码',
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `city_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市编码',
  `area` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '区',
  `area_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '区编码',
  `street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '详细地址',
  `postal_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '邮编',
  `address_label` tinyint(4) NULL DEFAULT 0 COMMENT '地址标签',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户地址' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_address
-- ----------------------------
INSERT INTO `user_address` VALUES (1, 1332985086563454976, '2020-12-03 12:32:07', '2020-12-03 12:32:07', 0, 0, 0, 0, '1', '11111111111', '上海市', '111', '上海市', '1111', '闵行区', '1111', '测试街道', '', 1);

SET FOREIGN_KEY_CHECKS = 1;
