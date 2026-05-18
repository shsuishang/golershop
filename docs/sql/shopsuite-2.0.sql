/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : modulith_open

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 05/05/2026 17:30:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account_user_analytics
-- ----------------------------
DROP TABLE IF EXISTS `account_user_analytics`;
CREATE TABLE `account_user_analytics`  (
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_spend` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '消费总额',
  `user_refund` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '退款总额',
  `user_order_buy_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效订单数量',
  `user_order_return_num` int(10) UNSIGNED NOT NULL COMMENT '退单数量',
  `user_product_buy_num` int(10) UNSIGNED NOT NULL COMMENT '购买数量',
  `user_product_return_num` int(10) UNSIGNED NOT NULL COMMENT '退货数量',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_base
-- ----------------------------
DROP TABLE IF EXISTS `account_user_base`;
CREATE TABLE `account_user_base`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_password` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `user_salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'salt值',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `user_account`(`user_account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户基本信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_user_base
-- ----------------------------
INSERT INTO `account_user_base` VALUES (10001, 'admin', 'ff64b03e55429443e29a3fff3abbe00b', '77c96449b196487d92214bd08121ac62', 1001);

-- ----------------------------
-- Table structure for account_user_bind_connect
-- ----------------------------
DROP TABLE IF EXISTS `account_user_bind_connect`;
CREATE TABLE `account_user_bind_connect`  (
  `bind_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '绑定标记',
  `bind_type` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '绑定类型(EMUN):1-mobile;  2-email;   13-weixin公众号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `bind_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
  `bind_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `bind_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户图标',
  `bind_openid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问编号',
  `bind_unionid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'unionid',
  `bind_active` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否激活(BOOL):0-未激活;1-激活',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`bind_id`) USING BTREE,
  UNIQUE INDEX `bind_type`(`bind_type`, `bind_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户绑定表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_block
-- ----------------------------
DROP TABLE IF EXISTS `account_user_block`;
CREATE TABLE `account_user_block`  (
  `user_black_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '黑名单编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `black_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拉黑用户编号',
  `user_black_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拉黑时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_black_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `black_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户黑名单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_card
-- ----------------------------
DROP TABLE IF EXISTS `account_user_card`;
CREATE TABLE `account_user_card`  (
  `card_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会员卡编号',
  `card_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '会员卡名称',
  `user_level_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员等级',
  `card_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '状态(BOOL):0-禁用;1-开启',
  `card_pirce` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '会费（元/期）',
  `card_market_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '原会费（元/期）',
  `card_save_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '预计年省金额',
  `card_year` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期年',
  `card_month` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期月',
  `card_day` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期日',
  `card_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '上线时间',
  `card_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `user_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '会员类型(ENUM):1-普通用户;2-扩展用户',
  `transport_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流模板编号',
  `card_bind_addr` bit(1) NOT NULL DEFAULT b'0' COMMENT '唯一地址(BOOL):0-否;1-是',
  `card_is_exp` bit(1) NOT NULL DEFAULT b'0' COMMENT '体验卡(BOOL):0-否;1-是',
  `card_exp_month` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '体验月',
  `card_exp_day` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '体验日',
  `card_exp_fee` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '体验费用',
  `card_exp_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '体验状态(BOOL):0-禁用;1-开启',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `activity_id_discount` int(11) NOT NULL DEFAULT 0 COMMENT '活动会员专享折扣编号',
  `activity_id_gitfbag` int(11) NOT NULL DEFAULT 0 COMMENT '活动会员开卡赠品编号',
  `activity_id_member_day` int(11) NOT NULL DEFAULT 0 COMMENT '活动会员日（实物）编号',
  `activity_id_birthday` int(11) NOT NULL DEFAULT 0 COMMENT '活动生日专享编号',
  `user_fx_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '允许分销(BOOL):1-启用分销;0-禁用分销',
  `user_fx_rate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '分销佣金比例',
  `card_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享标题',
  `card_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享描述',
  `card_keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '搜索关键字',
  `card_img` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分享图标',
  `card_fx_active` bit(1) NOT NULL DEFAULT b'0' COMMENT '开通推广员(BOOL):1-开通;0-不开通',
  `card_img_bg` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '背景图片',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`card_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户会员卡' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_channel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_channel`;
CREATE TABLE `account_user_channel`  (
  `user_channel_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '渠道编号',
  `user_channel_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道名称',
  `user_channel_prefix` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道前缀',
  `user_channel_enable` bit(11) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `user_channel_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否内置(BOOL):0-非内置;1-内置',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_channel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户渠道表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_channel_code
-- ----------------------------
DROP TABLE IF EXISTS `account_user_channel_code`;
CREATE TABLE `account_user_channel_code`  (
  `ucc_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '渠道编号',
  `ucc_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道编码',
  `ucc_path` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面地址',
  `ucc_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `user_channel_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '渠道编号',
  `ucc_used` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否使用(BOOL):0-未使用;1-已使用',
  `ucc_use_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '使用时间',
  `ucc_create_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `ucc_is_unique` bit(1) NOT NULL DEFAULT b'0' COMMENT '唯一使用(BOOL):0-不限制;1-可用一次',
  `ucc_use_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用次数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ucc_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户渠道邀请码表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_delivery_address
-- ----------------------------
DROP TABLE IF EXISTS `account_user_delivery_address`;
CREATE TABLE `account_user_delivery_address`  (
  `ud_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地址编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `ud_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `ud_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `ud_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `ud_telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `ud_province_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省编号',
  `ud_province` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `ud_city_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市编号',
  `ud_city` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市',
  `ud_county_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '县',
  `ud_county` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
  `ud_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `ud_postalcode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  `ud_tag_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址标签(ENUM):1001-家里;1002-公司',
  `ud_longitude` double NOT NULL DEFAULT 0 COMMENT '经度',
  `ud_latitude` double NOT NULL DEFAULT 0 COMMENT '纬读',
  `ud_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `ud_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否默认(BOOL):0-非默认;1-默认',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ud_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `ud_is_default`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户地址表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_distribution
-- ----------------------------
DROP TABLE IF EXISTS `account_user_distribution`;
CREATE TABLE `account_user_distribution`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级用户编号',
  `user_partner_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属城市合伙人',
  `user_team_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '团队人数',
  `user_province_team_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属省公司',
  `user_city_team_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属市公司',
  `user_county_team_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属区公司',
  `role_level_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '角色等级',
  `ucc_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '渠道编号',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动编号',
  `user_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `user_fans_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数量:冗余',
  `user_is_sp` bit(1) NOT NULL DEFAULT b'0' COMMENT '服务商(BOOL):0-否;1-是;',
  `user_is_da` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '区代理(BOOL):0-否;1-是为区Id;',
  `user_is_ca` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市代理(BOOL):0-否;1-是为市Id;',
  `user_is_pa` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省代理(BOOL):0-否;1-是为省Id;',
  `user_is_pt` bit(1) NOT NULL DEFAULT b'0' COMMENT '城市合伙人(BOOL):0-否;1-是;',
  `user_active` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否有效(ENUM):0-未生效;1-有效;2-审核失败',
  `user_active_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '有效审核备注',
  `user_voucher_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分销优惠券',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `user_time`(`user_time`) USING BTREE,
  INDEX `user_parent_id`(`user_parent_id`, `user_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '粉丝来源关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_friend
-- ----------------------------
DROP TABLE IF EXISTS `account_user_friend`;
CREATE TABLE `account_user_friend`  (
  `user_friend_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `friend_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友编号',
  `friend_note` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注名称',
  `user_friend_addtime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `friend_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '关注状态(ENUM):1-单向关注;2-双向关注',
  `friend_invite` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '邀请状态(ENUM):0-新邀请;2-处理完成后邀请',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_friend_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `friend_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户好友关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_group
-- ----------------------------
DROP TABLE IF EXISTS `account_user_group`;
CREATE TABLE `account_user_group`  (
  `group_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组名称',
  `group_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组类型(ENUM):0-临时组上限100人;  1-普通组上限300人; 2-VIP组 上限500人',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '管理员',
  `group_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '数量',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`group_id`) USING BTREE,
  INDEX `group_name`(`group_name`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友管理组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_group_rel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_group_rel`;
CREATE TABLE `account_user_group_rel`  (
  `group_rel_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `group_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '组名称',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`group_rel_id`) USING BTREE,
  UNIQUE INDEX `group_name`(`group_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友组关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_history
-- ----------------------------
DROP TABLE IF EXISTS `account_user_history`;
CREATE TABLE `account_user_history`  (
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_history_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '历史密码',
  `user_history_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '历史 IP',
  `user_history_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户密码历史信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_info
-- ----------------------------
DROP TABLE IF EXISTS `account_user_info`;
CREATE TABLE `account_user_info`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `user_state` mediumint(8) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态(ENUM):0-锁定;1-已激活;2-未激活;',
  `user_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码(mobile)',
  `user_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `user_gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '性别(ENUM):0-保密;1-男;  2-女;',
  `user_birthday` date NOT NULL DEFAULT '2000-01-01' COMMENT '生日(DATE)',
  `user_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱(email)',
  `user_level_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级编号',
  `user_realname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `user_idcard` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '身份证',
  `user_idcard_images` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '身份证图片(DTO)',
  `user_is_authentication` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '认证状态(ENUM):0-未认证;1-待审核;2-认证通过;3-认证失败',
  `tag_ids` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户标签(DOT)',
  `industry_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '行业编号',
  `user_from` smallint(5) UNSIGNED NOT NULL DEFAULT 2310 COMMENT '用户来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序;2315-公众号',
  `user_new` bit(1) NOT NULL DEFAULT b'1' COMMENT '新人标识(BOOL):0-不是;1-是',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `user_is_sale` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否销售员(BOOL):0-不是;1-是',
  `user_sale_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '销售员id',
  `user_wx_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信二维码',
  `user_customer_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加客户时间',
  `user_is_plus` bit(1) NOT NULL DEFAULT b'0' COMMENT 'PLUS会员(BOOL):0-不是;1-是',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_user_plus`(`user_is_plus`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户详细信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_user_info
-- ----------------------------
INSERT INTO `account_user_info` VALUES (10001, 'admin', '系统管理员', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231031/8aa85cf4119746f2be5854c8cb081ea5.png', 1, '13488888888', '+86', 2, '2023-10-26', 'suishang@2018.com', 1001, '随商', '341222199310239898', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231016/2cca497ab8344d588e353bdafe531992.png,https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231016/6a3e27ae488b4d408e4d454957755126.png', 2, '', 0, 2310, b'1', 1001, b'0', 0, '', 0, b'0');

-- ----------------------------
-- Table structure for account_user_invoice
-- ----------------------------
DROP TABLE IF EXISTS `account_user_invoice`;
CREATE TABLE `account_user_invoice`  (
  `user_invoice_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '发票编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `invoice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票抬头',
  `invoice_company_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '纳税人识别号',
  `invoice_content` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票内容',
  `invoice_is_company` bit(1) NOT NULL DEFAULT b'1' COMMENT '公司开票(BOOL):0-个人;1-公司',
  `invoice_is_electronic` bit(1) NOT NULL DEFAULT b'1' COMMENT '电子发票(ENUM):0-纸质发票;1-电子发票',
  `invoice_type` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发票类型(ENUM):1-普通发票;2-增值税专用发票',
  `invoice_datetime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `invoice_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位地址',
  `invoice_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位电话',
  `invoice_bankname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '开户银行',
  `invoice_bankaccount` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '银行账号',
  `invoice_contact_mobile` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人手机',
  `invoice_contact_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人邮箱',
  `invoice_is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否默认',
  `invoice_contact_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人',
  `invoice_contact_area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人地区',
  `invoice_contact_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票详细地址',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_invoice_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户发票管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_level
-- ----------------------------
DROP TABLE IF EXISTS `account_user_level`;
CREATE TABLE `account_user_level`  (
  `user_level_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '等级编号',
  `user_level_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '等级名称',
  `user_level_exp` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '升级经验值',
  `user_level_spend` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '累计消费',
  `user_level_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '等级图标',
  `user_level_rate` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '折扣率百分比',
  `user_level_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `user_level_is_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):0-否;1-是',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_level_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1004 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户等级表-平台' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_user_level
-- ----------------------------
INSERT INTO `account_user_level` VALUES (1001, 'vip1', 100, 0, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/a20b208b4da04b13af76d8f9703054ca.png', 100, 20230829143055, b'1', 1001);
INSERT INTO `account_user_level` VALUES (1002, 'vip2', 0, 2, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/ddd36d38be35474691ef1dedf807dcf0.png', 95, 20230829143058, b'1', 1001);
INSERT INTO `account_user_level` VALUES (1003, 'vip3', 0, 3, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230914/6bb643e537854d1ab9bd53e6f801183a.png', 90, 20230829140104, b'1', 1001);

-- ----------------------------
-- Table structure for account_user_level_log
-- ----------------------------
DROP TABLE IF EXISTS `account_user_level_log`;
CREATE TABLE `account_user_level_log`  (
  `user_level_log_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_level_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前等级ID',
  `user_level_pre_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '之前等级ID',
  `operate_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作备注',
  `operate_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人',
  `create_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_level_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户等级记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_login
-- ----------------------------
DROP TABLE IF EXISTS `account_user_login`;
CREATE TABLE `account_user_login`  (
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_active_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '激活时间',
  `user_reg_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '注册IP',
  `user_reg_date` date NOT NULL COMMENT '注册日期',
  `user_reg_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `user_count_login` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录次数',
  `user_lastlogin_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '上次登录IP',
  `user_lastlogin_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上次登录时间',
  `user_clientid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'APPCID(DOT)',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `user_reg_time`(`user_reg_time`) USING BTREE,
  INDEX `user_reg_date`(`user_reg_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户登录信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `account_user_login_history`;
CREATE TABLE `account_user_login_history`  (
  `user_login_history_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `user_login_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录时间',
  `user_login_state` bit(1) NOT NULL DEFAULT b'1' COMMENT '登录状态(BOOL):0-登录失败;1-登录成功',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_login_history_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `user_login_ip`(`user_login_ip`, `user_login_state`, `user_login_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户登录信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_message
-- ----------------------------
DROP TABLE IF EXISTS `account_user_message`;
CREATE TABLE `account_user_message`  (
  `message_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息编号',
  `message_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `message_kind` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息种类(ENUM):1-发送消息;2-接收消息',
  `user_other_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '相关用户:发送者或者接收者',
  `user_other_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '相关昵称:发送者或者接收者',
  `message_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息标题',
  `message_content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `message_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送时间',
  `message_is_read` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否读取(BOOL):0-未读;1-已读',
  `message_is_delete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除(BOOL):0-正常状态;1-删除状态',
  `message_type` smallint(5) UNSIGNED NOT NULL DEFAULT 2 COMMENT '消息类型(ENUM):1-系统消息;2-用户消息',
  `message_cat` enum('text','img','video','file','location','redpack') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包',
  `message_data_type` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息分类(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息',
  `message_data_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息数据:商品编号|订单编号',
  `message_length` mediumint(9) NOT NULL DEFAULT 0 COMMENT '消息长度',
  `message_w` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片宽度',
  `message_h` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片高度',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`message_id`) USING BTREE,
  INDEX `from_member_id`(`user_other_id`) USING BTREE,
  INDEX `to_member_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '短消息-聊天记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_sns
-- ----------------------------
DROP TABLE IF EXISTS `account_user_sns`;
CREATE TABLE `account_user_sns`  (
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_blog` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '微博数量',
  `user_friend` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友数量',
  `user_fans` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数量',
  `user_growth` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '成长值',
  `user_report` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否可以举报(BOOL):0-不可以;1-可以',
  `user_buy` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否可以购买商品(BOOL):0-不可以;1-可以',
  `user_comment` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否允许发表言论(BOOL):0-不可以;1-可以',
  `user_fans_store` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '推广店铺数量',
  `user_story` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子数量',
  `user_poster` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '背景图片',
  `user_sign` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户签名',
  `user_story_comment` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数量',
  `user_favorites_store` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏店铺',
  `user_favorites_item` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏商品',
  `user_favorites_brand` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏品牌',
  `user_story_collection` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏帖子',
  `user_story_like` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子点赞',
  `user_story_forward` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子转发',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户SNS信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_tag_base
-- ----------------------------
DROP TABLE IF EXISTS `account_user_tag_base`;
CREATE TABLE `account_user_tag_base`  (
  `tag_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签编码',
  `tag_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签标题',
  `tag_group_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类',
  `tag_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 255 COMMENT '配置排序:从小到大',
  `tag_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `tag_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`tag_id`) USING BTREE,
  INDEX `tag_title`(`tag_title`) USING BTREE,
  INDEX `tag_group_id`(`tag_group_id`, `tag_sort`, `tag_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_user_tag_base
-- ----------------------------
INSERT INTO `account_user_tag_base` VALUES (31, '新客户', 22, 10, b'0', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (32, '重点客户', 22, 20, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (33, '内部员工', 22, 50, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (34, '待召回', 22, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (35, '华东', 23, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (36, '华中', 23, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (37, '华南', 23, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (38, '华北', 23, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (39, '电子数码', 24, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (40, '美妆达人', 24, 255, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (42, '大客户', 22, 1, b'1', b'0', 1001);
INSERT INTO `account_user_tag_base` VALUES (44, '不活跃用户', 22, 0, b'1', b'0', 1001);

-- ----------------------------
-- Table structure for account_user_tag_group
-- ----------------------------
DROP TABLE IF EXISTS `account_user_tag_group`;
CREATE TABLE `account_user_tag_group`  (
  `tag_group_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分组编号',
  `tag_group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组名称',
  `tag_group_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 255 COMMENT '分组排序:从小到大',
  `tag_group_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):0-禁用;1-启用',
  `tag_group_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_time` timestamp(0) NOT NULL COMMENT '添加时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`tag_group_id`) USING BTREE,
  INDEX `index_name`(`tag_group_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签分组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_user_tag_group
-- ----------------------------
INSERT INTO `account_user_tag_group` VALUES (22, '客户级别', 10, b'1', b'0', '2023-09-14 08:22:05', '0000-00-00 00:00:00', 1001);
INSERT INTO `account_user_tag_group` VALUES (23, '客户地区', 20, b'1', b'0', '2023-09-14 08:24:12', '0000-00-00 00:00:00', 1001);
INSERT INTO `account_user_tag_group` VALUES (24, '购买类型', 30, b'1', b'0', '2023-09-14 08:25:49', '0000-00-00 00:00:00', 1001);

-- ----------------------------
-- Table structure for account_user_zone
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone`;
CREATE TABLE `account_user_zone`  (
  `zone_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '好友组编号',
  `zone_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '群组名称',
  `zone_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组类型(ENUM):0-临时组上限100人;  1-普通组上限300人; 2-VIP组 上限500人',
  `zone_permission` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请加入模式(ENUM): 0-默认直接加入; 1-需要身份验证; 2-私有群组',
  `zone_declared` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '群组公告',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员',
  `zone_bind_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '第三方群组编号',
  `zone_user_num` int(10) UNSIGNED NOT NULL DEFAULT 2 COMMENT '人数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`zone_id`) USING BTREE,
  UNIQUE INDEX `group_name`(`zone_name`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_zone_message
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone_message`;
CREATE TABLE `account_user_zone_message`  (
  `zone_message_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '短消息索引编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发送用户',
  `zone_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群组编号',
  `zone_message_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '短消息内容',
  `zone_message_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '短消息发送时间',
  `zone_message_is_delete` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '短消息状态(BOOL):0-正常状态;1-删除状态',
  `zone_message_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 2 COMMENT '消息类型(ENUM):1-系统消息;2-用户消息;3-私信',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`zone_message_id`) USING BTREE,
  INDEX `from_member_id`(`zone_id`) USING BTREE,
  INDEX `to_member_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组消息-聊天记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_user_zone_rel
-- ----------------------------
DROP TABLE IF EXISTS `account_user_zone_rel`;
CREATE TABLE `account_user_zone_rel`  (
  `zone_rel_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关系编号',
  `zone_id` int(10) UNSIGNED NOT NULL COMMENT '群组编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户',
  `zone_rel_permission` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请加入模式(ENUM): 0-加入; 1-待验证;',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`zone_rel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '群组用户关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_menu_base
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu_base`;
CREATE TABLE `admin_menu_base`  (
  `menu_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单编号',
  `menu_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单父编号',
  `menu_title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面网址',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件名称',
  `menu_path` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路由',
  `menu_component` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'Layout' COMMENT '组件路径',
  `menu_redirect` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '重定向',
  `menu_close` bit(1) NOT NULL DEFAULT b'1' COMMENT '允许关闭(BOOL):0-禁止;1-允许',
  `menu_hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否隐藏(BOOL):0-展示;1-隐藏',
  `menu_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `menu_class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '样式class',
  `menu_icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标设置',
  `menu_dot` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否红点(BOOL):0-隐藏;1-显示',
  `menu_bubble` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单标签',
  `menu_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '菜单排序',
  `menu_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '菜单类型(LIST):0-按钮;1-菜单',
  `menu_note` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_func` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '功能开启:设置config_key',
  `menu_role` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色类型(LIST):1-平台;2-商户;3-门店',
  `menu_param` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'url参数',
  `menu_permission` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识:后端地址',
  `menu_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `menu_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后更新时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `menu_path`(`menu_path`, `menu_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5355 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_menu_base
-- ----------------------------
INSERT INTO `admin_menu_base` VALUES (1, 0, '运营统计', '', 'Root', '/', 'Layout', '', b'1', b'0', b'1', '', 'home-3-line', b'0', '12', 1, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (2, 0, '系统管理', '', 'SysSetting', '/sys', 'Layout', '', b'1', b'0', b'1', '', 'settings-line', b'1', '', 20, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (101, 1, '运营首页', '', 'Index', '/index', '@/views/index/index', '', b'1', b'0', b'1', '', 'dashboard-line', b'1', '', 1, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (310, 0, '商城设置', '', 'StoreSetting', '/settings', 'Layout', '', b'1', b'0', b'1', '', 'settings-4-line', b'0', '', 19, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (360, 310, '站点设置', '', 'Site', '/site/1001', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'settings-4-line', b'0', '', 1, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4013, 2, '字典管理', '', 'Dict', '/dict', '@/views/sys/dict/index', '', b'1', b'0', b'1', '', 'article-line', b'0', '', 2, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4023, 2, '素材窗口', '', 'MaterialGalleryDialog', '/md', '@/views/sys/gallery/dialog', '', b'1', b'1', b'1', '', 'gallery-line', b'0', '', 249, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4029, 0, '用户管理', '', 'MenberManager', '/member', 'Layout', '', b'1', b'0', b'1', '', 'account-circle-line', b'1', '', 6, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4030, 4029, '用户管理', '', 'UserInfo', '/userInfo', '@/views/account/userInfo/index', '', b'1', b'0', b'1', '', 'user-3-line', b'0', '', 0, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4031, 4029, '会员等级', '', 'UserLevel', '/userLevel', '@/views/account/userLevel/index', '', b'1', b'0', b'1', '', 'vip-diamond-line', b'0', '', 0, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4033, 4052, '积分记录', '', 'userPointsHistory', '/userPointsHistory', '@/views/pay/userPointsHistory/index', '', b'1', b'0', b'1', '', 'gift-2-line', b'0', '', 0, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4039, 0, '订单管理', '', 'OrderManager', '/order', 'Layout', '', b'1', b'0', b'1', '', 'list-unordered', b'0', '', 4, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4040, 4043, '商品问答', '', 'productAskBase', '/productAskBase', '@/views/pt/productAskBase/index', '', b'1', b'1', b'1', '', 'customer-service-2-fill', b'0', '', 99, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4041, 4039, '在线订单', '', 'orderBase', '/orderBase', '@/views/trade/orderBase/index', '', b'1', b'0', b'1', '', 'list-unordered', b'0', '', 1, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4042, 4052, '发票管理', '', 'orderInvoice', '/orderInvoice', '@/views/trade/orderInvoice/index', '', b'1', b'0', b'1', '', 'ticket-2-line', b'0', '', 0, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4043, 0, '商品管理', '', 'PtManager', '/pt', 'Layout', '', b'1', b'0', b'1', '', 'product-hunt-line', b'0', '', 3, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4047, 4043, '商品评价', '', 'productComment', '/productComment', '@/views/pt/productComment/index', '', b'1', b'0', b'1', '', 'kakao-talk-line', b'0', '', 90, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4048, 0, '店铺管理', '', 'StoreManager', '/store', 'Layout', '', b'1', b'0', b'1', '', 'store-2-line', b'0', '', 5, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4049, 4048, '发货地址', '', 'storeShippingAddress', '/storeShippingAddress', '@/views/shop/storeShippingAddress/index', '', b'1', b'0', b'1', '', 'pin-distance-line', b'0', '', 30, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4052, 0, '财务管理', '', 'FinanceManager', '/finance', 'Layout', '', b'1', b'0', b'1', '', 'secure-payment-line', b'0', '', 10, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4053, 0, '报表管理', '', 'ReportsManager', '/reports', 'Layout', '', b'1', b'1', b'0', '', 'bar-chart-box-line', b'0', '', 12, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4056, 4048, '物流工具', '', 'storeTransportType', '/storeTransportType', '@/views/shop/storeTransportType/index', '', b'1', b'0', b'1', '', 'tools-fill', b'0', '', 50, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4057, 4361, '手机装修', '', 'DiyMobile', '/diy-mobile', '@/views/sys/page/mobile', '', b'1', b'0', b'1', '', 'brush-2-line', b'0', '', 20, 1, '', '', 1, '', '/manage/sys/pageBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4058, 0, '资讯管理', '', 'CmsManager', '/cms', 'Layout', '', b'1', b'0', b'1', '', 'newspaper-line', b'0', '', 18, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4059, 4058, '文章分类', '', 'articleCategory', '/articleCategory', '@/views/cms/articleCategory/index', '', b'1', b'0', b'1', '', 'bar-chart-horizontal-line', b'0', '', 0, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4062, 1, '用户统计', '', 'AnalyticsUser', '/analytics/user', '@/views/analytics/user', '', b'1', b'0', b'1', '', 'user-search-line', b'0', '', 20, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4065, 4030, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/account/userInfo/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4066, 4030, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/account/userInfo/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4067, 4030, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/account/userInfo/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4068, 4030, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/account/userInfo/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4070, 4030, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/account/userInfo/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4072, 4031, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/account/userLevel/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4073, 4031, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/account/userLevel/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4074, 4031, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/account/userLevel/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4075, 4031, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/account/userLevel/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4095, 2, '角色管理', '', 'UserRole', '/userRole', '@/views/admin/userRole/index', '', b'1', b'0', b'1', '', 'user-3-line', b'0', '', 90, 1, '', '', 1, '', '/manage/admin/userRole/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4096, 4095, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/admin/userRole/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4097, 4095, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/admin/userRole/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4098, 4095, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/admin/userRole/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4099, 4095, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/admin/userRole/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4109, 4013, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/dict/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4110, 4013, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/dict/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4111, 4013, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/dict/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4112, 4013, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/dict/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4117, 4013, '查询项目', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/dict/listItem', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4118, 4013, '添加项目', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/dict/addItem', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4119, 4013, '修改项目', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/dict/editItem', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4120, 4013, '删除项目', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/dict/removeItem', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4126, 4043, '商品分类', '', 'ProductCategory', '/productCategory', '@/views/pt/productCategory/index', '', b'1', b'0', b'1', '', 'list-check-2', b'0', '', 20, 1, '', '', 1, '', '/manage/pt/productCategory/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4127, 4126, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productCategory/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4128, 4126, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productCategory/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4129, 4126, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productCategory/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4130, 4126, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productCategory/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4132, 4126, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/pt/productCategory/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4134, 4043, '商品类型', '', 'ProductType', '/productType', '@/views/pt/productType/index', '', b'1', b'0', b'1', '', 'list-check', b'0', '', 30, 1, '', '', 1, '', '/manage/pt/productType/list', b'0', '2024-07-17 08:20:17', 1001);
INSERT INTO `admin_menu_base` VALUES (4135, 4134, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productType/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4136, 4134, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productType/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4137, 4134, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productType/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4138, 4134, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productType/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4142, 4043, '商品品牌', '', 'ProductBrand', '/productBrand', '@/views/pt/productBrand/index', '', b'1', b'0', b'1', '', 'menu-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pt/productBrand/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4143, 4142, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productBrand/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4144, 4142, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productBrand/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4145, 4142, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productBrand/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4146, 4142, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productBrand/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4148, 4142, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/pt/productBrand/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4150, 4043, '标签管理', '', 'ProductTag', '/productTag', '@/views/pt/productTag/index', '', b'1', b'0', b'1', '', 'hashtag', b'0', '', 60, 1, '', '', 1, '', '/manage/pt/productTag/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4151, 4150, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productTag/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4152, 4150, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productTag/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4153, 4150, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productTag/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4154, 4150, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productTag/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4158, 4043, '商品列表', '', 'ProductBase', '/productBase', '@/views/pt/productBase/index', '', b'1', b'0', b'1', '', 'shopping-basket-line', b'0', '', 10, 1, '', '', 1, '', '/manage/pt/productBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4159, 4158, '查询', '', '', '/productBase', 'Layout', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4161, 4158, '增改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productBase/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4162, 4158, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productBase/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4166, 4043, '规格管理', '', 'ProductSpec', '/productSpec', '@/views/pt/productSpec/index', '', b'1', b'0', b'1', '', 'grid-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pt/productSpec/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4167, 4166, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productSpec/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4168, 4166, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productSpec/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4169, 4166, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productSpec/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4170, 4166, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productSpec/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4172, 4166, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/pt/productSpec/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4174, 4043, '添加商品', '', 'SelectCategory', '/add', '@/views/pt/productBase/add', '', b'1', b'0', b'1', '', 'add-box-line', b'0', '', 0, 1, '', '', 1, '', '/manage/pt/productBase/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4175, 4058, '文章管理', '', 'ArticleBase', '/articleBase', '@/views/cms/articleBase/index', '', b'1', b'0', b'1', '', 'article-line', b'0', '', 20, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4176, 4058, '标签管理', '', 'ArticleTag', '/articleTag', '@/views/cms/articleTag/index', '', b'1', b'0', b'1', '', 'hashtag', b'0', '', 40, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4177, 4058, '文章评论', '', 'ArticleComment', '/articleComment', '@/views/cms/articleComment/index', '', b'1', b'0', b'1', '', 'message-fill', b'0', '', 50, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4178, 4052, '支付会员', '', 'userResource', '/userResource', '@/views/pay/userResource/index', '', b'1', b'0', b'1', '', 'account-circle-line', b'0', '', 50, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4180, 4059, '添加', '', '', '/articleCategory', 'Layout', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleCategory/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4181, 4059, '修改', '', '', '/articleCategory', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/cms/articleCategory/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4182, 4059, '删除', '', '', '/articleCategory', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/cms/articleCategory/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4186, 4177, '添加', '', '', '/articleComment', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleComment/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4187, 4177, '修改', '', '', '/articleComment', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/cms/articleComment/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4190, 4049, '编辑', '', '', '/storeShippingAddress', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/shop/storeShippingAddress/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4191, 4049, '删除', '', '', '/storeShippingAddress', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/shop/storeShippingAddress/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4192, 4178, '列表', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/pay/userResource/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4194, 4178, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/pay/userResource/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4196, 4177, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleComment/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4197, 4177, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/cms/articleComment/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4199, 4177, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/cms/articleComment/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4200, 4177, '批量删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 30, 0, '', '', 1, '', '/manage/cms/articleComment/removeBatch', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4201, 4176, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleTag/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4202, 4176, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/cms/articleTag/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4203, 4176, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/cms/articleTag/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4204, 4176, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/cms/articleTag/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4207, 4176, '批量删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 30, 0, '', '', 1, '', '/manage/cms/articleTag/removeBatch', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4208, 4059, '查询', '', '', '/articleCategory', 'Layout', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleCategory/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4212, 4175, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/cms/articleBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4213, 4175, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/cms/articleBase/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4214, 4175, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/cms/articleBase/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4215, 4175, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/cms/articleBase/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4217, 4175, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/cms/articleBase/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4218, 4175, '批量删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 30, 0, '', '', 1, '', '/manage/cms/articleBase/removeBatch', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4219, 2, '菜单设置', '', 'Menu', '/menu', '@/views/admin/menu/index', '', b'1', b'0', b'1', '', 'list-unordered', b'0', '', 93, 1, '', '', 1, '', '/manage/admin/menu/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4220, 4219, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/admin/menu/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4221, 4219, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/admin/menu/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4222, 4219, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/admin/menu/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4223, 4219, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/admin/menu/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4225, 4219, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/admin/menu/editState', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4227, 2, '素材管理', '', 'Gallery', '/gallery', '@/views/sys/gallery/index', '', b'1', b'0', b'1', '', 'gallery-line', b'0', '', 30, 1, '', '', 1, '', '/manage/sys/material/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4228, 4227, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/material/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4229, 4227, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/material/add', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4230, 4227, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/material/edit', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4231, 4227, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/material/remove', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4234, 4227, '批量删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 30, 0, '', '', 1, '', '/manage/sys/material/removeBatch', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4235, 2, '配置管理', '', 'Config', '/config', '@/views/sys/config/index', '', b'1', b'0', b'1', '', 'command-fill', b'0', '', 1, 1, '', '', 1, '', '/manage/sys/config/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4236, 4235, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/config/list', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4237, 4235, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/config/add', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4238, 4235, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/config/edit', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4239, 4235, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/config/remove', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4243, 4048, '物流管理', '', 'storeExpressLogistics', '/storeExpressLogistics', '@/views/shop/storeExpressLogistics/index', '', b'1', b'0', b'1', '', 'list-check', b'0', '', 20, 1, '', '', 1, '', '', b'0', '2023-05-20 21:55:38', 1001);
INSERT INTO `admin_menu_base` VALUES (4244, 4048, '快递公司', '', 'ExpressBase', '/expressBase', '@/views/sys/expressBase/index', '', b'1', b'0', b'1', '', 'truck-line', b'0', '', 10, 1, '', '', 1, '', '/manage/sys/expressBase/list', b'0', '2024-07-17 08:25:14', 1001);
INSERT INTO `admin_menu_base` VALUES (4245, 4244, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/expressBase/list', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4246, 4244, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/expressBase/add', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4247, 4244, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/expressBase/edit', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4248, 4244, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/expressBase/remove', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4250, 4244, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/expressBase/editState', b'0', '2023-11-02 10:53:31', 1001);
INSERT INTO `admin_menu_base` VALUES (4252, 4049, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/shop/storeShippingAddress/list', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4253, 4049, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/shop/storeShippingAddress/add', b'0', '2023-05-30 11:55:54', 1001);
INSERT INTO `admin_menu_base` VALUES (4257, 4056, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/shop/storeTransportType/list', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4258, 4056, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/shop/storeTransportType/add', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4259, 4056, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/shop/storeTransportType/edit', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4260, 4056, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/shop/storeTransportType/remove', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4264, 4243, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/shop/storeExpressLogistics/list', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4265, 4243, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/shop/storeExpressLogistics/add', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4266, 4243, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/shop/storeExpressLogistics/edit', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4267, 4243, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/shop/storeExpressLogistics/remove', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4269, 4243, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/shop/storeExpressLogistics/editState', b'0', '2023-11-02 10:53:31', 1001);
INSERT INTO `admin_menu_base` VALUES (4271, 2, '地区管理', '', 'DistrictBase', '/districtBase', '@/views/sys/districtBase/index', '', b'1', b'0', b'1', '', 'luggage-deposit-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/districtBase/list', b'0', '2023-05-20 21:50:16', 1001);
INSERT INTO `admin_menu_base` VALUES (4272, 4271, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/districtBase/list', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4273, 4271, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/districtBase/add', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4274, 4271, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/districtBase/edit', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4275, 4271, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/districtBase/remove', b'0', '2023-05-30 11:55:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4277, 4271, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/districtBase/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (4279, 4042, '添加', '', '', '', '', '0', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/trade/orderInvoice/add', b'0', '2023-05-21 11:07:49', 1001);
INSERT INTO `admin_menu_base` VALUES (4284, 4042, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/trade/orderInvoice/list', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4285, 4042, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 3, 0, '', '', 1, '', '/manage/trade/orderInvoice/edit', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4286, 4042, '更改状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 4, 0, '', '', 1, '', '/manage/trade/orderInvoice/editState', b'0', '2023-11-02 14:15:21', 1001);
INSERT INTO `admin_menu_base` VALUES (4287, 4040, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 0, 0, '', '', 1, '', '/manage/pt/productAskBase/add', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4289, 4040, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/pt/productAskBase/list', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4290, 4040, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 4, 0, '', '', 1, '', '/manage/pt/productAskBase/edit', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4291, 4040, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productAskBase/remove', b'0', '2023-05-29 18:06:25', 1001);
INSERT INTO `admin_menu_base` VALUES (4309, 4029, '用户标签', '', 'userTag', '/userTag', '@/views/account/userTag/index', '', b'1', b'0', b'1', '', 'price-tag-3-line', b'0', '', 50, 1, '', '', 1, '', '', b'1', '2023-12-07 17:46:10', 1001);
INSERT INTO `admin_menu_base` VALUES (4315, 4309, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 0, 0, '', '', 1, '', '/manage/account/userTagBase/add', b'1', '2023-12-07 17:46:10', 1001);
INSERT INTO `admin_menu_base` VALUES (4317, 4309, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/account/userTagBase/list', b'1', '2023-12-07 17:46:10', 1001);
INSERT INTO `admin_menu_base` VALUES (4318, 4309, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 4, 0, '', '', 1, '', '/manage/account/userTagBase/edit', b'1', '2023-12-07 17:46:10', 1001);
INSERT INTO `admin_menu_base` VALUES (4319, 4309, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/account/userTagBase/remove', b'1', '2023-12-07 17:46:10', 1001);
INSERT INTO `admin_menu_base` VALUES (4325, 310, '协议设置', '', 'Site5', '/site/1005', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'sticky-note-line', b'0', '', 15, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4326, 310, '支付设置', '', 'Site4', '/site/1004', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'alipay-line', b'0', '', 13, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4327, 310, '上传设置', '', 'Site2', '/site/1002', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'arrow-up-circle-line', b'0', '', 2, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4328, 310, '保障服务', '', 'ContractType', '/contractType', '@/views/sys/contractType/index', '', b'1', b'0', b'1', '', 'service-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/contractType/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4329, 4328, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/contractType/list', b'0', '2023-05-30 11:55:56', 1001);
INSERT INTO `admin_menu_base` VALUES (4330, 4328, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/contractType/add', b'0', '2023-05-30 11:55:56', 1001);
INSERT INTO `admin_menu_base` VALUES (4331, 4328, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/contractType/edit', b'0', '2023-05-30 11:55:56', 1001);
INSERT INTO `admin_menu_base` VALUES (4332, 4328, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/contractType/remove', b'0', '2023-05-30 11:55:56', 1001);
INSERT INTO `admin_menu_base` VALUES (4334, 4328, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/contractType/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (4336, 310, '运营设置', '', 'Site3', '/site/1003', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'command-line', b'0', '', 14, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4337, 4361, '用户反馈', '', 'FeedbackBase', '/feedbackBase', '@/views/sys/feedbackBase/index', '', b'1', b'0', b'1', '', 'feedback-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/feedbackBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4345, 310, '反馈分类', '', 'FeedbackCategory', '/feedbackCategory', '@/views/sys/feedbackCategory/index', '', b'1', b'0', b'1', '', 'list-settings-line', b'0', '', 61, 1, '', '', 1, '', '/manage/sys/feedbackCategory/list', b'0', '2023-05-30 12:03:58', 1001);
INSERT INTO `admin_menu_base` VALUES (4346, 4345, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/feedbackCategory/list', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4347, 4345, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/feedbackCategory/add', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4348, 4345, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/feedbackCategory/edit', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4349, 4345, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/feedbackCategory/remove', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4351, 4345, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/feedbackCategory/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (4353, 310, '反馈类型', '', 'FeedbackType', '/feedbackType', '@/views/sys/feedbackType/index', '', b'1', b'0', b'1', '', 'feedback-line', b'0', '', 60, 1, '', '', 1, '', '/manage/sys/feedbackType/list', b'0', '2023-05-30 12:01:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4354, 4353, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/feedbackType/list', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4355, 4353, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/feedbackType/add', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4356, 4353, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/feedbackType/edit', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4357, 4353, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/feedbackType/remove', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4359, 4353, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/feedbackType/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (4361, 0, '运营管理', '', 'OperManager', '/oper', 'Layout', '', b'1', b'0', b'1', '', 'opera-line', b'0', '', 2, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4362, 0, '推广管理', '', 'FxManager', '/fx', 'Layout', '', b'1', b'1', b'0', '', 'node-tree', b'0', '', 15, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4363, 2, '个人中心', '', 'Personal', '/personal', '@/views/account/userInfo/personal', '', b'1', b'1', b'1', '', 'gallery-line', b'0', '', 255, 1, '', '', 1, '', '', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4364, 4361, 'PC装修', '', 'DiyPc', '/diy-pc', '@/views/sys/page/pc', '', b'1', b'0', b'1', '', 'brush-2-line', b'0', '', 20, 1, '', '', 1, '', '/manage/sys/pageBase/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (4365, 310, '消息模板', '', 'MessageTemplate', '/messageTemplate', '@/views/sys/messageTemplate/index', '', b'1', b'0', b'1', '', 'message-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/messageTemplate/list', b'0', '2023-05-30 12:01:27', 1001);
INSERT INTO `admin_menu_base` VALUES (4366, 4365, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/messageTemplate/list', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4367, 4365, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/sys/messageTemplate/add', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4368, 4365, '修改', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/messageTemplate/edit', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4369, 4365, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/sys/messageTemplate/remove', b'0', '2023-05-30 11:55:57', 1001);
INSERT INTO `admin_menu_base` VALUES (4371, 4365, '设置状态', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/messageTemplate/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (4373, 4029, '用户消息', '', 'UserMessage', '/userMessage', '@/views/account/userMessage/index', '', b'1', b'0', b'1', '', 'chat-3-line', b'0', '', 50, 1, '', '', 1, '', '/manage/account/userMessage/list', b'0', '2023-05-30 12:07:01', 1001);
INSERT INTO `admin_menu_base` VALUES (4374, 4373, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/account/userMessage/list', b'0', '2023-05-30 12:05:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4375, 4373, '添加', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/account/userMessage/add', b'0', '2023-05-30 12:05:55', 1001);
INSERT INTO `admin_menu_base` VALUES (4663, 1, '交易统计', '', 'AnalyticsOrder', '/analytics/order', '@/views/analytics/order', '', b'1', b'0', b'1', '', 'money-cny-circle-line', b'0', '', 30, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2024-07-17 08:20:17', 1001);
INSERT INTO `admin_menu_base` VALUES (4664, 1, '交易统计', '', 'AnalyticsTrade', '/analytics/trade', '@/views/analytics/trade', '', b'1', b'1', b'0', '', 'product-hunt-line', b'0', '', 40, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2023-11-02 12:05:25', 1001);
INSERT INTO `admin_menu_base` VALUES (5028, 2, '操作日志', '', 'LogAction', '/logAction', '@/views/sys/logAction/index', '', b'1', b'0', b'1', '', 'history-line', b'0', '', 100, 1, '', '', 1, '', '/manage/sys/logAction/list', b'0', '2023-05-30 14:24:10', 1001);
INSERT INTO `admin_menu_base` VALUES (5036, 2, '错误日志', '', 'LogError', '/logError', '@/views/sys/logError/index', '', b'1', b'0', b'1', '', 'error-warning-line', b'0', '', 101, 1, '', '', 1, '', '/manage/sys/logError/list', b'0', '2023-05-30 14:23:01', 1001);
INSERT INTO `admin_menu_base` VALUES (5044, 2, '计划任务', '', 'CrontabBase', '/crontabBase', '@/views/sys/crontabBase/index', '', b'1', b'0', b'1', '', 'task-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/crontabBase/list', b'0', '2023-05-30 18:18:15', 1001);
INSERT INTO `admin_menu_base` VALUES (5045, 5044, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/crontabBase/list', b'0', '2023-05-30 18:14:44', 1001);
INSERT INTO `admin_menu_base` VALUES (5047, 5044, '修改', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/sys/crontabBase/edit', b'0', '2023-05-30 18:14:44', 1001);
INSERT INTO `admin_menu_base` VALUES (5050, 5044, '设置状态', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/sys/crontabBase/editState', b'0', '2023-11-02 10:51:50', 1001);
INSERT INTO `admin_menu_base` VALUES (5052, 4041, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/trade/orderBase/list', b'0', '2023-06-10 16:48:27', 1001);
INSERT INTO `admin_menu_base` VALUES (5053, 4041, '添加', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/trade/orderBase/add', b'0', '2023-06-10 16:48:27', 1001);
INSERT INTO `admin_menu_base` VALUES (5054, 4041, '修改', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/trade/orderBase/edit', b'0', '2023-06-10 16:48:27', 1001);
INSERT INTO `admin_menu_base` VALUES (5055, 4041, '删除', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/trade/orderBase/remove', b'0', '2023-06-10 16:48:27', 1001);
INSERT INTO `admin_menu_base` VALUES (5057, 4041, '设置状态', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5059, 310, '订单流程', '', 'Site50', '/site/5001', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'list-settings-line', b'0', '', 49, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-11-02 11:08:25', 1001);
INSERT INTO `admin_menu_base` VALUES (5060, 4039, '售后服务', '', 'OrderReturn', '/orderReturn', '@/views/trade/orderReturn/index', '', b'1', b'0', b'1', '', 'customer-service-2-fill', b'0', '', 12, 1, '', '', 1, '', '/manage/trade/orderReturn/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5061, 5060, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/trade/orderReturn/list', b'0', '2023-06-19 15:41:28', 1001);
INSERT INTO `admin_menu_base` VALUES (5063, 5060, '修改', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/trade/orderReturn/edit', b'0', '2023-06-19 15:41:28', 1001);
INSERT INTO `admin_menu_base` VALUES (5070, 4033, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/pay/userPointsHistory/list', b'0', '2023-06-29 14:04:16', 1001);
INSERT INTO `admin_menu_base` VALUES (5078, 4052, '收支明细', '', 'ConsumeRecord', '/consumeRecord', '@/views/pay/consumeRecord/index', '', b'1', b'0', b'1', '', 'money-cny-circle-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pay/consumeRecord/list', b'0', '2023-10-19 17:12:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5079, 5078, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pay/consumeRecord/list', b'0', '2023-06-30 11:40:09', 1001);
INSERT INTO `admin_menu_base` VALUES (5086, 4052, '充值记录', '', 'ConsumeDeposit', '/consumeDeposit', '@/views/pay/consumeDeposit/index', '', b'1', b'0', b'1', '', 'user-3-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pay/consumeDeposit/list', b'0', '2023-06-30 11:44:00', 1001);
INSERT INTO `admin_menu_base` VALUES (5087, 5086, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pay/consumeDeposit/list', b'0', '2023-06-30 11:40:37', 1001);
INSERT INTO `admin_menu_base` VALUES (5092, 5086, '设置状态', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/pay/consumeDeposit/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5094, 4052, '交易订单', '', 'ConsumeTrade', '/consumeTrade', '@/views/pay/consumeTrade/index', '', b'1', b'0', b'1', '', 'user-3-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pay/consumeTrade/list', b'0', '2023-10-19 17:13:28', 1001);
INSERT INTO `admin_menu_base` VALUES (5095, 5094, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pay/consumeTrade/list', b'0', '2023-06-30 11:41:17', 1001);
INSERT INTO `admin_menu_base` VALUES (5102, 1, '商品统计', '', 'AnalyticsProduct', '/analytics/product', '@/views/analytics/product', '', b'1', b'0', b'1', '', 'product-hunt-line', b'0', '', 10, 1, '', '', 1, '', '/manage/analytics/index/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5108, 2, '服务配置', '', 'Site60', '/site/6001', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'list-settings-line', b'0', '', 72, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5109, 310, '短信配置', '', 'Site70', '/site/7001', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'list-settings-line', b'0', '', 46, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-11-02 11:08:25', 1001);
INSERT INTO `admin_menu_base` VALUES (5110, 4047, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 0, 0, '', '', 1, '', '/manage/pt/productComment/list', b'0', '2023-07-20 14:42:37', 1001);
INSERT INTO `admin_menu_base` VALUES (5111, 310, '微信设置', '', 'Site80', '/site/8001', '@/views/sys/config/site', '', b'1', b'0', b'1', '', 'wechat-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-11-02 11:08:25', 1001);
INSERT INTO `admin_menu_base` VALUES (5119, 2, '授权代码', '', 'ReleaseWeb', '/release/web', '@/views/sys/release/web', '', b'1', b'0', b'1', '', 'folder-download-line', b'0', '', 75, 1, '', '', 1, '', '/manage/sys/release/download', b'0', '2023-11-02 11:57:21', 1001);
INSERT INTO `admin_menu_base` VALUES (5125, 4043, '辅助属性', '', 'ProductAssist', '/productAssist', '@/views/pt/productAssist/index', '', b'1', b'0', b'1', '', 'filter-2-line', b'0', '', 50, 1, '', '', 1, '', '/manage/pt/productAssist/list', b'0', '2023-09-07 11:03:22', 1001);
INSERT INTO `admin_menu_base` VALUES (5126, 5125, '查询', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/pt/productAssist/list', b'0', '2023-09-07 10:45:04', 1001);
INSERT INTO `admin_menu_base` VALUES (5127, 5125, '添加', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 5, 0, '', '', 1, '', '/manage/pt/productAssist/add', b'0', '2023-09-07 10:45:04', 1001);
INSERT INTO `admin_menu_base` VALUES (5128, 5125, '修改', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 10, 0, '', '', 1, '', '/manage/pt/productAssist/edit', b'0', '2023-09-07 10:45:04', 1001);
INSERT INTO `admin_menu_base` VALUES (5129, 5125, '删除', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 15, 0, '', '', 1, '', '/manage/pt/productAssist/remove', b'0', '2023-09-07 10:45:04', 1001);
INSERT INTO `admin_menu_base` VALUES (5131, 5125, '设置状态', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/pt/productAssist/editState', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5158, 4361, '页面导航', '', 'PagePcNav', '/pagePcNav', '@/views/sys/page/pagePcNav', '', b'1', b'0', b'1', '', 'guide-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/pagePcNav/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5166, 4361, '移动首页分类', '', 'PageMobileNav', '/pageMobileNav', '@/views/sys/pageMobileNav/index', '', b'1', b'1', b'1', '', 'user-3-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/pageMobileNav/list', b'0', '2023-11-03 11:12:47', 1001);
INSERT INTO `admin_menu_base` VALUES (5167, 5166, '装修', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/sys/diy/index', b'0', '2023-11-03 11:12:22', 1001);
INSERT INTO `admin_menu_base` VALUES (5174, 4361, '帮助导航', '', 'PcHelp', '/page/help', '@/views/sys/page/pcHelp', '', b'1', b'0', b'1', '', 'questionnaire-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/diy/index', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5175, 4361, '分类导航', '', 'PageCategoryNav', '/page/categoryNav', '@/views/sys/page/pageCategoryNav', '', b'1', b'0', b'1', '', 'navigation-line', b'0', '', 50, 1, '', '', 1, '', '/manage/sys/diy/index', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5176, 2, '清理缓存', '', 'cache', '/cache', '@/views/sys/cache/index', '', b'1', b'0', b'1', '', 'rocket-2-line', b'0', '', 73, 1, '', '', 1, '', '/manage/sys/cache/clean', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5177, 4337, '回复', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/sys/feedbackBase/edit', b'0', '2023-10-23 10:01:29', 1001);
INSERT INTO `admin_menu_base` VALUES (5178, 310, '推广设置', '', 'Promotion', '/promotion', '@/views/sys/config/promotion', '', b'1', b'0', b'1', '', 'team-line', b'0', '', 14, 1, '', '', 1, '', '/manage/sys/config/editSite', b'0', '2023-11-02 11:08:17', 1001);
INSERT INTO `admin_menu_base` VALUES (5179, 310, '退货原因', '', 'OrderReturnReason', '/orderReturnReason', '@/views/trade/orderReturnReason/index', '', b'1', b'0', b'1', '', 'list-ordered', b'0', '', 62, 1, '', '', 1, '', '', b'0', '2023-10-30 14:21:23', 1001);
INSERT INTO `admin_menu_base` VALUES (5180, 5179, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 0, 0, '', '', 1, '', '/manage/trade/orderReturnReason/list', b'0', '2023-10-30 14:22:01', 1001);
INSERT INTO `admin_menu_base` VALUES (5181, 5179, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/trade/orderReturnReason/add', b'0', '2023-10-30 14:22:28', 1001);
INSERT INTO `admin_menu_base` VALUES (5182, 5179, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/trade/orderReturnReason/edit', b'0', '2023-10-30 14:23:14', 1001);
INSERT INTO `admin_menu_base` VALUES (5183, 5179, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 3, 0, '', '', 1, '', '/manage/trade/orderReturnReason/remove', b'0', '2023-10-30 14:23:40', 1001);
INSERT INTO `admin_menu_base` VALUES (5184, 2, '系统用户', '', 'UserAdmin', '/userAdmin', '@/views/admin/userAdmin/index', '', b'1', b'0', b'1', '', 'admin-fill', b'0', '', 90, 1, '', '', 1, '', '', b'0', '2023-11-02 10:33:01', 1001);
INSERT INTO `admin_menu_base` VALUES (5185, 5184, '添加', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 1, 0, '', '', 1, '', '/manage/admin/userAdmin/add', b'0', '2023-11-02 10:42:43', 1001);
INSERT INTO `admin_menu_base` VALUES (5186, 5184, '编辑', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/admin/userAdmin/edit', b'0', '2023-11-02 10:42:44', 1001);
INSERT INTO `admin_menu_base` VALUES (5187, 5184, '删除', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 3, 0, '', '', 1, '', '/manage/admin/userAdmin/remove', b'0', '2023-11-02 10:42:45', 1001);
INSERT INTO `admin_menu_base` VALUES (5188, 5184, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 0, 0, '', '', 1, '', '/manage/admin/userAdmin/list', b'0', '2023-11-02 10:42:52', 1001);
INSERT INTO `admin_menu_base` VALUES (5199, 4030, '优惠券查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/shop/userVoucher/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5200, 4030, '订单查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/pay/consumeTrade/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5201, 4030, '积分查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/pay/userPointsHistory/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5202, 4030, '流水查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/pay/consumeRecord/list', b'0', '2023-12-07 17:45:41', 1001);
INSERT INTO `admin_menu_base` VALUES (5203, 4041, '取消', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/cancel', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5204, 4041, '审核', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/review', b'0', '2023-11-02 10:53:32', 1001);
INSERT INTO `admin_menu_base` VALUES (5205, 4041, '财务', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/finance', b'0', '2023-11-03 10:01:47', 1001);
INSERT INTO `admin_menu_base` VALUES (5206, 4041, '出库', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/picking', b'0', '2023-11-03 10:00:44', 1001);
INSERT INTO `admin_menu_base` VALUES (5207, 4041, '发货', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/shipping', b'0', '2023-11-03 10:00:47', 1001);
INSERT INTO `admin_menu_base` VALUES (5208, 4041, '收货', '', '', '', '', '', b'0', b'0', b'1', '', '', b'0', '', 25, 0, '', '', 1, '', '/manage/trade/orderBase/receive', b'0', '2023-11-03 10:01:43', 1001);
INSERT INTO `admin_menu_base` VALUES (5209, 4057, '装修', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/sys/diy/index', b'0', '2023-11-03 10:36:18', 1001);
INSERT INTO `admin_menu_base` VALUES (5210, 4364, '装修', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/sys/diy/index', b'0', '2023-11-03 10:37:24', 1001);
INSERT INTO `admin_menu_base` VALUES (5212, 4337, '查询', '', '', '', '', '', b'1', b'0', b'1', '', '', b'0', '', 50, 0, '', '', 1, '', '/manage/sys/feedbackBase/list', b'0', '2023-11-03 11:46:59', 1001);
INSERT INTO `admin_menu_base` VALUES (5354, 4047, '删除', '', '', '/productComment', '', '', b'1', b'0', b'1', '', '', b'0', '', 2, 0, '', '', 1, '', '/manage/pt/productComment/remove', b'1', '2024-06-29 15:10:31', 1001);

-- ----------------------------
-- Table structure for admin_user_admin
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_admin`;
CREATE TABLE `admin_user_admin`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色',
  `user_admin_ctime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_admin_utime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `user_is_superadmin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否超管',
  `role_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 9 COMMENT '角色编号(ENUM):0-用户;2-商家;3-门店;9-平台;',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属门店',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `user_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-否;1-是',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user_admin
-- ----------------------------
INSERT INTO `admin_user_admin` VALUES (10001, 1001, '2023-05-05 15:15:15', '2023-09-20 18:42:36', b'1', 9, 0, 0, 1001, b'1');

-- ----------------------------
-- Table structure for admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_role`;
CREATE TABLE `admin_user_role`  (
  `user_role_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色编号',
  `user_role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `user_role_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色标识',
  `menu_ids` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求列表(DOT)',
  `user_role_ctime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_role_utime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `user_role_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1005 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user_role
-- ----------------------------
INSERT INTO `admin_user_role` VALUES (1001, '管理员', 'admin', '1,101,5102,4062,4663,4361,4057,5209,4364,5210,4337,5177,5212,5158,5166,5167,5174,5175,4043,4174,4158,4159,4161,4162,4126,4127,4128,4129,4130,4132,4134,4135,4136,4137,4138,4142,4143,4144,4145,4146,4148,4166,4167,4168,4169,4170,4172,5125,5126,5127,5128,5129,5131,4150,4151,4152,4153,4154,4047,5110,5354,4040,4287,4289,4290,4291,4039,4041,5052,5053,5054,5055,5057,5203,5204,5205,5206,5207,5208,5245,5246,5247,5060,5061,5063,4048,4244,4245,4246,4247,4248,4250,4243,4264,4265,4266,4267,4269,4049,4252,4253,4190,4191,4056,4257,4258,4259,4260,4029,4030,4065,4066,4067,4068,4070,5199,5200,5201,5202,4031,4072,4073,4074,4075,4309,4315,4317,4318,4319,4373,4374,4375,4051,5008,5073,5074,5075,5076,5077,5124,5189,5190,5191,5192,5001,5002,5214,5215,5217,5218,5003,5004,5005,5006,5193,5194,5195,5196,5007,5009,5010,5011,5012,5014,5015,5016,5017,5018,5019,4052,4033,5070,4042,4279,4284,4285,4286,5248,4178,4192,4194,5078,5079,5086,5087,5092,5094,5095,5112,5113,5120,5121,5122,5123,5114,5115,5116,5117,5118,5219,5220,5271,5272,5273,5274,5225,5263,5264,5265,5266,5227,5279,5280,5281,5282,5228,5287,5288,5289,5290,5229,5295,5296,5297,5298,5230,5303,5304,5305,5306,5241,5311,5312,5313,5314,5242,5319,5320,5321,5322,5133,5134,5135,5136,5137,5142,5143,5150,5151,5153,4058,4059,4180,4208,4181,4182,4175,4212,4213,4214,4215,4217,4218,4176,4201,4202,4203,4204,4207,4177,4186,4196,4187,4197,4199,4200,310,360,4327,4326,4336,5178,4325,5109,5059,4328,4329,4330,4331,4332,4334,4365,4366,4367,4368,4369,4371,5111,4353,4354,4355,4356,4357,4359,4345,4346,4347,4348,4349,4351,5179,5180,5181,5182,5183,2,4235,4236,4237,4238,4239,4013,4109,4117,4110,4118,4111,4119,4112,4120,4227,4228,4229,4230,4231,4234,4271,4272,4273,4274,4275,4277,5044,5045,5047,5050,5108,5176,5119,4095,4096,4097,4098,4099,5184,5188,5185,5186,5187,4219,4220,4221,4222,4223,4225,5028,5036,5261,4023,4363', '2022-12-22 17:16:48', '2024-07-01 16:58:03', b'1', 0, 1001);
INSERT INTO `admin_user_role` VALUES (1002, '知识付费', 'edu', '1,101,5102,4062,4663,4361,4057,5209,4364,5210,4337,5177,5212,5158,5166,5167,5174,5175,4043,4174,4158,4159,4161,4162,4126,4127,4128,4129,4130,4132,4134,4135,4136,4137,4138,4142,4143,4144,4145,4146,4148,4166,4167,4168,4169,4170,4172,5125,5126,5127,5128,5129,5131,4150,4151,4152,4153,4154,4047,5110,5354,4040,4287,4289,4290,4291,4039,4041,5052,5053,5054,5055,5057,5203,5204,5205,5206,5207,5208,5245,5246,5247,5060,5061,5063,4048,4244,4245,4246,4247,4248,4250,4243,4264,4265,4266,4267,4269,4049,4252,4253,4190,4191,4056,4257,4258,4259,4260,4029,4030,4065,4066,4067,4068,4070,5199,5200,5201,5202,4031,4072,4073,4074,4075,4309,4315,4317,4318,4319,4373,4374,4375,4051,5008,5073,5074,5075,5076,5077,5124,5189,5190,5191,5192,5001,5002,5214,5215,5217,5218,5003,5004,5005,5006,5193,5194,5195,5196,5007,5009,5010,5011,5012,5014,5015,5016,5017,5018,5019,4052,4033,5070,4042,4279,4284,4285,4286,5248,4178,4192,4194,5078,5079,5086,5087,5092,5094,5095,5112,5113,5120,5121,5122,5123,5114,5115,5116,5117,5118,5219,5220,5271,5272,5273,5274,5225,5263,5264,5265,5266,5227,5279,5280,5281,5282,5228,5287,5288,5289,5290,5229,5295,5296,5297,5298,5230,5303,5304,5305,5306,5241,5311,5312,5313,5314,5242,5319,5320,5321,5322,5133,5134,5135,5136,5137,5142,5143,5150,5151,5153,4058,4059,4180,4208,4181,4182,4175,4212,4213,4214,4215,4217,4218,4176,4201,4202,4203,4204,4207,4177,4186,4196,4187,4197,4199,4200,310,360,4327,4326,4336,5178,4325,5109,5059,4328,4329,4330,4331,4332,4334,4365,4366,4367,4368,4369,4371,5111,4353,4354,4355,4356,4357,4359,4345,4346,4347,4348,4349,4351,5179,5180,5181,5182,5183,2,4235,4236,4237,4238,4239,4013,4109,4117,4110,4118,4111,4119,4112,4120,4227,4228,4229,4230,4231,4234,4271,4272,4273,4274,4275,4277,5044,5045,5047,5050,5108,5176,5119,4095,4096,4097,4098,4099,5184,5188,5185,5186,5187,4219,4220,4221,4222,4223,4225,5028,5036,5261,4023,4363', '2024-01-26 11:24:19', '2024-07-01 17:00:59', b'1', 0, 1001);
INSERT INTO `admin_user_role` VALUES (1003, '门店', 'chain', '1,101,5102,4062,4663,4361,4057,5209,4364,5210,4337,5177,5212,5158,5166,5167,5174,5175,4043,4174,4158,4159,4161,4162,4126,4127,4128,4129,4130,4132,4134,4135,4136,4137,4138,4142,4143,4144,4145,4146,4148,4166,4167,4168,4169,4170,4172,5125,5126,5127,5128,5129,5131,4150,4151,4152,4153,4154,4047,5110,5354,4040,4287,4289,4290,4291,4039,4041,5052,5053,5054,5055,5057,5203,5204,5205,5206,5207,5208,5245,5246,5247,5060,5061,5063,4048,4244,4245,4246,4247,4248,4250,4243,4264,4265,4266,4267,4269,4049,4252,4253,4190,4191,4056,4257,4258,4259,4260,4029,4030,4065,4066,4067,4068,4070,5199,5200,5201,5202,4031,4072,4073,4074,4075,4309,4315,4317,4318,4319,4373,4374,4375,4051,5008,5073,5074,5075,5076,5077,5124,5189,5190,5191,5192,5001,5002,5214,5215,5217,5218,5003,5004,5005,5006,5193,5194,5195,5196,5007,5009,5010,5011,5012,5014,5015,5016,5017,5018,5019,4052,4033,5070,4042,4279,4284,4285,4286,5248,4178,4192,4194,5078,5079,5086,5087,5092,5094,5095,5112,5113,5120,5121,5122,5123,5114,5115,5116,5117,5118,5219,5220,5271,5272,5273,5274,5225,5263,5264,5265,5266,5227,5279,5280,5281,5282,5228,5287,5288,5289,5290,5229,5295,5296,5297,5298,5230,5303,5304,5305,5306,5241,5311,5312,5313,5314,5242,5319,5320,5321,5322,5133,5134,5135,5136,5137,5142,5143,5150,5151,5153,4058,4059,4180,4208,4181,4182,4175,4212,4213,4214,4215,4217,4218,4176,4201,4202,4203,4204,4207,4177,4186,4196,4187,4197,4199,4200,310,360,4327,4326,4336,5178,4325,5109,5059,4328,4329,4330,4331,4332,4334,4365,4366,4367,4368,4369,4371,5111,4353,4354,4355,4356,4357,4359,4345,4346,4347,4348,4349,4351,5179,5180,5181,5182,5183,2,4235,4236,4237,4238,4239,4013,4109,4117,4110,4118,4111,4119,4112,4120,4227,4228,4229,4230,4231,4234,4271,4272,4273,4274,4275,4277,5044,5045,5047,5050,5108,5176,5119,4095,4096,4097,4098,4099,5184,5188,5185,5186,5187,4219,4220,4221,4222,4223,4225,5028,5036,5261,4023,4363', '2024-05-15 23:16:47', '2024-07-01 17:00:59', b'1', 0, 1001);
INSERT INTO `admin_user_role` VALUES (1004, '财务', 'finance', '1,101,5102,4062,4663,4361,4057,5209,4364,5210,4337,5177,5212,5158,5166,5167,5174,5175,4043,4174,4158,4159,4161,4162,4126,4127,4128,4129,4130,4132,4134,4135,4136,4137,4138,4142,4143,4144,4145,4146,4148,4166,4167,4168,4169,4170,4172,5125,5126,5127,5128,5129,5131,4150,4151,4152,4153,4154,4047,5110,5354,4040,4287,4289,4290,4291,4039,4041,5052,5053,5054,5055,5057,5203,5204,5205,5206,5207,5208,5245,5246,5247,5060,5061,5063,4048,4244,4245,4246,4247,4248,4250,4243,4264,4265,4266,4267,4269,4049,4252,4253,4190,4191,4056,4257,4258,4259,4260,4029,4030,4065,4066,4067,4068,4070,5199,5200,5201,5202,4031,4072,4073,4074,4075,4309,4315,4317,4318,4319,4373,4374,4375,4051,5008,5073,5074,5075,5076,5077,5124,5189,5190,5191,5192,5001,5002,5214,5215,5217,5218,5003,5004,5005,5006,5193,5194,5195,5196,5007,5009,5010,5011,5012,5014,5015,5016,5017,5018,5019,4052,4033,5070,4042,4279,4284,4285,4286,5248,4178,4192,4194,5078,5079,5086,5087,5092,5094,5095,5112,5113,5120,5121,5122,5123,5114,5115,5116,5117,5118,5219,5220,5271,5272,5273,5274,5225,5263,5264,5265,5266,5227,5279,5280,5281,5282,5228,5287,5288,5289,5290,5229,5295,5296,5297,5298,5230,5303,5304,5305,5306,5241,5311,5312,5313,5314,5242,5319,5320,5321,5322,5133,5134,5135,5136,5137,5142,5143,5150,5151,5153,4058,4059,4180,4208,4181,4182,4175,4212,4213,4214,4215,4217,4218,4176,4201,4202,4203,4204,4207,4177,4186,4196,4187,4197,4199,4200,310,360,4327,4326,4336,5178,4325,5109,5059,4328,4329,4330,4331,4332,4334,4365,4366,4367,4368,4369,4371,5111,4353,4354,4355,4356,4357,4359,4345,4346,4347,4348,4349,4351,5179,5180,5181,5182,5183,2,4235,4236,4237,4238,4239,4013,4109,4117,4110,4118,4111,4119,4112,4120,4227,4228,4229,4230,4231,4234,4271,4272,4273,4274,4275,4277,5044,5045,5047,5050,5108,5176,5119,4095,4096,4097,4098,4099,5184,5188,5185,5186,5187,4219,4220,4221,4222,4223,4225,5028,5036,5261,4023,4363', '2024-07-01 16:56:46', '2024-07-01 17:00:59', b'1', 0, 1001);

-- ----------------------------
-- Table structure for cms_article_analytics
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_analytics`;
CREATE TABLE `cms_article_analytics`  (
  `article_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_clicked` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击',
  `article_praise` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞',
  `article_comment_count` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章评论数目',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`article_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '新闻点赞数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_attachment
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_attachment`;
CREATE TABLE `cms_article_attachment`  (
  `attachment_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章作者编号',
  `attachment_time` int(10) UNSIGNED NOT NULL COMMENT '文章日期',
  `attachment_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `attachment_title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `attachment_remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章摘要',
  `attachment_password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章浏览密码',
  `attachment_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章别名',
  `attachment_link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章的url，链接地址',
  `attachment_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件path',
  `attachment_mime_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '上传的附件类型',
  `attachment_file` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件',
  `attachment_metadata` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `attachment_alt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章第一次添加',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`attachment_id`) USING BTREE,
  INDEX `type_status_date`(`attachment_time`, `attachment_id`) USING BTREE,
  INDEX `article_author`(`user_id`) USING BTREE,
  INDEX `article_name`(`attachment_name`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_base
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_base`;
CREATE TABLE `cms_article_base`  (
  `article_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `article_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章标题',
  `article_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章别名',
  `article_excerpt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章摘要',
  `article_content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容(HTML)',
  `article_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '调用网址:默认为本页面构造的网址',
  `category_id` smallint(5) UNSIGNED NOT NULL COMMENT '所属分类',
  `article_template` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章模板',
  `article_seo_title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SEO标题',
  `article_seo_keywords` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SEO关键字',
  `article_seo_description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SEO描述',
  `article_reply_flag` bit(1) NOT NULL DEFAULT b'1' COMMENT '启用问答(BOOL):0-否;1-是',
  `article_lang` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'cn' COMMENT '文章语言',
  `article_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '文章类型(ENUM):1-文章;2-公告',
  `article_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章排序',
  `article_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '文章状态(BOOL):0-关闭;1-启用',
  `article_add_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `article_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章图片',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章作者',
  `article_tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章标签(DOT):文章标签',
  `article_is_popular` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否热门(BOOL):0-否;1-是',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`article_id`) USING BTREE,
  UNIQUE INDEX `article_title`(`article_title`) USING BTREE,
  INDEX `category_id`(`category_id`, `article_sort`) USING BTREE,
  INDEX `category_id_3`(`category_id`, `article_status`, `article_sort`) USING BTREE,
  INDEX `category_id_2`(`category_id`, `article_type`, `article_status`, `article_sort`) USING BTREE,
  INDEX `article_sort`(`user_id`, `category_id`, `article_status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章内容' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_category
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_category`;
CREATE TABLE `cms_article_category`  (
  `category_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名称',
  `category_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级编号',
  `category_image_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `category_keywords` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类关键词',
  `category_desc` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类描述',
  `category_count` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内容数量',
  `category_template` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类模板',
  `category_alias` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类别名',
  `category_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '分类排序',
  `category_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):0-非内置;1-内置;',
  `category_is_leaf` bit(1) NOT NULL DEFAULT b'1' COMMENT '叶节点(BOOL):0-否;1-是;',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`category_id`) USING BTREE,
  UNIQUE INDEX `category_name`(`category_name`) USING BTREE,
  INDEX `category_parent_id`(`category_parent_id`, `category_order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_category_relationship
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_category_relationship`;
CREATE TABLE `cms_article_category_relationship`  (
  `category_relationship_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` mediumint(8) UNSIGNED NOT NULL,
  `category_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `category_relationship_order` mediumint(8) UNSIGNED NOT NULL DEFAULT 0,
  `article_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'publish' COMMENT '文章状态[ \'draft\' | \'publish\' | \'pending\'| \'future\' | \'private\' ]  ， 冗余字段，为了检索高效',
  `article_publish_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章发布时间|冗余字段 ->可以不加，通过后台程序更改',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`category_relationship_id`) USING BTREE,
  INDEX `article_id`(`article_id`) USING BTREE,
  INDEX `category_id`(`category_id`) USING BTREE,
  INDEX `article_status`(`article_status`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_comment
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_comment`;
CREATE TABLE `cms_article_comment`  (
  `comment_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论编号',
  `article_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章编号',
  `comment_user_ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论IP',
  `comment_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '评论时间',
  `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容 ',
  `comment_karma` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论karma值 ',
  `comment_approved` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '评论许可(ENUM):0-0;1-1;spam-spam\' ',
  `comment_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论代理:（浏览器，操作系统等） ',
  `comment_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论类型:（pingback|trackback)',
  `comment_parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论者编号',
  `comment_helpful_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有帮助数',
  `comment_is_show` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否显示(BOOL):1-显示;0-不显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `comment_post_id`(`article_id`) USING BTREE,
  INDEX `comment_approved_date_gmt`(`comment_approved`) USING BTREE,
  INDEX `comment_parent`(`comment_parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_comment_helpful
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_comment_helpful`;
CREATE TABLE `cms_article_comment_helpful`  (
  `comment_helpful_id` bigint(20) UNSIGNED NOT NULL COMMENT '编号',
  `comment_id` bigint(20) UNSIGNED NOT NULL COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户编号',
  `article_id` int(10) UNSIGNED NOT NULL COMMENT '文章编号',
  `comment_helpful_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_helpful_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`comment_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品评价有用投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_comment_reply`;
CREATE TABLE `cms_article_comment_reply`  (
  `comment_reply_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论回复编号',
  `comment_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论编号',
  `comment_reply_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复父编号',
  `article_id` int(10) UNSIGNED NOT NULL COMMENT '所属文章编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论回复编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复姓名',
  `user_id_to` int(10) UNSIGNED NOT NULL COMMENT '评论回复用户编号',
  `comment_reply_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复内容',
  `comment_reply_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论回复时间',
  `comment_reply_show_flag` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '问答是否显示',
  `comment_reply_islook` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已读(BOOL):1-已读;2-未读',
  `comment_reply_helpful_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有帮助数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_reply_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '问答回复表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_comment_reply_helpful
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_comment_reply_helpful`;
CREATE TABLE `cms_article_comment_reply_helpful`  (
  `comment_reply_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户编号',
  `article_id` int(10) UNSIGNED NOT NULL COMMENT '文章编号',
  `comment_helpful_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_reply_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`comment_reply_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章评价有用投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_group
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_group`;
CREATE TABLE `cms_article_group`  (
  `article_group_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `article_group_title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `article_group_lang` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'cn' COMMENT '语言',
  `article_group_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `article_group_logo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'logo',
  `article_group_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级分类编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`article_group_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '网站初始化内容分组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_id_number
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_id_number`;
CREATE TABLE `cms_article_id_number`  (
  `article_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `article_time` int(10) UNSIGNED NOT NULL COMMENT '文章日期',
  `article_order` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章的页面序号，不使用的话一般为0',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`article_id`) USING BTREE,
  INDEX `type_status_date`(`article_time`, `article_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章序号表-暂时不用，为了优化高效，使用redis等' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_read
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_read`;
CREATE TABLE `cms_article_read`  (
  `read_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '阅读编号',
  `article_id` mediumint(8) UNSIGNED NOT NULL COMMENT '文章编号',
  `read_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '阅读者ip',
  `read_time` date NOT NULL COMMENT '阅读时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`read_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章访问记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_recommend
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_recommend`;
CREATE TABLE `cms_article_recommend`  (
  `recommend_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '推荐编号',
  `article_id` mediumint(8) UNSIGNED NOT NULL COMMENT '文章编号',
  `recommend_type` enum('slider','hot') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'slider' COMMENT '推荐类型',
  `recommend_image_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '推荐图片地址',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`recommend_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推荐表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_reply
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_reply`;
CREATE TABLE `cms_article_reply`  (
  `article_reply_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论回复编号',
  `article_reply_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复父编号',
  `article_id` int(10) UNSIGNED NOT NULL COMMENT '所属文章编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论回复编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复姓名',
  `user_id_to` int(10) UNSIGNED NOT NULL COMMENT '评论回复用户编号',
  `user_name_to` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论回复用户名称',
  `article_reply_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复内容',
  `article_reply_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论回复时间',
  `article_reply_show_flag` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '问答是否显示',
  `article_reply_islook` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已读(BOOL):1-已读;2-未读',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`article_reply_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '问答回复表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_tag`;
CREATE TABLE `cms_article_tag`  (
  `tag_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签编号',
  `tag_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `tag_count` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内容数量',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cms_article_tag_relationship
-- ----------------------------
DROP TABLE IF EXISTS `cms_article_tag_relationship`;
CREATE TABLE `cms_article_tag_relationship`  (
  `tag_relationship_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` mediumint(8) UNSIGNED NOT NULL,
  `tag_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `tag_relationship_order` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`tag_relationship_id`) USING BTREE,
  INDEX `term_taxonomy_id`(`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章标签关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for invoicing_stock_bill
-- ----------------------------
DROP TABLE IF EXISTS `invoicing_stock_bill`;
CREATE TABLE `invoicing_stock_bill`  (
  `stock_bill_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '购货(退货)单编号',
  `stock_bill_checked` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否审核(BOOL):1-已审核;  0-未审核',
  `stock_bill_date` date NOT NULL COMMENT '单据日期',
  `stock_bill_modify_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `stock_bill_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `bill_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '业务类别purchase_type_id, sale_type_id(ENUM):2750-入库;2700-出库;2855-采购订单;2850-销售订单',
  `stock_transport_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存类型(ENUM)',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `warehouse_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属仓库',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '源单号码:一个订单一个出入库记录可以拆单',
  `stock_bill_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `employee_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '经办人',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '制单人',
  `stock_bill_other_money` decimal(20, 6) NOT NULL DEFAULT 0.000000 COMMENT '其它金额',
  `stock_bill_amount` decimal(20, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '单据金额',
  `stock_bill_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):1-有效; 0-无效',
  `stock_bill_src_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`stock_bill_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品出入库单据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for invoicing_stock_bill_item
-- ----------------------------
DROP TABLE IF EXISTS `invoicing_stock_bill_item`;
CREATE TABLE `invoicing_stock_bill_item`  (
  `stock_bill_item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '源单号码',
  `order_item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单商品表编号',
  `stock_bill_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '货品编号',
  `item_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `bill_item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品数量',
  `bill_item_unit_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '单价',
  `bill_item_subtotal` decimal(16, 2) NOT NULL DEFAULT 0.00 COMMENT '小计',
  `unit_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '单位编号',
  `warehouse_item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存量',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `warehouse_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属仓库',
  `stock_transport_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存类型(ENUM)',
  `bill_item_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `bill_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '业务类别(ENUM):2750-入库;2700-出库',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`stock_bill_item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '出入库单据item表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for invoicing_warehouse_base
-- ----------------------------
DROP TABLE IF EXISTS `invoicing_warehouse_base`;
CREATE TABLE `invoicing_warehouse_base`  (
  `warehouse_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '仓库编号',
  `warehouse_number` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '仓库编号',
  `warehouse_name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '仓库名称',
  `warehouse_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '仓库地址',
  `warehouse_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '仓库类型编号',
  `warehouse_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '仓库是否启用(BOOL):0-禁用;1-启用',
  `warehouse_is_default` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '仓库是否默认(BOOL):0-非默认;1-默认仓库',
  `warehouse_contact` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '仓库联系人',
  `warehouse_phone` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '仓库联系方式',
  `warehouse_delivery_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发货属性(ENUM):1-发货仓库; 0-备货仓库',
  `warehouse_delivery_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备货仓库编号(DOT)如果为发货仓库绑定的备货仓库编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`warehouse_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '仓库仓库表 ' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for invoicing_warehouse_item
-- ----------------------------
DROP TABLE IF EXISTS `invoicing_warehouse_item`;
CREATE TABLE `invoicing_warehouse_item`  (
  `warehouse_item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '仓库商品编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号（冗余）',
  `unit_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品单位id（冗余）',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'SKU',
  `category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类',
  `warehouse_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '仓库编号',
  `warehouse_item_cost` decimal(20, 6) NOT NULL DEFAULT 0.000000 COMMENT '成本',
  `warehouse_item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '数量',
  `warehouse_item_lock_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '锁定库存:实际库存-锁定库存=可售库存	',
  `warehouse_item_quantity_min` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最小库存',
  `warehouse_item_quantity_max` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最大库存',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`warehouse_item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品SKU存放详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for live_user_apply
-- ----------------------------
DROP TABLE IF EXISTS `live_user_apply`;
CREATE TABLE `live_user_apply`  (
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `wechat_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '微信号',
  `apply_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '审核状态(ENUM):1-提交；2-已审核通过；3-审核拒绝；',
  `apply_wx_send_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '向微信发送请求(BOOL):0-否;1-是;',
  `apply_time` bigint(20) UNSIGNED NOT NULL COMMENT '申请时间',
  `apply_verify_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核时间',
  `apply_verify_remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '审核文字',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '主播申请管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for live_user_room
-- ----------------------------
DROP TABLE IF EXISTS `live_user_room`;
CREATE TABLE `live_user_room`  (
  `room_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '房间编号',
  `room_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '房间名称',
  `room_slogan` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '口号',
  `room_pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '房间密码',
  `room_is_microphone` bit(1) NOT NULL DEFAULT b'0' COMMENT '连麦权限(BOOL):0-无;1-有权限',
  `room_pattern` tinyint(4) NOT NULL DEFAULT 1 COMMENT '直播模式(ENUM):1-视频模式;2-纯音模式',
  `room_logo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '房间LOGO',
  `room_poster` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'https://static.shopsuite.cn/xcxfile/appicon/live/video-img4.jpg' COMMENT '房间封面',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `room_state` bit(1) NOT NULL DEFAULT b'0' COMMENT '直播状态(BOOL):0-离线;1-直播中',
  `room_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '禁止状态(BOOL):0-禁止;1-允许',
  `room_stime` bigint(20) NOT NULL DEFAULT 0 COMMENT '直播时间',
  `room_etime` bigint(20) NOT NULL DEFAULT 0 COMMENT '结束时间',
  `room_video` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'https://txmov2.a.yximgs.com/bs2/newWatermark/MTYzNTczOTUwMjk_zh_4.mp4' COMMENT '房间视频',
  `item_ids` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '关联商品(DOT)',
  `item_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '重点商品',
  `live_tag_name` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '直播便签(DOT)',
  `live_virtual_click_num` int(11) NOT NULL DEFAULT 0 COMMENT '虚拟观看数',
  `live_virtual_like_num` int(11) NOT NULL DEFAULT 0 COMMENT '虚拟人气数',
  `live_click_num` int(11) NOT NULL DEFAULT 0 COMMENT '观看数',
  `live_like_num` int(11) NOT NULL DEFAULT 0 COMMENT '人气数',
  `room_recommend` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否推荐(ENUM):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`room_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '主播管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for marketing_activity_base
-- ----------------------------
DROP TABLE IF EXISTS `marketing_activity_base`;
CREATE TABLE `marketing_activity_base`  (
  `activity_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '活动编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `activity_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动名称',
  `activity_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动标题',
  `activity_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动说明',
  `activity_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动类型',
  `activity_starttime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动开始时间',
  `activity_endtime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动结束时间',
  `activity_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭',
  `activity_rule` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{\"requirement\":{\"buy\":{\"item\":[1,2,3],\"subtotal\":\"通过计算修正满足的条件\"}},\"rule\":[{\"total\":100,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}},{\"total\":200,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}}]}',
  `activity_effective_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已经参与数量',
  `activity_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '参与类型(ENUM):1-免费参与;2-积分参与;3-购买参与;4-分享参与',
  `activity_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '活动排序',
  `activity_is_finish` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动是否完成(ENUM):0-未完成;1-已完成;2-已解散(目前用于团购)',
  `activity_is_plus` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否PLUS专享:0-默认所有;1-PLUS专享',
  `activity_use_level` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '使用等级(DOT)',
  `activity_item_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动SKU(DOT):activity_rule中数据冗余',
  `activity_addtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`activity_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `activity_state`) USING BTREE,
  INDEX `activity_type_id`(`activity_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动表-通过插件实现\r\n当为拼团是activity_rule中的group_remain_quantity用于标识拼团剩余需要人数，如果用户登录了，需要查询出activity_groupbooking中参与该团的剩余情况' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for marketing_activity_item
-- ----------------------------
DROP TABLE IF EXISTS `marketing_activity_item`;
CREATE TABLE `marketing_activity_item`  (
  `activity_item_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品表编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `activity_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动类型编号',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类',
  `activity_item_starttime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `activity_item_endtime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
  `activity_item_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '活动价格',
  `activity_item_min_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '购买下限',
  `activity_item_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭',
  `activity_item_recommend` bit(1) NOT NULL DEFAULT b'0' COMMENT '推荐标志(BOOL):0-未推荐;1-已推荐',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`activity_item_id`) USING BTREE,
  INDEX `store_id`(`activity_id`) USING BTREE,
  INDEX `item_id`(`product_id`, `item_id`) USING BTREE,
  INDEX `item_id_2`(`item_id`, `activity_item_starttime`, `activity_item_endtime`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参与活动商品表-用户筛选计算' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for marketing_activity_type
-- ----------------------------
DROP TABLE IF EXISTS `marketing_activity_type`;
CREATE TABLE `marketing_activity_type`  (
  `activity_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '类型编号',
  `activity_type_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `activity_type_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型标题',
  `activity_type_icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型图标',
  `activity_type_introduce` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动介绍',
  `activity_type_prize` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动奖励(JSON):待扩展',
  `activity_type_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动排序',
  `activity_type_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-不启用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`activity_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for o2o_chain_base
-- ----------------------------
DROP TABLE IF EXISTS `o2o_chain_base`;
CREATE TABLE `o2o_chain_base`  (
  `chain_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '门店编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `chain_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '门店名称',
  `chain_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `chain_mobile` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '手机号码',
  `chain_telephone` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `chain_contacter` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `chain_lng` double NOT NULL DEFAULT 0 COMMENT '经度',
  `chain_lat` double NOT NULL DEFAULT 0 COMMENT '纬度',
  `chain_district_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `chain_district_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '地区编号',
  `chain_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '实体店铺详细地址',
  `chain_opening_hours` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '营业时间',
  `chain_close_hours` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '打烊时间',
  `chain_traffic_line` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '交通路线',
  `chain_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '门店图片',
  `chain_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `chain_category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店类型',
  `chain_status` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`chain_id`) USING BTREE,
  UNIQUE INDEX `chain_name`(`chain_name`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '门店表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for o2o_chain_category
-- ----------------------------
DROP TABLE IF EXISTS `o2o_chain_category`;
CREATE TABLE `o2o_chain_category`  (
  `chain_category_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `store_id` int(11) NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `chain_category_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '门店分类名称',
  `chain_category_pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '子级分类',
  `chain_category_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `chain_category_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`chain_category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '门店分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for o2o_chain_item
-- ----------------------------
DROP TABLE IF EXISTS `o2o_chain_item`;
CREATE TABLE `o2o_chain_item`  (
  `chain_item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商品编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商店编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `chain_item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
  `chain_item_unit_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '门店单价',
  `chain_item_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-下架;1-上架',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类',
  `store_category_ids` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分类编号(DOT)',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`chain_item_id`) USING BTREE,
  INDEX `chain_id`(`chain_id`) USING BTREE,
  INDEX `item_id`(`item_id`) USING BTREE,
  INDEX `category_id`(`category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '门店商品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for o2o_chain_user
-- ----------------------------
DROP TABLE IF EXISTS `o2o_chain_user`;
CREATE TABLE `o2o_chain_user`  (
  `chain_user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '门店用户编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属门店',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '店铺编号',
  `rights_group_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组编号(DOT)',
  `chain_user_is_admin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否管理员(BOOL):0-普通员工; 1-管理员',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`chain_user_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE,
  INDEX `chain_id`(`chain_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '门店用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_base_bank
-- ----------------------------
DROP TABLE IF EXISTS `pay_base_bank`;
CREATE TABLE `pay_base_bank`  (
  `bank_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '银行编号',
  `bank_name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行名称',
  `bank_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `bank_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `bank_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):1-启用;0-禁用',
  `settlement_account_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1004 COMMENT '账户类别(ENUM):1001-微信;1002-支付宝;1003-现金;1004-银行',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`bank_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '结算银行记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_card_history
-- ----------------------------
DROP TABLE IF EXISTS `pay_card_history`;
CREATE TABLE `pay_card_history`  (
  `card_history_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '卡片使用记录',
  `card_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡号',
  `card_type_id` smallint(5) UNSIGNED NOT NULL COMMENT '卡片编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `card_history_value` decimal(20, 6) NOT NULL DEFAULT 0.000000 COMMENT '改变额度',
  `user_recharge_card` decimal(20, 6) NOT NULL DEFAULT 0.000000 COMMENT '当前额度',
  `card_history_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '卡牌生成时间',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属订单编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `card_history_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `card_write_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡号',
  `card_is_use` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否已经使用(BOOL):0-未使用;1-已使用',
  `operation_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人员',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`card_history_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '卡片使用记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_card_info
-- ----------------------------
DROP TABLE IF EXISTS `pay_card_info`;
CREATE TABLE `pay_card_info`  (
  `card_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '卡片激活码',
  `card_password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡片密码',
  `card_type_id` smallint(5) UNSIGNED NOT NULL COMMENT '卡片编号',
  `card_fetch_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '领奖时间',
  `card_media_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '媒体编号',
  `server_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '领卡人的服务编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '领卡人账号',
  `card_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '卡牌生成时间',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `card_write_code` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '核销码',
  `card_type_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡片Logo',
  `order_is_paid` smallint(5) UNSIGNED NOT NULL DEFAULT 3010 COMMENT '付款状态(ENUM):3010-未付款;3013-已付款',
  `card_rule` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡片简介',
  `card_type_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `card_type_prize` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型奖品',
  `card_is_effective` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):0-无效;1-有效',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`card_code`) USING BTREE,
  INDEX `card_type_id`(`card_type_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '卡片信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_card_type
-- ----------------------------
DROP TABLE IF EXISTS `pay_card_type`;
CREATE TABLE `pay_card_type`  (
  `card_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '卡片类型编号',
  `card_type_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `card_type_prize` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型奖品',
  `card_type_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型描述',
  `card_type_starttime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效开始时间',
  `card_type_endtime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效结束时间',
  `card_type_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡片Logo',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `item_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '次卡可使用的商品(DOT)',
  `card_use_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用次数或数量',
  `card_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭',
  `pay_type` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '卡片类型(1:充值卡,2次卡,3券码,4权益包)',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`card_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '卡片基础信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_consume_combine
-- ----------------------------
DROP TABLE IF EXISTS `pay_consume_combine`;
CREATE TABLE `pay_consume_combine`  (
  `ctc_id` varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易编号',
  `order_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号(DOT)',
  `ctc_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ctc_id`) USING BTREE,
  INDEX `order_ids`(`order_ids`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '合并支付表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_consume_deposit
-- ----------------------------
DROP TABLE IF EXISTS `pay_consume_deposit`;
CREATE TABLE `pay_consume_deposit`  (
  `deposit_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付流水号',
  `deposit_no` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '商城支付编号',
  `deposit_trade_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易号:支付宝etc',
  `order_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户网站唯一订单号(DOT):合并支付则为多个订单号, 没有创建联合支付交易号',
  `payment_channel_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付渠道',
  `deposit_subject` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `deposit_payment_type` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付; ',
  `deposit_trade_status` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易状态',
  `deposit_seller_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖家户号:支付宝etc',
  `deposit_seller_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖家支付账号',
  `deposit_buyer_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家支付用户号',
  `deposit_buyer_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家支付账号',
  `currency_id` int(10) UNSIGNED NOT NULL DEFAULT 86 COMMENT '货币编号',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '￥' COMMENT '左符号',
  `deposit_total_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '交易金额',
  `deposit_quantity` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '购买数量',
  `deposit_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品单价',
  `deposit_body` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品描述',
  `deposit_is_total_fee_adjust` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否调整总价',
  `deposit_use_coupon` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否使用红包买家',
  `deposit_discount` decimal(6, 3) NOT NULL DEFAULT 0.000 COMMENT '折扣',
  `deposit_notify_time` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '通知时间',
  `deposit_notify_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '通知类型',
  `deposit_notify_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '通知校验编号',
  `deposit_sign_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '签名方式',
  `deposit_sign` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '签名',
  `deposit_extra_param` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '额外参数',
  `deposit_service` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付',
  `deposit_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态:0-默认; 1-接收正确数据处理完逻辑; 9-异常订单',
  `deposit_async` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否同步(BOOL):0-同步; 1-异步回调使用',
  `deposit_review` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收款确认(ENUM):0-未确认;1-已确认;2-驳回',
  `deposit_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否作废(BOOL):1-正常; 2-作废',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺:直接交易起作用',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属门店:直接交易起作用',
  `deposit_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `deposit_payment_voucher` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付凭证',
  `deposit_examine_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核时间',
  `deposit_examine_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核备注',
  PRIMARY KEY (`deposit_id`) USING BTREE,
  UNIQUE INDEX `deposit_no`(`deposit_no`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `order_id`(`order_id`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '支付表-支付回调callback使用-确认付款' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_consume_record
-- ----------------------------
DROP TABLE IF EXISTS `pay_consume_record`;
CREATE TABLE `pay_consume_record`  (
  `consume_record_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '支付流水号',
  `order_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户订单编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用编号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `currency_id` int(10) UNSIGNED NOT NULL DEFAULT 86 COMMENT '货币编号',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '￥' COMMENT '左符号',
  `record_total` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '金额',
  `record_money` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '金额:record_total-佣金',
  `record_commission_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '佣金:平台佣金针对销售收款',
  `record_distribution_commission_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '分销佣金:针对销售收款',
  `record_date` date NOT NULL COMMENT '年-月-日',
  `record_year` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年',
  `record_month` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '月',
  `record_day` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '日',
  `record_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `record_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `record_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '支付时间',
  `trade_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金; 1207-退货付款;1208-退货收款;1209-转账收款',
  `payment_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付; ',
  `payment_channel_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付渠道',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属门店',
  `payment_met_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消费类型(ENUM):1-余额支付; 2-充值卡支付; 3-积分支付; 4-信用支付; 5-红包支付',
  `record_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '状态(BOOL):1-已收款;0-作废',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`consume_record_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `trade_type_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `trade_type_id`) USING BTREE,
  INDEX `order_id`(`order_id`(191), `trade_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '交易明细表-账户收支明细-资金流水表-账户金额变化流水' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_consume_trade
-- ----------------------------
DROP TABLE IF EXISTS `pay_consume_trade`;
CREATE TABLE `pay_consume_trade`  (
  `consume_trade_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '交易订单编号',
  `trade_title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商户订单编号',
  `buyer_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `buyer_store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家是否有店铺',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `seller_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '卖家编号',
  `chain_id` int(10) UNSIGNED NOT NULL COMMENT '门店编号',
  `trade_is_paid` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态',
  `trade_type_id` smallint(5) UNSIGNED NOT NULL COMMENT '交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金;',
  `payment_channel_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付渠道',
  `trade_mode_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '交易模式(ENUM):1-担保交易;  2-直接交易',
  `recharge_level_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值编号',
  `currency_id` int(10) UNSIGNED NOT NULL DEFAULT 86 COMMENT '货币编号',
  `order_id_rows` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '￥' COMMENT '左符号',
  `order_payment_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '总付款额度: trade_payment_amount + trade_payment_money + trade_payment_recharge_card + trade_payment_points',
  `order_commission_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '平台交易佣金',
  `trade_payment_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实付金额:在线支付金额,此为订单默认需要支付额度。',
  `trade_payment_money` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '余额支付',
  `trade_payment_recharge_card` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '充值卡余额支付',
  `trade_payment_points` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '积分支付',
  `trade_payment_sp` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '众宝支付',
  `trade_payment_credit` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '信用支付',
  `trade_payment_redpack` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '红包支付',
  `trade_discount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '折扣优惠',
  `trade_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '总额虚拟:trade_order_amount + trade_discount',
  `trade_adjust_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '手工调整费用店铺优惠',
  `trade_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `trade_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `trade_create_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `trade_paid_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '付款时间',
  `trade_delete` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`consume_trade_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `trade_type_id`) USING BTREE,
  INDEX `buyer_id`(`buyer_id`, `trade_type_id`) USING BTREE,
  INDEX `trade_is_paid`(`trade_is_paid`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `payment_channel_id`(`payment_channel_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '交易订单表-强调唯一订单-充值则先创建充值订单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_consume_withdraw
-- ----------------------------
DROP TABLE IF EXISTS `pay_consume_withdraw`;
CREATE TABLE `pay_consume_withdraw`  (
  `withdraw_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `order_id` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属订单(DOT)',
  `return_id` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '退款单号(DOT)',
  `withdraw_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提现额度',
  `withdraw_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现状态(ENUM):0-申请中;1-提现通过;2-驳回;3-打款完成',
  `withdraw_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `withdraw_bank` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行',
  `withdraw_account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行账户',
  `withdraw_account_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '开户名称',
  `withdraw_fee` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提现手续费',
  `withdraw_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `withdraw_bankflow` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行流水账号',
  `withdraw_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作管理员',
  `withdraw_opertime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作时间',
  `withdraw_mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系手机',
  `withdraw_trans_state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `withdraw_mode` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现方式(ENUM):0-余额提现;1-佣金提现',
  `withdraw_invoice_method` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开票方式(ENUM):0-先票后款; 1-先款后票; 2-不开票;',
  `withdraw_invoice_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发票状态(ENUM):0-待上传; 1-已开票; 2-待审核; 3-已作废; 4-不开票;',
  `withdraw_invoice_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发票备注',
  `withdraw_invoice_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '绑定对应的发票号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`withdraw_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '提现申请表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_distribution_commission
-- ----------------------------
DROP TABLE IF EXISTS `pay_distribution_commission`;
CREATE TABLE `pay_distribution_commission`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `commission_amount` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '佣金总额:历史总额度',
  `commission_directseller_amount_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '销售员佣金',
  `commission_directseller_amount_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '二级销售员',
  `commission_directseller_amount_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '三级销售员',
  `commission_buy_amount_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '推广消费佣金',
  `commission_buy_amount_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '消费佣金',
  `commission_buy_amount_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '消费佣金',
  `commission_click_amount_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '本店流量佣金',
  `commission_click_amount_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '一级流量佣金',
  `commission_click_amount_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '二级流量佣金',
  `commission_reg_amount_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '本店注册佣金',
  `commission_reg_amount_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '一级注册佣金',
  `commission_reg_amount_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '二级注册佣金',
  `commission_settled` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '已经结算佣金',
  `commission_directseller_settled` decimal(15, 0) NOT NULL DEFAULT 0 COMMENT '销售员已经结算',
  `commission_buy_settled` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '推广员已经结算',
  `commission_buy_da` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '区代理收益',
  `commission_buy_ca` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '市代理收益',
  `commission_directseller_da` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '区代理收益',
  `commission_directseller_ca` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '市代理收益',
  `commission_buy_trade_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_buy_trade_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_buy_trade_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_buy_da_trade` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_buy_ca_trade` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_directseller_trade_0` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_directseller_trade_1` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_directseller_trade_2` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_directseller_da_trade` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_directseller_ca_trade` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '交易总额',
  `commission_partner_buy_trade` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '合伙人交易总额',
  `commission_partner_directseller_trade` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '合伙人交易总额',
  `commission_partner_deposit_trade` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '合伙人充值总额',
  `commission_distributor_amount` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '分销商收益',
  `commission_salesperson_amount` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '销售员收益',
  `commission_refund_amount` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '退款总佣金',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广员及收益表-用户赚取汇总' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_distribution_withdraw
-- ----------------------------
DROP TABLE IF EXISTS `pay_distribution_withdraw`;
CREATE TABLE `pay_distribution_withdraw`  (
  `duw_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员支付编号',
  `duw_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提现额度',
  `duw_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否成功(BOOL):0-申请中;1-提现通过',
  `duw_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作管理员',
  `duw_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`duw_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '佣金划转到余额申请表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_payment_type
-- ----------------------------
DROP TABLE IF EXISTS `pay_payment_type`;
CREATE TABLE `pay_payment_type`  (
  `payment_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `payment_type_code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代码名称',
  `payment_type_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付名称',
  `payment_type_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `payment_type_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `payment_type_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`payment_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '支付方式表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_recharge_level
-- ----------------------------
DROP TABLE IF EXISTS `pay_recharge_level`;
CREATE TABLE `pay_recharge_level`  (
  `recharge_level_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `recharge_level_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '额度名称',
  `recharge_level_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `recharge_level_value` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值额度',
  `recharge_level_gift` decimal(14, 2) NOT NULL DEFAULT 0.00 COMMENT '额外赠送',
  `recharge_level_validity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效期:按天计算',
  `recharge_level_rate` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '每日本金利息百分比',
  `recharge_level_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `recharge_level_repeat` tinyint(3) UNSIGNED NOT NULL COMMENT '重复购买(BOOL):0-购买一次;1-不限制',
  `recharge_level_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `recharge_level_duration` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '有效时长单位为年',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`recharge_level_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定额充值表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_trade_type
-- ----------------------------
DROP TABLE IF EXISTS `pay_trade_type`;
CREATE TABLE `pay_trade_type`  (
  `trade_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '状态编号',
  `trade_type_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单状态',
  `trade_type_text` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `trade_type_remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`trade_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '交易类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_bank_card
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_bank_card`;
CREATE TABLE `pay_user_bank_card`  (
  `user_bank_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '卡编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `bank_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '银行编号',
  `bank_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行名称',
  `settlement_account_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1004 COMMENT '账户类别(ENUM):1001-微信;1002-支付宝;1003-现金;1004-银行',
  `user_bank_card_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卡号账户名称',
  `user_bank_card_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行卡卡号',
  `user_bank_card_address` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '开户支行名称',
  `user_bank_card_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '银行卡添加时间',
  `user_bank_card_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行预留手机号',
  `province_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '开户省份编号',
  `city_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '开户城市编号',
  `type_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '卡的类型编号',
  `user_bank_default` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户当前所选默认提现卡',
  `user_bank_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否可用(BOOL):1-启用;0-禁用',
  `user_bank_begin_date` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '余额日期',
  `user_bank_amount_money` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '账户余额',
  `user_intl` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家区号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_bank_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '结算账户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_exp_history
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_exp_history`;
CREATE TABLE `pay_user_exp_history`  (
  `exp_log_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `exp_kind_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型(ENUM):1-获取;2-消耗',
  `exp_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `exp_log_value` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '可用经验',
  `user_exp` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前经验',
  `exp_log_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `exp_log_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `exp_log_date` timestamp(0) NOT NULL COMMENT '获得经验的日期',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`exp_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员经验日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_pay
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_pay`;
CREATE TABLE `pay_user_pay`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `user_pay_passwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付确认密码',
  `user_pay_salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'salt',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户基础信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_points_api_history
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_points_api_history`;
CREATE TABLE `pay_user_points_api_history`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付单编号',
  `offline_order_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用方自己生成的订单号',
  `out_order_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '微信或支付宝产生的订单号',
  `order_status` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单状态，详情参照：订单状态详解',
  `money` decimal(8, 2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
  `seller_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付商户编号',
  `buyer_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消费者编号',
  `payment_channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付渠道, 详情参照：支付渠道详解',
  `subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单简介',
  `processed` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '处理状态(BOOL):1-已处理; 0-未处理',
  `points_log_id` int(10) UNSIGNED NOT NULL COMMENT '对应日志记录',
  `payment_way` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付方式，详情参照：支付方式详解',
  `time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间',
  `total_amount` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单总额',
  `discount_amount` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '优惠金额',
  `pay_time` int(10) UNSIGNED NOT NULL COMMENT '支付时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `buyer_id`(`buyer_id`(191), `processed`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员积分日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_points_history
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_points_history`;
CREATE TABLE `pay_user_points_history`  (
  `points_log_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录编号',
  `points_kind_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型(ENUM):1-获取积分;2-消费积分;',
  `points_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `points_log_points` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '可用积分',
  `user_points` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '当前积分',
  `points_log_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `points_log_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `points_log_date` date NOT NULL COMMENT '积分日期',
  `user_id_other` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '交互会员',
  `points_log_state` bit(1) NOT NULL DEFAULT b'1' COMMENT '领取状态(BOOL):0-未领取;1-已领取',
  `ext_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联单号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`points_log_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `points_type_id`, `points_log_date`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员积分日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_user_resource
-- ----------------------------
DROP TABLE IF EXISTS `pay_user_resource`;
CREATE TABLE `pay_user_resource`  (
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `currency_id` int(10) UNSIGNED NOT NULL DEFAULT 86 COMMENT '货币编号',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '￥' COMMENT '左符号',
  `user_money` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '用户资金',
  `user_money_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '冻结资金:待结算余额',
  `user_settled_amount` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '已结算金额',
  `user_recharge_card` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '充值卡余额',
  `user_recharge_card_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '冻结充值卡:待结算',
  `user_points` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '积分',
  `user_points_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '冻结积分',
  `user_exp` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '平台总经验',
  `user_credit` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '可用信用',
  `user_credit_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '冻结额度',
  `user_credit_used` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '使用信用',
  `user_credit_total` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '信用额度',
  `user_margin` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '保证金',
  `user_redpack` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '红包额度',
  `user_redpack_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '红包冻结额度',
  `user_sp` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '积分2',
  `user_sp_frozen` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '冻结积分2',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户资源表-资金账户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pay_user_resource
-- ----------------------------
INSERT INTO `pay_user_resource` VALUES (10001, 86, '￥', 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 5.000000, 0.000000, 0, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 0, 1001);

-- ----------------------------
-- Table structure for pt_product_ask_base
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_ask_base`;
CREATE TABLE `pt_product_ask_base`  (
  `ask_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '咨询编号',
  `ask_type_id` int(10) UNSIGNED NOT NULL COMMENT '咨询类别编号',
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '商品编号',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '店铺编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `ask_question` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '咨询内容',
  `ask_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提问时间',
  `ask_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '答案',
  `ask_answer_time` timestamp(0) NOT NULL COMMENT '回答时间',
  `ask_answer_user_id` int(11) UNSIGNED ZEROFILL NOT NULL COMMENT '回复用户',
  `ask_answer_user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '回复昵称',
  `ask_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否回复(BOOL):0-未回复;1-已回复',
  `ask_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否现实(BOOL):0-不公开;1-公开',
  `ask_helpful` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ask_id`) USING BTREE,
  INDEX `product_id_2`(`product_id`) USING BTREE,
  INDEX `product_id`(`store_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `ask_question`(`ask_question`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品咨询' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_assist
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_assist`;
CREATE TABLE `pt_product_assist`  (
  `assist_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '辅助属性编号',
  `assist_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '辅助属性名称',
  `type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属类型编号',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号:只在后台快捷定位中起作用',
  `assist_item` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '辅助属性选项(DOT)',
  `assist_format` enum('select','checkbox','radio','text') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'checkbox' COMMENT '显示类型(LIST):select-select;checkbox-checkbox;radio-radio;text-text;',
  `assist_is_search` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否被搜索(BOOL):0-不搜索; 1-搜索',
  `assist_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '辅助属性排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`assist_id`) USING BTREE,
  INDEX `type_id`(`assist_sort`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品辅助属性表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_assist_index
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_assist_index`;
CREATE TABLE `pt_product_assist_index`  (
  `product_assist_index_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'product_id-assist_id',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品公共表编号',
  `assist_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '属性编号',
  `assist_item_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '属性值(DOT)',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_assist_index_id`) USING BTREE,
  UNIQUE INDEX `product_id`(`product_id`, `assist_id`) USING BTREE,
  INDEX `assist_id`(`assist_id`, `assist_item_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品与属性对应表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_assist_item
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_assist_item`;
CREATE TABLE `pt_product_assist_item`  (
  `assist_item_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '选项编号',
  `assist_item_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选项名称',
  `assist_id` mediumint(8) UNSIGNED NOT NULL COMMENT '属性编号',
  `assist_item_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '选项排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`assist_item_id`) USING BTREE,
  INDEX `assist_id`(`assist_id`, `assist_item_sort`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品辅助属性值表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_base
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_base`;
CREATE TABLE `pt_product_base`  (
  `product_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '产品编号',
  `product_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SPU货号:货号',
  `product_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品名称',
  `product_tips` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品卖点:商品广告词',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `product_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品主图',
  `product_video` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品视频 ',
  `transport_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '选择售卖区域:完成售卖区域及运费设置',
  `product_buy_limit` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '每人限购',
  `product_commission_rate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '平台佣金比率',
  `product_minimum_order` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '起订量',
  `product_file` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品文件',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品基础表-SPU表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_brand
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_brand`;
CREATE TABLE `pt_product_brand`  (
  `brand_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '品牌编号',
  `brand_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌名称',
  `brand_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌拼音',
  `brand_initial` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '首字母',
  `brand_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌描述',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类:一级分类即可',
  `brand_show_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '展示方式(ENUM):1-图片; 2-文字  | 在“全部品牌”页面的展示方式，如果设置为“图片”则显示该品牌的“品牌图片标识”，如果设置为“文字”则显示该品牌的“品牌名”',
  `brand_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌LOGO',
  `brand_recommend` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否推荐(BOOL):1-是; 0-否',
  `brand_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):1-启用; 0-禁用',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `brand_apply` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '品牌申请(ENUM):0-申请中; 1-通过 | 申请功能是会员使用，系统后台默认为1',
  `brand_bg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '背景图',
  `brand_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`brand_id`) USING BTREE,
  INDEX `category_id_2`(`category_id`) USING BTREE,
  INDEX `category_id`(`store_id`, `category_id`, `brand_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '品牌表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_category
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_category`;
CREATE TABLE `pt_product_category`  (
  `category_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `category_parent_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类父编号',
  `category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `category_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片',
  `type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '类型编号',
  `industry_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业编号(DOT)',
  `category_commission_rate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '分佣比例',
  `category_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `category_is_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-不显示;1-显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`category_id`) USING BTREE,
  INDEX `type_id`(`type_id`) USING BTREE,
  INDEX `category_order`(`category_sort`) USING BTREE,
  INDEX `category_name`(`category_name`) USING BTREE,
  INDEX `category_parent_id`(`category_parent_id`, `category_sort`, `category_is_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_product_category
-- ----------------------------
INSERT INTO `pt_product_category` VALUES (1001, 0, '默认分类', 'https://shopsuite.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230829/0c2b7826ae2943eca247bc53a4645648.png', 1001, '', 0.00, 50, b'1', 1001);

-- ----------------------------
-- Table structure for pt_product_comment
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_comment`;
CREATE TABLE `pt_product_comment`  (
  `comment_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评价编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `item_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品规格',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `store_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家姓名:user_nickname',
  `comment_points` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '获得积分:冗余，独立表记录',
  `comment_scores` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '评价星级:1-5积分',
  `comment_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评价内容',
  `comment_image` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论上传的图片(DOT)',
  `comment_helpful` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有帮助',
  `comment_nohelpful` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '无帮助',
  `comment_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
  `comment_is_anonymous` bit(1) NOT NULL DEFAULT b'0' COMMENT '匿名评价',
  `comment_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '评价信息的状态(BOOL): 1-正常显示; 0-禁止显示',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `comment_is_reply` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否回复(BOOL): 1-已回复; 0-未回复',
  PRIMARY KEY (`comment_id`) USING BTREE,
  UNIQUE INDEX `order_id`(`order_id`, `item_id`) USING BTREE,
  INDEX `goods_evaluation_ibfk_2`(`user_id`) USING BTREE,
  INDEX `goods_id`(`product_id`) USING BTREE,
  INDEX `item_id`(`item_id`) USING BTREE,
  INDEX `product_id`(`store_id`, `product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品评价表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_comment_helpful
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_comment_helpful`;
CREATE TABLE `pt_product_comment_helpful`  (
  `comment_helpful_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `comment_id` bigint(20) UNSIGNED NOT NULL COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户编号',
  `comment_helpful_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_helpful_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`comment_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品评价有用投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_comment_reply`;
CREATE TABLE `pt_product_comment_reply`  (
  `comment_reply_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论回复编号',
  `comment_id` bigint(20) UNSIGNED NOT NULL COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家评论者姓名',
  `user_id_to` int(10) UNSIGNED NOT NULL COMMENT '回复用户',
  `user_name_to` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '回复用户名称',
  `comment_reply_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复内容',
  `comment_reply_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论回复时间',
  `comment_reply_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '允许显示',
  `comment_reply_isadmin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员评价',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_reply_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`comment_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品评价回复表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_image
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_image`;
CREATE TABLE `pt_product_image`  (
  `product_image_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品图片编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号:product_id-color_id',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `color_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '系统默认颜色规格Id/spec_item_id, 如果没有则一条记录为0',
  `color_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格值',
  `item_image_default` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品主图',
  `item_image_other` varchar(2560) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '副图(DOT)',
  `item_video` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '视频地址',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_image_id`) USING BTREE,
  UNIQUE INDEX `p-color_id`(`product_id`, `color_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品图片表，按照颜色规格' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_index
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_index`;
CREATE TABLE `pt_product_index`  (
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号:定为SPU编号',
  `product_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SPU商家编码:货号',
  `product_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品名称',
  `product_name_index` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称索引关键字(DOT)',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `store_is_open` bit(1) NOT NULL DEFAULT b'1' COMMENT '店铺状态(BOOL):0-关闭;1-运营中',
  `store_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺',
  `store_category_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '店铺分类(DOT)',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类',
  `course_category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '课程分类',
  `type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型编号:冗余检索',
  `product_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存:冗余计算',
  `product_warn_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '预警数量',
  `brand_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '品牌编号',
  `product_service_type_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '售后服务(DOT)',
  `product_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售',
  `product_state_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '违规备注',
  `product_sale_district_ids` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '销售区域(DOT): district_id=1000全部区域',
  `product_verify_id` smallint(5) UNSIGNED NOT NULL DEFAULT 10 COMMENT '商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过',
  `product_verify_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核备注',
  `product_is_invoices` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否开票(BOOL): 1-是; 0-否',
  `product_is_return` bit(1) NOT NULL DEFAULT b'1' COMMENT '允许退换货(BOOL): 1-是; 0-否',
  `product_is_recommend` bit(1) NOT NULL DEFAULT b'0' COMMENT '商品推荐(BOOL):1-是; 0-否',
  `product_stock_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天',
  `kind_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品种类:1201-实物;1202-虚拟',
  `activity_type_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '参与活动(DOT)',
  `contract_type_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消费者保障(DOT):由店铺映射到商品',
  `product_assist_data` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET(\'1\', product_assist_data) ',
  `product_unit_price_min` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '最低单价',
  `product_unit_price_max` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '最高单价',
  `product_unit_points_min` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品积分',
  `product_unit_points_max` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品积分',
  `product_sale_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '销售数量',
  `product_favorite_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数量',
  `product_click` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击数量',
  `product_evaluation_num` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价次数',
  `product_region_district_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属区域(DOT)',
  `product_freight` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '运费:包邮为0',
  `product_tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品标签(DOT)',
  `store_is_selfsupport` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否自营(BOOL):1-自营;0-非自营',
  `product_sp_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '允许分销(BOOL):1-启用分销;0-禁用分销',
  `product_dist_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '三级分销允许分销(BOOL):1-启用分销;0-禁用分销',
  `product_add_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `product_sale_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上架时间:预设上架时间,可以动态修正状态',
  `product_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序:越小越靠前',
  `product_src_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '产品来源编号',
  `market_category_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属商圈(DOT)',
  `store_latitude` double NOT NULL DEFAULT 0 COMMENT '纬度',
  `store_longitude` double NOT NULL DEFAULT 0 COMMENT '经度',
  `product_is_video` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否视频(BOOL):1-有视频;0-无视频',
  `product_transport_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务',
  `product_plus_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用PLUS折扣(BOOL):1-启用;0-禁用',
  `product_is_lock` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改',
  `product_inventory_lock` smallint(5) UNSIGNED NOT NULL DEFAULT 1002 COMMENT '库存锁定(ENUM):1001-下单锁定;1002-支付锁定;',
  `product_from` smallint(5) UNSIGNED NOT NULL DEFAULT 1000 COMMENT '商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;',
  `industry_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业编号(DOT)',
  `product_level_membership` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '等级会员(ENUM):1-默认；2-自定义',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_id`) USING BTREE,
  INDEX `category_id`(`category_id`) USING BTREE,
  INDEX `type_id`(`type_id`) USING BTREE,
  INDEX `product_state_id`(`product_state_id`) USING BTREE,
  INDEX `product_verify_id`(`product_verify_id`) USING BTREE,
  INDEX `store_category_ids`(`store_category_ids`) USING BTREE,
  INDEX `activity_type_ids`(`activity_type_ids`) USING BTREE,
  INDEX `product_name`(`product_name`) USING BTREE,
  INDEX `product_add_time`(`product_add_time`) USING BTREE,
  INDEX `category_id_2`(`category_id`, `product_sale_num`, `product_add_time`) USING BTREE,
  INDEX `category_id_3`(`category_id`, `product_unit_price_min`, `product_add_time`) USING BTREE,
  INDEX `category_id_4`(`category_id`, `product_favorite_num`, `product_add_time`) USING BTREE,
  INDEX `product_sale_num`(`product_sale_num`, `product_add_time`) USING BTREE,
  INDEX `product_unit_price`(`product_unit_price_min`, `product_add_time`) USING BTREE,
  INDEX `product_favorite_num`(`product_favorite_num`, `product_add_time`) USING BTREE,
  INDEX `category_id_5`(`category_id`, `product_sale_num`, `product_add_time`) USING BTREE,
  INDEX `category_id_6`(`category_id`, `product_unit_price_min`, `product_add_time`) USING BTREE,
  INDEX `category_id_7`(`category_id`, `product_favorite_num`, `product_add_time`) USING BTREE,
  INDEX `product_state_id_2`(`product_verify_id`, `product_state_id`, `product_sale_time`) USING BTREE,
  INDEX `product_number`(`product_number`) USING BTREE,
  INDEX `product_assist_data`(`product_assist_data`(191)) USING BTREE,
  INDEX `product_name_index`(`product_name_index`(191)) USING BTREE,
  INDEX `product_name_index_2`(`product_name_index`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品索引表-不读取数据只读主键' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_info
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_info`;
CREATE TABLE `pt_product_info`  (
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `product_assist` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '属性(JSON) - 辅助属性及VAL',
  `product_spec` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格(JSON)-规格、规格值、goods_id  规格不需要全选就可以添加对应数据[{\'id\' : spec_id, \'name\' : spec_name, \'item\':[{\'id\' : spec_item_id, \'name\' : spec_item_name}, {\'id\' : spec_item_id, \'name\' : spec_item_name}]},{\'id\' : spec_id, \'name\' : spec_name, \'item\':[{\'id\' : spec_item_id, \'name\' : spec_item_name}, {\'id\' : spec_item_id, \'name\' : spec_item_name}]}]',
  `product_uniqid` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品SKU(JSON):{\'uniq_id\':[item_id, price, url]}',
  `product_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品描述',
  `product_meta_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Meta Tag 标题',
  `product_meta_description` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Meta Tag 描述',
  `product_meta_keyword` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Meta Tag 关键字',
  `spec_ids` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '规格类别编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_item
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_item`;
CREATE TABLE `pt_product_item`  (
  `item_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品编号-SKU编号',
  `item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '副标题(DOT):SKU名称',
  `item_index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '索引(DOT)',
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `color_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '颜色SKU，规格值',
  `item_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否为默认展示的商品，必须为item_enable',
  `item_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。',
  `item_barcode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条形码',
  `item_pack_spec` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '包装规格',
  `item_cost_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '成本价',
  `item_unit_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品价格',
  `item_market_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '市场价',
  `item_unit_points` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '积分价格',
  `item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品库存',
  `item_quantity_frozen` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品冻结库存',
  `item_warn_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存预警值',
  `item_spec` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '[]' COMMENT '商品规格序列化(JSON):{spec_id:spec_item_id, spec_id:spec_item_id, spec_id:spec_item_id}',
  `spec_item_ids` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品规格值编号',
  `item_enable` smallint(5) UNSIGNED NOT NULL DEFAULT 1002 COMMENT '是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售',
  `item_is_change` bit(1) NOT NULL DEFAULT b'0' COMMENT '被改动(BOOL):0-未改动;1-已改动分销使用',
  `item_weight` decimal(20, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '商品重量:KG',
  `item_volume` decimal(20, 2) NOT NULL DEFAULT 0.00 COMMENT '商品体积:立方米',
  `item_fx_commission` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '微小店分销佣金',
  `item_rebate` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '返利额度',
  `item_src_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '供应商SKU编号',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类',
  `course_category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '课程分类',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `item_level_discount` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品会员等级折扣(JSON)',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`item_id`) USING BTREE,
  INDEX `item_base_ibfk_1`(`product_id`, `color_id`, `item_enable`, `item_is_default`) USING BTREE,
  INDEX `item_src_id`(`item_src_id`, `store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品SKU表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_item_seq
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_item_seq`;
CREATE TABLE `pt_product_item_seq`  (
  `product_item_seq_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '编号:md5product_item_seq_val',
  `product_item_seq_val` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '编号值',
  `item_id` bigint(20) UNSIGNED NOT NULL COMMENT '商品编号-SKU编号',
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_item_seq_id`) USING BTREE,
  INDEX `item_base_ibfk_1`(`product_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'ITEM_ID/SKU唯一编号表。' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_kind
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_kind`;
CREATE TABLE `pt_product_kind`  (
  `kind_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品种类编号',
  `kind_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品种类名称',
  `kind_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品种类备注',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`kind_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品种类表-发布后不能修改  实物商品 （物流发货）  虚拟商品 （无需物流） ' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_log
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_log`;
CREATE TABLE `pt_product_log`  (
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `product_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售',
  `product_state_pre_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售',
  `product_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作备注',
  `operate_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '编辑时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_meta
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_meta`;
CREATE TABLE `pt_product_meta`  (
  `meta_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Meta id:data_id-meta_key',
  `data_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'product_id',
  `meta_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '键',
  `meta_value` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '值',
  `meta_datatype` enum('string','json','number','dot') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'string' COMMENT '数据类型',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`meta_id`) USING BTREE,
  INDEX `user_id`(`data_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`(191)) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品数据扩展表-要求买家留言可以在此处' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_pricing_policy
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_pricing_policy`;
CREATE TABLE `pt_product_pricing_policy`  (
  `policy_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品SPU',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `customer_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '客户',
  `policy_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-否;1-是',
  `policy_price` decimal(20, 2) NOT NULL DEFAULT 0.00 COMMENT '策略价格',
  `policy_discountrate` decimal(8, 6) NOT NULL DEFAULT 0.000000 COMMENT '真实折扣率:policy_price/item_unit_price',
  `policy_quantity_min` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '起订量',
  `policy_quantity_max` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '限订量',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`policy_id`) USING BTREE,
  UNIQUE INDEX `policy_item_id`(`item_id`, `store_id`, `user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '价格策略表-按客户定价' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_pricing_policy_level
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_pricing_policy_level`;
CREATE TABLE `pt_product_pricing_policy_level`  (
  `policy_level_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '策略编号',
  `item_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品SKU',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `customer_level_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '客户等级编号:决定默认折扣率',
  `policy_level_price` decimal(20, 6) NOT NULL DEFAULT 0.000000 COMMENT '策略价格',
  `policy_level_discountrate` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '真实折扣率:由policy_level_price决定',
  `policy_level_quantity_min` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '起订量',
  `policy_level_quantity_max` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '限订量',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`policy_level_id`) USING BTREE,
  UNIQUE INDEX `policy_item_id`(`store_id`, `customer_level_id`, `item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '价格策略表-按客户等级定价' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_pricing_policy_wholesale
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_pricing_policy_wholesale`;
CREATE TABLE `pt_product_pricing_policy_wholesale`  (
  `policy_wholesale_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '策略编号',
  `item_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品SKU',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `item_quantity` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最大数量',
  `policy_wholesale_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '阶梯价格',
  `policy_wholesale_group_price` decimal(10, 2) NOT NULL COMMENT '团购价格',
  `policy_wholesale_discountrate` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '折扣率:policy_wholesale_price决定',
  `policy_wholesale_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`policy_wholesale_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `item_id`, `policy_wholesale_price`) USING BTREE,
  INDEX `item_id`(`item_id`, `policy_wholesale_price`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '价格策略表-阶梯价' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_spec
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_spec`;
CREATE TABLE `pt_product_spec`  (
  `spec_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '规格编号',
  `spec_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格名称',
  `spec_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格注释',
  `spec_format` enum('text','image') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '显示类型(ENUM): text-文字; image-图片',
  `spec_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序:越小越靠前',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号:只在后台快捷定位中起作用',
  `spec_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`spec_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品规格表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_product_spec
-- ----------------------------
INSERT INTO `pt_product_spec` VALUES (1001, '颜色', '', 'image', 0, 1001, b'1', 1001);

-- ----------------------------
-- Table structure for pt_product_spec_item
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_spec_item`;
CREATE TABLE `pt_product_spec_item`  (
  `spec_item_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '规格值编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类编号',
  `spec_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '规格编号',
  `spec_item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格值名称',
  `spec_item_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序:越小越靠前',
  `spec_item_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-不显示;1-显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`spec_item_id`) USING BTREE,
  INDEX `spec_id`(`spec_id`) USING BTREE COMMENT '(null)',
  INDEX `store_id`(`store_id`, `spec_id`, `spec_item_name`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品规格值表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_state
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_state`;
CREATE TABLE `pt_product_state`  (
  `product_state_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售',
  `product_state_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品状态状态',
  `product_state_text_1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品状态',
  `product_state_text_2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品状态',
  `product_state_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_state_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1003 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品状态表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_product_state
-- ----------------------------
INSERT INTO `pt_product_state` VALUES (1000, 'PRODUCT_ILLEGAL', '违规下架禁售', '违规下架禁售', '', 1001);
INSERT INTO `pt_product_state` VALUES (1001, 'PRODUCT_NORMAL', '正常', '正常', '//时间控制的上下架状态，由后台守护进程修改上下架状态', 1001);
INSERT INTO `pt_product_state` VALUES (1002, 'PRODUCT_OFF_THE_SHELF', '下架', '卖家下架、无库存', '', 1001);

-- ----------------------------
-- Table structure for pt_product_tag
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_tag`;
CREATE TABLE `pt_product_tag`  (
  `product_tag_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签编号',
  `product_tag_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `product_tag_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_type
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_type`;
CREATE TABLE `pt_product_type`  (
  `type_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `type_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号-快捷定位功能',
  `spec_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格编号(DOT)',
  `brand_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌编号(DOT)',
  `assist_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '辅助属性(DOT)',
  `type_is_draft` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否草稿(ENUM):1-草稿;0-发布',
  `type_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-内置; 0-非内置',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`type_id`) USING BTREE,
  INDEX `type_is_draft`(`type_is_draft`) USING BTREE,
  INDEX `type_spec_ids`(`spec_ids`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_product_type
-- ----------------------------
INSERT INTO `pt_product_type` VALUES (1001, '默认通用', '', 1001, '', '', '', b'0', b'1', 1001);

-- ----------------------------
-- Table structure for pt_product_unit
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_unit`;
CREATE TABLE `pt_product_unit`  (
  `unit_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品计量单位编号',
  `parent_unit_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品计量单位父编号',
  `unit_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品计量单位名称',
  `unit_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品计量单位注释',
  `unit_rate` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '换算关系',
  `unit_guid` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'gu编号',
  `unit_is_base` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '基本单位',
  `unit_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品计量单位分类编号',
  `unit_is_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(ENUM):0-非内置; 1-系统内置 | 重量和长度单位，系统内置',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`unit_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品计量单位表-提供修改新增类别功能 - 530' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_unit_type
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_unit_type`;
CREATE TABLE `pt_product_unit_type`  (
  `unit_type_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品计量单位分类编号',
  `parent_unit_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品计量单位分类父编号',
  `unit_type_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品计量单位分类名称',
  `unit_type_remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品计量单位分类注释',
  `unit_type_is_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(ENUM):0-非内置; 1-系统内置 | 重量和长度单位，系统内置',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`unit_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品计量单位类别表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_valid_period
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_valid_period`;
CREATE TABLE `pt_product_valid_period`  (
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `product_valid_period` smallint(5) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '有效期:1001-长期有效;1002-自定义有效期;1003-购买起有效时长年单位',
  `product_validity_start` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `product_validity_end` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '失效时间',
  `product_validity_duration` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效时长单位为天',
  `product_valid_type` smallint(5) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '服务类型(ENUM):1001-到店服务;1002-上门服务',
  `product_service_date_flag` bit(1) NOT NULL DEFAULT b'0' COMMENT '填写预约日期(BOOL):0-否;1-是',
  `product_service_contactor_flag` bit(1) NOT NULL DEFAULT b'0' COMMENT '填写联系人(BOOL):0-否;1-是',
  `product_valid_refund_flag` bit(1) NOT NULL DEFAULT b'0' COMMENT '支持过期退款(BOOL):0-否;1-是',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '虚拟商品信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pt_product_verify
-- ----------------------------
DROP TABLE IF EXISTS `pt_product_verify`;
CREATE TABLE `pt_product_verify`  (
  `product_verify_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '审核状态编号',
  `product_verify_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核状态状态',
  `product_verify_text_1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核状态',
  `product_verify_text_2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核状态',
  `product_verify_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_verify_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3003 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品审核信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pt_product_verify
-- ----------------------------
INSERT INTO `pt_product_verify` VALUES (3000, 'PRODUCT_VERIFY_REFUSED', '审核未通过', '审核未通过', '', 1001);
INSERT INTO `pt_product_verify` VALUES (3001, 'PRODUCT_VERIFY_PASSED', '审核通过', '审核通过', '//审核通过', 1001);
INSERT INTO `pt_product_verify` VALUES (3002, 'PRODUCT_VERIFY_WAITING', '审核中', '审核中', '', 1001);

-- ----------------------------
-- Table structure for shop_store_analytics
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_analytics`;
CREATE TABLE `shop_store_analytics`  (
  `store_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺编号',
  `store_click_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺访问量',
  `store_favorite_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数量',
  `store_product_new_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最新商品',
  `store_product_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品数量',
  `store_product_normal_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '在售数量',
  `store_sales_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺销量',
  `store_evaluation_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价次数',
  `store_evaluation_rate` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 5.00 COMMENT '好评率:store_servicecredit/store_evaluation_num/5,使用综合选项服务评价',
  `store_desccredit` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 5.00 COMMENT '描述相符',
  `store_servicecredit` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 5.00 COMMENT '服务评价',
  `store_deliverycredit` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 5.00 COMMENT '物流评价',
  `store_trade_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 5.00 COMMENT '店铺成交总营业额:确认收货',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '店铺统计表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_base
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_base`;
CREATE TABLE `shop_store_base`  (
  `store_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺编号',
  `store_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `store_grade_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺等级',
  `store_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺logo',
  `store_latitude` double NOT NULL DEFAULT 0 COMMENT '纬度',
  `store_longitude` double NOT NULL DEFAULT 0 COMMENT '经度',
  `store_deliver_district_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配送区域(DOT)',
  `store_is_selfsupport` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否自营(ENUM): 1-自营;0-非自营',
  `store_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺',
  `store_is_open` bit(1) NOT NULL DEFAULT b'0' COMMENT '店铺状态(BOOL):0-关闭;  1-运营中',
  `store_category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺分类编号',
  `store_o2o_tags` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '免费服务(DOT)',
  `store_o2o_flag` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否O2O(BOOL):0-否;1-是',
  `store_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 3220 COMMENT '店铺资料信息状态(ENUM):3210-待完善资料; 3220-等待审核; 3230-资料审核没有通过; 3240-资料审核通过,待付款',
  `store_state_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺审核备注',
  `store_payment_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '付款状态(ENUM):0-未付款; 1-已付款待审核; 2-审核通过',
  `store_circle` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属商圈(DOT)',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `store_location` point NOT NULL COMMENT '坐标点',
  PRIMARY KEY (`store_id`) USING BTREE,
  INDEX `store_state_id`(`store_is_open`) USING BTREE,
  INDEX `store_name`(`store_name`) USING BTREE,
  INDEX `store_category_id`(`store_category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '店铺基础信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_category
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_category`;
CREATE TABLE `shop_store_category`  (
  `store_category_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺分类编号',
  `store_category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺分类名称',
  `store_category_deposit` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '保证金数额(元)',
  `store_category_parent_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺分类父编号',
  `store_category_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `store_category_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片',
  `store_category_is_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-不显示;1-显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`store_category_id`) USING BTREE,
  INDEX `category_order`(`store_category_sort`) USING BTREE,
  INDEX `category_name`(`store_category_name`) USING BTREE,
  INDEX `category_parent_id`(`store_category_parent_id`, `store_category_sort`, `store_category_is_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '市场分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_employee
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_employee`;
CREATE TABLE `shop_store_employee`  (
  `employee_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '卖家编号',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '店铺编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `rights_group_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖家组编号',
  `employee_is_admin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否管理员(BOOL):0-普通员工; 1-管理员',
  `employee_login_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间',
  `employee_is_kefu` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否客服(BOOL):0-非客服; 1-客服',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`employee_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `store_id`) USING BTREE,
  INDEX `store_id_2`(`store_id`, `employee_is_admin`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '卖家用户表—公司员工employee-通过user id启用用户中心' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_express_logistics
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_express_logistics`;
CREATE TABLE `shop_store_express_logistics`  (
  `logistics_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '物流编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `logistics_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '物流名称',
  `logistics_pinyin` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '物流拼音',
  `logistics_number` int(10) UNSIGNED NOT NULL COMMENT '公司编号',
  `logistics_state` bit(1) NOT NULL DEFAULT b'0' COMMENT '面单状态(BOOL):1-启用;0-禁用',
  `express_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '快递编号',
  `express_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递名称',
  `logistics_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否为默认(BOOL):1-默认;0-非默认',
  `logistics_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `logistics_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系手机',
  `logistics_contacter` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `logistics_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `logistics_fee` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '物流运费',
  `logistics_is_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):1-启用;0-禁用',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`logistics_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '物流 = shop_store_express' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_grade
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_grade`;
CREATE TABLE `shop_store_grade`  (
  `store_grade_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺等级编号',
  `store_grade_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `store_grade_fee` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '收费标准-收费标准，单：元/年，必须为数字，在会员开通或升级店铺时将显示在前台',
  `store_grade_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '申请说明',
  `store_grade_product_limit` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '可发布商品数 0:无限制',
  `store_grade_album_limit` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '可上传图片数',
  `store_grade_template` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '店铺可选模板',
  `store_grade_function_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '可用附加功能-function_editor_multimedia',
  `store_grade_sort` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '级别-数值越大表明级别越高',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`store_grade_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '店铺等级表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_info
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_info`;
CREATE TABLE `shop_store_info`  (
  `store_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺编号',
  `store_slogan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺口号',
  `store_banner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺banner',
  `store_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `store_tel` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖家电话',
  `store_free_shipping` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '免运费额度',
  `store_template` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'shop1' COMMENT '店铺绑定模板',
  `store_start_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '本次开始时间',
  `store_end_time` timestamp(0) NOT NULL COMMENT '有效期截止时间',
  `store_close_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关闭原因',
  `store_notice` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '购买须知',
  `store_opening_hours` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业时间',
  `store_close_hours` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '打烊时间',
  `store_area` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所在区域',
  `store_district_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '所属地区(DOT)',
  `store_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺详细地址',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `store_mobile_banners` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '移动端banner(JSON)',
  `hot_sale_color` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '热卖背景色',
  `mobile_top_banner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '移动端顶部banner',
  PRIMARY KEY (`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '店铺表-用户可以多店铺-需要分表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_shipping_address
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_shipping_address`;
CREATE TABLE `shop_store_shipping_address`  (
  `ss_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地址编号',
  `ss_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `ss_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `ss_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `ss_telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `ss_contacter` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人(未启用)',
  `ss_postalcode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮编',
  `ss_province_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省编号',
  `ss_province` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '省份',
  `ss_city_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市编号',
  `ss_city` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '市',
  `ss_county_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '县',
  `ss_county` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '县区',
  `ss_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址:不必重复填写地区',
  `ss_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `ss_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '默认地址(ENUM):0-否;1-是',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ss_id`) USING BTREE,
  INDEX `ss_is_default`(`store_id`, `ss_is_default`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '发货地址表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_transport_item
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_transport_item`;
CREATE TABLE `shop_store_transport_item`  (
  `transport_item_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `transport_type_id` mediumint(8) UNSIGNED NOT NULL COMMENT '自定义物流模板编号',
  `transport_item_default_num` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '默认数量',
  `transport_item_default_price` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '默认运费',
  `transport_item_add_num` decimal(6, 2) NOT NULL DEFAULT 1.00 COMMENT '增加数量',
  `transport_item_add_price` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '增加运费',
  `transport_item_city_ids` varchar(6000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '区域城市id(DOT):分区域全国都可售卖使用 * 代替id 价格需要根据重量等等计算物流费用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`transport_item_id`) USING BTREE,
  INDEX `temp_id`(`transport_type_id`) USING BTREE COMMENT '(null)',
  INDEX `transport_item_city_ids`(`transport_item_city_ids`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '自定义物流模板内容表-只处理区域及运费。' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_store_transport_type
-- ----------------------------
DROP TABLE IF EXISTS `shop_store_transport_type`;
CREATE TABLE `shop_store_transport_type`  (
  `transport_type_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '物流及售卖区域编号',
  `transport_type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模板名称',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '門店编号',
  `transport_type_pricing_method` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '计费规则(ENUM):1-按件数;2-按重量;3-按体积',
  `transport_type_freight_free` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '免运费额度',
  `transport_type_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):0-非内置;1-内置',
  `transport_type_free` bit(1) NOT NULL DEFAULT b'0' COMMENT '全免运费(BOOL):0-不全免;1-全免（不限制地区且免运费）',
  `transport_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否默认(BOOL):0-非默认;1-默认',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '编辑时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`transport_type_id`) USING BTREE,
  INDEX `user_id`(`store_id`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '自定义物流运费及售卖区域类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of shop_store_transport_type
-- ----------------------------
INSERT INTO `shop_store_transport_type` VALUES (1001, '全免', 0, 0, 1, 50.00, b'1', b'1', b'0', '2024-06-21 10:03:21', 1001);

-- ----------------------------
-- Table structure for shop_user_favorites_brand
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_favorites_brand`;
CREATE TABLE `shop_user_favorites_brand`  (
  `favorites_brand_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `brand_id` int(10) UNSIGNED NOT NULL COMMENT '品牌编号',
  `favorites_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`favorites_brand_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `brand_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏品牌' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_favorites_item
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_favorites_item`;
CREATE TABLE `shop_user_favorites_item`  (
  `favorites_item_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '收藏编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品SKU编号',
  `product_id` bigint(20) NOT NULL COMMENT '商品SPU编号',
  `favorites_item_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`favorites_item_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏的商品-根据SKU' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_favorites_store
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_favorites_store`;
CREATE TABLE `shop_user_favorites_store`  (
  `favorites_store_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '店铺编号',
  `favorites_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`favorites_store_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏的店铺' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_product_browse
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_product_browse`;
CREATE TABLE `shop_user_product_browse`  (
  `product_browse_id` bigint(20) UNSIGNED NOT NULL COMMENT '浏览编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `browse_date` date NOT NULL COMMENT '浏览日期',
  `browse_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '浏览时间',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_browse_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品浏览历史表-SPU-不应该直接存数据库' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_product_buy
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_product_buy`;
CREATE TABLE `shop_user_product_buy`  (
  `product_buy_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'SKU',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `product_buy_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '购买数量',
  `product_buy_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '浏览时间',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`product_buy_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `product_id`) USING BTREE,
  INDEX `product_id`(`product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户商品购买数量历史表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_product_stock
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_product_stock`;
CREATE TABLE `shop_user_product_stock`  (
  `ps_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'SKU编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员编号',
  `ps_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '购买数量',
  `ps_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '浏览时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ps_id`) USING BTREE,
  INDEX `product_id`(`product_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `product_id`, `item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户商品存货表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_product_stock_history
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_product_stock_history`;
CREATE TABLE `shop_user_product_stock_history`  (
  `upsh_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `upsh_kind_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型(ENUM):1-获取;2-消费;',
  `upsh_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型(ENUM):1-购买获得;2-赠送获得;3-送礼消耗;4-消费消耗;5-管理员操作;',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT 'SPU',
  `item_id` bigint(20) UNSIGNED NOT NULL COMMENT 'SKU',
  `upsh_num` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '操作数量',
  `ps_num` decimal(16, 6) NOT NULL DEFAULT 0.000000 COMMENT '原先数量',
  `upsh_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `upsh_desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `upsh_date` date NOT NULL COMMENT '积分日期',
  `user_id_other` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '交互会员',
  `upsh_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '领取状态(BOOL):0-未领取;1-已领取',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`upsh_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `upsh_type_id`, `upsh_date`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户商品存货记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_search_history
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_search_history`;
CREATE TABLE `shop_user_search_history`  (
  `search_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'md5.userid.keyword',
  `search_keyword` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '搜索关键词',
  `search_char_index` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '拼音索引',
  `search_nums` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '搜索次数',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '搜索用户',
  `search_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`search_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `search_time`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户搜索历史记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_sign
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_sign`;
CREATE TABLE `shop_user_sign`  (
  `sign_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '签到编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `sign_time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '签到时间',
  `sign_day` int(10) UNSIGNED NOT NULL COMMENT '签到天数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`sign_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户签到表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_voucher
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_voucher`;
CREATE TABLE `shop_user_voucher`  (
  `user_voucher_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '代金券编号',
  `activity_id` int(10) UNSIGNED NOT NULL COMMENT '代金券模版编号',
  `activity_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代金券名称',
  `voucher_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1501 COMMENT '代金券状态(ENUM):1501-未用;1502-已用;1503-过期;1504-收回',
  `user_voucher_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '代金券发放日期',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `user_voucher_activetime` timestamp(0) NOT NULL COMMENT '使用时间',
  `voucher_price` decimal(10, 2) NOT NULL COMMENT '优惠券可抵扣价格',
  `voucher_subtotal` decimal(10, 2) NOT NULL COMMENT '使用优惠券的订单金额',
  `voucher_start_date` bigint(20) NOT NULL DEFAULT 0 COMMENT '生效时间',
  `voucher_end_date` bigint(20) NOT NULL DEFAULT 0 COMMENT '失效时间',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺编号',
  `item_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '单品优惠商品编号(DOT)',
  `voucher_type` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '优惠券类型(ENUM): 0-普通优惠券;1-免拼券',
  `writeoff_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '线下活动提货码',
  `activity_rule` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{\"requirement\":{\"buy\":{\"item\":[1,2,3],\"subtotal\":\"通过计算修正满足的条件\"}},\"rule\":[{\"total\":100,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}},{\"total\":200,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}}]}',
  `src_order_item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源订单商品编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `voucher_source` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'normal' COMMENT '领取来源: normal-普通；open-PLUS开通；monthly-PLUS每月发放；manual-PLUS手动领取；',
  `voucher_period` varchar(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '券周期：YYYY-MM',
  `activity_is_plus` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否PLUS专享:0-所有;1-PLUS专享',
  `source_order_sn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源PLUS订单编号',
  PRIMARY KEY (`user_voucher_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE,
  INDEX `activity_id`(`activity_id`) USING BTREE,
  INDEX `idx_activity_user_source`(`activity_id`, `user_id`, `voucher_source`) USING BTREE,
  INDEX `idx_user_month`(`user_id`, `voucher_source`, `voucher_period`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户优惠券表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for shop_user_voucher_num
-- ----------------------------
DROP TABLE IF EXISTS `shop_user_voucher_num`;
CREATE TABLE `shop_user_voucher_num`  (
  `uvn_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '领取编号',
  `activity_id` int(10) UNSIGNED NOT NULL COMMENT '活动编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `uvn_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '领取数量',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`uvn_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `activity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户代金券领取数量表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_advertisement
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_advertisement`;
CREATE TABLE `sns_story_advertisement`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '广告id',
  `adv_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `adv_is_top` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '置顶(BOOL):0-正常;1-置顶',
  `adv_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '广告图片',
  `adv_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否显示:1-是; 0-否',
  `adv_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址链接',
  `adv_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `adv_open_mode` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '打开方式:1-弹出打开; 0-窗内打开。',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_base
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_base`;
CREATE TABLE `sns_story_base`  (
  `story_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `story_index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '搜索索引:检索使用',
  `story_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `story_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容(HTMl)',
  `story_file` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片(DOT)',
  `story_video` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '视频地址',
  `story_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型(ENUM): 1-文字; 2-图片; 3-音乐; 4-视频; 5-商品',
  `story_src_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '转发源',
  `story_time` bigint(20) UNSIGNED NOT NULL COMMENT '添加时间',
  `story_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态(ENUM):0-草稿;1-发布',
  `story_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否删除(BOOL):0-删除;1-有效',
  `story_privacy` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '隐私可见度(ENUM):0-所有人可见; 1-好友可见; 2-仅自己可见',
  `story_comment_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数',
  `story_forward_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '转发数',
  `story_like_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
  `story_collection_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子收藏数',
  `story_forward` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否可以转发(BOOL):0-不可转发;1-可以转发',
  `story_is_top` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '置顶(BOOL):0-正常;1-置顶',
  `story_brower_count` tinyint(4) NOT NULL DEFAULT 0 COMMENT '访问数',
  `item_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品编号(DOT)',
  `story_category_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号',
  `story_tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签(DOT)',
  `story_source` smallint(5) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '数据源(ENUM): 1001-帖子; 1002-商品评论; 1003-商品问答',
  `product_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '商品编号',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`story_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `story_type`, `story_status`, `story_enable`) USING BTREE,
  INDEX `user_id_2`(`user_id`, `story_status`, `story_enable`) USING BTREE,
  INDEX `user_id_3`(`user_id`, `story_source`, `story_status`, `story_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '动态信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_category
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_category`;
CREATE TABLE `sns_story_category`  (
  `story_category_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '版块编号',
  `story_category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块名称',
  `story_category_parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '所属父类',
  `story_category_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块图标',
  `story_category_bg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块背景',
  `story_category_keywords` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块关键词',
  `story_category_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块描述',
  `story_category_count` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版块内容数量',
  `story_category_template` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块模板',
  `story_category_alias` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '版块别名',
  `story_category_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '版块排序',
  `story_category_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(ENUM):0-非内置;1-内置;',
  `story_category_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '板块链接',
  `story_open_mode` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '打开方式:1-弹出打开; 0-窗内打开。',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`story_category_id`) USING BTREE,
  UNIQUE INDEX `story_category_name`(`story_category_name`) USING BTREE,
  INDEX `story_category_parent_id`(`story_category_parent_id`, `story_category_order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '版块' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_collection
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_collection`;
CREATE TABLE `sns_story_collection`  (
  `collection_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '动态信息编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `collection_time` bigint(20) NOT NULL COMMENT '收藏时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`collection_id`) USING BTREE,
  INDEX `story_id`(`story_id`, `user_id`, `collection_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_comment
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_comment`;
CREATE TABLE `sns_story_comment`  (
  `comment_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '原帖编号',
  `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `comment_state` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态(BOOL): 0-正常; 1-屏蔽',
  `to_user_id` int(10) UNSIGNED NOT NULL COMMENT '回复用户编号',
  `comment_like_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `comment_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '添加时间',
  `comment_is_show` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否显示: 0-隐藏; 1-显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `story_id`(`story_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_comment_helpful
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_comment_helpful`;
CREATE TABLE `sns_story_comment_helpful`  (
  `comment_helpful_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `comment_id` int(10) UNSIGNED NOT NULL COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '文章编号',
  `comment_helpful_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_helpful_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`comment_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评价有用投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_comment_reply`;
CREATE TABLE `sns_story_comment_reply`  (
  `comment_reply_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论回复编号',
  `comment_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论编号',
  `comment_reply_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复父编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '所属文章编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论回复编号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复姓名',
  `user_id_to` int(10) UNSIGNED NOT NULL COMMENT '评论回复用户编号',
  `comment_reply_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论回复内容',
  `comment_reply_time` bigint(20) NOT NULL COMMENT '评论回复时间',
  `comment_reply_show_flag` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '问答是否显示',
  `comment_reply_islook` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已读(BOOL):1-已读;2-未读',
  `comment_reply_helpful_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有帮助数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_reply_id`) USING BTREE,
  INDEX `comment_id`(`comment_id`, `comment_reply_show_flag`, `comment_reply_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '问答回复表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_comment_reply_helpful
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_comment_reply_helpful`;
CREATE TABLE `sns_story_comment_reply_helpful`  (
  `comment_reply_helpful_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `comment_reply_id` int(10) UNSIGNED NOT NULL COMMENT '评论编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击用户编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '文章编号',
  `comment_reply_helpful_time` bigint(20) NOT NULL COMMENT '评价时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`comment_reply_helpful_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_1`(`comment_reply_id`) USING BTREE,
  INDEX `goods_evaluation_comment_ibfk_2`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评价有用投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_forward
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_forward`;
CREATE TABLE `sns_story_forward`  (
  `forward_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '被转发帖子编号',
  `story_new_id` int(10) UNSIGNED NOT NULL COMMENT '转发后的新帖编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '原始帖子编号',
  `forward_time` bigint(20) UNSIGNED NOT NULL COMMENT '添加时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`forward_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '转播表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_like
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_like`;
CREATE TABLE `sns_story_like`  (
  `like_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `story_id` int(10) UNSIGNED NOT NULL COMMENT '动态信息编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '会员编号',
  `like_time` bigint(20) NOT NULL COMMENT '收藏时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`like_id`) USING BTREE,
  INDEX `story_id`(`story_id`, `user_id`, `like_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '点赞表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sns_story_timeline
-- ----------------------------
DROP TABLE IF EXISTS `sns_story_timeline`;
CREATE TABLE `sns_story_timeline`  (
  `timeline_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户编号',
  `story_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'story或者productcomment',
  `timeline_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动时间',
  `timeline_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型:1-story;2-商品评论',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`timeline_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `timeline_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户消息时间轴表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_access_history
-- ----------------------------
DROP TABLE IF EXISTS `sys_access_history`;
CREATE TABLE `sys_access_history`  (
  `access_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'SKU编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `access_client_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '唯一客户编号',
  `access_os` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  `access_browser_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器名称',
  `access_browser_version` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '浏览器版本',
  `access_spider` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '搜索引擎',
  `access_country` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `access_province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `access_city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市',
  `access_county` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '区',
  `access_lang` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '语言',
  `access_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问IP',
  `access_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `access_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问地址',
  `access_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '访问时间',
  `access_year` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年',
  `access_month` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '月',
  `access_day` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '日',
  `access_hour` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '时',
  `access_date` date NOT NULL COMMENT '年月日',
  `access_datetime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '时间',
  `access_refer_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `access_refer_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `access_mobile` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否手机',
  `access_pad` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否平板',
  `access_pc` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否PC',
  `access_device` smallint(5) UNSIGNED NOT NULL DEFAULT 3 COMMENT '终端(ENUM):1-Phone;2-Pad;3-Pc',
  `access_type` smallint(5) UNSIGNED NOT NULL DEFAULT 2310 COMMENT '终端来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序',
  `access_from` smallint(5) UNSIGNED NOT NULL DEFAULT 2320 COMMENT '终端来源(ENUM):2320-其它;2321-微信;2322-百度;2323-支付宝;2324-头条',
  `access_data` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求数据',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`access_id`) USING BTREE,
  INDEX `access_lang`(`access_lang`) USING BTREE,
  INDEX `access_country`(`access_country`, `access_province`, `access_city`, `access_county`) USING BTREE,
  INDEX `access_type`(`access_type`) USING BTREE,
  INDEX `access_from`(`access_from`) USING BTREE,
  INDEX `access_os`(`access_os`(191)) USING BTREE,
  INDEX `access_browser_name`(`access_browser_name`(191)) USING BTREE,
  INDEX `access_spider`(`access_spider`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 144 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '访问日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_config_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_config_base`;
CREATE TABLE `sys_config_base`  (
  `config_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置编码',
  `config_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置标题',
  `config_datatype` enum('readonly','number','text','textarea','array','password','radio','checkbox','select','icon','date','datetime','image','images','file','files','ueditor','area') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '配置类型(ENUM):readonly-只读文本;number-数字;text-单行文本;textarea-多行文本;array-数组;password-密码;radio-单选框;checkbox-复选框;select-下拉框;icon-字体图标;date-日期;datetime-时间;image-单张图片;images-多张图片;file-单个文件;files-多个文件;ueditor-富文本编辑器;area-地区选择',
  `config_options` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置项',
  `config_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置值',
  `config_type_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属分类',
  `config_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置注释',
  `config_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 255 COMMENT '配置排序:从小到大',
  `config_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `config_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`config_key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统参数设置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config_base
-- ----------------------------
INSERT INTO `sys_config_base` VALUES ('Access-Control-Allow-Headers', 'security', 'radio', '1:测试一\n2:测试二\n3:测试三', '1:测试一\n2:测试二\n3:测试三', 1425, '', 254, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('Access-Control-Allow-Methods', 'security', 'text', '', '*', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('Access-Control-Allow-Origin', 'security', 'text', '', '*', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('adv_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_app_cert_path', '应用证书路径', 'text', '', '/upload/files/zfb/appCertPublicKey_2021004121611368.crt', 1401, '支付宝应用证书路径', 4, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_app_id', '支付应用Appid', 'text', '', '2021004121611368', 1401, '', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_callback_url', 'connect', 'readonly', '', 'https://test.suteshop.com/account/api/connect/alipay.php', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_cert_path', '支付宝证书路径', 'text', '', '/upload/files/zfb/alipayCertPublicKey_RSA2.crt', 1401, '支付宝证书路径', 5, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_enable', '是否启用', 'radio', '0:禁用|1:启用', '1', 1401, '登录支付宝商家(地址：https://b.alipay.com，需要配置ip白名单以及回调地址回调地址：https://demo.shopsuite.cn)', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_logo', '支付宝图标', 'image', '', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231017/fae091baa9ea40e0b79e76bc0273aa3e.png', 1401, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_root_cert_path', '支付宝证书根路径', 'text', '', '/upload/files/zfb/alipayRootCert.crt', 1401, '支付宝证书根路径', 6, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_rsa_private_key', '支付私钥', 'text', '', '', 1401, '支付私钥（开放平台助手生成，RSA2 PKCS1格式）', 3, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('alipay_rsa_public_key', '支付公钥', 'text', '', '', 1401, '支付公钥（开放平台网站支付宝公钥，不是应用公钥，RSA2格式）', 2, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_access_key_id', '阿里云accessId', 'text', '', '', 35, '', 13, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_access_key_secret', '阿里云 accessSecret', 'text', '', '', 35, '', 13, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_bucket', '阿里云bucket', 'text', '', 'kuteshop', 35, '', 234, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_default_dir', '上传目录', 'readonly', '', 'modulithshop', 35, '上传目录', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_endpoint', '阿里云endpoint', 'text', '', 'oss-accelerate.aliyuncs.com', 35, '', 14, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_record_callback', '阿里云直播回放', 'text', '', '/mobile/sns/Live/aliRecordCallback', 35, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('aliyun_regionid', '阿里云regionid', 'text', '', 'cn-shanghai', 35, '', 11, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('apple_status', 'connect', 'text', '', '1', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_android_last_note', 'app', 'text', '', '1、修复bug \\n 2、机型适配', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_android_last_url', 'app', 'text', '', 'https://p.shopsuite.cn/zhongcaiedu/Android.apk', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_android_last_version', 'app', 'text', '', '3.0.3', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_ios_last_note', 'app', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_ios_last_url', 'app', 'text', '', 'https://apps.apple.com/cn/app/%E9%9A%8F%E5%95%86%E4%B9%90%E8%B4%AD/id1512547393', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_ios_last_version', 'app', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('app_member_center', '', 'text', '', '{\"page_type\":3,\"page_tpl\":0,\"Id\":100,\"PageTitle\":\"个人中心\",\"AppId\":8,\"PageCode\":\"{\\\"type\\\":2,\\\"list\\\":[{\\\"id\\\":1,\\\"name\\\":\\\"我的拼团\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#DB384C\\\",\\\"icon\\\":\\\"icon-gouwu\\\",\\\"url\\\":\\\"/activity/fightgroup/order\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"FightGrp\\\"},{\\\"id\\\":36,\\\"name\\\":\\\"售后服务\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#44afa4\\\",\\\"icon\\\":\\\"zc zc-tuihuanhuo\\\",\\\"url\\\":\\\"/member/member/returnlist\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"service\\\"},{\\\"id\\\":4,\\\"name\\\":\\\"我的砍价\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#ffc333\\\",\\\"icon\\\":\\\"icon-kanjia\\\",\\\"url\\\":\\\"/activity/cutprice/userlist\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"CutPrice\\\"},{\\\"id\\\":44,\\\"name\\\":\\\"签到\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#ffc333\\\",\\\"icon\\\":\\\"icon-edit\\\",\\\"url\\\":\\\"/member/member/sign\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"MemSign\\\"},{\\\"id\\\":6,\\\"name\\\":\\\"会员中心\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#ffc333\\\",\\\"icon\\\":\\\"icon-zuanshi\\\",\\\"url\\\":\\\"/member/member/task\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"MemGrade\\\"},{\\\"id\\\":107,\\\"name\\\":\\\"商品收藏\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#56ABE4\\\",\\\"icon\\\":\\\"icon-liwu\\\",\\\"url\\\":\\\"/member/member/favorites\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"FavProd\\\"},{\\\"id\\\":108,\\\"name\\\":\\\"我的足迹\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#56ABE4\\\",\\\"icon\\\":\\\"zc zc-zuji\\\",\\\"url\\\":\\\"/member/member/browse\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"FavProd\\\"},{\\\"id\\\":8,\\\"name\\\":\\\"收货地址\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#1BC2A6\\\",\\\"icon\\\":\\\"icon-shouhuodizhi\\\",\\\"url\\\":\\\"/member/address/list\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"UserAddress\\\"},{\\\"id\\\":120,\\\"name\\\":\\\"开票信息\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#1BC2A6\\\",\\\"icon\\\":\\\"zc-caiwukaipiao\\\",\\\"url\\\":\\\"/member/invoice/list\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"UserInvoice\\\"},{\\\"id\\\":120,\\\"name\\\":\\\"我的发票\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#1BC2A6\\\",\\\"icon\\\":\\\"zc-kaipiao\\\",\\\"url\\\":\\\"/member/invoice/order\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"OrderInvoice\\\"},{\\\"id\\\":21,\\\"name\\\":\\\"推广中心\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#327eac\\\",\\\"icon\\\":\\\"zc zc-fenxiao\\\",\\\"url\\\":\\\"/member/fans/index\\\",\\\"isShow\\\":false,\\\"FeatureKey\\\":\\\"fenxiao\\\"},{\\\"id\\\":32,\\\"name\\\":\\\"帮助\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#ac8dd5\\\",\\\"icon\\\":\\\"zc zc-bangzhu\\\",\\\"url\\\":\\\"/pagesub/article/list\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"Help\\\"},{\\\"id\\\":11,\\\"name\\\":\\\"清除缓存\\\",\\\"cat\\\":1,\\\"color\\\":\\\"#DB384C\\\",\\\"icon\\\":\\\"zc zc-qingchuhuancun\\\",\\\"url\\\":\\\"\\\",\\\"isShow\\\":true,\\\"FeatureKey\\\":\\\"CleanCacheKey\\\"}],\\\"titleType\\\":1}\"}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_content_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_content_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_content_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_keywords', 'seo', 'text', '', '{keyword}文章', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('article_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('baidu_app_id', 'ai', 'text', '', '18108099', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('baidu_client_ak', '浏览器端访问应用（AK）', 'text', '', 'h7rkDvchP7iOJBnX1OccQNL66G5aZBh6', 1433, '', 4, b'1', b'0', 1001);
INSERT INTO `sys_config_base` VALUES ('baidu_client_id', '浏览器端应用编号', 'text', '', '', 1433, '百度地图应用编号', 3, b'1', b'0', 1001);
INSERT INTO `sys_config_base` VALUES ('baidu_secret_id', 'ai', 'text', '', 'su3NxVz3yaBd698vugM9Fmzv', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('baidu_secret_key', 'ai', 'text', '', 'uaW4p1AXHqPQRLqOB3PnfrPcAA8uo7yT', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_content_description', 'seo', 'text', '', 'n{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_content_keywords', 'seo', 'text', '', '123123', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_content_title', 'seo', 'text', '', 'e{sitename}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('brand_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('cache_enable', '启用缓存', 'radio', '0:关闭|1:开启', '0', 1, '', 27, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('cache_expire', '缓存时长', 'number', '', '0', 1, '配置全局缓存时间（秒），默认留空为永久缓存', 27, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('cancellation_agreement', '注销协议', 'ueditor', '', '<p>平台注销协议</p>', 32, '', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('category_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('category_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('category_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('chain_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('chat_global', '默认连接', 'radio', '0:禁用|1:启用', '0', 1423, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('circle_used_protocols_description', 'protocols', 'text', '', '<p>&nbsp; &nbsp;请您仔细阅读本须知的全部内容（特别是以<strong>粗体下划线</strong>标注的内容）。如果您不同意本须知的任意内容，您应当停止使用本产品。</p>\r\n\r\n<p>1、&ldquo;圈子&rdquo;是本公司开设的一个供商城用户（以下简称&ldquo;用户&rdquo;或&ldquo;您&rdquo;）交流购物体验等信息的网络社区。<strong>您使用&ldquo;圈子&rdquo;产品需遵守本须知，并遵守本公司公布的操作流程和规则。</strong></p>\r\n\r\n<p>2、&ldquo;圈子&rdquo;产品的功能和产品提供方式由本公司自行决定，后续本公司可能调整产品名称和产品运行的域名等。<strong>本须知适用于&ldquo;圈子&rdquo;产品的调整、改进版本和附加功能。</strong></p>\r\n\r\n<p>3、您可以通过本产品创建网络关系圈子，其他感兴趣的用户可以加入您创建的圈子。您应当遵守任何适用的法律之规定，并自觉尊重和维护其他参与者的合法权利。您不得以任何形式开展违法活动、侵犯他人合法权益、损害本公司或其他公司的合法利益，否则您需为此自行承担法律责任。<strong>您同意本公司无需为产品使用者的违法或侵权等行为承担任何责任。</strong></p>\r\n\r\n<p>4、您同意并保证通过本产品上传、发布的文字、图片等全部信息素材符合相关法律的规定。您保证素材内容以及素材所含链接指向的内容的合法性和正当性，不侵犯他人的肖像权、名誉权、知识产权、隐私权等合法权益，也不会侵犯法人或其他团体的商业秘密等合法权益。</p>\r\n\r\n<p>5、<strong>您使用本产品可能需要提供关于您的个人资料、肖像、联系方式等个人信息。您了解并同意，在使用本产品过程中关于您的个人信息可能会通过网络等渠道进行传播。</strong></p>\r\n\r\n<p>6、您通过本产品上传、发布素材，即意味着<strong>您同意向本公司提供免费的、永久性的、不可撤销的权利和许可，使本公司可以在全球范围内复制、发行、展示、演绎和通过信息网络等渠道使用您上传的素材和信息</strong>，例如将您提供的图片发布于活动页面或平面媒体中。</p>\r\n\r\n<p>7、本公司无法事先了解您上传素材的真实性和合法性。如您上传的素材被发现不适宜展示或遭受他人举报或投诉的，本公司有权立即删除或进行屏蔽从而停止该素材的继续传播。<strong>如果您违反本须知的内容、有关协议或规则等的，本公司有权删除相关素材并有权拒绝您继续使用产品，届时您无权要求本公司进行补偿或赔偿。</strong></p>\r\n\r\n<p>8、您使用本产品应同时遵守《用户服务协议》、本公司公布的各项规则以及本公司发布的关于本产品的特别规则和制度。</p>\r\n\r\n<p>9、第三方可能通过分公司其他产品或本产品知悉并使用您上传的素材、个人信息或进而侵犯您的合法权利。<strong>本公司提醒您注意和谨防网络诈骗以及其他可能对您不利的行动和信息，但本公司对第三方的侵权、违法等行为不承担赔偿等法律责任。</strong>您承诺合法使用并善待其他用户上传的素材和信息。</p>\r\n\r\n<p>10、您应自行对上传的素材进行备份。本公司可能按照本须知删除或屏蔽素材，相关系统亦可能遭受网络攻击或网络故障，类似或其他情况均可能使您上传的素材丢失或故障，对此本公司将尽力避免但不做任何保证。</p>\r\n\r\n<p>11、<strong>如您因使用本产品与本公司发生纠纷的，您同意由本公司住所地人民法院管辖审理。</strong></p>\r\n\r\n<p>12、<strong>本公司有权更新、修改本须知以及有关规则、流程等相关文件的内容，本公司在法律允许的范围内负责对本须知进行说明和解释</strong>。如您对修改存有异议，您有权选择不再继续使用本产品，但您在此前的行为仍受本须知以及相关文件最新的修改版本的约束。</p>\r\n', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('client_cache', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('client_cache_expire', 'plantform', 'radio', '0:禁用|1:启用', '3600', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('closed_reason', '关闭原因', 'textarea', '', '维护', 1, '', 240, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_encryption', 'secure', 'text', '', '2WRZC8FkF8PT7376IX2snJNbuAmzcNCPVwSEYD0tYC14hbIvGIRFxPRwmCVckEaRZSztIE6yGz86gKUD1Q8VeqtHUJH9zVMad9EWUFxjspcPfJprrpiYHXrW7qb2wHSCsRH9RAAJT32QYwVuIbHs97ZUgRtK04eKsi9DzPr0HmfsAi2UPsvpEZy6oRfSTLlTpXeGrvUfbIqnNGSxjiQD6IjAJBxG9zlpB0onZzoYeIMdM5mMXntmc1yX3qs6NsvaNcBCCPo6Trp5jOuww2BKPWD6cNgiPCMlr8zw979THypkqy9acj3LBtjP1K2Q8DAdRUB7J2vqAs3IoiebWe1N7RYgyWhiBwQU0g61nYUvNF7HsaKA4uySjIQ5wB7ay7kQYmQs4cA29aKoDgbuJKdIhSrnsa64Aj5z1ciU0IqZgEE90fSXIRiXYP6daeJm4tNSi43fhGoVSMW3WP5RJ6ZRjxNVeiUO5bhdeaSmlGs6gh4m6aeVR8NiPkjR3sa8DPHP17XWl8S61S2HA2HWNKne3Plv38NvWBuFXewYqWrJFR2W1XJepwtNEOuTWt18pbDYROZ0HvyDMJInW1fdQeLmGlcFLlQJ5qnEMD42AlUap1Jfp4jI3JSllPvyl9yL2Slp0xNuLPhjghg2txJ5k9FECBnmLVuiigvWFCtXJxR6bvebgpbW2BMRapOjeFPeB4hPFbk5FS3aZXvZ4rqjHHtBfwBNBUmM70Y8h8Iwx52LpJGJgtXztdVpccITCy9PfTRT806rkX3aAOFPtsMGn3dRTpIJkygR0Nd7oC4wHQw8eOZ1OEeAjBu1ygUDcLfr1RyZAZFYWLiN3LjrnLSjJOITUnxGOVUspBHwspar9CghmX34CItEje3ziyApMu7dS3HsY67yUZItcJwsgD40dBSS8JxZCm1ZSaOHWMJcbyn2ud6hOyuRzXLIkJQa6opYnfCWUOcvrRZdZuV9ILD0P2YMtpmjmsfmtbFCjKZNy0LsdOVABCA5YrkvSVWTEPdkjYQK', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_error_display', 'error_reporting', 'text', '', 'on', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_error_filename', 'error_reporting', 'text', '', 'error.log', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_error_log', 'error_reporting', 'text', '', 'on', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_maintenance', 'secure', 'text', '', '0', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_meta_description', 'meta', 'text', '', 'QianyiShop  中文专业版，成都光大网络科技有限公司出品，产品官方网站  ', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_meta_keyword', 'meta', 'text', '', 'QianyiShop  中文专业版，成都光大网络科技有限公司出品，产品官方网站 ', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_meta_title', 'meta', 'text', '', 'QianyiShop 中文专业版', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_robots', 'secure', 'text', '', 'abot\r\ndbot\r\nebot\r\nhbot\r\nkbot\r\nlbot\r\nmbot\r\nnbot\r\nobot\r\npbot\r\nrbot\r\nsbot\r\ntbot\r\nvbot\r\nybot\r\nzbot\r\nbot.\r\nbot/\r\n_bot\r\n.bot\r\n/bot\r\n-bot\r\n:bot\r\n(bot\r\ncrawl\r\nslurp\r\nspider\r\nseek\r\naccoona\r\nacoon\r\nadressendeutschland\r\nah-ha.com\r\nahoy\r\naltavista\r\nananzi\r\nanthill\r\nappie\r\narachnophilia\r\narale\r\naraneo\r\naranha\r\narchitext\r\naretha\r\narks\r\nasterias\r\natlocal\r\natn\r\natomz\r\naugurfind\r\nbackrub\r\nbannana_bot\r\nbaypup\r\nbdfetch\r\nbig brother\r\nbiglotron\r\nbjaaland\r\nblackwidow\r\nblaiz\r\nblog\r\nblo.\r\nbloodhound\r\nboitho\r\nbooch\r\nbradley\r\nbutterfly\r\ncalif\r\ncassandra\r\nccubee\r\ncfetch\r\ncharlotte\r\nchurl\r\ncienciaficcion\r\ncmc\r\ncollective\r\ncomagent\r\ncombine\r\ncomputingsite\r\ncsci\r\ncurl\r\ncusco\r\ndaumoa\r\ndeepindex\r\ndelorie\r\ndepspid\r\ndeweb\r\ndie blinde kuh\r\ndigger\r\nditto\r\ndmoz\r\ndocomo\r\ndownload express\r\ndtaagent\r\ndwcp\r\nebiness\r\nebingbong\r\ne-collector\r\nejupiter\r\nemacs-w3 search engine\r\nesther\r\nevliya celebi\r\nezresult\r\nfalcon\r\nfelix ide\r\nferret\r\nfetchrover\r\nfido\r\nfindlinks\r\nfireball\r\nfish search\r\nfouineur\r\nfunnelweb\r\ngazz\r\ngcreep\r\ngenieknows\r\ngetterroboplus\r\ngeturl\r\nglx\r\ngoforit\r\ngolem\r\ngrabber\r\ngrapnel\r\ngralon\r\ngriffon\r\ngromit\r\ngrub\r\ngulliver\r\nhamahakki\r\nharvest\r\nhavindex\r\nhelix\r\nheritrix\r\nhku www octopus\r\nhomerweb\r\nhtdig\r\nhtml index\r\nhtml_analyzer\r\nhtmlgobble\r\nhubater\r\nhyper-decontextualizer\r\nia_archiver\r\nibm_planetwide\r\nichiro\r\niconsurf\r\niltrovatore\r\nimage.kapsi.net\r\nimagelock\r\nincywincy\r\nindexer\r\ninfobee\r\ninformant\r\ningrid\r\ninktomisearch.com\r\ninspector web\r\nintelliagent\r\ninternet shinchakubin\r\nip3000\r\niron33\r\nisraeli-search\r\nivia\r\njack\r\njakarta\r\njavabee\r\njetbot\r\njumpstation\r\nkatipo\r\nkdd-explorer\r\nkilroy\r\nknowledge\r\nkototoi\r\nkretrieve\r\nlabelgrabber\r\nlachesis\r\nlarbin\r\nlegs\r\nlibwww\r\nlinkalarm\r\nlink validator\r\nlinkscan\r\nlockon\r\nlwp\r\nlycos\r\nmagpie\r\nmantraagent\r\nmapoftheinternet\r\nmarvin/\r\nmattie\r\nmediafox\r\nmediapartners\r\nmercator\r\nmerzscope\r\nmicrosoft url control\r\nminirank\r\nmiva\r\nmj12\r\nmnogosearch\r\nmoget\r\nmonster\r\nmoose\r\nmotor\r\nmultitext\r\nmuncher\r\nmuscatferret\r\nmwd.search\r\nmyweb\r\nnajdi\r\nnameprotect\r\nnationaldirectory\r\nnazilla\r\nncsa beta\r\nnec-meshexplorer\r\nnederland.zoek\r\nnetcarta webmap engine\r\nnetmechanic\r\nnetresearchserver\r\nnetscoop\r\nnewscan-online\r\nnhse\r\nnokia6682/\r\nnomad\r\nnoyona\r\nnutch\r\nnzexplorer\r\nobjectssearch\r\noccam\r\nomni\r\nopen text\r\nopenfind\r\nopenintelligencedata\r\norb search\r\nosis-project\r\npack rat\r\npageboy\r\npagebull\r\npage_verifier\r\npanscient\r\nparasite\r\npartnersite\r\npatric\r\npear.\r\npegasus\r\nperegrinator\r\npgp key agent\r\nphantom\r\nphpdig\r\npicosearch\r\npiltdownman\r\npimptrain\r\npinpoint\r\npioneer\r\npiranha\r\nplumtreewebaccessor\r\npogodak\r\npoirot\r\npompos\r\npoppelsdorf\r\npoppi\r\npopular iconoclast\r\npsycheclone\r\npublisher\r\npython\r\nrambler\r\nraven search\r\nroach\r\nroad runner\r\nroadhouse\r\nrobbie\r\nrobofox\r\nrobozilla\r\nrules\r\nsalty\r\nsbider\r\nscooter\r\nscoutjet\r\nscrubby\r\nsearch.\r\nsearchprocess\r\nsemanticdiscovery\r\nsenrigan\r\nsg-scout\r\nshai\'hulud\r\nshark\r\nshopwiki\r\nsidewinder\r\nsift\r\nsilk\r\nsimmany\r\nsite searcher\r\nsite valet\r\nsitetech-rover\r\nskymob.com\r\nsleek\r\nsmartwit\r\nsna-\r\nsnappy\r\nsnooper\r\nsohu\r\nspeedfind\r\nsphere\r\nsphider\r\nspinner\r\nspyder\r\nsteeler/\r\nsuke\r\nsuntek\r\nsupersnooper\r\nsurfnomore\r\nsven\r\nsygol\r\nszukacz\r\ntach black widow\r\ntarantula\r\ntempleton\r\n/teoma\r\nt-h-u-n-d-e-r-s-t-o-n-e\r\ntheophrastus\r\ntitan\r\ntitin\r\ntkwww\r\ntoutatis\r\nt-rex\r\ntutorgig\r\ntwiceler\r\ntwisted\r\nucsd\r\nudmsearch\r\nurl check\r\nupdated\r\nvagabondo\r\nvalkyrie\r\nverticrawl\r\nvictoria\r\nvision-search\r\nvolcano\r\nvoyager/\r\nvoyager-hc\r\nw3c_validator\r\nw3m2\r\nw3mir\r\nwalker\r\nwallpaper\r\nwanderer\r\nwauuu\r\nwavefire\r\nweb core\r\nweb hopper\r\nweb wombat\r\nwebbandit\r\nwebcatcher\r\nwebcopy\r\nwebfoot\r\nweblayers\r\nweblinker\r\nweblog monitor\r\nwebmirror\r\nwebmonkey\r\nwebquest\r\nwebreaper\r\nwebsitepulse\r\nwebsnarf\r\nwebstolperer\r\nwebvac\r\nwebwalk\r\nwebwatch\r\nwebwombat\r\nwebzinger\r\nwhizbang\r\nwhowhere\r\nwild ferret\r\nworldlight\r\nwwwc\r\nwwwster\r\nxenu\r\nxget\r\nxift\r\nxirq\r\nyandex\r\nyanga\r\nyeti\r\nyodao\r\nzao\r\nzippp\r\nzyborg', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('config_secure', 'secure', 'text', '', '0', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('Content-Security-Policy', 'security', 'text', '', '', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('copyright', '版权信息', 'text', '', '© Copyright 2012-2022 MallSuite版权所有', 1, '', 22, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('credit_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('date_format', '日期格式', 'text', '', 'Y-m-d', 1, '', 100, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('default_image', '默认图片', 'image', '1', 'http://mallsuite.oss-accelerate.aliyuncs.com/mall/images/media/store/1266/20220707/d88059d3cd4e4d158b53b0f61daff241.png', 1, '', 13, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('default_shipping_district', '默认区域', 'area', '', '310000,310100,310104', 39, '消费者默认收货区域，商品详情页显示及O2O默认城市显示', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('dinghuo_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('distribution_protocols_description', 'protocols', 'text', '', '<p>121212</p>', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('distribution_user_auto_active', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('donwload_app_enable', 'frontend', 'text', '', 'true', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('edu_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_addr', 'email', 'text', '', 'chy254228@163.com', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_debug', 'email', 'text', '', '2', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_fromname', 'email', 'text', '', '随商', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_host', 'email', 'text', '', 'smtp.163.com', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_id', 'email', 'text', '', 'chy254228', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_pass', 'email', 'text', '', 'RVIWCIHTKNTDKVAY', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_port', 'email', 'text', '', '465', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_secure', 'email', 'text', '', '2', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('email_test', 'email', 'text', '', 'chenhaiyin25@163.com', 12, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('esearch_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('evaluation_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('exp_consume_max', '每单最多增加', 'number', '', '500', 1424, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('exp_consume_rate', '消费额与增加比例', 'number', '', '0.1', 1424, '消费额与赠送比例。 获得额度 = 消费额度 * 赠送比例。 如果填写100， 用户消费1元，则赠送100经验值！', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('exp_evaluate_good', '评论增加', 'number', '', '50', 1424, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('exp_login', '登录增加', 'number', '', '50', 1424, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('exp_reg', '注册增加', 'number', '', '0', 1424, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('facebook_app_id', 'connect', 'text', '', '1aaa', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('facebook_app_key', 'connect', 'text', '', '1aa', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('facebook_status', 'connect', 'text', '', '0', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('fb_app_id', 'connect', 'text', '', 'afdfdaf', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('fb_app_key', 'connect', 'text', '', 'fdfadf', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('fb_callback_url', 'connect', 'text', '', '', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('fb_status', 'connect', 'text', '', '2', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('friendly_url_flag', '系统参数', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('frontend_enable', 'frontend', 'text', '', 'mp_alipay,mp_baidu,wechat,mp_weixin,app,h5', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('fx_level_config', 'distribution', 'text', '', '[{\"user_level_fans_store\":0,\"user_level_rate\":100,\"user_level_fans_user\":0,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2021-07-30 17:57:16\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v1\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon1.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":1,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1001,\"plantform_fx_cps_rate_1\":\"10\",\"plantform_fx_cps_rate_2\":\"5\",\"plantform_fx_cps_rate_3\":\"5\"},{\"user_level_fans_store\":0,\"user_level_rate\":100,\"user_level_fans_user\":0,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2021-12-29 11:32:16\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v2\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon2.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":10,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1002,\"plantform_fx_cps_rate_1\":\"10\",\"plantform_fx_cps_rate_2\":\"5\",\"plantform_fx_cps_rate_3\":\"5\"},{\"user_level_fans_store\":0,\"user_level_rate\":90,\"user_level_fans_user\":0,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2021-07-30 18:37:29\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v3\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon3.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":110,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1003,\"plantform_fx_cps_rate_1\":\"15\",\"plantform_fx_cps_rate_2\":\"10\",\"plantform_fx_cps_rate_3\":\"5\"},{\"user_level_fans_store\":0,\"user_level_rate\":100,\"user_level_fans_user\":0,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2021-04-09 15:22:01\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v4\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon4.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":11110,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1004,\"plantform_fx_cps_rate_1\":\"15\",\"plantform_fx_cps_rate_2\":\"10\",\"plantform_fx_cps_rate_3\":\"5\"},{\"user_level_fans_store\":0,\"user_level_rate\":100,\"user_level_fans_user\":0,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2019-12-24 19:50:47\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v5\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon5.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":1111000,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1005,\"plantform_fx_cps_rate_1\":\"15\",\"plantform_fx_cps_rate_2\":\"10\",\"plantform_fx_cps_rate_3\":\"5\"},{\"user_level_fans_store\":0,\"user_level_rate\":100,\"user_level_fans_user\":100,\"user_level_exp_reduction\":0,\"user_level_annual_fee\":0,\"user_level_time\":\"2022-04-24 12:24:44\",\"user_level_spend\":0,\"user_level_is_buildin\":0,\"user_level_product_id\":\"0\",\"user_level_name\":\"v6\",\"user_level_logo\":\"https://test.shopsuite.cn/account/static/src/default/images/user_level/vip_icon_un.png\",\"user_level_fans_vip\":0,\"user_level_fx_rate\":100,\"user_level_validity\":0,\"user_level_exp\":3333333,\"user_level_privilege\":\"\",\"user_level_fans_team\":0,\"user_level_id\":1006,\"plantform_fx_cps_rate_1\":\"20\",\"plantform_fx_cps_rate_2\":\"15\",\"plantform_fx_cps_rate_3\":\"10\"}]', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_content_description', 'seo', 'text', '', '团购{sitename}{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_content_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_content_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gb_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('google_app_id', 'connect', 'text', '', '11', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('google_app_key', 'connect', 'text', '', '11', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('google_callback_url', 'connect', 'text', '', 'https://test.suteshop.com/account.php/Connect_Google/callback', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('google_status', 'connect', 'text', '', '2', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('gzipcompress', '系统参数', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('hall_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('huawei_access_key', '华为access_key', 'text', '', '', 37, '', 32, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('huawei_bucket', '华为bucket', 'text', '', '', 37, '', 35, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('huawei_domain', '华为domain', 'text', '', '', 37, '', 31, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('huawei_endpoint', '华为endpoint', 'text', '', '', 37, '', 34, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('huawei_secret_key', '华为secret_key', 'text', '', '', 37, '', 33, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('icp_number', '备案号', 'text', '', '沪ICP备18022949号-1-1', 1, '', 20, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('im_enable', '是否启用', 'radio', '0:禁用|1:启用', '0', 1423, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('im_offline_msg', 'im', 'text', '', '我不在线，欢迎留言！', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_banner_pc_enable', 'pop', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_banner_pc_html', 'pop', 'text', '', '  <div id=\"login\" style=\" background: #090359;\">\n    <div class=\"login_img\">\n    <a>  <img src=\"http://mallsuite.oss-accelerate.aliyuncs.com/mall/images/media/plantform/20220727/98b4dc4a587a4fa6a95337d47bc9f064.jpg\" alt=\"\" /></a>\n\n    </div>\n  </div>', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_pc_enable', 'pop', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_pc_img', 'pop', 'text', '', 'http://mallsuite.oss-accelerate.aliyuncs.com/mall/images/media/plantform/20220727/bf83063fe4bd45a2a1b8462e12af7dac.png', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_pc_index', 'pop', 'text', '', '10111', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_pc_url', 'pop', 'text', '', 'https://demo.mallsuite.cn/pc/#/index', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_wap_enable', 'pop', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_wap_img', 'pop', 'text', '', 'https://test.suteshop.com/image.php/shop/data/upload/media/user/10001/image/20210122/1611305867399088.png', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('index_pop_wap_url', 'pop', 'text', '', '/pagesub/diy-page/diy-page?id=1611141316254', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('info_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('invoice_enable', '发票功能启用', 'radio', '0:禁用|1:开启', '', 8, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('invoice_special_enable', '专用发票功能开启|关闭', 'radio', '0:禁用|1:开启', '1,2,3,4,5,6,7', 8, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('invoice_tips', '发票提示', 'text', '', '1.订单支付后365日内可申请发票。  2.开票满200包邮，低于200需要自付邮费。  3.申请增值税普通发票预计1周内可寄出快递，若遇月底可能会延迟7个工作日左右。申请增值税专用发票1个月内可寄出快递。', 1427, 'pc提交订单页发票文字', 255, b'1', b'0', 1001);
INSERT INTO `sys_config_base` VALUES ('invoicing_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('kefu_type_id', '小程序客服类型', 'radio', '1:微信客服（公众号或小程序客服）', '1', 1423, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('kuaidi100_app_id', 'logistics', 'text', '', '15F0A847ADF70BC261A45019B5F3DAA7', 10, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('kuaidi100_app_key', 'logistics', 'text', '', 'sfhiJkEO8429', 10, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('kuaidiniao_app_key', '快递鸟App Key', 'text', '', 'fb39b19b-a139-43f5-820b-40fcd75279ae', 10, '', 20, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('kuaidiniao_e_business_id', '快递鸟Business Id', 'text', '', '1745287', 10, '尚无Business Id,<a href=\"https://www.kdniao.com/reg\" target=\"_blank\">点击申请</a>', 10, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('language_id', '默认语言', 'text', '', '1001', 1, '', 98, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('licence_key', '授权证书', 'text', '', '2+WJU0eC2hmoby4Fx3PHTu1dT1IRV5wMIp24otk29WtRThXgdzWPSyHdl0oeNkFNATwI8m+EXBJPAsd0Bo/WpyVkPzFOz7WAh1AMmYmFN1tI7nMCTWC2UmyP+cBuZukq0AalVfbMyzb9ll+t5LGwGD44DKS3CdsEIhqJEr1JUFCsj/D92lK0XB2pyZgyJCnQxDj2dapPbgsNXzsePheIaqQ3v2YHGViQ11ypUJDCwqO57HcNQLLXW0P0v2tCdGzkY9oi3G4Z1mC2ob15cim4baKgnkHBjg5RgvHqkNd1PC+ePeiwHDdDN7nToM3rOFFzt84+G4twZcBI335Gjyu/yg==', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('loadctrl', '系统参数', 'text', '', '80', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('logistics_channel', 'logistics', 'text', '', '1', 10, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('minify', 'security', 'text', '', '0', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('mobile_status', 'connect', 'text', '', '1', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('money_pay_enable', '是否启用', 'radio', '0:禁用|1:启用', '1', 1406, '余额支付请选择开启或关闭', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('money_pay_logo', '余额支付图标', 'image', '', '', 1406, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('mp_id', '公众号原始ID', 'text', '', 'gh_8c3e1109ecfd', 4, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('multilang_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('multishop_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('o2o_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('offline_pay_enable', '是否启用', 'radio', '0:禁用|1:启用', '0', 1422, '线下支付请选择开启或关闭', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('offline_pay_logo', '线下支付图标', 'image', '', '', 1422, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('open_store_description', 'protocols', 'text', '', '<p>使用本公司服务所须遵守的条款和条件。<br />\r\n<br />\r\n1.用户资格<br />\r\n本公司的服务仅向适用法律下能够签订具有法律约束力的合同的个人提供并仅由其使用。在不限制前述规定的前提下，本公司的服务不向18周岁以下或被临时或无限期中止的用户提供。如您不合资格，请勿使用本公司的服务。此外，您的账户（包括信用评价）和用户名不得向其他方转让或出售。另外，本公司保留根据其意愿中止或终止您的账户的权利。<br />\r\n<br />\r\n2.您的资料（包括但不限于所添加的任何商品）不得：<br />\r\n*具有欺诈性、虚假、不准确或具误导性；<br />\r\n*侵犯任何第三方著作权、专利权、商标权、商业秘密或其他专有权利或发表权或隐私权；<br />\r\n*违反任何适用的法律或法规（包括但不限于有关出口管制、消费者保护、不正当竞争、刑法、反歧视或贸易惯例/公平贸易法律的法律或法规）；<br />\r\n*有侮辱或者诽谤他人，侵害他人合法权益的内容；<br />\r\n*有淫秽、色情、赌博、暴力、凶杀、恐怖或者教唆犯罪的内容；<br />\r\n*包含可能破坏、改变、删除、不利影响、秘密截取、未经授权而接触或征用任何系统、数据或个人资料的任何病毒、特洛依木马、蠕虫、定时炸弹、删除蝇、复活节彩蛋、间谍软件或其他电脑程序；<br />\r\n<br />\r\n3.违约<br />\r\n如发生以下情形，本公司可能限制您的活动、立即删除您的商品、向本公司社区发出有关您的行为的警告、发出警告通知、暂时中止、无限期地中止或终止您的用户资格及拒绝向您提供服务：<br />\r\n(a)您违反本协议或纳入本协议的文件；<br />\r\n(b)本公司无法核证或验证您向本公司提供的任何资料；<br />\r\n(c)本公司相信您的行为可能对您、本公司用户或本公司造成损失或法律责任。<br />\r\n<br />\r\n4.责任限制<br />\r\n本公司、本公司的关联公司和相关实体或本公司的供应商在任何情况下均不就因本公司的网站、本公司的服务或本协议而产生或与之有关的利润损失或任何特别、间接或后果性的损害（无论以何种方式产生，包括疏忽）承担任何责任。您同意您就您自身行为之合法性单独承担责任。您同意，本公司和本公司的所有关联公司和相关实体对本公司用户的行为的合法性及产生的任何结果不承担责任。<br />\r\n<br />\r\n5.无代理关系<br />\r\n用户和本公司是独立的合同方，本协议无意建立也没有创立任何代理、合伙、合营、雇员与雇主或特许经营关系。本公司也不对任何用户及其网上交易行为做出明示或默许的推荐、承诺或担保。<br />\r\n<br />\r\n6.一般规定<br />\r\n本协议在所有方面均受中华人民共和国法律管辖。本协议的规定是可分割的，如本协议任何规定被裁定为无效或不可执行，该规定可被删除而其余条款应予以执行。</p>\r\n', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('order_autocancel_time', '订单取消时长', 'text', '', '10', 40, '消费者下单后，超过时长未支付，则订单自动取消。单位(小时),0为不设置自动取消', 30, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('order_autofinish_time', '自动收货时间', 'number', '', '7', 40, '从商家发货时间算起，系统自动收货时间,单位(天),0为不设置自动收货', 49, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('order_refund_flag', '退款方式', 'radio', '0:禁用|1:开启', '1', 8, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('order_rsync_time', 'supply', 'text', '', '0.0001', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('page_pc_help', '', 'text', '', '[{\"items\":[{\"id\":\"\",\"title\":\"一标题1\",\"link\":\"/product/lists?category_id=2043\",\"switch\":true,\"index\":0},{\"id\":\"\",\"title\":\"一标题2\",\"link\":\"/product/lists\",\"switch\":true,\"index\":1}],\"index\":0,\"cat_name\":\"分类一\"},{\"cat_name\":\"分类二\",\"switch\":true,\"items\":[{\"id\":\"\",\"title\":\"二标题1\",\"link\":\"/product/lists\",\"switch\":true,\"index\":0}],\"index\":1}]', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('paid_member_agreement', '付费会员协议', 'ueditor', '', '', 33, '', 33, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('pc_enable', 'frontend', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_app_url', 'plantform', 'radio', '0:禁用|1:启用', 'https://test.suteshop.com/index.php', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_auto_withdraw', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_enable', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_rate_1', 'distribution', 'text', '', '5', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_rate_2', 'distribution', 'text', '', '3', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_rate_2_enable', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_rate_3', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_cps_rate_3_enable', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_directseller_customer_exptime', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_directseller_exptime_type', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_directseller_rel_exptime', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_gift_point', '推广粉丝赠送积分', 'number', '', '100', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_is_verify', 'distribution', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_settle_time_type', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_sku_enable', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_userlevel_flag', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_fx_withdraw_min_amount', 'distribution', 'text', '', '10', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plantform_rebate_enable', 'distribution', 'text', '', 'false', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('plugin_rows', '系统参数', 'text', '', '{}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_consume_max', '每订单最多赠送积分', 'number', '', '500', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_consume_rate', '消费额赠送积分比例', 'number', '', '0.5', 7, '消费额赠送积分比例。 获得积分额度 = 消费额度 * 赠送比例。 如果填写100， 用户消费1元，则赠送100积分！', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_evaluate_good', '订单商品评论', 'number', '', '10', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_login', '会员每天签到', 'number', '', '5', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_pay_enable', '积分支付', 'radio', '0:禁用|1:启用', '0', 1413, '', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_pay_logo', '积分支付图标', 'image', '', '', 1413, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_protocols', '积分协议', 'ueditor', '', '', 34, '', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_reg', '注册送积分', 'number', '', '30', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('points_vaue_rate', '积分价值', 'number', '', '0.01', 7, '1 积分 = N 元人民币', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_content_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_content_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_content_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('point_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('pricing_policy', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('private_key', '私钥', 'text', '', 'MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMFPa+v52FkSUXvcUnrGI/XzW3EpZRI0s9BCWJ3oNQmEYA5luWW5p8h0uadTIoTyYweFPdH4hveyxlwmS7oefvbIdiP+o+QIYW/R4Wjsb4Yl8MhR4PJqUE3RCy6IT9fM8ckG4kN9ECs6Ja8fQFc6/mSl5dJczzJO3k1rWMBhKJD/AgMBAAECgYEAucMakH9dWeryhrYoRHcXo4giPVJsH9ypVt4KzmOQY/7jV7KFQK3x//27UoHfUCak51sxFw9ek7UmTPM4HjikA9LkYeE7S381b4QRvFuf3L6IbMP3ywJnJ8pPr2l5SqQ00W+oKv+w/VmEsyUHr+k4Z+4ik+FheTkVWp566WbqFsECQQDjYaMcaKw3j2Zecl8T6eUe7fdaRMIzp/gcpPMfT/9rDzIQk+7ORvm1NI9AUmFv/FAlfpuAMrdL2n7p9uznWb7RAkEA2aP934kbXg5bdV0R313MrL+7WTK/qdcYxATUbMsMuWWQBoS5irrt80WCZbG48hpocJavLNjbtrjmUX3CuJBmzwJAOJg8uP10n/+ZQzjEYXh+BszEHDuw+pp8LuT/fnOy5zrJA0dO0RjpXijO3vuiNPVgHXT9z1LQPJkNrb5ACPVVgQJBALPeb4uV0bNrJDUb5RB4ghZnIxv18CcaqNIft7vuGCcFBAIPIRTBprR+RuVq+xHDt3sNXdsvom4h49+Hky1b0ksCQBBwUtVaqH6ztCtwUF1j2c/Zcrt5P/uN7IHAd44K0gIJc1+Csr3qPG+G2yoqRM8KVqLI8Z2ZYn9c+AvEE+L9OQY=', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_salenum_flag', '显示销量', 'radio', '0:不显示|1:显示', '1', 39, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_show_type', 'plantform', 'radio', '0:禁用|1:启用', 'SKU', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_spec_edit', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_verify_flag', '商品是否审核', 'radio', '0:不审核|1:需要审核', '0', 39, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('product_virtual_salenum', '虚拟销量', 'number', '', '10', 39, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('public_key', '公钥', 'text', '', 'MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDBT2vr+dhZElF73FJ6xiP181txKWUSNLPQQlid6DUJhGAOZblluafIdLmnUyKE8mMHhT3R+Ib3ssZcJku6Hn72yHYj/qPkCGFv0eFo7G+GJfDIUeDyalBN0QsuiE/XzPHJBuJDfRArOiWvH0BXOv5kpeXSXM8yTt5Na1jAYSiQ/wIDAQAB', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('purchase_activity_description', 'protocols', 'text', '', '<p>一、团购的所有权和运作权归本公司。</p><p>二、本公司有权在必要时修改本协议，本协议一旦发生变更，将会在相关页面上公布。如果您不同意所改动的内容，您应主动停止使用团购服务。如果您继续使用服务，则视为接受本协议的变更。</p><p>三、如发生下列任何一种情形，本公司有权中断或终止向您提供的服务而无需通知您：</p><p>1、 您提供的个人资料不真实；</p><p>2、您违反本协议的规定；</p><p>3、 按照政府主管部门的监管要求；</p><p>4、本公司认为您的行为违反团购服务性质或需求的特殊情形。</p><p>四、尽管本协议可能另有其他规定，本公司仍然可以随时终止本协议。</p><p>五、本公司终止本协议的权利不会妨害本公司可能拥有的在本协议终止前因您违反本协议或本公司本应享有的任何其他权利。</p><p>六、您理解并完全接受，本公司有权自行对团购资源作下线处理。</p>', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('qq_app_id', 'connect', 'text', '', '101960965', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('qq_app_key', 'connect', 'text', '', '20bbd077f0ff3f343949c47b0fd5168a', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('qq_callback_url', 'connect', 'text', '', 'https://test.suteshop.com/account/api/connect/qq.php', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('qq_status', 'connect', 'text', '', '2', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('redpacket_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_ca_bp', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_ca_fee', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_da_bp', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_da_fee', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_privacy_protocols', '隐私政策', 'ueditor', '', '<p><span style=\"color: rgb(96, 98, 102);\">隐私政策内容</span></p>', 31, '', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_protocols', '注册服务协议', 'ueditor', '', '', 30, '', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_protocols_description', 'protocols', 'text', '', '<p>特别提醒用户认真阅读本《用户服务协议及App隐私协议》(下称《协议》) 中各条款。除非您接受本《协议》条款，否则您无权使用本网站提供的相关服务。您的使用行为将视为对本《协议》的接受，并同意接受本《协议》各项条款的约束。&nbsp;<br />\r\n<br />\r\n<strong>一、定义</strong></p>\r\n\r\n<ol>\r\n	<li>&quot;用户&quot;指符合本协议所规定的条件，同意遵守本网站各种规则、条款（包括但不限于本协议），并使用本网站的个人或机构。</li>\r\n	<li>&quot;卖家&quot;是指在本网站上出售物品的用户。&quot;买家&quot;是指在本网站购买物品的用户。</li>\r\n	<li>&quot;成交&quot;指买家根据卖家所刊登的交易要求，在特定时间内提出最优的交易条件，因而取得依其提出的条件购买该交易物品的权利。</li>\r\n</ol>\r\n\r\n<p><strong>二、用户资格</strong><br />\r\n<br />\r\n只有符合下列条件之一的人员或实体才能申请成为本网站用户，可以使用本网站的服务。</p>\r\n\r\n<ol>\r\n	<li>年满十八岁，并具有民事权利能力和民事行为能力的自然人；</li>\r\n	<li>未满十八岁，但监护人（包括但不仅限于父母）予以书面同意的自然人；</li>\r\n	<li>根据中国法律或设立地法律、法规和/或规章成立并合法存在的公司、企事业单位、社团组织和其他组织。</li>\r\n</ol>\r\n\r\n<p><br />\r\n无民事行为能力人、限制民事行为能力人以及无经营或特定经营资格的组织不当注册为本网站用户或超过其民事权利或行为能力范围从事交易的，其与本网站之间的协议自始无效，本网站一经发现，有权立即注销该用户，并追究其使用本网站&quot;服务&quot;的一切法律责任。<br />\r\n&nbsp;</p>\r\n\r\n<p><strong>三、用户个人隐私信息保护</strong></p>\r\n\r\n<ol>\r\n	<li>如果本网站发现或收到他人举报或投诉用户违反本协议约定的，本网站有权不经通知随时对相关内容，包括但不限于用户资料、发贴记录进行审查、删除，并视情节轻重对违规账号处以包括但不限于警告、账号封禁 、设备封禁 、功能封禁 的处罚，且通知用户处理结果。</li>\r\n	<li>因违反用户协议被封禁的用户，可以自行与本网站联系。其中，被实施功能封禁的用户会在封禁期届满后自动恢复被封禁功能。被封禁用户可提交申诉，本网站将对申诉进行审查，并自行合理判断决定是否变更处罚措施。</li>\r\n	<li>用户理解并同意，本网站有权依合理判断对违反有关法律法规或本协议规定的行为进行处罚，对违法违规的任何用户采取适当的法律行动，并依据法律法规保存有关信息向有关部门报告等，用户应承担由此而产生的一切法律责任。</li>\r\n	<li>用户理解并同意，因用户违反本协议约定，导致或产生的任何第三方主张的任何索赔、要求或损失，包括合理的律师费，用户应当赔偿本网站与合作公司、关联公司，并使之免受损害。</li>\r\n</ol>\r\n\r\n<p><br />\r\n<strong>四、用户的权利和义务</strong></p>\r\n\r\n<ol>\r\n	<li>用户有权根据本协议的规定及本网站发布的相关规则，利用本网站网上交易平台登录物品、发布交易信息、查询物品信息、购买物品、与其他用户订立物品买卖合同、在本网站社区发帖、参加本网站的有关活动及有权享受本网站提供的其他的有关资讯及信息服务。</li>\r\n	<li>用户有权根据需要更改密码和交易密码。用户应对以该用户名进行的所有活动和事件负全部责任。</li>\r\n	<li>用户有义务确保向本网站提供的任何资料、注册信息真实准确，包括但不限于真实姓名、身份证号、联系电话、地址、邮政编码等。保证本网站及其他用户可以通过上述联系方式与自己进行联系。同时，用户也有义务在相关资料实际变更时及时更新有关注册资料。</li>\r\n	<li>用户不得以任何形式擅自转让或授权他人使用自己在本网站的用户账号。</li>\r\n	<li>用户有义务确保在本网站网上交易平台上登录物品、发布的交易信息真实、准确，无误导性。</li>\r\n	<li>用户不得在本网站网上交易平台买卖国家禁止销售的或限制销售的物品、不得买卖侵犯他人知识产权或其他合法权益的物品，也不得买卖违背社会公共利益或公共道德的物品。</li>\r\n	<li>用户不得在本网站发布各类违法或违规信息。包括但不限于物品信息、交易信息、社区帖子、物品留言，店铺留言，评价内容等。</li>\r\n	<li>用户在本网站交易中应当遵守诚实信用原则，不得以干预或操纵物品价格等不正当竞争方式扰乱网上交易秩序，不得从事与网上交易无关的不当行为，不得在交易平台上发布任何违法信息。</li>\r\n	<li>用户不应采取不正当手段（包括但不限于虚假交易、互换好评等方式）提高自身或他人信用度，或采用不正当手段恶意评价其他用户，降低其他用户信用度。</li>\r\n	<li>用户承诺自己在使用本网站网上交易平台实施的所有行为遵守国家法律、法规和本网站的相关规定以及各种社会公共利益或公共道德。对于任何法律后果的发生，用户将以自己的名义独立承担所有相应的法律责任。</li>\r\n	<li>用户在本网站网上交易过程中如与其他用户因交易产生纠纷，可以请求本网站从中予以协调。用户如发现其他用户有违法或违反本协议的行为，可以向本网站举报。如用户因网上交易与其他用户产生诉讼的，用户有权通过司法部门要求本网站提供相关资料。</li>\r\n	<li>用户应自行承担因交易产生的相关费用，并依法纳税。</li>\r\n	<li>未经本网站书面允许，用户不得将本网站资料以及在交易平台上所展示的任何信息以复制、修改、翻译等形式制作衍生作品、分发或公开展示。</li>\r\n	<li>用户同意接收来自本网站的信息，包括但不限于活动信息、交易信息、促销信息等。</li>\r\n</ol>\r\n\r\n<p><br />\r\n<br />\r\n<strong>五、 本网站的权利和义务</strong></p>\r\n\r\n<ol>\r\n	<li>本网站不是传统意义上的&quot;拍卖商&quot;，仅为用户提供一个信息交流、进行物品买卖的平台，充当买卖双方之间的交流媒介，而非买主或卖主的代理商、合伙 人、雇员或雇主等经营关系人。公布在本网站上的交易物品是用户自行上传进行交易的物品，并非本网站所有。对于用户刊登物品、提供的信息或参与竞标的过程， 本网站均不加以监视或控制，亦不介入物品的交易过程，包括运送、付款、退款、瑕疵担保及其它交易事项，且不承担因交易物品存在品质、权利上的瑕疵以及交易 方履行交易协议的能力而产生的任何责任，对于出现在拍卖上的物品品质、安全性或合法性，本网站均不予保证。</li>\r\n	<li>本网站有义务在现有技术水平的基础上努力确保整个网上交易平台的正常运行，尽力避免服务中断或将中断时间限制在最短时间内，保证用户网上交易活动的顺利进行。</li>\r\n	<li>本网站有义务对用户在注册使用本网站网上交易平台中所遇到的问题及反映的情况及时作出回复。</li>\r\n	<li>本网站有权对用户的注册资料进行查阅，对存在任何问题或怀疑的注册资料，本网站有权发出通知询问用户并要求用户做出解释、改正，或直接做出处罚、删除等处理。</li>\r\n	<li>用 户因在本网站网上交易与其他用户产生纠纷的，用户通过司法部门或行政部门依照法定程序要求本网站提供相关资料，本网站将积极配合并提供有关资料；用户将纠 纷告知本网站，或本网站知悉纠纷情况的，经审核后，本网站有权通过电子邮件及电话联系向纠纷双方了解纠纷情况，并将所了解的情况通过电子邮件互相通知对 方。</li>\r\n	<li>因网上交易平台的特殊性，本网站没有义务对所有用户的注册资料、所有的交易行为以及与交易有关的其他事项进行事先审查，但如发生以下情形，本网站有权限制用户的活动、向用户核实有关资料、发出警告通知、暂时中止、无限期地中止及拒绝向该用户提供服务：\r\n	<ul>\r\n		<li>用户违反本协议或因被提及而纳入本协议的文件；</li>\r\n		<li>存在用户或其他第三方通知本网站，认为某个用户或具体交易事项存在违法或不当行为，并提供相关证据，而本网站无法联系到该用户核证或验证该用户向本网站提供的任何资料；</li>\r\n		<li>存在用户或其他第三方通知本网站，认为某个用户或具体交易事项存在违法或不当行为，并提供相关证据。本网站以普通非专业交易者的知识水平标准对相关内容进行判别，可以明显认为这些内容或行为可能对本网站用户或本网站造成财务损失或法律责任。</li>\r\n	</ul>\r\n	</li>\r\n	<li>在反网络欺诈行动中，本着保护广大用户利益的原则，当用户举报自己交易可能存在欺诈而产生交易争议时，本网站有权通过表面判断暂时冻结相关用户账号，并有权核对当事人身份资料及要求提供交易相关证明材料。</li>\r\n	<li>根据国家法律法规、本协议的内容和本网站所掌握的事实依据，可以认定用户存在违法或违反本协议行为以及在本网站交易平台上的其他不当行为，本网站有权在本网站交易平台及所在网站上以网络发布形式公布用户的违法行为，并有权随时作出删除相关信息，而无须征得用户的同意。</li>\r\n	<li>本 网站有权在不通知用户的前提下删除或采取其他限制性措施处理下列信息：包括但不限于以规避费用为目的；以炒作信用为目的；存在欺诈等恶意或虚假内容；与网 上交易无关或不是以交易为目的；存在恶意竞价或其他试图扰乱正常交易秩序因素；该信息违反公共利益或可能严重损害本网站和其他用户合法利益的。</li>\r\n	<li>用 户授予本网站独家的、全球通用的、永久的、免费的信息许可使用权利，本网站有权对该权利进行再授权，依此授权本网站有权(全部或部份地) 使用、复制、修订、改写、发布、翻译、分发、执行和展示用户公示于网站的各类信息或制作其派生作品，以现在已知或日后开发的任何形式、媒体或技术，将上述 信息纳入其他作品内。</li>\r\n</ol>\r\n\r\n<p><br />\r\n<br />\r\n<strong>六、服务的中断和终止</strong></p>\r\n\r\n<ol>\r\n	<li>在 本网站未向用户收取相关服务费用的情况下，本网站可自行全权决定以任何理由 (包括但不限于本网站认为用户已违反本协议的字面意义和精神，或用户在超过180天内未登录本网站等) 终止对用户的服务，并不再保存用户在本网站的全部资料（包括但不限于用户信息、商品信息、交易信息等）。同时本网站可自行全权决定，在发出通知或不发出通 知的情况下，随时停止提供全部或部分服务。服务终止后，本网站没有义务为用户保留原用户资料或与之相关的任何信息，或转发任何未曾阅读或发送的信息给用户 或第三方。此外，本网站不就终止对用户的服务而对用户或任何第三方承担任何责任。</li>\r\n	<li>如用户向本网站提出注销本网站注册用户身份，需经本网站审核同意，由本网站注销该注册用户，用户即解除与本网站的协议关系，但本网站仍保留下列权利：\r\n	<ul>\r\n		<li>用户注销后，本网站有权保留该用户的资料,包括但不限于以前的用户资料、店铺资料、商品资料和交易记录等。</li>\r\n		<li>用户注销后，如用户在注销前在本网站交易平台上存在违法行为或违反本协议的行为，本网站仍可行使本协议所规定的权利。</li>\r\n	</ul>\r\n	</li>\r\n	<li>如存在下列情况，本网站可以通过注销用户的方式终止服务：\r\n	<ul>\r\n		<li>在用户违反本协议相关规定时，本网站有权终止向该用户提供服务。本网站将在中断服务时通知用户。但如该用户在被本网站终止提供服务后，再一次直接或间接或以他人名义注册为本网站用户的，本网站有权再次单方面终止为该用户提供服务；</li>\r\n		<li>一旦本网站发现用户注册资料中主要内容是虚假的，本网站有权随时终止为该用户提供服务；</li>\r\n		<li>本协议终止或更新时，用户未确认新的协议的。</li>\r\n		<li>其它本网站认为需终止服务的情况。</li>\r\n	</ul>\r\n	</li>\r\n	<li>因用户违反相关法律法规或者违反本协议规定等原因而致使本网站中断、终止对用户服务的，对于服务中断、终止之前用户交易行为依下列原则处理：\r\n	<ul>\r\n		<li>本网站有权决定是否在中断、终止对用户服务前将用户被中断或终止服务的情况和原因通知用户交易关系方，包括但不限于对该交易有意向但尚未达成交易的用户,参与该交易竞价的用户，已达成交易要约用户。</li>\r\n		<li>服务中断、终止之前，用户已经上传至本网站的物品尚未交易或交易尚未完成的，本网站有权在中断、终止服务的同时删除此项物品的相关信息。</li>\r\n		<li>服务中断、终止之前，用户已经就其他用户出售的具体物品作出要约，但交易尚未结束，本网站有权在中断或终止服务的同时删除该用户的相关要约和信息。</li>\r\n	</ul>\r\n	</li>\r\n	<li>本网站若因用户的行为（包括但不限于刊登的商品、在本网站社区发帖等）侵害了第三方的权利或违反了相关规定，而受到第三方的追偿或受到主管机关的处分时，用户应赔偿本网站因此所产生的一切损失及费用。</li>\r\n	<li>对违反相关法律法规或者违反本协议规定，且情节严重的用户，本网站有权终止该用户的其它服务。</li>\r\n</ol>\r\n\r\n<p><br />\r\n<br />\r\n<strong>七、协议的修订</strong><br />\r\n<br />\r\n本协议可由本网站随时修订，并将修订后的协议公告于本网站之上，修订后的条款内容自公告时起生效，并成为本协议的一部分。用户若在本协议修改之后，仍继续使用本网站，则视为用户接受和自愿遵守修订后的协议。本网站行使修改或中断服务时，不需对任何第三方负责。<br />\r\n<br />\r\n<strong>八、 本网站的责任范围&nbsp;</strong><br />\r\n<br />\r\n当用户接受该协议时，用户应明确了解并同意∶</p>\r\n\r\n<ol>\r\n	<li>是否经由本网站下载或取得任何资料，由用户自行考虑、衡量并且自负风险，因下载任何资料而导致用户电脑系统的任何损坏或资料流失，用户应负完全责任。</li>\r\n	<li>用户经由本网站取得的建议和资讯，无论其形式或表现，绝不构成本协议未明示规定的任何保证。</li>\r\n	<li>基于以下原因而造成的利润、商誉、使用、资料损失或其它无形损失，本网站不承担任何直接、间接、附带、特别、衍生性或惩罚性赔偿（即使本网站已被告知前款赔偿的可能性）：\r\n	<ul>\r\n		<li>本网站的使用或无法使用。</li>\r\n		<li>经由或通过本网站购买或取得的任何物品，或接收之信息，或进行交易所随之产生的替代物品及服务的购买成本。</li>\r\n		<li>用户的传输或资料遭到未获授权的存取或变更。</li>\r\n		<li>本网站中任何第三方之声明或行为。</li>\r\n		<li>本网站其它相关事宜。</li>\r\n	</ul>\r\n	</li>\r\n	<li>本网站只是为用户提供一个交易的平台，对于用户所刊登的交易物品的合法性、真实性及其品质，以及用户履行交易的能力等，本网站一律不负任何担保责任。用户如果因使用本网站，或因购买刊登于本网站的任何物品，而受有损害时，本网站不负任何补偿或赔偿责任。</li>\r\n	<li>本 网站提供与其它互联网上的网站或资源的链接，用户可能会因此连结至其它运营商经营的网站，但不表示本网站与这些运营商有任何关系。其它运营商经营的网站均 由各经营者自行负责，不属于本网站控制及负责范围之内。对于存在或来源于此类网站或资源的任何内容、广告、产品或其它资料，本网站亦不予保证或负责。因使 用或依赖任何此类网站或资源发布的或经由此类网站或资源获得的任何内容、物品或服务所产生的任何损害或损失，本网站不负任何直接或间接的责任。</li>\r\n</ol>\r\n\r\n<p><br />\r\n<br />\r\n<strong>九、不可抗力</strong><br />\r\n<br />\r\n因不可抗力或者其他意外事件，使得本协议的履行不可能、不必要或者无意义的，双方均不承担责任。本合同所称之不可抗力意指不能预见、不能避免并不能克服的 客观情况，包括但不限于战争、台风、水灾、火灾、雷击或地震、罢工、暴动、法定疾病、黑客攻击、网络病毒、电信部门技术管制、政府行为或任何其它自然或人 为造成的灾难等客观情况。<br />\r\n<br />\r\n<strong>十、争议解决方式</strong></p>\r\n\r\n<ol>\r\n	<li>本协议及其修订本的有效性、履行和与本协议及其修订本效力有关的所有事宜，将受中华人民共和国法律管辖，任何争议仅适用中华人民共和国法律。</li>\r\n	<li>因 使用本网站服务所引起与本网站的任何争议，均应提交深圳仲裁委员会按照该会届时有效的仲裁规则进行仲裁。相关争议应单独仲裁，不得与任何其它方的争议在任 何仲裁中合并处理，该仲裁裁决是终局，对各方均有约束力。如果所涉及的争议不适于仲裁解决，用户同意一切争议由人民法院管辖。</li>\r\n</ol>\r\n', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_seller_fee', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_sp_bp', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('reg_sp_fee', 'distribution', 'text', '', '3000', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('rewrite_open', '系统参数', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('saas_status', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sale_prize_ca_rate', 'distribution', 'text', '', '3', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sale_prize_da_rate', 'distribution', 'text', '', '2', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sale_prize_sp_enable', 'distribution', 'text', '', 'false', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sale_prize_sp_rate_1', 'distribution', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('scope', 'paotui', 'text', '', '10', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sc_order_process', '订单流程', 'checkbox', '2010:待支付|2020:待出库|2030:待发货|2040:待收货|2060:交易完成', '2010,2030,2040,2060,2020', 41, '注意：如果有未处理的订单不要修改这里的订单处理流程！', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sc_return_process', '退单单流程', 'checkbox', '3100:提交退单|3105:退单审核|3110:退货收货确认|3115:退款确认|3120:收款确认|3125:完成', '3100,3105,3110,3125,3115,3120', 41, '注意：如果有未处理的售后单不要修改这里的处理流程！', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('search_hot_words', '热门搜索默认词', 'array', '', '鞋子, 电脑, 数码办公, 衬衫, 香水, 手机, 夏装', 5, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('security_csp', 'security', 'text', '', 'default-src \'self\' \'unsafe-eval\'', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('service_app_key', '服务Key', 'text', '', '1111', 20, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('service_user_id', '服务编号', 'text', '', '1000', 20, '', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('setting_email', 'setting', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('setting_phone', 'setting', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('shop_app_id', 'shop', 'text', '', '102', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('shop_app_key', 'shop', 'text', '', 'KKOJIASaFR0Aw08PV1fka', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('shop_app_url', 'shop', 'text', '', 'https://test.suteshop.com/index.php', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('show_cancel_time', '订单取消倒计时', 'radio', '0:否|1:是', '1', 40, '开启后，显示订单取消倒计时。', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sidecart_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_address', '联系地址', 'text', '', '上海市松江区莘砖公路258号棕榈广场32号楼7楼701-1', 1, '', 15, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_admin_logo', '后台Logo', 'image', '', 'https://shopsuite.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230829/0c2b7826ae2943eca247bc53a4645648.png', 1, '菜单展开左上角logo,建议尺寸：128 * 128px', 10, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_company_name', '公司名称', 'text', '', '随商科技', 1, '', 14, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_description', 'seo', 'text', '', '{site_meta_description} 网上超市，最经济实惠的网上购物商城，用鼠标逛超市，不用排队，方便实惠送上门，网上购物新生活。', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_email', '电子邮件', 'text', '', 'wangcm@shopsuite.cn', 1, '', 17, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_icon', 'Favorites Icon', 'image', '', 'https://shopsuite.oss-accelerate.aliyuncs.com/media/user/10472/image/20210422/1619057337196344.ico', 1, '64 * 64, ICO图片', 8, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_im', 'plantform', 'radio', '0:禁用|1:启用', '', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_is_https', '启用https', 'radio', '0:关闭|1:开启', '1', 1, '', 199, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_keywords', 'seo', 'text', '', '{site_meta_keyword} 网上超市，网上商城，网络购物，进口食品，美容护理，母婴玩具，厨房清洁用品，家用电器，手机数码，电脑软件办公用品，家居生活，服饰内衣，营养保健，钟表珠宝，饰品箱包，汽车生活，图书音像，礼品卡，药品，医疗器械，隐形眼镜等，1号店。', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_meta_description', '网站描述', 'textarea', '', 'mallsuite网上商城系统套件是基于Java技术的企业级电子商务平台系统，以其安全稳定、强大易用、高效专业等优势赢得了用户的广泛好评。环球软件为大、中、小企业提供一个安全、高效、强大的电子商务解决方案，协助企业快速构建、部署和管理其电子商务平台，拓展企业销售渠道，突显电子商务商业价值，致力于推动Java技术和电子商务行业的发展而不断努力。', 1, '', 4, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_meta_keyword', '网站关键词', 'text', '', '网店系统,商城系统,网上商城系统,购物系统,网上购物系统,电子商务系统、订货系统、B2C单用户商城、BBC多用户商城、商城小程序、微商城、微分销、商城ERP系统', 1, '', 3, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_mobile_logo', '移动端LOGO', 'image', '', 'https://shopsuite.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230818/753a0487e15b49a097438d887424e823.png', 1, '移动端登录LOGO，建议尺寸：128 * 128 px', 11, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_name', '网站全称', 'text', '', 'ModulithShop', 1, '网站名称很多地方会显示的，建议认真填写', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_pc_logo', 'PC网站LOGO', 'image', '', 'http://shopsuite.oss-accelerate.aliyuncs.com/mall/images/media/plantform/20220824/33971f5e8288417e90fcce72fc2cb0ba.png', 1, '尺寸 255 * 64', 9, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_slogan', 'slogan', 'text', '', '技术为本、服务至上', 1, '', 7, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_status', '站点状态', 'radio', '0:关闭|1:开启', '0', 1, '站点开始|关闭（用于升级等临时关闭）', 239, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_tel', '联系电话', 'text', '', '19921140807', 1, '', 16, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_title', 'seo', 'text', '', '{site_name} - 订货系统、B2C单用户商城、BBC多用户商城、商城小程序、微商城、微分销、商城ERP系统', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('site_version', '站点版本', 'text', '', '2.0.269', 1, '', 200, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sms_type', '短信类型', 'radio', '0:内置短信|1:阿里云|2:腾讯云|3:华为云', '0', 13, '内置短信，需要开通随商服务账号。', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sns_description', 'seo', 'text', '', '{description}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sns_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sns_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sns_title', 'seo', 'text', '', '{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sphinx_enable', 'sphinx', 'text', '', '0', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sp_vaue_rate', '扩展，未使用', 'number', '', '0.01', 7, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('statistics_code', '第三方流量统计代码', 'textarea', '', '', 1, '', 21, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('stock_warning', '警戒库存', 'number', '', '5', 40, '警戒库存提醒值', 60, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('store_description', 'seo', 'text', '', '', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('store_fx_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('store_keywords', 'seo', 'text', '', '{keyword}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('store_supplier_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('store_title', 'seo', 'text', '', '随商科技-{name}', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('Strict-Transport-Security', 'security', 'text', '', 'max-age=600;includeSubDomains', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('subsite_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('suggest_search_words', '默认搜索', 'array', '', '王者荣耀,和平精英,手机', 5, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('supplier_market_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('supply_transfer_flag', 'supply', 'text', '', '1', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('sw_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('task_pay_protocols_description', 'protocols', 'text', '', '<p>1222</p>\r\n', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('task_protocols_description', 'protocols', 'text', '', '<p>111</p>\r\n', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_app_id', '腾讯AppId', 'text', '', '32133213', 36, '', 21, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_default_dir', '腾讯上传目录', 'text', '', '324', 36, '', 26, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_endpoint', '腾讯endpoint', 'text', '', '4234', 36, '', 24, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_region_region', '腾讯region', 'text', '', '43242344324', 36, '', 25, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_secret_id', '腾讯secretId', 'text', '', '43244324', 36, '', 22, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tengxun_secret_key', '腾讯secretKey', 'text', '', '儿4', 36, '', 23, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('text_filter', '过滤词汇', 'textarea', '', 'admin,2', 3, '多个词之间使用\',\'分割', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('time_format', '时间格式', 'text', '', 'H时i分s秒', 1, '', 101, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('time_zone_id', '默认时区', 'text', '', '+08:00', 1, '', 99, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('token_key', '登录Token', 'text', '', 'ULgNsWJ8rPjRtnjzX/Gv2RGS80Ksnm/ZaLpvIL+NrBg=', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('tour_diy_protocols_description', 'protocols', 'text', '', '<p>1111</p>', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('transferPla', 'paotui', 'text', '', '5', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('transferSto', 'paotui', 'text', '', '6', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_app_id', 'push', 'text', '', 'BSI8xYmHyT76ngmcxFQDy9', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_app_key', 'push', 'text', '', 'nhdgI308yP7HKfiauhbji9', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_app_secret', 'push', 'text', '', 'cTBzsJbcZSAOeBXIREr4y2', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_master_secret', 'push', 'text', '', 'ybdUcVJJKw5bAcQB4DpC0A', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_packagename', 'push', 'text', '', 'com.suisung.shopsuite.bbc', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('unipush_packagename_ios', 'push', 'text', '', 'com.suisung.shopsuite.bbc', 14, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('upload_file_ext', '文件类型', 'checkbox', 'rar:rar|zip:zip|tar:tar|gz:gz|7z:7z|bz2:bz2|cab:cab|iso:iso|doc:doc|docx:docx|xls:xls|xlsx:xlsx|csv:csv|ppt:ppt|pptx:pptx|pdf:pdf|txt:txt|md:md|xml:xml|html:html|pem:pem', ',zip,md,xlsx,rar,doc,pem,7z,iso,cab,tar,html,bz2,csv,pptx,txt,gz,pdf,ppt', 2, '', 12, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('upload_image_ext', '图片类型', 'checkbox', 'jpg:jpg|jpeg:jpeg|png:png|gif:gif|webp:webp|bmp:bmp|ico:ico', ',1,jpg,bmp,gif,png,jpeg,webp,ico', 2, '', 10, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('upload_max_filesize', '上传大小', 'number', '', '49', 2, '当前服务器环境，最大允许上传 %d MB 的文件(1024 KB = 1MB)。', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('upload_type', '存储类型', 'radio', '0:本地存储|1:阿里云|2:腾讯云|3:华为云', '0', 2, '', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('upload_video_ext', '视频类型', 'checkbox', 'mp4:mp4|mkv:mkv|webm:webm|mpeg:mpeg|mpg:mpg|ogg:ogg|ogv:ogv|mov:mov|wmv:wmv|wav:wav|flv:flv|swf:swf|avi:avi|rm:rm|rmvb:rmvb|mp3:mp3|mid:mid', ',14,wmv,wav,mid,mpg,swf,mkv,mpeg,ogv,rm,webm,ogg,mp4,flv,mp3,rmvb,mov,avi', 2, '', 11, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('user_level_price_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('user_level_rate_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('user_no_avatar', '默认头像', 'image', '', 'http://mallsuite.oss-accelerate.aliyuncs.com/mall/images/media/store/1266/20220707/d88059d3cd4e4d158b53b0f61daff241.png', 1, '256 * 256, 图片', 13, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('virtual_enable', 'plantform', 'radio', '0:禁用|1:启用', '1', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('voucher_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('warehouse_enable', 'plantform', 'radio', '0:禁用|1:启用', '0', 21, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_app_app_id', '移动应用APPID', 'text', '', 'wx449456ef15998b79', 23, '微信开放平台（微信登录、支付都需要开通此配置）', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_app_app_secret', '移动应用AppSecret', 'text', '', '6aabbd8af1093e4015f1d2d0a73e8fb4', 23, '微信开放平台（微信登录、支付都需要开通此配置）', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_app_id', '公众号APPID', 'text', '', 'wx449456ef15998b79', 4, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_app_secret', '公众号AppSecret', 'text', '', '6aabbd8af1093e4015f1d2d0a73e8fb4', 4, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_connect_auto', '微信自动登录', 'radio', '0:否|1:是', '0', 11, '微信公众号及小程序中，自动授权登录', 20, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_mp_qrcode', '公众号二维码', 'image', '', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20240625/b8437a9950fd4816b5341dac99ff8b92.png', 1, '', 13, b'1', b'0', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_apiclient_cert', '微信支付证书', 'text', '', 'C:/upload/files/20230919/3b206c09c73d44f39af70efc39b2245a.pem', 1403, '微信支付证书，在微信商家平台中可以下载！文件名一般为apiclient_cert.pem', 5, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_apiclient_key', '微信支付证书密钥', 'text', '', 'C:/upload/files/20230919/c4f75e821e614ba6b8640d6955b29b4c.pem', 1403, '微信支付证书密钥，在微信商家平台中可以下载！文件名一般为apiclient_key.pem', 7, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_enable', '是否启用', 'radio', '0:禁用|1:启用', '1', 1403, '登录微信商户(地址：https://pay.weixin.qq.com，支付授权目录、回调链接：https://demo.shopsuite.cn； http,https最好都配置)。微信平台证书默认放入${upload.dir}中，命名为：wx_cert.pem， 例如： upload/files/wx_cert.pem', 0, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_logo', '微信支付图标', 'image', '', 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231008/115cde6d982840c5aee2962ece359fab.jpg', 1403, '', 10, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_mchid', '商户编号', 'text', '', '1504498771', 1403, '受理商ID,身份标识', 1, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_serial_no', '证书序列号', 'text', '', '25082384A7D5651BE9E50650167286CCE29A5B42', 1403, '「商户API证书」的「证书序列号」，可以在证书管理里面查看', 2, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_pay_v3_key', 'v3支付KEY', 'text', '', 'YZnt8B766xMYV1HzuLfQBoflWWzngkNL', 1403, '', 4, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_status', '是否开启', 'text', '', '1', 4, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_xcx_app_id', 'AppId', 'text', '', 'wx9d19d00b58d7878e', 22, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('wechat_xcx_app_secret', 'AppSecret', 'text', '', 'e456bb4d7da41a488a59bed2c8a4068c', 22, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weibo_app_id', 'connect', 'text', '', '748829303', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weibo_app_key', 'connect', 'text', '', '75a6d3939e44d16082bf94fede1113fe', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weibo_status', 'connect', 'text', '', '2', 11, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weixin_callback_url', 'connect', 'text', '', 'https://test.suteshop.com/account.php?ctl=Connect_Weixin&met=callback', 11, '', 11, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weixin_pc_app_id', '网站应用APPID', 'text', '', 'wx2ed622ec8d825617', 29, 'pc端用户扫码登录使用:开放平台网站应用的APPID', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weixin_pc_app_key', '网站应用AppSecret', 'text', '', '10540c54a4b1391182aae40f6ada6dba', 29, 'pc端用户扫码登录使用', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('weixin_status', '微信登录', 'radio', '0:禁用|1:启用', '1', 11, '微信公众号、微信小程序、PC网站、App', 10, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('withdraw_fee_rate', '提现佣金比例', 'text', '', '0', 8, '提现佣金比例，请填写百分比符号%前面的数字', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('withdraw_min_amount', '提现最低金额', 'text', '', '100', 8, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('withdraw_monthday', '佣金提现日期', 'text', '', '1,2,3,10,21,25,30,31,16', 8, '佣金提现日期，填写\'1,2\'代表只有每月1日,2日可以申请提现', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('withdraw_protocols_description', 'protocols', 'text', '', '<p><u>11111</u></p>', 0, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('withdraw_received_day', '客户确认收货后N天可提现', 'text', '', '0.001', 8, '客户收货确认后N天可提现，如果填写7，则代表客户确认收货7天后，这笔订单金额商家可申请提现', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('X-Content-Type-Options', 'security', 'text', '', 'nosniff', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('X-Frame-Options', 'security', 'text', '', 'SAMEORIGIN', 1425, '', 255, b'1', b'1', 1001);
INSERT INTO `sys_config_base` VALUES ('X-XSS-Protection', 'security', 'text', '', '1; mode=block', 1425, '', 255, b'1', b'1', 1001);

-- ----------------------------
-- Table structure for sys_config_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_config_type`;
CREATE TABLE `sys_config_type`  (
  `config_type_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分组编号',
  `config_type_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组名称',
  `config_type_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 255 COMMENT '分组排序:从小到大',
  `config_type_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):0-禁用;1-启用',
  `config_type_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `create_time` timestamp(0) NOT NULL COMMENT '添加时间',
  `config_type_module` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属模块(ENUM):1001-站点设置;1002-上传设置;1003-运营设置;1004-财务设置;',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`config_type_id`) USING BTREE,
  INDEX `index_name`(`config_type_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1428 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '配置分组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config_type
-- ----------------------------
INSERT INTO `sys_config_type` VALUES (1, '站点设置', 1, b'1', b'1', '2023-08-05 15:43:39', '0000-00-00 00:00:00', 1001, 1001);
INSERT INTO `sys_config_type` VALUES (2, '上传设置', 50, b'1', b'1', '2023-05-29 08:48:34', '0000-00-00 00:00:00', 1002, 1001);
INSERT INTO `sys_config_type` VALUES (3, '过滤词汇', 3, b'1', b'1', '2023-05-28 12:18:42', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (4, '公众号H5', 4, b'1', b'0', '2023-08-05 16:49:20', '0000-00-00 00:00:00', 8001, 1001);
INSERT INTO `sys_config_type` VALUES (5, '搜索设置', 5, b'0', b'1', '2023-08-05 15:47:05', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (7, '积分设置', 7, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (8, '财务设置', 8, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1101, 1001);
INSERT INTO `sys_config_type` VALUES (10, '物流设置', 10, b'1', b'0', '2022-11-10 16:53:01', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (11, '互联登录', 11, b'1', b'0', '2022-11-10 16:53:01', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (12, '邮件设置', 12, b'1', b'0', '2022-11-10 16:53:01', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (13, '短信设置', 13, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 7001, 1001);
INSERT INTO `sys_config_type` VALUES (14, '推送设置', 14, b'1', b'0', '2022-11-10 16:53:01', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (20, '服务设置', 20, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 6001, 1001);
INSERT INTO `sys_config_type` VALUES (21, '功能设置', 21, b'1', b'0', '2023-07-19 15:25:32', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (22, '微信小程序', 4, b'1', b'0', '2023-08-05 16:16:42', '0000-00-00 00:00:00', 8001, 1001);
INSERT INTO `sys_config_type` VALUES (23, '开放移动应用', 4, b'1', b'0', '2023-08-05 16:27:46', '0000-00-00 00:00:00', 8001, 1001);
INSERT INTO `sys_config_type` VALUES (29, '开放网站应用', 4, b'1', b'0', '2023-08-05 16:27:08', '0000-00-00 00:00:00', 8001, 1001);
INSERT INTO `sys_config_type` VALUES (30, '注册协议', 41, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1005, 1001);
INSERT INTO `sys_config_type` VALUES (31, '隐私政策', 42, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1005, 1001);
INSERT INTO `sys_config_type` VALUES (32, '注销协议', 43, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1005, 1001);
INSERT INTO `sys_config_type` VALUES (33, '付费会员协议', 44, b'0', b'1', '2023-10-08 11:54:44', '0000-00-00 00:00:00', 1005, 1001);
INSERT INTO `sys_config_type` VALUES (34, '积分协议', 45, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1005, 1001);
INSERT INTO `sys_config_type` VALUES (35, '阿里云Oss', 51, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1002, 1001);
INSERT INTO `sys_config_type` VALUES (36, '腾讯云Oss', 52, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1002, 1001);
INSERT INTO `sys_config_type` VALUES (37, '华为云Oss', 53, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1002, 1001);
INSERT INTO `sys_config_type` VALUES (39, '商品设置', 3, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (40, '交易设置', 3, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (41, '订单流程', 50, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 5001, 1001);
INSERT INTO `sys_config_type` VALUES (1401, '支付宝支付', 32, b'1', b'1', '2023-06-29 12:29:22', '0000-00-00 00:00:00', 1004, 1001);
INSERT INTO `sys_config_type` VALUES (1403, '微信支付', 31, b'1', b'1', '2023-06-29 12:29:59', '0000-00-00 00:00:00', 1004, 1001);
INSERT INTO `sys_config_type` VALUES (1406, '余额支付', 33, b'1', b'1', '2023-06-29 12:30:49', '0000-00-00 00:00:00', 1004, 1001);
INSERT INTO `sys_config_type` VALUES (1413, '积分支付', 35, b'1', b'1', '2023-08-05 15:43:32', '0000-00-00 00:00:00', 1004, 1001);
INSERT INTO `sys_config_type` VALUES (1422, '线下支付', 34, b'1', b'1', '2023-06-29 12:31:17', '0000-00-00 00:00:00', 1004, 1001);
INSERT INTO `sys_config_type` VALUES (1423, '客服类型', 255, b'1', b'0', '2023-08-05 17:00:31', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (1424, '升级经验值', 255, b'1', b'0', '2023-08-06 07:48:09', '0000-00-00 00:00:00', 1003, 1001);
INSERT INTO `sys_config_type` VALUES (1425, '安全设置', 255, b'1', b'0', '2023-09-10 08:42:50', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (1426, '用户管理', 255, b'1', b'0', '2023-10-08 18:01:12', '0000-00-00 00:00:00', 0, 1001);
INSERT INTO `sys_config_type` VALUES (1427, '发票设置', 255, b'1', b'0', '2023-10-27 10:14:35', '0000-00-00 00:00:00', 0, 1001);

-- ----------------------------
-- Table structure for sys_contract_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_contract_type`;
CREATE TABLE `sys_contract_type`  (
  `contract_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '保障编号',
  `contract_type_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '保障名称',
  `contract_type_desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '保障简写',
  `contract_type_text` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '保障描述',
  `contract_type_deposit` decimal(8, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '保证金',
  `contract_type_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '项目图标',
  `contract_type_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '说明网址',
  `contract_type_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '保障排序',
  `contract_type_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否开启(BOOL):0-关闭;1-开启',
  `contract_type_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL): 0-非内置;1-系统内置',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`contract_type_id`) USING BTREE,
  INDEX `contract_type_enable`(`contract_type_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1007 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消费者保障服务表-店铺加入' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_contract_type
-- ----------------------------
INSERT INTO `sys_contract_type` VALUES (1001, '七天退货', '', '<p>无理由自动同意退货</p>', 0.00, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231225/d0968915338e43088b5839822988dec7.png', '', 100, b'1', b'1', 1001);
INSERT INTO `sys_contract_type` VALUES (1002, '品质承诺', '', '品质向买家作出承诺商品为正品。', 1000.00, 'https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211027/1635298741611488.png', '', 255, b'1', b'1', 1001);
INSERT INTO `sys_contract_type` VALUES (1003, '破损补寄', '', '<p>运输中出现破损可申请补寄。</p>', 0.00, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231225/e258208f5e714cbb8a704d08654bc551.png', '', 255, b'1', b'0', 1001);
INSERT INTO `sys_contract_type` VALUES (1004, '急速物流', '', '<p>保障买家良好的购物物流体验。</p>', 0.00, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231225/88780275192d47aabf68bee1a20e1eb3.png', '', 255, b'1', b'0', 1001);
INSERT INTO `sys_contract_type` VALUES (1006, '不支持退换', '', '<p>不支持退换</p>', 0.00, 'https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20231019/59538399a915445e933ab56f27ad405f.png', '', 255, b'0', b'1', 1001);

-- ----------------------------
-- Table structure for sys_crontab_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_crontab_base`;
CREATE TABLE `sys_crontab_base`  (
  `crontab_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务编号',
  `crontab_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `crontab_file` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务脚本',
  `crontab_last_exe_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上次执行时间',
  `crontab_next_exe_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下次执行时间',
  `crontab_minute` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '*' COMMENT '分钟(LIST):*-每分; 0-0;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30; 31-31; 32-32; 33-33; 34-34; 35-35; 36-36; 37-37; 38-38; 39-39; 40-40; 41-41; 42-42; 43-43; 44-44; 45-45; 46-46; 47-47; 48-48; 49-49; 50-50; 51-51; 52-52; 53-53; 54-54; 55-55; 56-56; 57-57; 58-58; 59-59',
  `crontab_hour` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '*' COMMENT '小时(LIST):*-任意; 0-0;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23',
  `crontab_day` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '*' COMMENT '每天(LIST):*-任意;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30; 31-31',
  `crontab_month` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '*' COMMENT '每月(LIST):*-任意;1-1月; 2-2月; 3-3月; 4-4月; 5-5月; 6-6月; 7-7月; 8-8月; 9-9月; 10-10月; 11-11月; 12-12月',
  `crontab_week` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '*' COMMENT '每周(LIST):*-每周;0-周日; 1-周一;2-周二;3-周三;4-周四;5-周五;6-周六',
  `crontab_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用; 1-启用',
  `crontab_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否内置(BOOL):0-否; 1-是',
  `crontab_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务备注',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`crontab_id`) USING BTREE,
  INDEX `crontab_enable`(`crontab_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '计划任务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_crontab_base
-- ----------------------------
INSERT INTO `sys_crontab_base` VALUES (1001, '更新订单状态', 'UpdateOrderStatusJob', 0, 0, '*', '*', '*', '*', '?', b'1', b'0', '', 1001);

-- ----------------------------
-- Table structure for sys_currency_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_currency_base`;
CREATE TABLE `sys_currency_base`  (
  `currency_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号:全球各国的电话区号',
  `currency_title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '货币名称',
  `currency_lang` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言符号',
  `currency_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '左符号',
  `currency_symbol_right` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '右符号',
  `currency_decimal_place` bit(1) NOT NULL DEFAULT b'0' COMMENT '小数位',
  `currency_exchange_rate` decimal(15, 8) UNSIGNED NOT NULL COMMENT '汇率',
  `currency_status` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-停用; 1-启用',
  `currency_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `currency_is_default` bit(1) NOT NULL DEFAULT b'0' COMMENT '默认汇率(BOOL):0-否; 1-默认',
  `currency_default_lang` bit(1) NOT NULL DEFAULT b'0' COMMENT '默认语言(BOOL):0-否; 1-默认',
  `currency_is_standard` bit(1) NOT NULL DEFAULT b'0' COMMENT '默认UI(BOOL):0-否; 1-默认',
  `currency_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`currency_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 995 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '货币设置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_currency_base
-- ----------------------------
INSERT INTO `sys_currency_base` VALUES (1, '美国', 'en-US', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbd2fde.png', '$', 'USD', b'1', 0.00000000, b'1', '2024-07-15 18:39:31', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (7, '俄罗斯', 'ru-RU', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affba229e.png', 'Rbl.', 'RUB', b'0', 74.68080000, b'1', '2024-07-15 18:45:48', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (30, '希腊语', 'el-GR', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbec621.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:46:45', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (31, '荷兰语', 'nl-NL', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbd7633.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:47:53', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (33, '法国', 'fr-FR', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affb9fb8a.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:44:52', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (34, '西班牙', 'es-ES', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affb9d481.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:44:13', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (35, 'Denmark', 'da-DK', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbe9741.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:38:18', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (39, '意大利', 'it-IT', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbba16a.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:45:30', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (40, '罗马尼亚', 'ro-RO', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affb9ccab.png', 'L.', 'ROL', b'0', 0.24700000, b'1', '2024-07-15 18:46:06', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (44, 'Britain', 'en-GB', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc9b6c.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:43:50', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (48, '波兰语', 'pl-PL', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbb5356.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:49:05', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (49, '德国', 'de-DE', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc0ecf.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:44:34', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (52, '墨西哥', 'es-MX', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affb9d481.png', 'Mex.$', 'MXN', b'0', 22.00100000, b'1', '2024-07-15 18:41:42', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (60, 'Malay', 'ms-MY', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc939c.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:41:24', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (62, '印度尼西亚', 'id-ID', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbefcd4.png', 'Rps.', 'IDR', b'0', 14530.00000000, b'1', '2024-07-15 18:43:21', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (66, '泰国', 'th-TH', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affba41e4.png', 'Tc.', 'THB', b'0', 31.19000000, b'1', '2024-07-15 18:41:24', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (81, '日本', 'ja-JP', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affba1acd.png', 'J.￥', 'JPY', b'0', 104.27000000, b'1', '2024-07-15 18:40:45', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (82, '한국.', 'ko-KR', 'https://static.shopsuite.cn/shopsuite/flags/square/ko_KR.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:41:08', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (84, '越南', 'vi-VN', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbaedc0.png', 'D.', 'VND', b'0', 23170.00000000, b'1', '2024-07-15 18:42:44', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (86, '中国', 'zh-CN', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc4d56.png', '¥', 'CNY', b'1', 1.00000000, b'1', '2024-07-15 18:38:16', b'1', b'1', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (90, '土耳其', '', '', 'TL.', 'TRY', b'0', 6.96250000, b'1', '2020-07-31 14:41:36', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (95, 'မြန်မာ', 'my-MM', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affba323e.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:40:45', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (327, 'ҚазақстанName', 'kk-KZ', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc2e0f.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:40:12', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (351, '葡萄牙', 'pt-PT', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbf0c70.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:45:10', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (358, '芬兰语', 'fi-FI', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbdd3f1.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:47:04', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (371, '拉脱维亚语', 'lv-LV', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbe29e1.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:47:23', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (380, '乌克兰语', 'uk-UA', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbd5ec1.png', '€', 'EUR', b'0', 8.14280000, b'1', '2024-07-15 18:49:23', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (381, 'Република Србија', 'sr-RS', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affb99601.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:38:39', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (852, '中国香港', 'zh-HK', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbc4d56.png', '¥', 'CNY', b'1', 1.00000000, b'1', '2024-07-15 18:38:27', b'1', b'1', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (856, 'ເມືອງລາວ', 'lo-LA', 'https://static.shopsuite.cn/shopsuite/flags/square/lo_LA.png', '', '', b'0', 0.00000000, b'1', '2024-07-15 19:41:29', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (964, '伊拉克共和国', '', '', 'ID', 'IQD', b'0', 0.27300000, b'1', '2020-07-31 13:50:14', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (965, '科威特国', '', '', 'K.D.', 'KWT', b'0', 0.08300000, b'1', '2020-07-31 13:51:29', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (966, '沙特阿拉伯王国', 'ar-SA', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbcc280.png', 'S.R.', 'SAR', b'0', 1.00000000, b'1', '2024-07-15 18:42:11', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (968, '阿曼', '', '', 'RO.', 'OMR', b'0', 0.10500000, b'1', '2020-07-31 13:52:16', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (971, '阿拉伯联合酋长国', '', '', 'د.إ', 'AED', b'0', 1.00000000, b'1', '2020-07-31 18:20:54', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (973, '巴林王国', 'az-AZ', 'https://static.shopsuite.cn/shopsuite/flags/square/5322affbd2044.png', 'BD.', 'BHD', b'0', 0.10300000, b'1', '2024-07-15 18:46:25', b'0', b'0', b'0', 50, 1001);
INSERT INTO `sys_currency_base` VALUES (994, '阿塞拜疆语', '', '', '€', 'EUR', b'0', 8.14280000, b'1', '2020-09-14 14:04:28', b'0', b'0', b'0', 50, 1001);

-- ----------------------------
-- Table structure for sys_dict_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_base`;
CREATE TABLE `sys_dict_base`  (
  `dict_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键编号',
  `dict_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 125 COMMENT '显示顺序:从小到大',
  `dict_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典备注',
  `dict_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `dict_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`dict_id`) USING BTREE,
  INDEX `name`(`dict_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_dict_item
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_item`;
CREATE TABLE `sys_dict_item`  (
  `dict_item_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典项编号',
  `dict_item_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典项名称',
  `dict_item_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典项值',
  `dict_item_datatype` enum('readonly','number','text','textarea','array','password','radio','checkbox','select','icon','date','datetime','image','images','file','files','ueditor','area') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '配置类型(ENUM):readonly-只读文本;number-数字;text-单行文本;textarea-多行文本;array-数组;password-密码;radio-单选框;checkbox-复选框;select-下拉框;icon-字体图标;date-日期;datetime-时间;image-单张图片;images-多张图片;file-单个文件;files-多个文件;ueditor-富文本编辑器;area-地区选择',
  `dict_item_required` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否必须(BOOL):0-非必填;1-必填',
  `dict_item_placeholder` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'placeholder',
  `dict_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `dict_item_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否使用(BOOL):0-未用;1-使用',
  `dict_item_note` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `dict_item_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '显示顺序',
  `dict_item_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `dict_item_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '系统内置(BOOL):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`dict_item_id`) USING BTREE,
  UNIQUE INDEX `title`(`dict_item_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典项表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_district_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_district_base`;
CREATE TABLE `sys_district_base`  (
  `district_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地区编号',
  `district_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地区名称',
  `district_parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级编号',
  `district_level` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '地区等级',
  `district_citycode` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '区号',
  `district_zipcode` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮编',
  `district_lng` decimal(11, 8) NOT NULL DEFAULT 0.00000000 COMMENT '经度',
  `district_lat` decimal(11, 8) NOT NULL DEFAULT 0.00000000 COMMENT '维度',
  `district_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '地区排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`district_id`) USING BTREE,
  INDEX `upid`(`district_parent_id`, `district_sort`) USING BTREE COMMENT '(null)',
  INDEX `district_parent_id`(`district_name`(191), `district_level`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 659012 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '地区表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_district_base
-- ----------------------------
INSERT INTO `sys_district_base` VALUES (110000, '北京市', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110100, '北京市', 110000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110101, '东城区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110102, '西城区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110105, '朝阳区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110106, '丰台区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110107, '石景山区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110108, '海淀区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110109, '门头沟区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110111, '房山区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110112, '通州区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110113, '顺义区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110114, '昌平区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110115, '大兴区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110116, '怀柔区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110117, '平谷区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110118, '密云区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (110119, '延庆区', 110100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120000, '天津市', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120100, '天津市', 120000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120101, '和平区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120102, '河东区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120103, '河西区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120104, '南开区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120105, '河北区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120106, '红桥区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120110, '东丽区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120111, '西青区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120112, '津南区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120113, '北辰区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120114, '武清区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120115, '宝坻区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120116, '滨海新区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120117, '宁河区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120118, '静海区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (120119, '蓟州区', 120100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130000, '河北省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130100, '石家庄市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130102, '长安区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130104, '桥西区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130105, '新华区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130107, '井陉矿区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130108, '裕华区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130109, '藁城区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130110, '鹿泉区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130111, '栾城区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130121, '井陉县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130123, '正定县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130125, '行唐县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130126, '灵寿县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130127, '高邑县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130128, '深泽县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130129, '赞皇县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130130, '无极县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130131, '平山县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130132, '元氏县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130133, '赵县', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130171, '石家庄高新技术产业开发区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130172, '石家庄循环化工园区', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130181, '辛集市', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130183, '晋州市', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130184, '新乐市', 130100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130200, '唐山市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130202, '路南区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130203, '路北区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130204, '古冶区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130205, '开平区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130207, '丰南区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130208, '丰润区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130209, '曹妃甸区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130224, '滦南县', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130225, '乐亭县', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130227, '迁西县', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130229, '玉田县', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130271, '河北唐山芦台经济开发区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130272, '唐山市汉沽管理区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130273, '唐山高新技术产业开发区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130274, '河北唐山海港经济开发区', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130281, '遵化市', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130283, '迁安市', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130284, '滦州市', 130200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130300, '秦皇岛市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130302, '海港区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130303, '山海关区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130304, '北戴河区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130306, '抚宁区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130321, '青龙满族自治县', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130322, '昌黎县', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130324, '卢龙县', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130371, '秦皇岛市经济技术开发区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130372, '北戴河新区', 130300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130400, '邯郸市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130402, '邯山区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130403, '丛台区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130404, '复兴区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130406, '峰峰矿区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130407, '肥乡区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130408, '永年区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130423, '临漳县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130424, '成安县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130425, '大名县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130426, '涉县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130427, '磁县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130430, '邱县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130431, '鸡泽县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130432, '广平县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130433, '馆陶县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130434, '魏县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130435, '曲周县', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130471, '邯郸经济技术开发区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130473, '邯郸冀南新区', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130481, '武安市', 130400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130500, '邢台市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130502, '襄都区', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130503, '信都区', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130505, '任泽区', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130506, '南和区', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130522, '临城县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130523, '内丘县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130524, '柏乡县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130525, '隆尧县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130528, '宁晋县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130529, '巨鹿县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130530, '新河县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130531, '广宗县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130532, '平乡县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130533, '威县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130534, '清河县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130535, '临西县', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130571, '河北邢台经济开发区', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130581, '南宫市', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130582, '沙河市', 130500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130600, '保定市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130602, '竞秀区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130606, '莲池区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130607, '满城区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130608, '清苑区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130609, '徐水区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130623, '涞水县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130624, '阜平县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130626, '定兴县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130627, '唐县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130628, '高阳县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130629, '容城县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130630, '涞源县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130631, '望都县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130632, '安新县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130633, '易县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130634, '曲阳县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130635, '蠡县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130636, '顺平县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130637, '博野县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130638, '雄县', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130671, '保定高新技术产业开发区', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130672, '保定白沟新城', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130681, '涿州市', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130682, '定州市', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130683, '安国市', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130684, '高碑店市', 130600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130700, '张家口市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130702, '桥东区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130703, '桥西区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130705, '宣化区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130706, '下花园区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130708, '万全区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130709, '崇礼区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130722, '张北县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130723, '康保县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130724, '沽源县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130725, '尚义县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130726, '蔚县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130727, '阳原县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130728, '怀安县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130730, '怀来县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130731, '涿鹿县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130732, '赤城县', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130771, '张家口经济开发区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130772, '张家口市察北管理区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130773, '张家口市塞北管理区', 130700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130800, '承德市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130802, '双桥区', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130803, '双滦区', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130804, '鹰手营子矿区', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130821, '承德县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130822, '兴隆县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130824, '滦平县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130825, '隆化县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130826, '丰宁满族自治县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130827, '宽城满族自治县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130828, '围场满族蒙古族自治县', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130871, '承德高新技术产业开发区', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130881, '平泉市', 130800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130900, '沧州市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130902, '新华区', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130903, '运河区', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130921, '沧县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130922, '青县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130923, '东光县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130924, '海兴县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130925, '盐山县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130926, '肃宁县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130927, '南皮县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130928, '吴桥县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130929, '献县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130930, '孟村回族自治县', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130971, '河北沧州经济开发区', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130972, '沧州高新技术产业开发区', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130973, '沧州渤海新区', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130981, '泊头市', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130982, '任丘市', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130983, '黄骅市', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (130984, '河间市', 130900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131000, '廊坊市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131002, '安次区', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131003, '广阳区', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131022, '固安县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131023, '永清县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131024, '香河县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131025, '大城县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131026, '文安县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131028, '大厂回族自治县', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131071, '廊坊经济技术开发区', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131081, '霸州市', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131082, '三河市', 131000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131100, '衡水市', 130000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131102, '桃城区', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131103, '冀州区', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131121, '枣强县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131122, '武邑县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131123, '武强县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131124, '饶阳县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131125, '安平县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131126, '故城县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131127, '景县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131128, '阜城县', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131171, '河北衡水高新技术产业开发区', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131172, '衡水滨湖新区', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (131182, '深州市', 131100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140000, '山西省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140100, '太原市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140105, '小店区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140106, '迎泽区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140107, '杏花岭区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140108, '尖草坪区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140109, '万柏林区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140110, '晋源区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140121, '清徐县', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140122, '阳曲县', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140123, '娄烦县', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140171, '山西转型综合改革示范区', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140181, '古交市', 140100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140200, '大同市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140212, '新荣区', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140213, '平城区', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140214, '云冈区', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140215, '云州区', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140221, '阳高县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140222, '天镇县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140223, '广灵县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140224, '灵丘县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140225, '浑源县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140226, '左云县', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140271, '山西大同经济开发区', 140200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140300, '阳泉市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140302, '城区', 140300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140303, '矿区', 140300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140311, '郊区', 140300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140321, '平定县', 140300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140322, '盂县', 140300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140400, '长治市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140403, '潞州区', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140404, '上党区', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140405, '屯留区', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140406, '潞城区', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140423, '襄垣县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140425, '平顺县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140426, '黎城县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140427, '壶关县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140428, '长子县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140429, '武乡县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140430, '沁县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140431, '沁源县', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140471, '山西长治高新技术产业园区', 140400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140500, '晋城市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140502, '城区', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140521, '沁水县', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140522, '阳城县', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140524, '陵川县', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140525, '泽州县', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140581, '高平市', 140500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140600, '朔州市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140602, '朔城区', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140603, '平鲁区', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140621, '山阴县', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140622, '应县', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140623, '右玉县', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140671, '山西朔州经济开发区', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140681, '怀仁市', 140600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140700, '晋中市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140702, '榆次区', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140703, '太谷区', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140721, '榆社县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140722, '左权县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140723, '和顺县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140724, '昔阳县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140725, '寿阳县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140727, '祁县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140728, '平遥县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140729, '灵石县', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140781, '介休市', 140700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140800, '运城市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140802, '盐湖区', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140821, '临猗县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140822, '万荣县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140823, '闻喜县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140824, '稷山县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140825, '新绛县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140826, '绛县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140827, '垣曲县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140828, '夏县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140829, '平陆县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140830, '芮城县', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140881, '永济市', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140882, '河津市', 140800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140900, '忻州市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140902, '忻府区', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140921, '定襄县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140922, '五台县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140923, '代县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140924, '繁峙县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140925, '宁武县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140926, '静乐县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140927, '神池县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140928, '五寨县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140929, '岢岚县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140930, '河曲县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140931, '保德县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140932, '偏关县', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140971, '五台山风景名胜区', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (140981, '原平市', 140900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141000, '临汾市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141002, '尧都区', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141021, '曲沃县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141022, '翼城县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141023, '襄汾县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141024, '洪洞县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141025, '古县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141026, '安泽县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141027, '浮山县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141028, '吉县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141029, '乡宁县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141030, '大宁县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141031, '隰县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141032, '永和县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141033, '蒲县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141034, '汾西县', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141081, '侯马市', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141082, '霍州市', 141000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141100, '吕梁市', 140000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141102, '离石区', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141121, '文水县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141122, '交城县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141123, '兴县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141124, '临县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141125, '柳林县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141126, '石楼县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141127, '岚县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141128, '方山县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141129, '中阳县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141130, '交口县', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141181, '孝义市', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (141182, '汾阳市', 141100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150000, '内蒙古自治区', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150100, '呼和浩特市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150102, '新城区', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150103, '回民区', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150104, '玉泉区', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150105, '赛罕区', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150121, '土默特左旗', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150122, '托克托县', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150123, '和林格尔县', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150124, '清水河县', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150125, '武川县', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150172, '呼和浩特经济技术开发区', 150100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150200, '包头市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150202, '东河区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150203, '昆都仑区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150204, '青山区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150205, '石拐区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150206, '白云鄂博矿区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150207, '九原区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150221, '土默特右旗', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150222, '固阳县', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150223, '达尔罕茂明安联合旗', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150271, '包头稀土高新技术产业开发区', 150200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150300, '乌海市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150302, '海勃湾区', 150300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150303, '海南区', 150300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150304, '乌达区', 150300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150400, '赤峰市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150402, '红山区', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150403, '元宝山区', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150404, '松山区', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150421, '阿鲁科尔沁旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150422, '巴林左旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150423, '巴林右旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150424, '林西县', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150425, '克什克腾旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150426, '翁牛特旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150428, '喀喇沁旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150429, '宁城县', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150430, '敖汉旗', 150400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150500, '通辽市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150502, '科尔沁区', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150521, '科尔沁左翼中旗', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150522, '科尔沁左翼后旗', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150523, '开鲁县', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150524, '库伦旗', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150525, '奈曼旗', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150526, '扎鲁特旗', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150571, '通辽经济技术开发区', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150581, '霍林郭勒市', 150500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150600, '鄂尔多斯市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150602, '东胜区', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150603, '康巴什区', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150621, '达拉特旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150622, '准格尔旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150623, '鄂托克前旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150624, '鄂托克旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150625, '杭锦旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150626, '乌审旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150627, '伊金霍洛旗', 150600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150700, '呼伦贝尔市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150702, '海拉尔区', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150703, '扎赉诺尔区', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150721, '阿荣旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150722, '莫力达瓦达斡尔族自治旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150723, '鄂伦春自治旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150724, '鄂温克族自治旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150725, '陈巴尔虎旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150726, '新巴尔虎左旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150727, '新巴尔虎右旗', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150781, '满洲里市', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150782, '牙克石市', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150783, '扎兰屯市', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150784, '额尔古纳市', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150785, '根河市', 150700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150800, '巴彦淖尔市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150802, '临河区', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150821, '五原县', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150822, '磴口县', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150823, '乌拉特前旗', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150824, '乌拉特中旗', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150825, '乌拉特后旗', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150826, '杭锦后旗', 150800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150900, '乌兰察布市', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150902, '集宁区', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150921, '卓资县', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150922, '化德县', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150923, '商都县', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150924, '兴和县', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150925, '凉城县', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150926, '察哈尔右翼前旗', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150927, '察哈尔右翼中旗', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150928, '察哈尔右翼后旗', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150929, '四子王旗', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (150981, '丰镇市', 150900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152200, '兴安盟', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152201, '乌兰浩特市', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152202, '阿尔山市', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152221, '科尔沁右翼前旗', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152222, '科尔沁右翼中旗', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152223, '扎赉特旗', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152224, '突泉县', 152200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152500, '锡林郭勒盟', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152501, '二连浩特市', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152502, '锡林浩特市', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152522, '阿巴嘎旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152523, '苏尼特左旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152524, '苏尼特右旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152525, '东乌珠穆沁旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152526, '西乌珠穆沁旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152527, '太仆寺旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152528, '镶黄旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152529, '正镶白旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152530, '正蓝旗', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152531, '多伦县', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152571, '乌拉盖管委会', 152500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152900, '阿拉善盟', 150000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152921, '阿拉善左旗', 152900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152922, '阿拉善右旗', 152900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152923, '额济纳旗', 152900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (152971, '内蒙古阿拉善高新技术产业开发区', 152900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210000, '辽宁省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210100, '沈阳市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210102, '和平区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210103, '沈河区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210104, '大东区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210105, '皇姑区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210106, '铁西区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210111, '苏家屯区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210112, '浑南区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210113, '沈北新区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210114, '于洪区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210115, '辽中区', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210123, '康平县', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210124, '法库县', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210181, '新民市', 210100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210200, '大连市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210202, '中山区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210203, '西岗区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210204, '沙河口区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210211, '甘井子区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210212, '旅顺口区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210213, '金州区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210214, '普兰店区', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210224, '长海县', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210281, '瓦房店市', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210283, '庄河市', 210200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210300, '鞍山市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210302, '铁东区', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210303, '铁西区', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210304, '立山区', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210311, '千山区', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210321, '台安县', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210323, '岫岩满族自治县', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210381, '海城市', 210300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210400, '抚顺市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210402, '新抚区', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210403, '东洲区', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210404, '望花区', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210411, '顺城区', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210421, '抚顺县', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210422, '新宾满族自治县', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210423, '清原满族自治县', 210400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210500, '本溪市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210502, '平山区', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210503, '溪湖区', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210504, '明山区', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210505, '南芬区', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210521, '本溪满族自治县', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210522, '桓仁满族自治县', 210500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210600, '丹东市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210602, '元宝区', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210603, '振兴区', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210604, '振安区', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210624, '宽甸满族自治县', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210681, '东港市', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210682, '凤城市', 210600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210700, '锦州市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210702, '古塔区', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210703, '凌河区', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210711, '太和区', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210726, '黑山县', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210727, '义县', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210781, '凌海市', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210782, '北镇市', 210700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210800, '营口市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210802, '站前区', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210803, '西市区', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210804, '鲅鱼圈区', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210811, '老边区', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210881, '盖州市', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210882, '大石桥市', 210800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210900, '阜新市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210902, '海州区', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210903, '新邱区', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210904, '太平区', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210905, '清河门区', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210911, '细河区', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210921, '阜新蒙古族自治县', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (210922, '彰武县', 210900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211000, '辽阳市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211002, '白塔区', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211003, '文圣区', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211004, '宏伟区', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211005, '弓长岭区', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211011, '太子河区', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211021, '辽阳县', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211081, '灯塔市', 211000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211100, '盘锦市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211102, '双台子区', 211100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211103, '兴隆台区', 211100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211104, '大洼区', 211100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211122, '盘山县', 211100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211200, '铁岭市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211202, '银州区', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211204, '清河区', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211221, '铁岭县', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211223, '西丰县', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211224, '昌图县', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211281, '调兵山市', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211282, '开原市', 211200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211300, '朝阳市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211302, '双塔区', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211303, '龙城区', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211321, '朝阳县', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211322, '建平县', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211324, '喀喇沁左翼蒙古族自治县', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211381, '北票市', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211382, '凌源市', 211300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211400, '葫芦岛市', 210000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211402, '连山区', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211403, '龙港区', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211404, '南票区', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211421, '绥中县', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211422, '建昌县', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (211481, '兴城市', 211400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220000, '吉林省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220100, '长春市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220102, '南关区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220103, '宽城区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220104, '朝阳区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220105, '二道区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220106, '绿园区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220112, '双阳区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220113, '九台区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220122, '农安县', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220171, '长春经济技术开发区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220172, '长春净月高新技术产业开发区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220173, '长春高新技术产业开发区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220174, '长春汽车经济技术开发区', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220182, '榆树市', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220183, '德惠市', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220184, '公主岭市', 220100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220200, '吉林市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220202, '昌邑区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220203, '龙潭区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220204, '船营区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220211, '丰满区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220221, '永吉县', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220271, '吉林经济开发区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220272, '吉林高新技术产业开发区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220273, '吉林中国新加坡食品区', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220281, '蛟河市', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220282, '桦甸市', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220283, '舒兰市', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220284, '磐石市', 220200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220300, '四平市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220302, '铁西区', 220300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220303, '铁东区', 220300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220322, '梨树县', 220300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220323, '伊通满族自治县', 220300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220382, '双辽市', 220300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220400, '辽源市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220402, '龙山区', 220400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220403, '西安区', 220400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220421, '东丰县', 220400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220422, '东辽县', 220400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220500, '通化市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220502, '东昌区', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220503, '二道江区', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220521, '通化县', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220523, '辉南县', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220524, '柳河县', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220581, '梅河口市', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220582, '集安市', 220500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220600, '白山市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220602, '浑江区', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220605, '江源区', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220621, '抚松县', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220622, '靖宇县', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220623, '长白朝鲜族自治县', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220681, '临江市', 220600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220700, '松原市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220702, '宁江区', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220721, '前郭尔罗斯蒙古族自治县', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220722, '长岭县', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220723, '乾安县', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220771, '吉林松原经济开发区', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220781, '扶余市', 220700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220800, '白城市', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220802, '洮北区', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220821, '镇赉县', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220822, '通榆县', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220871, '吉林白城经济开发区', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220881, '洮南市', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (220882, '大安市', 220800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222400, '延边朝鲜族自治州', 220000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222401, '延吉市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222402, '图们市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222403, '敦化市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222404, '珲春市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222405, '龙井市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222406, '和龙市', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222424, '汪清县', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (222426, '安图县', 222400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230000, '黑龙江省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230100, '哈尔滨市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230102, '道里区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230103, '南岗区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230104, '道外区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230108, '平房区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230109, '松北区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230110, '香坊区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230111, '呼兰区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230112, '阿城区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230113, '双城区', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230123, '依兰县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230124, '方正县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230125, '宾县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230126, '巴彦县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230127, '木兰县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230128, '通河县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230129, '延寿县', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230183, '尚志市', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230184, '五常市', 230100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230200, '齐齐哈尔市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230202, '龙沙区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230203, '建华区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230204, '铁锋区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230205, '昂昂溪区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230206, '富拉尔基区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230207, '碾子山区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230208, '梅里斯达斡尔族区', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230221, '龙江县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230223, '依安县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230224, '泰来县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230225, '甘南县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230227, '富裕县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230229, '克山县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230230, '克东县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230231, '拜泉县', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230281, '讷河市', 230200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230300, '鸡西市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230302, '鸡冠区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230303, '恒山区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230304, '滴道区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230305, '梨树区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230306, '城子河区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230307, '麻山区', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230321, '鸡东县', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230381, '虎林市', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230382, '密山市', 230300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230400, '鹤岗市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230402, '向阳区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230403, '工农区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230404, '南山区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230405, '兴安区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230406, '东山区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230407, '兴山区', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230421, '萝北县', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230422, '绥滨县', 230400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230500, '双鸭山市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230502, '尖山区', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230503, '岭东区', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230505, '四方台区', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230506, '宝山区', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230521, '集贤县', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230522, '友谊县', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230523, '宝清县', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230524, '饶河县', 230500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230600, '大庆市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230602, '萨尔图区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230603, '龙凤区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230604, '让胡路区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230605, '红岗区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230606, '大同区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230621, '肇州县', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230622, '肇源县', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230623, '林甸县', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230624, '杜尔伯特蒙古族自治县', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230671, '大庆高新技术产业开发区', 230600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230700, '伊春市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230717, '伊美区', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230718, '乌翠区', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230719, '友好区', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230722, '嘉荫县', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230723, '汤旺县', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230724, '丰林县', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230725, '大箐山县', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230726, '南岔县', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230751, '金林区', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230781, '铁力市', 230700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230800, '佳木斯市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230803, '向阳区', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230804, '前进区', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230805, '东风区', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230811, '郊区', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230822, '桦南县', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230826, '桦川县', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230828, '汤原县', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230881, '同江市', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230882, '富锦市', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230883, '抚远市', 230800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230900, '七台河市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230902, '新兴区', 230900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230903, '桃山区', 230900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230904, '茄子河区', 230900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (230921, '勃利县', 230900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231000, '牡丹江市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231002, '东安区', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231003, '阳明区', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231004, '爱民区', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231005, '西安区', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231025, '林口县', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231071, '牡丹江经济技术开发区', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231081, '绥芬河市', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231083, '海林市', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231084, '宁安市', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231085, '穆棱市', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231086, '东宁市', 231000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231100, '黑河市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231102, '爱辉区', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231123, '逊克县', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231124, '孙吴县', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231181, '北安市', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231182, '五大连池市', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231183, '嫩江市', 231100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231200, '绥化市', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231202, '北林区', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231221, '望奎县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231222, '兰西县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231223, '青冈县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231224, '庆安县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231225, '明水县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231226, '绥棱县', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231281, '安达市', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231282, '肇东市', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (231283, '海伦市', 231200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232700, '大兴安岭地区', 230000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232701, '漠河市', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232721, '呼玛县', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232722, '塔河县', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232761, '加格达奇区', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232762, '松岭区', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232763, '新林区', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (232764, '呼中区', 232700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310000, '上海市', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310100, '上海市', 310000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310101, '黄浦区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310104, '徐汇区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310105, '长宁区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310106, '静安区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310107, '普陀区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310109, '虹口区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310110, '杨浦区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310112, '闵行区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310113, '宝山区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310114, '嘉定区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310115, '浦东新区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310116, '金山区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310117, '松江区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310118, '青浦区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310120, '奉贤区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (310151, '崇明区', 310100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320000, '江苏省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320100, '南京市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320102, '玄武区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320104, '秦淮区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320105, '建邺区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320106, '鼓楼区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320111, '浦口区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320113, '栖霞区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320114, '雨花台区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320115, '江宁区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320116, '六合区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320117, '溧水区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320118, '高淳区', 320100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320200, '无锡市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320205, '锡山区', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320206, '惠山区', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320211, '滨湖区', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320213, '梁溪区', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320214, '新吴区', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320281, '江阴市', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320282, '宜兴市', 320200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320300, '徐州市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320302, '鼓楼区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320303, '云龙区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320305, '贾汪区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320311, '泉山区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320312, '铜山区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320321, '丰县', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320322, '沛县', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320324, '睢宁县', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320371, '徐州经济技术开发区', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320381, '新沂市', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320382, '邳州市', 320300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320400, '常州市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320402, '天宁区', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320404, '钟楼区', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320411, '新北区', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320412, '武进区', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320413, '金坛区', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320481, '溧阳市', 320400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320500, '苏州市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320505, '虎丘区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320506, '吴中区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320507, '相城区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320508, '姑苏区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320509, '吴江区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320571, '苏州工业园区', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320581, '常熟市', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320582, '张家港市', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320583, '昆山市', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320585, '太仓市', 320500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320600, '南通市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320612, '通州区', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320613, '崇川区', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320614, '海门区', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320623, '如东县', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320671, '南通经济技术开发区', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320681, '启东市', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320682, '如皋市', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320685, '海安市', 320600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320700, '连云港市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320703, '连云区', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320706, '海州区', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320707, '赣榆区', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320722, '东海县', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320723, '灌云县', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320724, '灌南县', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320771, '连云港经济技术开发区', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320772, '连云港高新技术产业开发区', 320700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320800, '淮安市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320803, '淮安区', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320804, '淮阴区', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320812, '清江浦区', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320813, '洪泽区', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320826, '涟水县', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320830, '盱眙县', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320831, '金湖县', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320871, '淮安经济技术开发区', 320800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320900, '盐城市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320902, '亭湖区', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320903, '盐都区', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320904, '大丰区', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320921, '响水县', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320922, '滨海县', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320923, '阜宁县', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320924, '射阳县', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320925, '建湖县', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320971, '盐城经济技术开发区', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (320981, '东台市', 320900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321000, '扬州市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321002, '广陵区', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321003, '邗江区', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321012, '江都区', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321023, '宝应县', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321071, '扬州经济技术开发区', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321081, '仪征市', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321084, '高邮市', 321000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321100, '镇江市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321102, '京口区', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321111, '润州区', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321112, '丹徒区', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321171, '镇江新区', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321181, '丹阳市', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321182, '扬中市', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321183, '句容市', 321100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321200, '泰州市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321202, '海陵区', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321203, '高港区', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321204, '姜堰区', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321271, '泰州医药高新技术产业开发区', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321281, '兴化市', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321282, '靖江市', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321283, '泰兴市', 321200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321300, '宿迁市', 320000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321302, '宿城区', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321311, '宿豫区', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321322, '沭阳县', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321323, '泗阳县', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321324, '泗洪县', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (321371, '宿迁经济技术开发区', 321300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330000, '浙江省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330100, '杭州市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330102, '上城区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330105, '拱墅区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330106, '西湖区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330108, '滨江区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330109, '萧山区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330110, '余杭区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330111, '富阳区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330112, '临安区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330113, '临平区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330114, '钱塘区', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330122, '桐庐县', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330127, '淳安县', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330182, '建德市', 330100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330200, '宁波市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330203, '海曙区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330205, '江北区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330206, '北仑区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330211, '镇海区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330212, '鄞州区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330213, '奉化区', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330225, '象山县', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330226, '宁海县', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330281, '余姚市', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330282, '慈溪市', 330200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330300, '温州市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330302, '鹿城区', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330303, '龙湾区', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330304, '瓯海区', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330305, '洞头区', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330324, '永嘉县', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330326, '平阳县', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330327, '苍南县', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330328, '文成县', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330329, '泰顺县', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330381, '瑞安市', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330382, '乐清市', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330383, '龙港市', 330300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330400, '嘉兴市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330402, '南湖区', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330411, '秀洲区', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330421, '嘉善县', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330424, '海盐县', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330481, '海宁市', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330482, '平湖市', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330483, '桐乡市', 330400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330500, '湖州市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330502, '吴兴区', 330500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330503, '南浔区', 330500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330521, '德清县', 330500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330522, '长兴县', 330500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330523, '安吉县', 330500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330600, '绍兴市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330602, '越城区', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330603, '柯桥区', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330604, '上虞区', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330624, '新昌县', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330681, '诸暨市', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330683, '嵊州市', 330600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330700, '金华市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330702, '婺城区', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330703, '金东区', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330723, '武义县', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330726, '浦江县', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330727, '磐安县', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330781, '兰溪市', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330782, '义乌市', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330783, '东阳市', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330784, '永康市', 330700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330800, '衢州市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330802, '柯城区', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330803, '衢江区', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330822, '常山县', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330824, '开化县', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330825, '龙游县', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330881, '江山市', 330800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330900, '舟山市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330902, '定海区', 330900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330903, '普陀区', 330900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330921, '岱山县', 330900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (330922, '嵊泗县', 330900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331000, '台州市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331002, '椒江区', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331003, '黄岩区', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331004, '路桥区', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331022, '三门县', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331023, '天台县', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331024, '仙居县', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331081, '温岭市', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331082, '临海市', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331083, '玉环市', 331000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331100, '丽水市', 330000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331102, '莲都区', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331121, '青田县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331122, '缙云县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331123, '遂昌县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331124, '松阳县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331125, '云和县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331126, '庆元县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331127, '景宁畲族自治县', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (331181, '龙泉市', 331100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340000, '安徽省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340100, '合肥市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340102, '瑶海区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340103, '庐阳区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340104, '蜀山区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340111, '包河区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340121, '长丰县', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340122, '肥东县', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340123, '肥西县', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340124, '庐江县', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340171, '合肥高新技术产业开发区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340172, '合肥经济技术开发区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340173, '合肥新站高新技术产业开发区', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340181, '巢湖市', 340100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340200, '芜湖市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340202, '镜湖区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340207, '鸠江区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340209, '弋江区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340210, '湾沚区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340212, '繁昌区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340223, '南陵县', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340271, '芜湖经济技术开发区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340272, '安徽芜湖三山经济开发区', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340281, '无为市', 340200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340300, '蚌埠市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340302, '龙子湖区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340303, '蚌山区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340304, '禹会区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340311, '淮上区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340321, '怀远县', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340322, '五河县', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340323, '固镇县', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340371, '蚌埠市高新技术开发区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340372, '蚌埠市经济开发区', 340300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340400, '淮南市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340402, '大通区', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340403, '田家庵区', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340404, '谢家集区', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340405, '八公山区', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340406, '潘集区', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340421, '凤台县', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340422, '寿县', 340400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340500, '马鞍山市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340503, '花山区', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340504, '雨山区', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340506, '博望区', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340521, '当涂县', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340522, '含山县', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340523, '和县', 340500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340600, '淮北市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340602, '杜集区', 340600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340603, '相山区', 340600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340604, '烈山区', 340600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340621, '濉溪县', 340600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340700, '铜陵市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340705, '铜官区', 340700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340706, '义安区', 340700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340711, '郊区', 340700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340722, '枞阳县', 340700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340800, '安庆市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340802, '迎江区', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340803, '大观区', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340811, '宜秀区', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340822, '怀宁县', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340825, '太湖县', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340826, '宿松县', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340827, '望江县', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340828, '岳西县', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340871, '安徽安庆经济开发区', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340881, '桐城市', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (340882, '潜山市', 340800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341000, '黄山市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341002, '屯溪区', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341003, '黄山区', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341004, '徽州区', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341021, '歙县', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341022, '休宁县', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341023, '黟县', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341024, '祁门县', 341000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341100, '滁州市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341102, '琅琊区', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341103, '南谯区', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341122, '来安县', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341124, '全椒县', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341125, '定远县', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341126, '凤阳县', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341171, '中新苏滁高新技术产业开发区', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341172, '滁州经济技术开发区', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341181, '天长市', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341182, '明光市', 341100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341200, '阜阳市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341202, '颍州区', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341203, '颍东区', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341204, '颍泉区', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341221, '临泉县', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341222, '太和县', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341225, '阜南县', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341226, '颍上县', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341271, '阜阳合肥现代产业园区', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341272, '阜阳经济技术开发区', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341282, '界首市', 341200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341300, '宿州市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341302, '埇桥区', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341321, '砀山县', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341322, '萧县', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341323, '灵璧县', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341324, '泗县', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341371, '宿州马鞍山现代产业园区', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341372, '宿州经济技术开发区', 341300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341500, '六安市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341502, '金安区', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341503, '裕安区', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341504, '叶集区', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341522, '霍邱县', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341523, '舒城县', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341524, '金寨县', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341525, '霍山县', 341500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341600, '亳州市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341602, '谯城区', 341600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341621, '涡阳县', 341600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341622, '蒙城县', 341600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341623, '利辛县', 341600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341700, '池州市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341702, '贵池区', 341700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341721, '东至县', 341700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341722, '石台县', 341700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341723, '青阳县', 341700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341800, '宣城市', 340000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341802, '宣州区', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341821, '郎溪县', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341823, '泾县', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341824, '绩溪县', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341825, '旌德县', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341871, '宣城市经济开发区', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341881, '宁国市', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (341882, '广德市', 341800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350000, '福建省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350100, '福州市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350102, '鼓楼区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350103, '台江区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350104, '仓山区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350105, '马尾区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350111, '晋安区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350112, '长乐区', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350121, '闽侯县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350122, '连江县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350123, '罗源县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350124, '闽清县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350125, '永泰县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350128, '平潭县', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350181, '福清市', 350100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350200, '厦门市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350203, '思明区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350205, '海沧区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350206, '湖里区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350211, '集美区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350212, '同安区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350213, '翔安区', 350200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350300, '莆田市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350302, '城厢区', 350300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350303, '涵江区', 350300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350304, '荔城区', 350300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350305, '秀屿区', 350300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350322, '仙游县', 350300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350400, '三明市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350404, '三元区', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350405, '沙县区', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350421, '明溪县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350423, '清流县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350424, '宁化县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350425, '大田县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350426, '尤溪县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350428, '将乐县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350429, '泰宁县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350430, '建宁县', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350481, '永安市', 350400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350500, '泉州市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350502, '鲤城区', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350503, '丰泽区', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350504, '洛江区', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350505, '泉港区', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350521, '惠安县', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350524, '安溪县', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350525, '永春县', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350526, '德化县', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350527, '金门县', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350581, '石狮市', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350582, '晋江市', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350583, '南安市', 350500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350600, '漳州市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350602, '芗城区', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350603, '龙文区', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350604, '龙海区', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350605, '长泰区', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350622, '云霄县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350623, '漳浦县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350624, '诏安县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350626, '东山县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350627, '南靖县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350628, '平和县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350629, '华安县', 350600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350700, '南平市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350702, '延平区', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350703, '建阳区', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350721, '顺昌县', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350722, '浦城县', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350723, '光泽县', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350724, '松溪县', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350725, '政和县', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350781, '邵武市', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350782, '武夷山市', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350783, '建瓯市', 350700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350800, '龙岩市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350802, '新罗区', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350803, '永定区', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350821, '长汀县', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350823, '上杭县', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350824, '武平县', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350825, '连城县', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350881, '漳平市', 350800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350900, '宁德市', 350000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350902, '蕉城区', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350921, '霞浦县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350922, '古田县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350923, '屏南县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350924, '寿宁县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350925, '周宁县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350926, '柘荣县', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350981, '福安市', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (350982, '福鼎市', 350900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360000, '江西省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360100, '南昌市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360102, '东湖区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360103, '西湖区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360104, '青云谱区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360111, '青山湖区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360112, '新建区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360113, '红谷滩区', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360121, '南昌县', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360123, '安义县', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360124, '进贤县', 360100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360200, '景德镇市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360202, '昌江区', 360200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360203, '珠山区', 360200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360222, '浮梁县', 360200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360281, '乐平市', 360200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360300, '萍乡市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360302, '安源区', 360300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360313, '湘东区', 360300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360321, '莲花县', 360300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360322, '上栗县', 360300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360323, '芦溪县', 360300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360400, '九江市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360402, '濂溪区', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360403, '浔阳区', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360404, '柴桑区', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360423, '武宁县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360424, '修水县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360425, '永修县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360426, '德安县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360428, '都昌县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360429, '湖口县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360430, '彭泽县', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360481, '瑞昌市', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360482, '共青城市', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360483, '庐山市', 360400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360500, '新余市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360502, '渝水区', 360500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360521, '分宜县', 360500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360600, '鹰潭市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360602, '月湖区', 360600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360603, '余江区', 360600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360681, '贵溪市', 360600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360700, '赣州市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360702, '章贡区', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360703, '南康区', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360704, '赣县区', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360722, '信丰县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360723, '大余县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360724, '上犹县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360725, '崇义县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360726, '安远县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360728, '定南县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360729, '全南县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360730, '宁都县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360731, '于都县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360732, '兴国县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360733, '会昌县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360734, '寻乌县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360735, '石城县', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360781, '瑞金市', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360783, '龙南市', 360700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360800, '吉安市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360802, '吉州区', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360803, '青原区', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360821, '吉安县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360822, '吉水县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360823, '峡江县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360824, '新干县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360825, '永丰县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360826, '泰和县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360827, '遂川县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360828, '万安县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360829, '安福县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360830, '永新县', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360881, '井冈山市', 360800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360900, '宜春市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360902, '袁州区', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360921, '奉新县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360922, '万载县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360923, '上高县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360924, '宜丰县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360925, '靖安县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360926, '铜鼓县', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360981, '丰城市', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360982, '樟树市', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (360983, '高安市', 360900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361000, '抚州市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361002, '临川区', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361003, '东乡区', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361021, '南城县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361022, '黎川县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361023, '南丰县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361024, '崇仁县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361025, '乐安县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361026, '宜黄县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361027, '金溪县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361028, '资溪县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361030, '广昌县', 361000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361100, '上饶市', 360000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361102, '信州区', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361103, '广丰区', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361104, '广信区', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361123, '玉山县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361124, '铅山县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361125, '横峰县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361126, '弋阳县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361127, '余干县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361128, '鄱阳县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361129, '万年县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361130, '婺源县', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (361181, '德兴市', 361100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370000, '山东省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370100, '济南市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370102, '历下区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370103, '市中区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370104, '槐荫区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370105, '天桥区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370112, '历城区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370113, '长清区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370114, '章丘区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370115, '济阳区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370116, '莱芜区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370117, '钢城区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370124, '平阴县', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370126, '商河县', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370171, '济南高新技术产业开发区', 370100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370200, '青岛市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370202, '市南区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370203, '市北区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370211, '黄岛区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370212, '崂山区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370213, '李沧区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370214, '城阳区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370215, '即墨区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370271, '青岛高新技术产业开发区', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370281, '胶州市', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370283, '平度市', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370285, '莱西市', 370200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370300, '淄博市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370302, '淄川区', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370303, '张店区', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370304, '博山区', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370305, '临淄区', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370306, '周村区', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370321, '桓台县', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370322, '高青县', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370323, '沂源县', 370300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370400, '枣庄市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370402, '市中区', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370403, '薛城区', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370404, '峄城区', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370405, '台儿庄区', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370406, '山亭区', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370481, '滕州市', 370400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370500, '东营市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370502, '东营区', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370503, '河口区', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370505, '垦利区', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370522, '利津县', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370523, '广饶县', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370571, '东营经济技术开发区', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370572, '东营港经济开发区', 370500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370600, '烟台市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370602, '芝罘区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370611, '福山区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370612, '牟平区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370613, '莱山区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370614, '蓬莱区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370671, '烟台高新技术产业开发区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370672, '烟台经济技术开发区', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370681, '龙口市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370682, '莱阳市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370683, '莱州市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370685, '招远市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370686, '栖霞市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370687, '海阳市', 370600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370700, '潍坊市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370702, '潍城区', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370703, '寒亭区', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370704, '坊子区', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370705, '奎文区', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370724, '临朐县', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370725, '昌乐县', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370772, '潍坊滨海经济技术开发区', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370781, '青州市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370782, '诸城市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370783, '寿光市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370784, '安丘市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370785, '高密市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370786, '昌邑市', 370700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370800, '济宁市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370811, '任城区', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370812, '兖州区', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370826, '微山县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370827, '鱼台县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370828, '金乡县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370829, '嘉祥县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370830, '汶上县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370831, '泗水县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370832, '梁山县', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370871, '济宁高新技术产业开发区', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370881, '曲阜市', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370883, '邹城市', 370800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370900, '泰安市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370902, '泰山区', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370911, '岱岳区', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370921, '宁阳县', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370923, '东平县', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370982, '新泰市', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (370983, '肥城市', 370900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371000, '威海市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371002, '环翠区', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371003, '文登区', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371071, '威海火炬高技术产业开发区', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371072, '威海经济技术开发区', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371073, '威海临港经济技术开发区', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371082, '荣成市', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371083, '乳山市', 371000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371100, '日照市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371102, '东港区', 371100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371103, '岚山区', 371100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371121, '五莲县', 371100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371122, '莒县', 371100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371171, '日照经济技术开发区', 371100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371300, '临沂市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371302, '兰山区', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371311, '罗庄区', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371312, '河东区', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371321, '沂南县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371322, '郯城县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371323, '沂水县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371324, '兰陵县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371325, '费县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371326, '平邑县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371327, '莒南县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371328, '蒙阴县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371329, '临沭县', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371371, '临沂高新技术产业开发区', 371300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371400, '德州市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371402, '德城区', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371403, '陵城区', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371422, '宁津县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371423, '庆云县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371424, '临邑县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371425, '齐河县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371426, '平原县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371427, '夏津县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371428, '武城县', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371471, '德州天衢新区', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371481, '乐陵市', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371482, '禹城市', 371400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371500, '聊城市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371502, '东昌府区', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371503, '茌平区', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371521, '阳谷县', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371522, '莘县', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371524, '东阿县', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371525, '冠县', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371526, '高唐县', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371581, '临清市', 371500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371600, '滨州市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371602, '滨城区', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371603, '沾化区', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371621, '惠民县', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371622, '阳信县', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371623, '无棣县', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371625, '博兴县', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371681, '邹平市', 371600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371700, '菏泽市', 370000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371702, '牡丹区', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371703, '定陶区', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371721, '曹县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371722, '单县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371723, '成武县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371724, '巨野县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371725, '郓城县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371726, '鄄城县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371728, '东明县', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371771, '菏泽经济技术开发区', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (371772, '菏泽高新技术开发区', 371700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410000, '河南省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410100, '郑州市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410102, '中原区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410103, '二七区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410104, '管城回族区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410105, '金水区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410106, '上街区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410108, '惠济区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410122, '中牟县', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410171, '郑州经济技术开发区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410172, '郑州高新技术产业开发区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410173, '郑州航空港经济综合实验区', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410181, '巩义市', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410182, '荥阳市', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410183, '新密市', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410184, '新郑市', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410185, '登封市', 410100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410200, '开封市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410202, '龙亭区', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410203, '顺河回族区', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410204, '鼓楼区', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410205, '禹王台区', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410212, '祥符区', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410221, '杞县', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410222, '通许县', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410223, '尉氏县', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410225, '兰考县', 410200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410300, '洛阳市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410302, '老城区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410303, '西工区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410304, '瀍河回族区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410305, '涧西区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410307, '偃师区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410308, '孟津区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410311, '洛龙区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410323, '新安县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410324, '栾川县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410325, '嵩县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410326, '汝阳县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410327, '宜阳县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410328, '洛宁县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410329, '伊川县', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410371, '洛阳高新技术产业开发区', 410300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410400, '平顶山市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410402, '新华区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410403, '卫东区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410404, '石龙区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410411, '湛河区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410421, '宝丰县', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410422, '叶县', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410423, '鲁山县', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410425, '郏县', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410471, '平顶山高新技术产业开发区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410472, '平顶山市城乡一体化示范区', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410481, '舞钢市', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410482, '汝州市', 410400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410500, '安阳市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410502, '文峰区', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410503, '北关区', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410505, '殷都区', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410506, '龙安区', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410522, '安阳县', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410523, '汤阴县', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410526, '滑县', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410527, '内黄县', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410571, '安阳高新技术产业开发区', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410581, '林州市', 410500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410600, '鹤壁市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410602, '鹤山区', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410603, '山城区', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410611, '淇滨区', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410621, '浚县', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410622, '淇县', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410671, '鹤壁经济技术开发区', 410600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410700, '新乡市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410702, '红旗区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410703, '卫滨区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410704, '凤泉区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410711, '牧野区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410721, '新乡县', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410724, '获嘉县', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410725, '原阳县', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410726, '延津县', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410727, '封丘县', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410771, '新乡高新技术产业开发区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410772, '新乡经济技术开发区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410773, '新乡市平原城乡一体化示范区', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410781, '卫辉市', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410782, '辉县市', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410783, '长垣市', 410700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410800, '焦作市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410802, '解放区', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410803, '中站区', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410804, '马村区', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410811, '山阳区', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410821, '修武县', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410822, '博爱县', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410823, '武陟县', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410825, '温县', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410871, '焦作城乡一体化示范区', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410882, '沁阳市', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410883, '孟州市', 410800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410900, '濮阳市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410902, '华龙区', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410922, '清丰县', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410923, '南乐县', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410926, '范县', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410927, '台前县', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410928, '濮阳县', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410971, '河南濮阳工业园区', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (410972, '濮阳经济技术开发区', 410900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411000, '许昌市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411002, '魏都区', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411003, '建安区', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411024, '鄢陵县', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411025, '襄城县', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411071, '许昌经济技术开发区', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411081, '禹州市', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411082, '长葛市', 411000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411100, '漯河市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411102, '源汇区', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411103, '郾城区', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411104, '召陵区', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411121, '舞阳县', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411122, '临颍县', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411171, '漯河经济技术开发区', 411100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411200, '三门峡市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411202, '湖滨区', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411203, '陕州区', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411221, '渑池县', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411224, '卢氏县', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411271, '河南三门峡经济开发区', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411281, '义马市', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411282, '灵宝市', 411200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411300, '南阳市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411302, '宛城区', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411303, '卧龙区', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411321, '南召县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411322, '方城县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411323, '西峡县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411324, '镇平县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411325, '内乡县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411326, '淅川县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411327, '社旗县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411328, '唐河县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411329, '新野县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411330, '桐柏县', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411371, '南阳高新技术产业开发区', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411372, '南阳市城乡一体化示范区', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411381, '邓州市', 411300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411400, '商丘市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411402, '梁园区', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411403, '睢阳区', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411421, '民权县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411422, '睢县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411423, '宁陵县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411424, '柘城县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411425, '虞城县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411426, '夏邑县', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411471, '豫东综合物流产业聚集区', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411472, '河南商丘经济开发区', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411481, '永城市', 411400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411500, '信阳市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411502, '浉河区', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411503, '平桥区', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411521, '罗山县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411522, '光山县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411523, '新县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411524, '商城县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411525, '固始县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411526, '潢川县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411527, '淮滨县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411528, '息县', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411571, '信阳高新技术产业开发区', 411500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411600, '周口市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411602, '川汇区', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411603, '淮阳区', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411621, '扶沟县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411622, '西华县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411623, '商水县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411624, '沈丘县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411625, '郸城县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411627, '太康县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411628, '鹿邑县', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411671, '河南周口经济开发区', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411681, '项城市', 411600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411700, '驻马店市', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411702, '驿城区', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411721, '西平县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411722, '上蔡县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411723, '平舆县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411724, '正阳县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411725, '确山县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411726, '泌阳县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411727, '汝南县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411728, '遂平县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411729, '新蔡县', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (411771, '河南驻马店经济开发区', 411700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (419000, '省直辖县', 410000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (419001, '济源市', 419000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420000, '湖北省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420100, '武汉市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420102, '江岸区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420103, '江汉区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420104, '硚口区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420105, '汉阳区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420106, '武昌区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420107, '青山区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420111, '洪山区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420112, '东西湖区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420113, '汉南区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420114, '蔡甸区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420115, '江夏区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420116, '黄陂区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420117, '新洲区', 420100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420200, '黄石市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420202, '黄石港区', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420203, '西塞山区', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420204, '下陆区', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420205, '铁山区', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420222, '阳新县', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420281, '大冶市', 420200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420300, '十堰市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420302, '茅箭区', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420303, '张湾区', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420304, '郧阳区', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420322, '郧西县', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420323, '竹山县', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420324, '竹溪县', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420325, '房县', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420381, '丹江口市', 420300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420500, '宜昌市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420502, '西陵区', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420503, '伍家岗区', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420504, '点军区', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420505, '猇亭区', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420506, '夷陵区', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420525, '远安县', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420526, '兴山县', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420527, '秭归县', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420528, '长阳土家族自治县', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420529, '五峰土家族自治县', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420581, '宜都市', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420582, '当阳市', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420583, '枝江市', 420500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420600, '襄阳市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420602, '襄城区', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420606, '樊城区', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420607, '襄州区', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420624, '南漳县', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420625, '谷城县', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420626, '保康县', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420682, '老河口市', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420683, '枣阳市', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420684, '宜城市', 420600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420700, '鄂州市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420702, '梁子湖区', 420700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420703, '华容区', 420700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420704, '鄂城区', 420700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420800, '荆门市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420802, '东宝区', 420800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420804, '掇刀区', 420800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420822, '沙洋县', 420800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420881, '钟祥市', 420800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420882, '京山市', 420800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420900, '孝感市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420902, '孝南区', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420921, '孝昌县', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420922, '大悟县', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420923, '云梦县', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420981, '应城市', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420982, '安陆市', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (420984, '汉川市', 420900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421000, '荆州市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421002, '沙市区', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421003, '荆州区', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421022, '公安县', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421024, '江陵县', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421071, '荆州经济技术开发区', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421081, '石首市', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421083, '洪湖市', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421087, '松滋市', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421088, '监利市', 421000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421100, '黄冈市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421102, '黄州区', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421121, '团风县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421122, '红安县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421123, '罗田县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421124, '英山县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421125, '浠水县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421126, '蕲春县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421127, '黄梅县', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421171, '龙感湖管理区', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421181, '麻城市', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421182, '武穴市', 421100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421200, '咸宁市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421202, '咸安区', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421221, '嘉鱼县', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421222, '通城县', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421223, '崇阳县', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421224, '通山县', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421281, '赤壁市', 421200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421300, '随州市', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421303, '曾都区', 421300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421321, '随县', 421300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (421381, '广水市', 421300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422800, '恩施土家族苗族自治州', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422801, '恩施市', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422802, '利川市', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422822, '建始县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422823, '巴东县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422825, '宣恩县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422826, '咸丰县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422827, '来凤县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (422828, '鹤峰县', 422800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (429000, '省直辖县', 420000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (429004, '仙桃市', 429000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (429005, '潜江市', 429000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (429006, '天门市', 429000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (429021, '神农架林区', 429000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430000, '湖南省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430100, '长沙市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430102, '芙蓉区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430103, '天心区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430104, '岳麓区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430105, '开福区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430111, '雨花区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430112, '望城区', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430121, '长沙县', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430181, '浏阳市', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430182, '宁乡市', 430100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430200, '株洲市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430202, '荷塘区', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430203, '芦淞区', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430204, '石峰区', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430211, '天元区', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430212, '渌口区', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430223, '攸县', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430224, '茶陵县', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430225, '炎陵县', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430281, '醴陵市', 430200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430300, '湘潭市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430302, '雨湖区', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430304, '岳塘区', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430321, '湘潭县', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430371, '湖南湘潭高新技术产业园区', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430372, '湘潭昭山示范区', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430373, '湘潭九华示范区', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430381, '湘乡市', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430382, '韶山市', 430300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430400, '衡阳市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430405, '珠晖区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430406, '雁峰区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430407, '石鼓区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430408, '蒸湘区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430412, '南岳区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430421, '衡阳县', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430422, '衡南县', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430423, '衡山县', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430424, '衡东县', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430426, '祁东县', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430471, '衡阳综合保税区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430472, '湖南衡阳高新技术产业园区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430473, '湖南衡阳松木经济开发区', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430481, '耒阳市', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430482, '常宁市', 430400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430500, '邵阳市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430502, '双清区', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430503, '大祥区', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430511, '北塔区', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430522, '新邵县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430523, '邵阳县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430524, '隆回县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430525, '洞口县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430527, '绥宁县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430528, '新宁县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430529, '城步苗族自治县', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430581, '武冈市', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430582, '邵东市', 430500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430600, '岳阳市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430602, '岳阳楼区', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430603, '云溪区', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430611, '君山区', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430621, '岳阳县', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430623, '华容县', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430624, '湘阴县', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430626, '平江县', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430671, '岳阳市屈原管理区', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430681, '汨罗市', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430682, '临湘市', 430600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430700, '常德市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430702, '武陵区', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430703, '鼎城区', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430721, '安乡县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430722, '汉寿县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430723, '澧县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430724, '临澧县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430725, '桃源县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430726, '石门县', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430771, '常德市西洞庭管理区', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430781, '津市市', 430700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430800, '张家界市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430802, '永定区', 430800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430811, '武陵源区', 430800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430821, '慈利县', 430800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430822, '桑植县', 430800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430900, '益阳市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430902, '资阳区', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430903, '赫山区', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430921, '南县', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430922, '桃江县', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430923, '安化县', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430971, '益阳市大通湖管理区', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430972, '湖南益阳高新技术产业园区', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (430981, '沅江市', 430900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431000, '郴州市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431002, '北湖区', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431003, '苏仙区', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431021, '桂阳县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431022, '宜章县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431023, '永兴县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431024, '嘉禾县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431025, '临武县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431026, '汝城县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431027, '桂东县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431028, '安仁县', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431081, '资兴市', 431000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431100, '永州市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431102, '零陵区', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431103, '冷水滩区', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431122, '东安县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431123, '双牌县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431124, '道县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431125, '江永县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431126, '宁远县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431127, '蓝山县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431128, '新田县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431129, '江华瑶族自治县', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431171, '永州经济技术开发区', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431173, '永州市回龙圩管理区', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431181, '祁阳市', 431100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431200, '怀化市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431202, '鹤城区', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431221, '中方县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431222, '沅陵县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431223, '辰溪县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431224, '溆浦县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431225, '会同县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431226, '麻阳苗族自治县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431227, '新晃侗族自治县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431228, '芷江侗族自治县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431229, '靖州苗族侗族自治县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431230, '通道侗族自治县', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431271, '怀化市洪江管理区', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431281, '洪江市', 431200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431300, '娄底市', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431302, '娄星区', 431300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431321, '双峰县', 431300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431322, '新化县', 431300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431381, '冷水江市', 431300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (431382, '涟源市', 431300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433100, '湘西土家族苗族自治州', 430000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433101, '吉首市', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433122, '泸溪县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433123, '凤凰县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433124, '花垣县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433125, '保靖县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433126, '古丈县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433127, '永顺县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (433130, '龙山县', 433100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440000, '广东省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440100, '广州市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440103, '荔湾区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440104, '越秀区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440105, '海珠区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440106, '天河区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440111, '白云区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440112, '黄埔区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440113, '番禺区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440114, '花都区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440115, '南沙区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440117, '从化区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440118, '增城区', 440100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440200, '韶关市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440203, '武江区', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440204, '浈江区', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440205, '曲江区', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440222, '始兴县', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440224, '仁化县', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440229, '翁源县', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440232, '乳源瑶族自治县', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440233, '新丰县', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440281, '乐昌市', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440282, '南雄市', 440200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440300, '深圳市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440303, '罗湖区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440304, '福田区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440305, '南山区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440306, '宝安区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440307, '龙岗区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440308, '盐田区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440309, '龙华区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440310, '坪山区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440311, '光明区', 440300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440400, '珠海市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440402, '香洲区', 440400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440403, '斗门区', 440400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440404, '金湾区', 440400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440500, '汕头市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440507, '龙湖区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440511, '金平区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440512, '濠江区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440513, '潮阳区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440514, '潮南区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440515, '澄海区', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440523, '南澳县', 440500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440600, '佛山市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440604, '禅城区', 440600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440605, '南海区', 440600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440606, '顺德区', 440600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440607, '三水区', 440600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440608, '高明区', 440600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440700, '江门市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440703, '蓬江区', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440704, '江海区', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440705, '新会区', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440781, '台山市', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440783, '开平市', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440784, '鹤山市', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440785, '恩平市', 440700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440800, '湛江市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440802, '赤坎区', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440803, '霞山区', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440804, '坡头区', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440811, '麻章区', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440823, '遂溪县', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440825, '徐闻县', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440881, '廉江市', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440882, '雷州市', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440883, '吴川市', 440800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440900, '茂名市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440902, '茂南区', 440900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440904, '电白区', 440900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440981, '高州市', 440900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440982, '化州市', 440900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (440983, '信宜市', 440900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441200, '肇庆市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441202, '端州区', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441203, '鼎湖区', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441204, '高要区', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441223, '广宁县', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441224, '怀集县', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441225, '封开县', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441226, '德庆县', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441284, '四会市', 441200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441300, '惠州市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441302, '惠城区', 441300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441303, '惠阳区', 441300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441322, '博罗县', 441300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441323, '惠东县', 441300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441324, '龙门县', 441300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441400, '梅州市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441402, '梅江区', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441403, '梅县区', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441422, '大埔县', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441423, '丰顺县', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441424, '五华县', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441426, '平远县', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441427, '蕉岭县', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441481, '兴宁市', 441400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441500, '汕尾市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441502, '城区', 441500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441521, '海丰县', 441500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441523, '陆河县', 441500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441581, '陆丰市', 441500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441600, '河源市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441602, '源城区', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441621, '紫金县', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441622, '龙川县', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441623, '连平县', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441624, '和平县', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441625, '东源县', 441600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441700, '阳江市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441702, '江城区', 441700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441704, '阳东区', 441700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441721, '阳西县', 441700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441781, '阳春市', 441700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441800, '清远市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441802, '清城区', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441803, '清新区', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441821, '佛冈县', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441823, '阳山县', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441825, '连山壮族瑶族自治县', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441826, '连南瑶族自治县', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441881, '英德市', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441882, '连州市', 441800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (441900, '东莞市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (442000, '中山市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445100, '潮州市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445102, '湘桥区', 445100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445103, '潮安区', 445100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445122, '饶平县', 445100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445200, '揭阳市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445202, '榕城区', 445200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445203, '揭东区', 445200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445222, '揭西县', 445200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445224, '惠来县', 445200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445281, '普宁市', 445200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445300, '云浮市', 440000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445302, '云城区', 445300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445303, '云安区', 445300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445321, '新兴县', 445300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445322, '郁南县', 445300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (445381, '罗定市', 445300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450000, '广西壮族自治区', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450100, '南宁市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450102, '兴宁区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450103, '青秀区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450105, '江南区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450107, '西乡塘区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450108, '良庆区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450109, '邕宁区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450110, '武鸣区', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450123, '隆安县', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450124, '马山县', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450125, '上林县', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450126, '宾阳县', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450181, '横州市', 450100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450200, '柳州市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450202, '城中区', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450203, '鱼峰区', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450204, '柳南区', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450205, '柳北区', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450206, '柳江区', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450222, '柳城县', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450223, '鹿寨县', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450224, '融安县', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450225, '融水苗族自治县', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450226, '三江侗族自治县', 450200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450300, '桂林市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450302, '秀峰区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450303, '叠彩区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450304, '象山区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450305, '七星区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450311, '雁山区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450312, '临桂区', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450321, '阳朔县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450323, '灵川县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450324, '全州县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450325, '兴安县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450326, '永福县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450327, '灌阳县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450328, '龙胜各族自治县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450329, '资源县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450330, '平乐县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450332, '恭城瑶族自治县', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450381, '荔浦市', 450300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450400, '梧州市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450403, '万秀区', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450405, '长洲区', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450406, '龙圩区', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450421, '苍梧县', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450422, '藤县', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450423, '蒙山县', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450481, '岑溪市', 450400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450500, '北海市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450502, '海城区', 450500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450503, '银海区', 450500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450512, '铁山港区', 450500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450521, '合浦县', 450500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450600, '防城港市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450602, '港口区', 450600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450603, '防城区', 450600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450621, '上思县', 450600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450681, '东兴市', 450600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450700, '钦州市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450702, '钦南区', 450700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450703, '钦北区', 450700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450721, '灵山县', 450700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450722, '浦北县', 450700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450800, '贵港市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450802, '港北区', 450800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450803, '港南区', 450800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450804, '覃塘区', 450800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450821, '平南县', 450800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450881, '桂平市', 450800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450900, '玉林市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450902, '玉州区', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450903, '福绵区', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450921, '容县', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450922, '陆川县', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450923, '博白县', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450924, '兴业县', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (450981, '北流市', 450900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451000, '百色市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451002, '右江区', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451003, '田阳区', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451022, '田东县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451024, '德保县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451026, '那坡县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451027, '凌云县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451028, '乐业县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451029, '田林县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451030, '西林县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451031, '隆林各族自治县', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451081, '靖西市', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451082, '平果市', 451000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451100, '贺州市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451102, '八步区', 451100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451103, '平桂区', 451100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451121, '昭平县', 451100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451122, '钟山县', 451100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451123, '富川瑶族自治县', 451100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451200, '河池市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451202, '金城江区', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451203, '宜州区', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451221, '南丹县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451222, '天峨县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451223, '凤山县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451224, '东兰县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451225, '罗城仫佬族自治县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451226, '环江毛南族自治县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451227, '巴马瑶族自治县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451228, '都安瑶族自治县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451229, '大化瑶族自治县', 451200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451300, '来宾市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451302, '兴宾区', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451321, '忻城县', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451322, '象州县', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451323, '武宣县', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451324, '金秀瑶族自治县', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451381, '合山市', 451300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451400, '崇左市', 450000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451402, '江州区', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451421, '扶绥县', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451422, '宁明县', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451423, '龙州县', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451424, '大新县', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451425, '天等县', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (451481, '凭祥市', 451400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460000, '海南省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460100, '海口市', 460000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460105, '秀英区', 460100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460106, '龙华区', 460100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460107, '琼山区', 460100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460108, '美兰区', 460100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460200, '三亚市', 460000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460202, '海棠区', 460200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460203, '吉阳区', 460200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460204, '天涯区', 460200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460205, '崖州区', 460200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460300, '三沙市', 460000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460321, '西沙群岛', 460300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460322, '南沙群岛', 460300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460323, '中沙群岛的岛礁及其海域', 460300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (460400, '儋州市', 460000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469000, '省直辖县', 460000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469001, '五指山市', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469002, '琼海市', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469005, '文昌市', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469006, '万宁市', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469007, '东方市', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469021, '定安县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469022, '屯昌县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469023, '澄迈县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469024, '临高县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469025, '白沙黎族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469026, '昌江黎族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469027, '乐东黎族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469028, '陵水黎族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469029, '保亭黎族苗族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (469030, '琼中黎族苗族自治县', 469000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500000, '重庆市', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500100, '重庆市', 500000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500101, '万州区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500102, '涪陵区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500103, '渝中区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500104, '大渡口区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500105, '江北区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500106, '沙坪坝区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500107, '九龙坡区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500108, '南岸区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500109, '北碚区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500110, '綦江区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500111, '大足区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500112, '渝北区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500113, '巴南区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500114, '黔江区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500115, '长寿区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500116, '江津区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500117, '合川区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500118, '永川区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500119, '南川区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500120, '璧山区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500151, '铜梁区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500152, '潼南区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500153, '荣昌区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500154, '开州区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500155, '梁平区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500156, '武隆区', 500100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500200, '重庆市', 500000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500229, '城口县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500230, '丰都县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500231, '垫江县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500233, '忠县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500235, '云阳县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500236, '奉节县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500237, '巫山县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500238, '巫溪县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500240, '石柱土家族自治县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500241, '秀山土家族苗族自治县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500242, '酉阳土家族苗族自治县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (500243, '彭水苗族土家族自治县', 500200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510000, '四川省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510100, '成都市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510104, '锦江区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510105, '青羊区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510106, '金牛区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510107, '武侯区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510108, '成华区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510112, '龙泉驿区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510113, '青白江区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510114, '新都区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510115, '温江区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510116, '双流区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510117, '郫都区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510118, '新津区', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510121, '金堂县', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510129, '大邑县', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510131, '蒲江县', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510181, '都江堰市', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510182, '彭州市', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510183, '邛崃市', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510184, '崇州市', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510185, '简阳市', 510100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510300, '自贡市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510302, '自流井区', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510303, '贡井区', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510304, '大安区', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510311, '沿滩区', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510321, '荣县', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510322, '富顺县', 510300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510400, '攀枝花市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510402, '东区', 510400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510403, '西区', 510400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510411, '仁和区', 510400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510421, '米易县', 510400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510422, '盐边县', 510400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510500, '泸州市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510502, '江阳区', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510503, '纳溪区', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510504, '龙马潭区', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510521, '泸县', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510522, '合江县', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510524, '叙永县', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510525, '古蔺县', 510500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510600, '德阳市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510603, '旌阳区', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510604, '罗江区', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510623, '中江县', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510681, '广汉市', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510682, '什邡市', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510683, '绵竹市', 510600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510700, '绵阳市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510703, '涪城区', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510704, '游仙区', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510705, '安州区', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510722, '三台县', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510723, '盐亭县', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510725, '梓潼县', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510726, '北川羌族自治县', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510727, '平武县', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510781, '江油市', 510700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510800, '广元市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510802, '利州区', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510811, '昭化区', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510812, '朝天区', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510821, '旺苍县', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510822, '青川县', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510823, '剑阁县', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510824, '苍溪县', 510800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510900, '遂宁市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510903, '船山区', 510900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510904, '安居区', 510900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510921, '蓬溪县', 510900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510923, '大英县', 510900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (510981, '射洪市', 510900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511000, '内江市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511002, '市中区', 511000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511011, '东兴区', 511000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511024, '威远县', 511000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511025, '资中县', 511000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511083, '隆昌市', 511000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511100, '乐山市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511102, '市中区', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511111, '沙湾区', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511112, '五通桥区', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511113, '金口河区', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511123, '犍为县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511124, '井研县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511126, '夹江县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511129, '沐川县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511132, '峨边彝族自治县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511133, '马边彝族自治县', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511181, '峨眉山市', 511100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511300, '南充市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511302, '顺庆区', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511303, '高坪区', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511304, '嘉陵区', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511321, '南部县', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511322, '营山县', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511323, '蓬安县', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511324, '仪陇县', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511325, '西充县', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511381, '阆中市', 511300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511400, '眉山市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511402, '东坡区', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511403, '彭山区', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511421, '仁寿县', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511423, '洪雅县', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511424, '丹棱县', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511425, '青神县', 511400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511500, '宜宾市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511502, '翠屏区', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511503, '南溪区', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511504, '叙州区', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511523, '江安县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511524, '长宁县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511525, '高县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511526, '珙县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511527, '筠连县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511528, '兴文县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511529, '屏山县', 511500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511600, '广安市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511602, '广安区', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511603, '前锋区', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511621, '岳池县', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511622, '武胜县', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511623, '邻水县', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511681, '华蓥市', 511600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511700, '达州市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511702, '通川区', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511703, '达川区', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511722, '宣汉县', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511723, '开江县', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511724, '大竹县', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511725, '渠县', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511781, '万源市', 511700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511800, '雅安市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511802, '雨城区', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511803, '名山区', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511822, '荥经县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511823, '汉源县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511824, '石棉县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511825, '天全县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511826, '芦山县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511827, '宝兴县', 511800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511900, '巴中市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511902, '巴州区', 511900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511903, '恩阳区', 511900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511921, '通江县', 511900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511922, '南江县', 511900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (511923, '平昌县', 511900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (512000, '资阳市', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (512002, '雁江区', 512000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (512021, '安岳县', 512000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (512022, '乐至县', 512000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513200, '阿坝藏族羌族自治州', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513201, '马尔康市', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513221, '汶川县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513222, '理县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513223, '茂县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513224, '松潘县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513225, '九寨沟县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513226, '金川县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513227, '小金县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513228, '黑水县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513230, '壤塘县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513231, '阿坝县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513232, '若尔盖县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513233, '红原县', 513200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513300, '甘孜藏族自治州', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513301, '康定市', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513322, '泸定县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513323, '丹巴县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513324, '九龙县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513325, '雅江县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513326, '道孚县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513327, '炉霍县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513328, '甘孜县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513329, '新龙县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513330, '德格县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513331, '白玉县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513332, '石渠县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513333, '色达县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513334, '理塘县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513335, '巴塘县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513336, '乡城县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513337, '稻城县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513338, '得荣县', 513300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513400, '凉山彝族自治州', 510000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513401, '西昌市', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513402, '会理市', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513422, '木里藏族自治县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513423, '盐源县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513424, '德昌县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513426, '会东县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513427, '宁南县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513428, '普格县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513429, '布拖县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513430, '金阳县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513431, '昭觉县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513432, '喜德县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513433, '冕宁县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513434, '越西县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513435, '甘洛县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513436, '美姑县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (513437, '雷波县', 513400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520000, '贵州省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520100, '贵阳市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520102, '南明区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520103, '云岩区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520111, '花溪区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520112, '乌当区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520113, '白云区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520115, '观山湖区', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520121, '开阳县', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520122, '息烽县', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520123, '修文县', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520181, '清镇市', 520100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520200, '六盘水市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520201, '钟山区', 520200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520203, '六枝特区', 520200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520204, '水城区', 520200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520281, '盘州市', 520200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520300, '遵义市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520302, '红花岗区', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520303, '汇川区', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520304, '播州区', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520322, '桐梓县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520323, '绥阳县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520324, '正安县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520325, '道真仡佬族苗族自治县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520326, '务川仡佬族苗族自治县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520327, '凤冈县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520328, '湄潭县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520329, '余庆县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520330, '习水县', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520381, '赤水市', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520382, '仁怀市', 520300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520400, '安顺市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520402, '西秀区', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520403, '平坝区', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520422, '普定县', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520423, '镇宁布依族苗族自治县', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520424, '关岭布依族苗族自治县', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520425, '紫云苗族布依族自治县', 520400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520500, '毕节市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520502, '七星关区', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520521, '大方县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520523, '金沙县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520524, '织金县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520525, '纳雍县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520526, '威宁彝族回族苗族自治县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520527, '赫章县', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520581, '黔西市', 520500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520600, '铜仁市', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520602, '碧江区', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520603, '万山区', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520621, '江口县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520622, '玉屏侗族自治县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520623, '石阡县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520624, '思南县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520625, '印江土家族苗族自治县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520626, '德江县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520627, '沿河土家族自治县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (520628, '松桃苗族自治县', 520600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522300, '黔西南布依族苗族自治州', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522301, '兴义市', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522302, '兴仁市', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522323, '普安县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522324, '晴隆县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522325, '贞丰县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522326, '望谟县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522327, '册亨县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522328, '安龙县', 522300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522600, '黔东南苗族侗族自治州', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522601, '凯里市', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522622, '黄平县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522623, '施秉县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522624, '三穗县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522625, '镇远县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522626, '岑巩县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522627, '天柱县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522628, '锦屏县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522629, '剑河县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522630, '台江县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522631, '黎平县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522632, '榕江县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522633, '从江县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522634, '雷山县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522635, '麻江县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522636, '丹寨县', 522600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522700, '黔南布依族苗族自治州', 520000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522701, '都匀市', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522702, '福泉市', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522722, '荔波县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522723, '贵定县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522725, '瓮安县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522726, '独山县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522727, '平塘县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522728, '罗甸县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522729, '长顺县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522730, '龙里县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522731, '惠水县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (522732, '三都水族自治县', 522700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530000, '云南省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530100, '昆明市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530102, '五华区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530103, '盘龙区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530111, '官渡区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530112, '西山区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530113, '东川区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530114, '呈贡区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530115, '晋宁区', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530124, '富民县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530125, '宜良县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530126, '石林彝族自治县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530127, '嵩明县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530128, '禄劝彝族苗族自治县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530129, '寻甸回族彝族自治县', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530181, '安宁市', 530100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530300, '曲靖市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530302, '麒麟区', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530303, '沾益区', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530304, '马龙区', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530322, '陆良县', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530323, '师宗县', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530324, '罗平县', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530325, '富源县', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530326, '会泽县', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530381, '宣威市', 530300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530400, '玉溪市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530402, '红塔区', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530403, '江川区', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530423, '通海县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530424, '华宁县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530425, '易门县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530426, '峨山彝族自治县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530427, '新平彝族傣族自治县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530428, '元江哈尼族彝族傣族自治县', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530481, '澄江市', 530400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530500, '保山市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530502, '隆阳区', 530500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530521, '施甸县', 530500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530523, '龙陵县', 530500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530524, '昌宁县', 530500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530581, '腾冲市', 530500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530600, '昭通市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530602, '昭阳区', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530621, '鲁甸县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530622, '巧家县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530623, '盐津县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530624, '大关县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530625, '永善县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530626, '绥江县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530627, '镇雄县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530628, '彝良县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530629, '威信县', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530681, '水富市', 530600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530700, '丽江市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530702, '古城区', 530700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530721, '玉龙纳西族自治县', 530700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530722, '永胜县', 530700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530723, '华坪县', 530700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530724, '宁蒗彝族自治县', 530700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530800, '普洱市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530802, '思茅区', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530821, '宁洱哈尼族彝族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530822, '墨江哈尼族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530823, '景东彝族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530824, '景谷傣族彝族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530825, '镇沅彝族哈尼族拉祜族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530826, '江城哈尼族彝族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530827, '孟连傣族拉祜族佤族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530828, '澜沧拉祜族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530829, '西盟佤族自治县', 530800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530900, '临沧市', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530902, '临翔区', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530921, '凤庆县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530922, '云县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530923, '永德县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530924, '镇康县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530925, '双江拉祜族佤族布朗族傣族自治县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530926, '耿马傣族佤族自治县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (530927, '沧源佤族自治县', 530900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532300, '楚雄彝族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532301, '楚雄市', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532302, '禄丰市', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532322, '双柏县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532323, '牟定县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532324, '南华县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532325, '姚安县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532326, '大姚县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532327, '永仁县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532328, '元谋县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532329, '武定县', 532300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532500, '红河哈尼族彝族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532501, '个旧市', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532502, '开远市', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532503, '蒙自市', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532504, '弥勒市', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532523, '屏边苗族自治县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532524, '建水县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532525, '石屏县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532527, '泸西县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532528, '元阳县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532529, '红河县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532530, '金平苗族瑶族傣族自治县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532531, '绿春县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532532, '河口瑶族自治县', 532500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532600, '文山壮族苗族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532601, '文山市', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532622, '砚山县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532623, '西畴县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532624, '麻栗坡县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532625, '马关县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532626, '丘北县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532627, '广南县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532628, '富宁县', 532600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532800, '西双版纳傣族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532801, '景洪市', 532800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532822, '勐海县', 532800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532823, '勐腊县', 532800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532900, '大理白族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532901, '大理市', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532922, '漾濞彝族自治县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532923, '祥云县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532924, '宾川县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532925, '弥渡县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532926, '南涧彝族自治县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532927, '巍山彝族回族自治县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532928, '永平县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532929, '云龙县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532930, '洱源县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532931, '剑川县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (532932, '鹤庆县', 532900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533100, '德宏傣族景颇族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533102, '瑞丽市', 533100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533103, '芒市', 533100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533122, '梁河县', 533100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533123, '盈江县', 533100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533124, '陇川县', 533100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533300, '怒江傈僳族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533301, '泸水市', 533300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533323, '福贡县', 533300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533324, '贡山独龙族怒族自治县', 533300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533325, '兰坪白族普米族自治县', 533300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533400, '迪庆藏族自治州', 530000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533401, '香格里拉市', 533400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533422, '德钦县', 533400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (533423, '维西傈僳族自治县', 533400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540000, '西藏自治区', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540100, '拉萨市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540102, '城关区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540103, '堆龙德庆区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540104, '达孜区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540121, '林周县', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540122, '当雄县', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540123, '尼木县', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540124, '曲水县', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540127, '墨竹工卡县', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540171, '格尔木藏青工业园区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540172, '拉萨经济技术开发区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540173, '西藏文化旅游创意园区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540174, '达孜工业园区', 540100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540200, '日喀则市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540202, '桑珠孜区', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540221, '南木林县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540222, '江孜县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540223, '定日县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540224, '萨迦县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540225, '拉孜县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540226, '昂仁县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540227, '谢通门县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540228, '白朗县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540229, '仁布县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540230, '康马县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540231, '定结县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540232, '仲巴县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540233, '亚东县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540234, '吉隆县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540235, '聂拉木县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540236, '萨嘎县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540237, '岗巴县', 540200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540300, '昌都市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540302, '卡若区', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540321, '江达县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540322, '贡觉县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540323, '类乌齐县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540324, '丁青县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540325, '察雅县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540326, '八宿县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540327, '左贡县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540328, '芒康县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540329, '洛隆县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540330, '边坝县', 540300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540400, '林芝市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540402, '巴宜区', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540421, '工布江达县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540422, '米林县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540423, '墨脱县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540424, '波密县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540425, '察隅县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540426, '朗县', 540400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540500, '山南市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540502, '乃东区', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540521, '扎囊县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540522, '贡嘎县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540523, '桑日县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540524, '琼结县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540525, '曲松县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540526, '措美县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540527, '洛扎县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540528, '加查县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540529, '隆子县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540530, '错那县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540531, '浪卡子县', 540500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540600, '那曲市', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540602, '色尼区', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540621, '嘉黎县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540622, '比如县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540623, '聂荣县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540624, '安多县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540625, '申扎县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540626, '索县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540627, '班戈县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540628, '巴青县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540629, '尼玛县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (540630, '双湖县', 540600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542500, '阿里地区', 540000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542521, '普兰县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542522, '札达县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542523, '噶尔县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542524, '日土县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542525, '革吉县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542526, '改则县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (542527, '措勤县', 542500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610000, '陕西省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610100, '西安市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610102, '新城区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610103, '碑林区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610104, '莲湖区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610111, '灞桥区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610112, '未央区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610113, '雁塔区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610114, '阎良区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610115, '临潼区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610116, '长安区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610117, '高陵区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610118, '鄠邑区', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610122, '蓝田县', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610124, '周至县', 610100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610200, '铜川市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610202, '王益区', 610200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610203, '印台区', 610200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610204, '耀州区', 610200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610222, '宜君县', 610200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610300, '宝鸡市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610302, '渭滨区', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610303, '金台区', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610304, '陈仓区', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610305, '凤翔区', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610323, '岐山县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610324, '扶风县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610326, '眉县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610327, '陇县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610328, '千阳县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610329, '麟游县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610330, '凤县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610331, '太白县', 610300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610400, '咸阳市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610402, '秦都区', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610403, '杨陵区', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610404, '渭城区', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610422, '三原县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610423, '泾阳县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610424, '乾县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610425, '礼泉县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610426, '永寿县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610428, '长武县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610429, '旬邑县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610430, '淳化县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610431, '武功县', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610481, '兴平市', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610482, '彬州市', 610400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610500, '渭南市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610502, '临渭区', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610503, '华州区', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610522, '潼关县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610523, '大荔县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610524, '合阳县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610525, '澄城县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610526, '蒲城县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610527, '白水县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610528, '富平县', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610581, '韩城市', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610582, '华阴市', 610500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610600, '延安市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610602, '宝塔区', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610603, '安塞区', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610621, '延长县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610622, '延川县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610625, '志丹县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610626, '吴起县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610627, '甘泉县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610628, '富县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610629, '洛川县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610630, '宜川县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610631, '黄龙县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610632, '黄陵县', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610681, '子长市', 610600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610700, '汉中市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610702, '汉台区', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610703, '南郑区', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610722, '城固县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610723, '洋县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610724, '西乡县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610725, '勉县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610726, '宁强县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610727, '略阳县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610728, '镇巴县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610729, '留坝县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610730, '佛坪县', 610700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610800, '榆林市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610802, '榆阳区', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610803, '横山区', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610822, '府谷县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610824, '靖边县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610825, '定边县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610826, '绥德县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610827, '米脂县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610828, '佳县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610829, '吴堡县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610830, '清涧县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610831, '子洲县', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610881, '神木市', 610800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610900, '安康市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610902, '汉滨区', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610921, '汉阴县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610922, '石泉县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610923, '宁陕县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610924, '紫阳县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610925, '岚皋县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610926, '平利县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610927, '镇坪县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610929, '白河县', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (610981, '旬阳市', 610900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611000, '商洛市', 610000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611002, '商州区', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611021, '洛南县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611022, '丹凤县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611023, '商南县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611024, '山阳县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611025, '镇安县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (611026, '柞水县', 611000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620000, '甘肃省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620100, '兰州市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620102, '城关区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620103, '七里河区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620104, '西固区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620105, '安宁区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620111, '红古区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620121, '永登县', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620122, '皋兰县', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620123, '榆中县', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620171, '兰州新区', 620100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620200, '嘉峪关市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620201, '市辖区', 620200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620300, '金昌市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620302, '金川区', 620300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620321, '永昌县', 620300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620400, '白银市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620402, '白银区', 620400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620403, '平川区', 620400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620421, '靖远县', 620400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620422, '会宁县', 620400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620423, '景泰县', 620400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620500, '天水市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620502, '秦州区', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620503, '麦积区', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620521, '清水县', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620522, '秦安县', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620523, '甘谷县', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620524, '武山县', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620525, '张家川回族自治县', 620500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620600, '武威市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620602, '凉州区', 620600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620621, '民勤县', 620600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620622, '古浪县', 620600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620623, '天祝藏族自治县', 620600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620700, '张掖市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620702, '甘州区', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620721, '肃南裕固族自治县', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620722, '民乐县', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620723, '临泽县', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620724, '高台县', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620725, '山丹县', 620700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620800, '平凉市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620802, '崆峒区', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620821, '泾川县', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620822, '灵台县', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620823, '崇信县', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620825, '庄浪县', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620826, '静宁县', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620881, '华亭市', 620800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620900, '酒泉市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620902, '肃州区', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620921, '金塔县', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620922, '瓜州县', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620923, '肃北蒙古族自治县', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620924, '阿克塞哈萨克族自治县', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620981, '玉门市', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (620982, '敦煌市', 620900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621000, '庆阳市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621002, '西峰区', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621021, '庆城县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621022, '环县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621023, '华池县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621024, '合水县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621025, '正宁县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621026, '宁县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621027, '镇原县', 621000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621100, '定西市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621102, '安定区', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621121, '通渭县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621122, '陇西县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621123, '渭源县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621124, '临洮县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621125, '漳县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621126, '岷县', 621100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621200, '陇南市', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621202, '武都区', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621221, '成县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621222, '文县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621223, '宕昌县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621224, '康县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621225, '西和县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621226, '礼县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621227, '徽县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (621228, '两当县', 621200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622900, '临夏回族自治州', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622901, '临夏市', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622921, '临夏县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622922, '康乐县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622923, '永靖县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622924, '广河县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622925, '和政县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622926, '东乡族自治县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (622927, '积石山保安族东乡族撒拉族自治县', 622900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623000, '甘南藏族自治州', 620000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623001, '合作市', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623021, '临潭县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623022, '卓尼县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623023, '舟曲县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623024, '迭部县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623025, '玛曲县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623026, '碌曲县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (623027, '夏河县', 623000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630000, '青海省', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630100, '西宁市', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630102, '城东区', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630103, '城中区', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630104, '城西区', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630105, '城北区', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630106, '湟中区', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630121, '大通回族土族自治县', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630123, '湟源县', 630100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630200, '海东市', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630202, '乐都区', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630203, '平安区', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630222, '民和回族土族自治县', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630223, '互助土族自治县', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630224, '化隆回族自治县', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (630225, '循化撒拉族自治县', 630200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632200, '海北藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632221, '门源回族自治县', 632200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632222, '祁连县', 632200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632223, '海晏县', 632200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632224, '刚察县', 632200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632300, '黄南藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632301, '同仁市', 632300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632322, '尖扎县', 632300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632323, '泽库县', 632300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632324, '河南蒙古族自治县', 632300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632500, '海南藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632521, '共和县', 632500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632522, '同德县', 632500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632523, '贵德县', 632500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632524, '兴海县', 632500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632525, '贵南县', 632500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632600, '果洛藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632621, '玛沁县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632622, '班玛县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632623, '甘德县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632624, '达日县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632625, '久治县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632626, '玛多县', 632600, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632700, '玉树藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632701, '玉树市', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632722, '杂多县', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632723, '称多县', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632724, '治多县', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632725, '囊谦县', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632726, '曲麻莱县', 632700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632800, '海西蒙古族藏族自治州', 630000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632801, '格尔木市', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632802, '德令哈市', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632803, '茫崖市', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632821, '乌兰县', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632822, '都兰县', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632823, '天峻县', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (632857, '大柴旦行政委员会', 632800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640000, '宁夏回族自治区', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640100, '银川市', 640000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640104, '兴庆区', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640105, '西夏区', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640106, '金凤区', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640121, '永宁县', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640122, '贺兰县', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640181, '灵武市', 640100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640200, '石嘴山市', 640000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640202, '大武口区', 640200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640205, '惠农区', 640200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640221, '平罗县', 640200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640300, '吴忠市', 640000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640302, '利通区', 640300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640303, '红寺堡区', 640300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640323, '盐池县', 640300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640324, '同心县', 640300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640381, '青铜峡市', 640300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640400, '固原市', 640000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640402, '原州区', 640400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640422, '西吉县', 640400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640423, '隆德县', 640400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640424, '泾源县', 640400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640425, '彭阳县', 640400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640500, '中卫市', 640000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640502, '沙坡头区', 640500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640521, '中宁县', 640500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (640522, '海原县', 640500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650000, '新疆维吾尔自治区', 0, 1, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650100, '乌鲁木齐市', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650102, '天山区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650103, '沙依巴克区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650104, '新市区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650105, '水磨沟区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650106, '头屯河区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650107, '达坂城区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650109, '米东区', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650121, '乌鲁木齐县', 650100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650200, '克拉玛依市', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650202, '独山子区', 650200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650203, '克拉玛依区', 650200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650204, '白碱滩区', 650200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650205, '乌尔禾区', 650200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650400, '吐鲁番市', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650402, '高昌区', 650400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650421, '鄯善县', 650400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650422, '托克逊县', 650400, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650500, '哈密市', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650502, '伊州区', 650500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650521, '巴里坤哈萨克自治县', 650500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (650522, '伊吾县', 650500, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652300, '昌吉回族自治州', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652301, '昌吉市', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652302, '阜康市', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652323, '呼图壁县', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652324, '玛纳斯县', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652325, '奇台县', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652327, '吉木萨尔县', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652328, '木垒哈萨克自治县', 652300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652700, '博尔塔拉蒙古自治州', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652701, '博乐市', 652700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652702, '阿拉山口市', 652700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652722, '精河县', 652700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652723, '温泉县', 652700, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652800, '巴音郭楞蒙古自治州', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652801, '库尔勒市', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652822, '轮台县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652823, '尉犁县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652824, '若羌县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652825, '且末县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652826, '焉耆回族自治县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652827, '和静县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652828, '和硕县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652829, '博湖县', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652871, '库尔勒经济技术开发区', 652800, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652900, '阿克苏地区', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652901, '阿克苏市', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652902, '库车市', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652922, '温宿县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652924, '沙雅县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652925, '新和县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652926, '拜城县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652927, '乌什县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652928, '阿瓦提县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (652929, '柯坪县', 652900, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653000, '克孜勒苏柯尔克孜自治州', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653001, '阿图什市', 653000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653022, '阿克陶县', 653000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653023, '阿合奇县', 653000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653024, '乌恰县', 653000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653100, '喀什地区', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653101, '喀什市', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653121, '疏附县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653122, '疏勒县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653123, '英吉沙县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653124, '泽普县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653125, '莎车县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653126, '叶城县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653127, '麦盖提县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653128, '岳普湖县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653129, '伽师县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653130, '巴楚县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653131, '塔什库尔干塔吉克自治县', 653100, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653200, '和田地区', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653201, '和田市', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653221, '和田县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653222, '墨玉县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653223, '皮山县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653224, '洛浦县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653225, '策勒县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653226, '于田县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (653227, '民丰县', 653200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654000, '伊犁哈萨克自治州', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654002, '伊宁市', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654003, '奎屯市', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654004, '霍尔果斯市', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654021, '伊宁县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654022, '察布查尔锡伯自治县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654023, '霍城县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654024, '巩留县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654025, '新源县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654026, '昭苏县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654027, '特克斯县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654028, '尼勒克县', 654000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654200, '塔城地区', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654201, '塔城市', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654202, '乌苏市', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654203, '沙湾市', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654221, '额敏县', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654224, '托里县', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654225, '裕民县', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654226, '和布克赛尔蒙古自治县', 654200, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654300, '阿勒泰地区', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654301, '阿勒泰市', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654321, '布尔津县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654322, '富蕴县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654323, '福海县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654324, '哈巴河县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654325, '青河县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (654326, '吉木乃县', 654300, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659000, '自治区直辖县', 650000, 2, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659001, '石河子市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659002, '阿拉尔市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659003, '图木舒克市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659004, '五家渠市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659005, '北屯市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659006, '铁门关市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659007, '双河市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659008, '可克达拉市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659009, '昆玉市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659010, '胡杨河市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);
INSERT INTO `sys_district_base` VALUES (659011, '新星市', 659000, 3, '', '', 0.00000000, 0.00000000, 0, 1001);

-- ----------------------------
-- Table structure for sys_express_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_express_base`;
CREATE TABLE `sys_express_base`  (
  `express_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '快递编号',
  `express_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递名称',
  `express_pinyin` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递编码',
  `express_pinyin_100` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递公司100',
  `express_site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递官网',
  `express_is_fav` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否常用',
  `express_order` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '快递排序:越小越靠前',
  `express_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '启用状态(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`express_id`) USING BTREE,
  INDEX `express_name`(`express_enable`, `express_order`) USING BTREE,
  INDEX `express_name_2`(`express_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '快递表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_feedback_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_feedback_base`;
CREATE TABLE `sys_feedback_base`  (
  `feedback_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '反馈编号',
  `feedback_category_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '分类编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `feedback_question` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '反馈问题:在这里描述您遇到的问题',
  `feedback_question_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面链接',
  `feedback_question_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '反馈问题',
  `feedback_question_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '反馈时间',
  `feedback_question_answer_time` timestamp(0) NULL DEFAULT NULL COMMENT '回复时间',
  `feedback_question_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '举报状态(BOOL):0-未处理;1-已处理',
  `feedback_question_result` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '举报处理结果(ENUM):1-无效举报;2-恶意举报;3-有效举报',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复人员',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`feedback_id`) USING BTREE,
  INDEX `plantform_feedback_ibfk_1`(`feedback_category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台反馈表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_feedback_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_feedback_category`;
CREATE TABLE `sys_feedback_category`  (
  `feedback_category_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `feedback_category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名称',
  `feedback_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1002 COMMENT '反馈类型',
  `feedback_category_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`feedback_category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台反馈主题分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_feedback_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_feedback_type`;
CREATE TABLE `sys_feedback_type`  (
  `feedback_type_id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '类型编号',
  `feedback_type_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型名称',
  `feedback_type_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `feedback_type_genus` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '所属身份(ENUM):1-会员所属;2-经销商所属',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`feedback_type_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台反馈类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_filter_keyword
-- ----------------------------
DROP TABLE IF EXISTS `sys_filter_keyword`;
CREATE TABLE `sys_filter_keyword`  (
  `filter_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '敏感词编号',
  `filter_find` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '查找敏感词',
  `filter_replace` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '替换内容',
  `filter_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '过滤类型(ENUM):1-banned禁止;2-替换replace',
  `filter_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '添加时间',
  `filter_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否系统内置(ENUM): 0-非内置;1-系统内置',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `filter_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用(ENUM): 0-未启用; 1-已启用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`filter_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '敏感词过滤-启用api' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_ip_ban
-- ----------------------------
DROP TABLE IF EXISTS `sys_ip_ban`;
CREATE TABLE `sys_ip_ban`  (
  `ban_ip` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `ban_reason` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '原因',
  `ban_unlock_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '解锁时间',
  `ban_lock_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '锁定时间',
  `ban_autorelease` bit(1) NOT NULL DEFAULT b'1' COMMENT '自动打开(BOOL):1-自动解锁; 0-人工解锁',
  `ban_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '锁定状态(ENUM):1-锁定中; 2-已解锁',
  `ban_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '锁定类型(ENUM):1-禁止访问; 2-禁止注册',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ban_ip`) USING BTREE,
  INDEX `ban_ip`(`ban_lock_time`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'IP禁止表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_lang_meta
-- ----------------------------
DROP TABLE IF EXISTS `sys_lang_meta`;
CREATE TABLE `sys_lang_meta`  (
  `meta_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Meta id = data_id-meta_key',
  `data_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '翻译内容Md5',
  `meta_key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属语言',
  `meta_ori` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '源',
  `meta_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '值',
  `meta_datatype` enum('string','json','number') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'string' COMMENT '数据类型',
  `meta_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '时间',
  `meta_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(BOOL):0-非内置;1-内置',
  `table_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表明',
  `primary_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键',
  `column_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字段',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '所属店铺',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`meta_id`) USING BTREE,
  INDEX `meta_key`(`meta_key`) USING BTREE,
  INDEX `data_id`(`data_id`, `meta_key`) USING BTREE,
  INDEX `meta_ori`(`meta_ori`(191), `store_id`, `meta_key`, `table_name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户数据翻译扩展表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_lang_standard
-- ----------------------------
DROP TABLE IF EXISTS `sys_lang_standard`;
CREATE TABLE `sys_lang_standard`  (
  `zh_CN` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '中文',
  `zh_TW` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '繁体中文',
  `en_GB` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '英文',
  `th_TH` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '泰语',
  `es_MX` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '西班牙',
  `ar_SA` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '阿拉伯',
  `vi_VN` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '越南语',
  `tr_TR` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '土耳其',
  `ja_JP` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '日语',
  `id_ID` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '印尼语',
  `de_DE` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '德语',
  `fr_FR` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '法语',
  `pt_PT` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '葡萄牙',
  `it_IT` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '意大利',
  `ru_RU` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '俄罗斯',
  `ro_RO` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '罗马尼亚',
  `az_AZ` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '阿塞拜疆语',
  `el_GR` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '希腊语',
  `fi_FI` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '芬兰语',
  `lv_LV` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '拉脱维亚语',
  `nl_NL` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '荷兰语',
  `da_DK` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '丹麦语',
  `sr_RS` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '塞尔维亚语',
  `pl_PL` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '波兰语',
  `uk_UA` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '乌克兰语',
  `kk_KZ` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '哈萨克斯坦',
  `my_MM` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '缅甸语',
  `ko_KR` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '韩语',
  `ms_MY` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '马来语',
  `time` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '时间',
  `is_imp` bit(1) NOT NULL DEFAULT b'0' COMMENT '重要文字',
  `is_used` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):1-启用;0-禁用',
  `frontend` bit(1) NOT NULL DEFAULT b'0' COMMENT '前端启用(BOOL):1-启用;0-禁用',
  `backend` bit(1) NOT NULL DEFAULT b'0' COMMENT '后端启用(BOOL):1-启用;0-禁用',
  `java` bit(1) NOT NULL DEFAULT b'0' COMMENT '服务端启用(BOOL):1-启用;0-禁用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`zh_CN`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '语言翻译表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_log_action
-- ----------------------------
DROP TABLE IF EXISTS `sys_log_action`;
CREATE TABLE `sys_log_action`  (
  `log_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `user_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '玩家编号',
  `user_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色账户',
  `user_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `log_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求名称',
  `action_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '行为id:protocal_id -> rights_id',
  `action_type_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作类型编号:right_parent_id',
  `log_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求接口',
  `log_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `log_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求的参数',
  `log_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日志IP',
  `log_date` date NOT NULL COMMENT '日志日期',
  `log_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '记录时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`log_id`) USING BTREE,
  INDEX `player_id`(`user_id`) USING BTREE COMMENT '(null)',
  INDEX `log_date`(`log_date`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户行为日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_log_error
-- ----------------------------
DROP TABLE IF EXISTS `sys_log_error`;
CREATE TABLE `sys_log_error`  (
  `log_error_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志编号',
  `log_error_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '日志类型',
  `log_error_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日志名称',
  `log_error_line` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '日志文件',
  `log_error_time` bigint(20) UNSIGNED NOT NULL COMMENT '日志时间',
  `log_error_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日志内容',
  `log_error_date` date NOT NULL COMMENT '日志日期',
  `log_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '日志时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`log_error_id`) USING BTREE,
  INDEX `log_date`(`log_error_date`) USING BTREE COMMENT '(null)',
  INDEX `log_error_type`(`log_error_type`) USING BTREE COMMENT '(null)',
  INDEX `log_error_name`(`log_error_name`(191)) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统错误日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_market_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_market_category`;
CREATE TABLE `sys_market_category`  (
  `market_category_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '市场分类编号',
  `category_parent_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市场分类父编号',
  `category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市场分类名称',
  `category_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片:客户群不同',
  `category_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `category_level` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '车型分类等级',
  `category_is_leaf` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否叶节点',
  `category_is_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):0-不显示;1-显示',
  `market_latitude` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '纬度',
  `market_longitude` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '经度',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`market_category_id`) USING BTREE,
  INDEX `category_order`(`category_order`) USING BTREE,
  INDEX `category_name`(`category_name`) USING BTREE,
  INDEX `category_parent_id`(`category_parent_id`, `category_order`, `category_is_enable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '市场分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_material_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_material_base`;
CREATE TABLE `sys_material_base`  (
  `material_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '素材编号',
  `material_number` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件md5',
  `gallery_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `material_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件URL',
  `material_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材来源',
  `material_sort` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '素材排序',
  `material_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材path:本地存储',
  `material_type` enum('video','other','image','audio','document') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'image' COMMENT '素材类型',
  `material_image_h` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '素材高度',
  `material_image_w` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '素材宽度',
  `material_size` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '素材大小',
  `material_mime_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材类型',
  `material_alt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材alt',
  `material_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材标题',
  `material_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材描述',
  `material_duration` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '素材时长:（音频/视频）',
  `material_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '素材日期',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`material_id`) USING BTREE,
  INDEX `gallery_id`(`gallery_id`, `material_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '素材表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_material_base
-- ----------------------------
INSERT INTO `sys_material_base` VALUES (2, '', 1001, 0, 0, 'https://shopsuite.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230829/0c2b7826ae2943eca247bc53a4645648.png', '', 0, './resource/public/uploads/videos/10001/20240704/d2gh5qszl1x8ovpcmo.png', 'image', 0, 0, 20637, 'image/png', '资源 2@2x.png', '资源 2@2x.png', '', '', '2024-07-04 12:37:15', 0);

-- ----------------------------
-- Table structure for sys_material_gallery
-- ----------------------------
DROP TABLE IF EXISTS `sys_material_gallery`;
CREATE TABLE `sys_material_gallery`  (
  `gallery_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类编号',
  `gallery_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `gallery_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类描述',
  `gallery_num` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '素材数量',
  `gallery_is_default` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否默认(BOOL):0-否;1-是',
  `gallery_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类排序',
  `gallery_type` enum('video','other','image','audio','document') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'image' COMMENT '分类类型',
  `gallery_cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类封面',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`gallery_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '素材分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_material_gallery
-- ----------------------------
INSERT INTO `sys_material_gallery` VALUES (1001, '系统使用', '', 0, 0, 10, 'image', '', 0, 0, 0);

-- ----------------------------
-- Table structure for sys_message_template
-- ----------------------------
DROP TABLE IF EXISTS `sys_message_template`;
CREATE TABLE `sys_message_template`  (
  `message_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模板编号',
  `message_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模板编码',
  `message_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模板名称',
  `message_email_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮件标题',
  `message_email_content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮件内容',
  `message_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '站内消息',
  `message_sms` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '短信内容',
  `message_app` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'APP内容',
  `message_type` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息类型(ENUM):1-用户;2-商家;3-平台;',
  `message_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '站内通知(BOOL):0-禁用;1-启用',
  `message_sms_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '短息通知(BOOL):0-禁用;1-启用',
  `message_email_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '邮件通知(BOOL):0-禁用;1-启用',
  `message_wechat_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '微信通知(BOOL):0-禁用;1-启用',
  `message_xcx_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '小程序通知(BOOL):0-禁用;1-启用',
  `message_app_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT 'APP推送(BOOL):0-禁用;1-启用',
  `message_sms_force` bit(1) NOT NULL DEFAULT b'0' COMMENT '手机短信(BOOL):0-不强制;1-强制',
  `message_email_force` bit(1) NOT NULL DEFAULT b'0' COMMENT '邮件(BOOL):0-不强制;1-强制',
  `message_app_force` bit(1) NOT NULL DEFAULT b'0' COMMENT 'APP(BOOL):0-不强制;1-强制',
  `message_force` bit(1) NOT NULL DEFAULT b'0' COMMENT '站内信(BOOL):0-不强制;1-强制',
  `message_category` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息分组(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息',
  `message_order` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '消息排序',
  `message_tpl_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模板编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`message_id`) USING BTREE,
  INDEX `message_type`(`message_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息模板表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_number_seq
-- ----------------------------
DROP TABLE IF EXISTS `sys_number_seq`;
CREATE TABLE `sys_number_seq`  (
  `prefix` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前缀',
  `number` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '增长值',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`prefix`) USING BTREE,
  UNIQUE INDEX `prefix`(`prefix`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '编号管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_number_seq
-- ----------------------------
INSERT INTO `sys_number_seq` VALUES ('product_id', 10000, 1001);

-- ----------------------------
-- Table structure for sys_option_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_option_base`;
CREATE TABLE `sys_option_base`  (
  `option_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '选项编号',
  `option_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选项类型',
  `option_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选项名称',
  `option_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '选项排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`option_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '选项表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_page_base
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_base`;
CREATE TABLE `sys_page_base`  (
  `page_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '页面编号',
  `page_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面名称',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `page_buildin` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否内置(BOOL):0-否;1-是',
  `page_type` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型(ENUM):1-WAP;2-PC;3-APP',
  `page_tpl` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '页面布局模板',
  `app_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属APP',
  `page_code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '页面代码',
  `page_nav` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '导航数据',
  `page_config` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面配置',
  `page_share_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分享标题',
  `page_share_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分享图片',
  `page_qrcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分享二维码',
  `page_index` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否首页(BOOL):0-非首页;1-首页',
  `page_gb` bit(1) NOT NULL DEFAULT b'0' COMMENT '拼团首页(BOOL):0-非首页;1-首页',
  `page_activity` bit(1) NOT NULL DEFAULT b'0' COMMENT '活动首页(BOOL):0-非首页;1-首页',
  `page_point` bit(1) NOT NULL DEFAULT b'0' COMMENT '积分首页(BOOL):0-非首页;1-首页',
  `page_gbs` bit(1) NOT NULL DEFAULT b'0' COMMENT '团购首页(BOOL):0-非首页;1-首页',
  `page_package` bit(1) NOT NULL DEFAULT b'0' COMMENT '组合套餐(BOOL):0-非首页;1-首页',
  `page_pfgb` bit(1) NOT NULL DEFAULT b'0' COMMENT '批发团购首页(BOOL):0-非首页;1-首页',
  `page_sns` bit(1) NOT NULL DEFAULT b'0' COMMENT '社区(BOOL):0-非首页;1-首页',
  `page_article` bit(1) NOT NULL DEFAULT b'0' COMMENT '资讯(BOOL):0-非首页;1-首页',
  `page_zerobuy` bit(1) NOT NULL DEFAULT b'0' COMMENT '零元购区(BOOL):0-否;1-是',
  `page_higharea` bit(1) NOT NULL DEFAULT b'0' COMMENT '高额返区(BOOL):0-否;1-是',
  `page_taday` bit(1) NOT NULL DEFAULT b'0' COMMENT '今日爆款(BOOL):0-否;1-是',
  `page_everyday` bit(1) NOT NULL DEFAULT b'0' COMMENT '每日好店(BOOL):0-否;1-是',
  `page_secondkill` bit(1) NOT NULL DEFAULT b'0' COMMENT '整点秒杀(BOOL):0-否;1-是',
  `page_secondday` bit(1) NOT NULL DEFAULT b'0' COMMENT '天天秒淘(BOOL):0-否;1-是',
  `page_rura` bit(1) NOT NULL DEFAULT b'0' COMMENT '设置土特产(BOOL):0-否;1-是',
  `page_likeyou` bit(1) NOT NULL DEFAULT b'0' COMMENT '用户页banner(BOOL):0-否;1-是',
  `page_exchange` bit(1) NOT NULL DEFAULT b'0' COMMENT '兑换专区(BOOL):0-否;1-是',
  `page_new` bit(1) NOT NULL DEFAULT b'0' COMMENT '新品首发(BOOL):0-否;1-是',
  `page_newperson` bit(1) NOT NULL DEFAULT b'0' COMMENT '新人优惠(BOOL):0-否;1-是',
  `page_upgrade` bit(1) NOT NULL DEFAULT b'0' COMMENT '升级VIP(BOOL):0-否;1-是',
  `page_message` bit(1) NOT NULL DEFAULT b'0' COMMENT '信息发布(BOOL):0-否;1-是',
  `page_release` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否发布(BOOL):0-否;1-是',
  `industry_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业编号集合',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`page_id`) USING BTREE,
  INDEX `store_id`(`app_id`, `store_id`, `page_type`) USING BTREE,
  INDEX `page_type`(`page_type`, `page_index`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1695182055304 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_page_base
-- ----------------------------
INSERT INTO `sys_page_base` VALUES (1638151353229, 'pc商城首页_2021', 0, 0, b'0', 2, 0, 0, '', '', '', '', '', '', b'1', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'1', '', 0);
INSERT INTO `sys_page_base` VALUES (1695182055303, '审核专题', 0, 0, b'0', 3, 0, 0, '[{\"id\":1700550641146,\"eltmType\":\"3\",\"bgColor\":\"#0185b2\",\"paddingTop\":\"0\",\"paddingRight\":\"0\",\"paddingBottom\":\"0\",\"paddingLeft\":\"0\",\"borderTopLeftRadius\":\"0\",\"borderTopRightRadius\":\"0\",\"borderBottomLeftRadius\":\"0\",\"borderBottomRightRadius\":\"0\",\"eltm3\":{\"showNbg\":1,\"align\":\"1\",\"padding\":0,\"width\":\"375\",\"height\":\"181\",\"borderRadius\":\"0\",\"progress\":true,\"color\":\"#ffffff\",\"activeColor\":\"#000000\",\"autoplay\":true,\"data\":[{\"id\":1700557994137,\"did\":3,\"ids\":\"\",\"name\":\"CHANEL香奈儿柔和净肤洁面乳150ml\",\"ItemSalePrice\":490,\"path\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20220224/1645670522349034.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"ProductTips\":\"新品上市\",\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1700557995524,\"did\":10,\"ids\":\"\",\"name\":\"Apple/苹果 13 英寸 Macbook Air Apple M1芯片\",\"ItemSalePrice\":7199,\"path\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/10472/image/20210901/1630492845720352.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0}]}},{\"id\":1700559260453,\"eltmType\":\"7\",\"bgColor\":\"#ffffff\",\"paddingTop\":0,\"paddingRight\":0,\"paddingBottom\":0,\"paddingLeft\":0,\"borderTopLeftRadius\":0,\"borderTopRightRadius\":0,\"borderBottomLeftRadius\":0,\"borderBottomRightRadius\":0,\"eltm7\":{\"column\":\"5\",\"bgColor\":\"#fff\",\"paddingTop\":20,\"paddingRight\":10,\"paddingBottom\":20,\"paddingLeft\":10,\"border\":false,\"data\":[{\"id\":1700559266664,\"did\":16,\"ids\":\"\",\"name\":\"美妆护肤\",\"ItemSalePrice\":0,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231121/cf1f3c86c4754381b27f87dbcfc7a96a.png\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":\"4\",\"AppUrl\":\"/pagesub/product/list?category_id=1269&cname=护肤\",\"UserLimit\":0,\"OrderCount\":0},{\"id\":1700559274341,\"did\":13,\"ids\":\"\",\"name\":\"拼团活动\",\"ItemSalePrice\":0,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231121/343dd2df544d44ab8e24cc9a5ac45bb9.gif\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":\"4\",\"AppUrl\":\"/activity/fightgroup/list\",\"UserLimit\":0,\"OrderCount\":0},{\"id\":1700559275370,\"did\":36,\"ids\":\"\",\"name\":\"限时折扣\",\"ItemSalePrice\":0,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231121/eadcff5b6010401e9a1595e673fa676c.gif\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":\"4\",\"AppUrl\":\"/pagesub/product/list?activity_type_id=1103&cname=限时折扣\",\"UserLimit\":0,\"OrderCount\":0},{\"id\":1700559276092,\"did\":34,\"ids\":\"\",\"name\":\"积分商城\",\"ItemSalePrice\":0,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231121/0cdf60770c854aa6aa4a45d6f859e240.gif\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":\"4\",\"AppUrl\":\"/integral/integral/integral\",\"UserLimit\":0,\"OrderCount\":0},{\"id\":1700559276870,\"did\":9,\"ids\":\"\",\"name\":\"领券中心\",\"ItemSalePrice\":0,\"path\":\"https://static.shopsuite.cn/xcxfile/appicon/icon9.png\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":\"4\",\"AppUrl\":\"/activity/coupon/list\",\"UserLimit\":0,\"OrderCount\":0}]}},{\"id\":1695182067667,\"eltmType\":\"4\",\"bgColor\":\"#ffffff\",\"paddingTop\":0,\"paddingRight\":0,\"paddingBottom\":0,\"paddingLeft\":0,\"borderTopLeftRadius\":0,\"borderTopRightRadius\":0,\"borderBottomLeftRadius\":0,\"borderBottomRightRadius\":0,\"eltm4\":{\"shadow\":true,\"btnType\":0,\"listTyle\":\"2\",\"isPrice\":true,\"isProductTips\":true,\"btnColor\":\"#DB384C\",\"btnText\":\"购买\",\"priceColor\":\"#DB384C\",\"btnFontColor\":\"#fff\",\"data\":[{\"id\":1695182072183,\"did\":95,\"ids\":\"\",\"name\":\"Calvin Klein卡文克莱男士皮革钱包\",\"ItemSalePrice\":300,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230918/1a68f6a7b39847f58b1c322ee7453e84.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182077798,\"did\":94,\"ids\":\"\",\"name\":\"OCCA大容量纯PC拉杆行李箱 托运箱\",\"ItemSalePrice\":800,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230918/08ab31c8993e404897266b1373db6414.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182082694,\"did\":91,\"ids\":\"\",\"name\":\"飞利浦27英寸IPS显示器家用办公屏\",\"ItemSalePrice\":659,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230918/32182910e2c4487e9a3c3a0203ce16ce.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182087175,\"did\":88,\"ids\":\"\",\"name\":\"林氏木业床头收纳柜原木白色四抽\",\"ItemSalePrice\":139,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230918/05d75c8e523d439aa06dec8136b620e7.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182094398,\"did\":86,\"ids\":\"\",\"name\":\"琥珀原木收纳柜桌面收纳\",\"ItemSalePrice\":49,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230918/5fde11d72222445f856c93b5daa83eb8.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182101639,\"did\":83,\"ids\":\"\",\"name\":\"秋季新款圆头拼接玛丽珍单鞋\",\"ItemSalePrice\":369,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230915/203bf87519fd4700a419ec659331613c.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182109752,\"did\":68,\"ids\":\"\",\"name\":\"HLA/海澜之家条纹长袖衬衫\",\"ItemSalePrice\":268,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230915/e44d40c5a12b4a06840c58c1995e17f4.jpg\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0},{\"id\":1695182122183,\"did\":62,\"ids\":\"\",\"name\":\"美的一级能效3匹冷暖变频客厅家用壁挂式空调\",\"ItemSalePrice\":5999,\"path\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/image/media/plantform/20230915/09d2ec59268c4139bdd6a79030da571b.png\",\"pathBg\":\"\",\"flexNum\":0,\"selectType\":1,\"UserLimit\":0,\"OrderCount\":0}]}},{\"id\":1695182180323,\"eltmType\":\"5\",\"bgColor\":\"\",\"paddingTop\":0,\"paddingRight\":0,\"paddingBottom\":0,\"paddingLeft\":0,\"borderTopLeftRadius\":0,\"borderTopRightRadius\":0,\"borderBottomLeftRadius\":0,\"borderBottomRightRadius\":0,\"eltm5\":{\"height\":20}}]', '{\"window\":{\"navigationBarBackgroundColor\":\"#ffffff\",\"navigationBarTextStyle\":\"black\",\"navigationBarTitleText\":\"\",\"backgroundColor\":\"#f8f8f8\",\"backgroundTextStyle\":\"dark\"},\"tabBar\":{\"color\":\"#999999\",\"selectedColor\":\"#DB384C\",\"backgroundColor\":\"#ffffff\",\"borderStyle\":\"white\",\"position\":\"bottom\",\"list\":[{\"pagePath\":\"pages/index/index\",\"iconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/aa07e66f-d4fb-4a2b-bc4c-67ed9aace69b.png\",\"selectedIconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/2840118f-dd01-471a-a749-d910905cefc3.png\",\"text\":\"首页\"},{\"pagePath\":\"pages/category/category\",\"iconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/16ff7864-3513-4ff8-bf12-a9a9808d39e5.png\",\"selectedIconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/cdb5e3df-9d40-42ef-a46a-686dac547fd8.png\",\"text\":\"分类\"},{\"pagePath\":\"pages/cart/cart\",\"iconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/63d2f0b9-7059-4b06-b170-d3b93a749f99.png\",\"selectedIconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/ea018028-20c0-4f1d-a914-3bfc55c8f13f.png\",\"text\":\"购物车\"},{\"pagePath\":\"pages/UserCenter/UserCenter\",\"iconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/50737e26-52e6-4c10-858a-ae2c025c2ebc.png\",\"selectedIconPath\":\"https://static.shopsuite.cn/xcxfile/img003/2017082515/462eb69a-b3ec-446c-9997-c9f1e3212bcb.png\",\"text\":\"我的\"}]}}', '{\"BackgroundObj\":{\"type\":\"1\",\"bgColor\":\"#ffffff\",\"pathColor\":\"#f8f8f8\",\"path\":\"\"}}', '', '', '', b'1', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', b'0', '', 0);

-- ----------------------------
-- Table structure for sys_page_category_nav
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_category_nav`;
CREATE TABLE `sys_page_category_nav`  (
  `category_nav_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `category_nav_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `category_nav_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片',
  `category_ids` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推荐分类(DOT)',
  `item_ids` varchar(5120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推荐商品(DOT)',
  `brand_ids` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推荐品牌(DOT)',
  `category_nav_adv` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '[]' COMMENT '广告数据(JSON)',
  `category_nav_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '模板分类(ENUM):1-分类模板1;2-商品模板',
  `category_nav_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `category_nav_enable` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否启用(BOOL):0-不显示;1-显示',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`category_nav_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'PC分类导航表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_page_mobile_nav
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_mobile_nav`;
CREATE TABLE `sys_page_mobile_nav`  (
  `category_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品类别编号-一级分类使用',
  `category_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名称',
  `category_alias` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类别名',
  `category_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片',
  `page_id` bigint(20) UNSIGNED NOT NULL DEFAULT 1 COMMENT '页面编号',
  `category_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`category_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '移动端首页分类使用' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_page_module
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_module`;
CREATE TABLE `sys_page_module`  (
  `pm_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pm_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
  `page_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `pm_color` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '颜色',
  `pm_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所在页面',
  `module_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '模版',
  `pm_utime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `pm_order` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序',
  `pm_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否显示',
  `pm_html` varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块html代码',
  `pm_json` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '模块JSON代码(JSON)',
  `pm_position` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'column_left:content_top',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`pm_id`) USING BTREE,
  INDEX `page_id`(`page_id`, `pm_enable`, `pm_order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6676 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面模块表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_page_module
-- ----------------------------
INSERT INTO `sys_page_module` VALUES (650, '', 1638151353229, 11791, '', '', '1104', '0000-00-00 00:00:00', 1, b'1', '', '{\"banner_nav1\":{\"name\":\"<h2>品牌闪购</h2>麦吉丽<span>·</span>得宝<span>·</span>美赞臣<span>·</span>美的\"},\"banner_nav2\":{\"name\":\"<b class=\'hot_pic\'></b><h2>频道广场</h2>拍拍二手·</span>女装馆\"},\"banner_nav3\":{\"name\":\"<b class=\'zhuan_pic\'></b><h2>精选热点</h2>电子产品<span>.</span>护肤产品<span>·</span>潮流家电<span></span><span></span>\"},\"nb1_1\":{\"url\":\"/product/lists?category_id=1086\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640781363297770.png\",\"imgSize\":\"290x260\"},\"nb1_2\":{\"url\":\"/product/lists?category_id=1283\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640678149692703.png\",\"imgSize\":\"290x170\"},\"nb1_3\":{\"url\":\"/product/lists?category_id=1007\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640781149517607.png\",\"imgSize\":\"400x363\"},\"nb1_4\":{\"url\":\"/product/detail?item_id=115\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640785766619118.png\",\"imgSize\":\"120x47\"},\"nb1_5\":{\"name\":\"美的智能家电狂欢购\",\"desc\":\"至高24期免息\"},\"nb1_6\":{\"url\":\"/product/lists?category_id=1006\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640679511494281.png\",\"imgSize\":\"290x260\"},\"nb1_7\":{\"url\":\"/product/lists?category_id=1008\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640871056897252.png\",\"imgSize\":\"290x170\"},\"nb1_8\":{\"url\":\"/product/lists?category_id=1002\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640832506559289.png\",\"imgSize\":\"190x400\"},\"nb2_1\":{\"url\":\"/product/lists?category_id=2127\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640682052177249.png\",\"imgSize\":\"294x320\"},\"nb2_2\":{\"url\":\"/product/lists?category_id=2042\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640682057255539.png\",\"imgSize\":\"294x320\"},\"nb2_3\":{\"url\":\"/product/lists?category_id=1007\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640682615844097.png\",\"imgSize\":\"294x320\"},\"nb2_4\":{\"url\":\"/product/lists?category_id=1292\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211228/1640682388708882.png\",\"imgSize\":\"294x320\"},\"nb3_1\":{\"url\":\"/product/detail?item_id=50\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749559922238.png\",\"imgSize\":\"326x215\"},\"nb3_2\":{\"url\":\"/product/detail?item_id=46\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749565442063.png\",\"imgSize\":\"326x215\"},\"nb3_3\":{\"url\":\"/product/detail?item_id=40\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749571131763.png\",\"imgSize\":\"326x215\"},\"nb3_4\":{\"url\":\"/product/detail?item_id=48\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749576241625.png\",\"imgSize\":\"326x215\"},\"nb3_5\":{\"url\":\"/product/detail?item_id=40\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749583174322.png\",\"imgSize\":\"326x215\"},\"nb3_6\":{\"url\":\"/product/detail?item_id=40\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640749590301558.png\",\"imgSize\":\"326x215\"},\"nb3_7\":{\"url\":\"/product/lists?category_id=2127\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211229/1640772734398197.png\",\"imgSize\":\"190x480\"},\"module_id\":\"1104\",\"pm_id\":650,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (659, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 4, b'1', '', '{\"banner1\":{\"url\":\"/product/lists?category_id=1008\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640850680691258.png\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":659,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (661, '', 1638151353229, 11791, '', '', '1011', '2023-07-17 17:09:32', 5, b'1', '', '{\"pm_id\":661,\"banner4\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640865739845411.png\",\"url\":\"/product/detail?item_id=3\",\"name\":\"资生堂SKII\",\"desc\":\"平衡水油柔嫩肌肤\"},\"img_items2\":[{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640952826330503.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1008\"},{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955545462234.png\",\"url\":\"/product/lists?category_id=1008\"}],\"img_items3\":[{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640951596530870.png\",\"url\":\"https://demo.modulithshop.cn/product/detail?item_id=30\"},{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640951601717331.png\",\"url\":\"/product/detail?item_id=30\"}],\"img_items1\":[{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20220105/1641365252810834.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1008\"},{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640872306536720.png\",\"url\":\"/product/lists?category_id=1008\"}],\"banner2\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640865722962833.png\",\"url\":\"/product/detail?item_id=99\",\"name\":\"Olay/玉兰油\",\"desc\":\"淡斑小白瓶美白精华液\"},\"banner3\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640865674139598.png\",\"url\":\"/product/detail?item_id=179\",\"name\":\"怡丽丝尔水乳\",\"desc\":\"抗初老维持肌肤弹力\"},\"setting\":{\"classSelect\":\"floor1,floor2,floor3,floor4,floor5,floor6,floor7,floor8,floor9\",\"tabs\":{\"productItemLimit\":\"6\"},\"showFloorName\":\"0\",\"className\":\"floor2\"},\"banner1\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640865705881676.png\",\"url\":\"/product/detail?item_id=99\",\"name\":\"美妆护肤\",\"desc\":\"抗氧化面部保湿精油\"},\"pm_enable\":true,\"module_id\":\"1011\",\"links\":[{\"link\":\"/product/lists?category_id=1008\",\"name\":\"卡姿兰\"},{\"link\":\"/product/lists?category_id=1008\",\"name\":\"范思哲\"},{\"link\":\"/product/lists?category_id=1008\",\"name\":\"百雀羚\"},{\"link\":\"/product/lists?category_id=1008\",\"name\":\"自然堂\"},{\"link\":\"/product/lists?category_id=1008\",\"name\":\"珀莱雅\"},{\"link\":\"/product/lists?category_id=1008\",\"name\":\"薇诺娜\"}],\"floor\":{\"imgSize\":\"38x38\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211230/1640870274370510.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1008\",\"name\":\"美妆护肤\"}}', '', 0);
INSERT INTO `sys_page_module` VALUES (664, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 2, b'1', '', '{\"banner1\":{\"url\":\"\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640938551514028.png\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":664,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (665, '', 1638151353229, 11791, '', '', '1011', '2023-07-17 17:33:03', 3, b'1', '', '{\"pm_id\":665,\"banner4\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640945165889342.png\",\"url\":\"/product/lists?category_id=1007\",\"name\":\"韩国大宇加湿器\",\"desc\":\"纯净型无雾 贴近自然\"},\"img_items2\":[{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955627165765.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1007\"},{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955788289443.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1007\"}],\"img_items3\":[{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640949771391089.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1007\"},{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640949777657827.png\",\"url\":\"/product/lists?category_id=1007\"}],\"img_items1\":[{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640952514200860.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=1007\"},{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640952525631973.png\",\"url\":\"//test.suteshop.com/Product/detail/item_id/1025\"}],\"banner2\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640944630850143.png\",\"url\":\"/product/lists?category_id=1007\",\"name\":\"美的除螨仪\",\"desc\":\"旋风气旋 深层除螨\"},\"banner3\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640944620946159.png\",\"url\":\"/product/lists?category_id=1007\",\"name\":\"CTT烘干机\",\"desc\":\"双向翻转 全自动干衣\"},\"setting\":{\"classSelect\":\"floor1,floor2,floor3,floor4,floor5,floor6,floor7,floor8,floor9\",\"tabs\":{\"productItemLimit\":\"6\"},\"showFloorName\":\"1\",\"className\":\"floor7\"},\"banner1\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640944641312197.png\",\"url\":\"/product/lists?category_id=1007\",\"name\":\"科沃斯集尘王扫\",\"desc\":\"12期免息分期 送天猫精灵\"},\"pm_enable\":true,\"module_id\":\"1011\",\"links\":[{\"link\":\"/product/lists?category_id=1007\",\"name\":\"美的\"},{\"link\":\"/product/lists?category_id=1007\",\"name\":\"九阳\"},{\"link\":\"/product/lists?category_id=1007\",\"name\":\"飞利浦\"},{\"link\":\"/product/lists?category_id=1007\",\"name\":\"苏泊尔\"},{\"link\":\"/product/lists?category_id=1007\",\"name\":\"荣事达\"},{\"link\":\"/product/lists?category_id=1007\",\"name\":\"科沃斯\"}],\"floor\":{\"imgSize\":\"38x38\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640938873534470.png\",\"url\":\"/product/lists?category_id=1007\",\"name\":\"家用电器\"}}', '', 0);
INSERT INTO `sys_page_module` VALUES (666, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 8, b'1', '', '{\"banner1\":{\"url\":\"\",\"picimg\":\"https:\\/\\/shopsuite.oss-accelerate.aliyuncs.com\\/media\\/user\\/11791\\/image\\/20211231\\/1640953314620031.png\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":\"666\",\"pm_enable\":\"1\"}', '', 0);
INSERT INTO `sys_page_module` VALUES (667, '', 1638151353229, 11791, '', '', '1011', '2023-07-17 17:15:02', 9, b'1', '', '{\"pm_id\":667,\"banner4\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640954214169691.png\",\"url\":\"/product/detail?item_id=90\",\"name\":\"新生儿衣服\",\"desc\":\"纯棉三角爬 可开档\"},\"img_items2\":[{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955023644977.png\",\"url\":\"https://demo.modulithshop.cn/product/detail?item_id=105\"},{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955013708512.png\",\"url\":\"https://demo.modulithshop.cn/product/detail?item_id=105\"}],\"img_items3\":[{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955136654995.png\",\"url\":\"https://demo.modulithshop.cn/product/detail?item_id=90\"},{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955124149121.png\",\"url\":\"/product/detail?item_id=90\"}],\"img_items1\":[{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20220105/1641354911583188.png\",\"url\":\"https://test.suteshop.com/i/9304\"},{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955265173363.png\",\"url\":\"/product/detail?item_id=89\"}],\"banner2\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640954208670368.png\",\"url\":\"/product/detail?item_id=102\",\"name\":\"新生儿摇椅\",\"desc\":\"沉浸式哄睡体验\"},\"banner3\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640954198702737.png\",\"url\":\"/product/detail?item_id=90\",\"name\":\"新生儿衣服\",\"desc\":\"婴幼儿A类标准\"},\"setting\":{\"classSelect\":\"floor1,floor2,floor3,floor4,floor5,floor6,floor7,floor8,floor9\",\"tabs\":{\"productItemLimit\":\"6\"},\"showFloorName\":\"1\",\"className\":\"floor9\"},\"banner1\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640954203349605.png\",\"url\":\"/product/detail?item_id=56\",\"name\":\"新生儿礼盒\",\"desc\":\"创意梦幻手拎礼盒\"},\"pm_enable\":true,\"module_id\":\"1011\",\"links\":[{\"link\":\"/product/lists?category_id=1002\",\"name\":\"马克珍妮\"},{\"link\":\"/product/lists?category_id=1002\",\"name\":\"戴维贝拉\"},{\"link\":\"/product/lists?category_id=1002\",\"name\":\"小雅象\"},{\"link\":\"/product/lists?category_id=1002\",\"name\":\"木木屋\"},{\"link\":\"/product/lists?category_id=1002\",\"name\":\"布鲁可\"},{\"link\":\"/product/lists?category_id=1002\",\"name\":\"艾惟诺\"}],\"floor\":{\"imgSize\":\"38x38\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640953076930470.png\",\"url\":\"product/detail?item_id=90\",\"name\":\"母婴玩具\"}}', '', 0);
INSERT INTO `sys_page_module` VALUES (668, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 6, b'1', '', '{\"banner1\":{\"url\":\"/product/lists?category_id=2133\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640955981889073.png\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":668,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (669, '', 1638151353229, 11791, '', '', '1011', '2023-07-17 17:12:37', 7, b'1', '', '{\"pm_id\":669,\"banner4\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959178470711.png\",\"url\":\"/product/detail?item_id=387\",\"name\":\"高钙全脂奶粉\",\"desc\":\"含A2蛋白质\"},\"img_items2\":[{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640960196387969.png\",\"url\":\"https://demo.modulithshop.cn/product/detail?item_id=402\"},{\"imgSize\":\"400x430\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640960204976465.png\",\"url\":\"/product/detail?item_id=297\"}],\"img_items3\":[{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959565780376.png\",\"url\":\"//test.suteshop.com/Product/detail/item_id/1025\"},{\"imgSize\":\"390x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959193389326.png\",\"url\":\"/product/detail?item_id=382\"}],\"img_items1\":[{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959578203525.png\",\"url\":\"//test.suteshop.com/Product/detail/item_id/1025\"},{\"imgSize\":\"200x200\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959183486391.png\",\"url\":\"https://test.suteshop.com/i/9599\"}],\"banner2\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959171874068.png\",\"url\":\"/product/detail?item_id=402\",\"name\":\"修正左旋肉碱\",\"desc\":\"肥瘦身燃脂减脂非酵素\"},\"banner3\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959159510852.png\",\"url\":\"/product/detail?item_id=416\",\"name\":\"浓缩磷脂软胶囊 \",\"desc\":\"健康从血管开始\"},\"setting\":{\"classSelect\":\"floor1,floor2,floor3,floor4,floor5,floor6,floor7,floor8,floor9\",\"tabs\":{\"productItemLimit\":\"6\"},\"showFloorName\":\"1\",\"className\":\"floor5\"},\"banner1\":{\"imgSize\":\"190x210\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640959164698674.png\",\"url\":\"/product/detail?item_id=323\",\"name\":\"玛咖胶囊\",\"desc\":\"原装进口有机黑玛卡粉\"},\"pm_enable\":true,\"module_id\":\"1011\",\"links\":[{\"link\":\"/product/detail?item_id=323\",\"name\":\"康恩贝\"},{\"link\":\"/product/detail?item_id=298\",\"name\":\"同仁堂\"},{\"link\":\"/product/detail?item_id=341\",\"name\":\"青源堂\"},{\"link\":\"/product/detail?item_id=387\",\"name\":\"燕之坊\"},{\"link\":\"/product/detail?item_id=387\",\"name\":\"东阿阿胶\"},{\"link\":\"/product/detail?item_id=447\",\"name\":\"汤成倍健\"}],\"floor\":{\"imgSize\":\"38x38\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640957732151037.png\",\"url\":\"https://demo.modulithshop.cn/product/lists?category_id=2133\",\"name\":\"医药保健\"}}', '', 0);
INSERT INTO `sys_page_module` VALUES (671, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 10, b'1', '', '{\"banner1\":{\"url\":\"/product/lists\",\"picimg\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231107/7475eb913745480ea9a976d3c38c582d.jpg\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":671,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (672, '', 1638151353229, 11791, '', '', '1101', '0000-00-00 00:00:00', 13, b'0', '', '{\"banner1\":{\"url\":\"\",\"picimg\":\"https://shopsuite.oss-accelerate.aliyuncs.com/media/user/11791/image/20211231/1640960385108620.png\",\"name\":\"横幅名称\",\"imgSize\":\"1200x120\"},\"module_id\":\"1101\",\"pm_id\":672,\"pm_enable\":false}', '', 0);
INSERT INTO `sys_page_module` VALUES (801, '', 1638151353229, 0, '', '', '3001', '2023-07-17 17:56:21', 0, b'1', '', '{\"banner2\":{\"imgSize\":\"210x210\",\"picimg\":\"//static.shopsuite.cn/pagepreview/data/01_210x210.png\",\"url\":\"//test.shopsuite.cn/\",\"name\":\"FURNITURE\"},\"setting\":{\"showMegamenus\":0},\"banner1\":{\"imgSize\":\"210x210\",\"picimg\":\"//static.shopsuite.cn/pagepreview/data/01_210x210.png\",\"url\":\"//test.shopsuite.cn/\",\"name\":\"FURNITURE\"},\"items\":[{\"imgSize\":\"1920x453\",\"picimg\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231107/188cb569a8c047bca2c8562fcd91e251.jpg\",\"url\":\"/product/lists?category_id=1001\",\"bgColor\":\"#fff\",\"name\":\"商品信息\"},{\"url\":\"\",\"picimg\":\"https://kuteshop.oss-accelerate.aliyuncs.com/modulithshop/guest/image/plantform/20231107/c22e11693cad4226a02df98e2b6115e3.jpg\",\"name\":\"\",\"imgSize\":\"1920x453\",\"bgColor\":\"#fff\"}],\"module_id\":\"3001\",\"pm_id\":801,\"pm_enable\":true}', '', 0);
INSERT INTO `sys_page_module` VALUES (6675, '', 1638151353229, 0, '', '', '1103', '2023-07-24 18:18:25', 1, b'0', '', '{\"banner1\":{\"imgSize\":\"370x200\",\"name\":\"横幅名称1\",\"picimg\":\"//static.shopsuite.cn/demodata/assets/data/option14/b1.jpg\",\"url\":\"/product/detail?item_id=139\"},\"banner2\":{\"imgSize\":\"370x370\",\"name\":\"横幅名称\",\"picimg\":\"//static.shopsuite.cn/demodata/assets/data/option14/b2.jpg\",\"url\":\"/product/detail?item_id=143\"},\"banner3\":{\"imgSize\":\"370x360\",\"name\":\"横幅名称\",\"picimg\":\"//static.shopsuite.cn/demodata/assets/data/option14/b3.jpg\",\"url\":\"/product/detail?item_id=144\"},\"banner4\":{\"imgSize\":\"370x370\",\"name\":\"横幅名称\",\"picimg\":\"//static.shopsuite.cn/demodata/assets/data/option14/b4.jpg\",\"url\":\"/product/detail?item_id=92\"},\"banner5\":{\"imgSize\":\"370x200\",\"name\":\"横幅名称\",\"picimg\":\"//static.shopsuite.cn/demodata/assets/data/option14/b5.jpg\",\"url\":\"/product/detail?item_id=31\"},\"module_id\":\"1103\",\"pm_enable\":false,\"pm_id\":6675,\"setting\":{}}', '', 0);

-- ----------------------------
-- Table structure for sys_page_pc_nav
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_pc_nav`;
CREATE TABLE `sys_page_pc_nav`  (
  `nav_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '导航编号',
  `nav_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别(ENUM):0-自定义导航;1-商品分类;2-文章导航;3-活动导航',
  `nav_item_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类别内容编号',
  `nav_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '导航标题',
  `nav_url` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '导航链接(HTML)',
  `nav_position` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '导航位置(ENUM):0-头部;1-中部;2-底部',
  `nav_target_blank` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否以新窗口打开(BOOL):1-是; 0-否',
  `nav_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '导航图片',
  `nav_dropdown_menu` varchar(15000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '导航下拉内容(HTML)',
  `nav_order` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `nav_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用(BOOL):1-是; 0-否',
  `nav_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(ENUM):1-是; 0-否',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`nav_id`) USING BTREE,
  INDEX `nav_enable`(`nav_enable`, `nav_order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面导航表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_page_user_form
-- ----------------------------
DROP TABLE IF EXISTS `sys_page_user_form`;
CREATE TABLE `sys_page_user_form`  (
  `user_form_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '页面编号',
  `page_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属店铺',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_form_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '提交数据',
  `form_question_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表单数据',
  `user_form_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '提交时间',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商家编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_form_id`) USING BTREE,
  INDEX `page_id`(`page_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_plantform_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_plantform_notice`;
CREATE TABLE `sys_plantform_notice`  (
  `notice_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '公告编号',
  `notice_titile` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告标题',
  `notice_content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告内容(HTML)',
  `notice_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '推送时间',
  `notice_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否启用(BOOL):1-启用;0-禁用',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_user_reg_source
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_reg_source`;
CREATE TABLE `sys_user_reg_source`  (
  `user_reg_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '注册编号',
  `user_reg_date` date NOT NULL COMMENT '日期编号',
  `user_reg_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户数',
  `user_reg_mobile` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '手机用户数',
  `user_reg_pad` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '平板用户数',
  `user_reg_pc` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'PC用户数',
  `user_reg_h5` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'H5用户数',
  `user_reg_app` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'APP用户数',
  `user_reg_mp` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小程序用户数',
  `user_reg_wechat` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '微信用户数',
  `user_reg_baidu` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '百度用户数',
  `user_reg_alipay` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付宝用户数',
  `user_reg_toutiao` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '头条用户数',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`user_reg_id`) USING BTREE,
  INDEX `access_os`(`user_reg_num`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户来源终端表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_wechat_tplmsg
-- ----------------------------
DROP TABLE IF EXISTS `sys_wechat_tplmsg`;
CREATE TABLE `sys_wechat_tplmsg`  (
  `tplmsg_id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模版编号',
  `tplmsg_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模版标题',
  `tplmsg_type_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '模版类型(LIST):1-订单提醒; 2-支付提醒;3-发货提醒',
  `tplmsg_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信消息模板标题',
  `tplmsg_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模版库编号',
  `tplmsg_tpl_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模版库编号',
  `tplmsg_enable` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用(BOOL):0-禁用;1-启用',
  `tplmsg_is_buildin` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系统内置(BOOL):0-否;1-是',
  `tplmsg_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `tplmsg_sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商城编号',
  `message_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息模板id:关联消息模板',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`tplmsg_id`) USING BTREE,
  INDEX `erpbuilder_base_customer_ibfk_1`(`tplmsg_type_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息模板表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_chain_code
-- ----------------------------
DROP TABLE IF EXISTS `trade_chain_code`;
CREATE TABLE `trade_chain_code`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单商品编号',
  `chain_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '虚拟码:pickupcode',
  `chain_code_status` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '虚拟码状态(ENUM): 0-未使用; 1-已使用; 2-冻结',
  `chain_code_usetime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '虚拟兑换码使用时间',
  `virtual_service_date` date NOT NULL,
  `virtual_service_time` timestamp(0) NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `product_validity_end` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '失效时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `chain_code`(`chain_code`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE,
  INDEX `chain_id`(`chain_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '虚拟兑换码门店自提兑换码表pickupcode' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_distribution_generated_commission
-- ----------------------------
DROP TABLE IF EXISTS `trade_distribution_generated_commission`;
CREATE TABLE `trade_distribution_generated_commission`  (
  `ugc_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '贡献佣金编号:user_id + \'-\' + level',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '推广粉丝',
  `user_parent_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级推广员编号',
  `ugc_level` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '推广层级: 1父  ,2祖父, 记录不变，如果关系更变，则增加其它记录',
  `ugc_amount` decimal(16, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '产生佣金',
  `ugc_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单数量',
  `user_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`ugc_id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`, `user_parent_id`, `ugc_level`) USING BTREE COMMENT '(null)'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台推广粉丝产生佣金汇总表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_distribution_order
-- ----------------------------
DROP TABLE IF EXISTS `trade_distribution_order`;
CREATE TABLE `trade_distribution_order`  (
  `uo_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单收益编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号:上级ID,获取佣金推广员',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `uo_buy_commission` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '推广员佣金',
  `uo_directseller_commission` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '销售员佣金',
  `buyer_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `uo_level` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级',
  `uo_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '时间',
  `uo_active` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否有效(BOOL):0-未生效;1-有效',
  `uo_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `uo_is_paid` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否有效(BOOL):0-未支付;1-已支付',
  `uo_paytime` bigint(20) NOT NULL DEFAULT 0 COMMENT '支付时间',
  `uo_receivetime` bigint(20) NOT NULL DEFAULT 0 COMMENT '收货时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`uo_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `uo_level`, `uo_time`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广订单收益详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_distribution_order_item
-- ----------------------------
DROP TABLE IF EXISTS `trade_distribution_order_item`;
CREATE TABLE `trade_distribution_order_item`  (
  `uoi_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单收益编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品SKU',
  `uoi_buy_commission` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '推广员佣金',
  `uoi_directseller_commission` decimal(15, 6) NOT NULL DEFAULT 0.000000 COMMENT '销售员佣金',
  `uoi_distributor_commission` decimal(15, 0) NOT NULL DEFAULT 0 COMMENT '分销商收益-本店销售获取差价',
  `buyer_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `uoi_level` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级',
  `uoi_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '时间',
  `uoi_active` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否有效(BOOL):0-未生效;1-有效',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `uoi_is_paid` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否有效(BOOL):0-未支付;1-已支付',
  `uoi_receivetime` bigint(20) NOT NULL DEFAULT 0 COMMENT '收货时间',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`uoi_id`) USING BTREE,
  INDEX `user_id`(`user_id`, `uoi_level`, `uoi_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广订单商品收益详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_base
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_base`;
CREATE TABLE `trade_order_base`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `order_product_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品原价总和:商品发布原价',
  `order_payment_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '应付金额:order_product_amount - order_discount_amount + order_shipping_fee - order_voucher_price - order_points_fee - order_adjust_fee',
  `currency_id` int(10) UNSIGNED NOT NULL DEFAULT 86 COMMENT '货币编号',
  `currency_symbol_left` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '￥' COMMENT '左符号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `store_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家昵称',
  `order_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 2010 COMMENT '订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `order_state_id`(`order_state_id`) USING BTREE,
  INDEX `buyer_user_id`(`user_id`, `order_state_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `order_state_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单详细信息-检索不分表也行，cache' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_comment
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_comment`;
CREATE TABLE `trade_order_comment`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '卖家店铺编号-冗余',
  `store_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家姓名',
  `comment_points` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '获得积分-冗余，独立表记录',
  `comment_scores` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '评价星级1-5积分',
  `comment_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评价内容',
  `comment_image` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论上传的图片：|分割多张图片',
  `comment_helpful` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有帮助',
  `comment_nohelpful` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '无帮助',
  `comment_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
  `comment_is_anonymous` bit(1) NOT NULL DEFAULT b'0' COMMENT '匿名评价',
  `comment_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '评价信息的状态(BOOL): 1-正常显示; 0-禁止显示',
  `comment_store_desc_credit` tinyint(3) UNSIGNED NOT NULL DEFAULT 5 COMMENT '描述相符评分 - order_buyer_evaluation_status , 评价状态改变后不需要再次评论，根据订单走',
  `comment_store_service_credit` tinyint(3) UNSIGNED NOT NULL DEFAULT 5 COMMENT '服务态度评分 - order_buyer_evaluation_status',
  `comment_store_delivery_credit` tinyint(3) UNSIGNED NOT NULL DEFAULT 5 COMMENT '发货速度评分 - order_buyer_evaluation_status',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_id`) USING BTREE,
  UNIQUE INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `goods_evaluation_ibfk_2`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单店铺评价表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_data
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_data`;
CREATE TABLE `trade_order_data`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单描述',
  `order_delay_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '延迟时间,默认为0 - 收货确认',
  `delivery_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送方式',
  `delivery_time_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '配送时间:要求，不限、周一~周五、周末等等',
  `delivery_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '配送日期',
  `delivery_istimer` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否定时配送(BOOL):0-不定时;1-定时',
  `order_message` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家订单留言',
  `order_item_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品总价格/商品金额, 不包含运费',
  `order_discount_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '折扣价格/优惠总金额',
  `order_adjust_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '手工调整费用店铺优惠',
  `order_points_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '积分抵扣费用',
  `order_sale_person_discount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售员折扣额度',
  `order_shipping_fee_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '运费价格/运费金额',
  `order_shipping_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实际运费金额-卖家可修改',
  `voucher_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '代金券id/优惠券/返现:发放选择使用',
  `voucher_number` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '代金券编码',
  `voucher_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '代金券面额',
  `redpacket_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '红包id-平台代金券',
  `redpacket_number` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '红包编码',
  `redpacket_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '红包面额',
  `order_redpacket_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '红包抵扣订单金额',
  `order_resource_ext1` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '第二需要支付资源例如积分',
  `order_resource_ext2` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '众宝',
  `order_resource_ext3` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '金宝',
  `trade_payment_money` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '余额支付',
  `trade_payment_recharge_card` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '充值卡支付',
  `trade_payment_credit` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '信用支付',
  `order_refund_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退款状态:0-是无退款;1-是部分退款;2-是全部退款',
  `order_refund_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款金额:申请额度',
  `order_refund_agree_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款金额:同意额度',
  `order_refund_agree_cash` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '已同意退的现金',
  `order_refund_agree_points` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '已退的积分额度',
  `order_refund_agree_credit` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '已退的信用额度',
  `order_return_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退货状态(ENUM):0-是无退货;1-是部分退货;2-是全部退货',
  `order_return_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退货数量',
  `order_return_ids` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退货单编号s(DOT):冗余',
  `order_commission_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '平台交易佣金',
  `order_commission_fee_refund` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '交易佣金-退款',
  `order_points_add` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '订单赠送积分',
  `order_points_add_refund` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '已退订单赠送积分',
  `order_activity_data` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '促销信息',
  `order_cancel_identity` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单取消者身份(ENUM):1-买家; 2-卖家; 3-系统',
  `order_cancel_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单取消原因',
  `order_cancel_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '订单取消时间',
  `order_bp_add` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '赠送资源2',
  `order_rebate` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '订单返利',
  `buyer_mobile` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '手机号码',
  `buyer_contacter` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `activity_manhui_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '满返优惠券活动id(DOT)',
  `activity_double_points_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动-多倍积分id',
  `order_double_points_add` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '活动-多倍积分',
  `activity_voucher_id` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '满返用户优惠券id(DOT)',
  `order_activity_manhui_state` smallint(5) UNSIGNED NOT NULL DEFAULT 1000 COMMENT '满返优惠券发放状态(ENUM):1000-无需发放;1001-待发放; 1002-已发放; 1003-发放异常',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `order_plus_discount_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '订单PLUS优惠金额',
  `order_level_discount_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '会员等级优惠金额',
  `plus_points_multiple` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'PLUS积分倍数',
  `plus_shipping_discount` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'PLUS运费优惠金额',
  `points_value_rate` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '积分价值比例',
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `order_activity_manhui_state`(`order_activity_manhui_state`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单详细信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_delivery_address
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_delivery_address`;
CREATE TABLE `trade_order_delivery_address`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `da_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `da_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `da_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `da_telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `da_province_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '省编号',
  `da_province` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `da_city_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '市编号',
  `da_city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市',
  `da_county_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '县',
  `da_county` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
  `da_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `da_postalcode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  `da_tag_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '地址标签：家里、公司等等',
  `da_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `da_longitude` double(20, 0) NOT NULL DEFAULT 0 COMMENT '经度',
  `da_latitude` double(20, 0) NOT NULL DEFAULT 0 COMMENT '纬读',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收货地址表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_info
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_info`;
CREATE TABLE `trade_order_info`  (
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单标题',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '卖家店铺编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `kind_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1201 COMMENT '订单种类(ENUM): 1201-实物 ; 1202-教育类 ; 1203-电子卡券  ; 1204-其它',
  `order_sale_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单销售员',
  `order_return_hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '退货按钮隐藏(BOOL): 1-是; 0-否',
  `order_invoice_hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '开票按钮隐藏(BOOL): 1-是; 0-否',
  `order_push_wx_shipping` tinyint(1) NOT NULL DEFAULT 0 COMMENT '推送微信发货是否成功',
  `order_lock_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '锁定状态(BOOL):0-正常;1-锁定,退款退货',
  `order_is_settlemented` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单是否结算(ENUM):0-未结算; 1-已结算; 2-待结算',
  `order_settlement_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '订单结算时间',
  `order_buyer_is_settlemented` bit(1) NOT NULL DEFAULT b'0' COMMENT '订单采购是否结算(BOOL):0-未结算; 1-已结算',
  `order_buyer_settlement_time` timestamp(0) NOT NULL COMMENT '订单结算时间',
  `order_buyer_evaluation_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家针对订单对店铺评价(ENUM): 0-未评价;1-已评价;  2-已过期未评价',
  `order_seller_evaluation_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '卖家评价状态(ENUM):0-未评价;1-已评价;  2-已过期未评价',
  `order_buyer_hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '买家删除(BOOL): 1-是; 0-否',
  `order_shop_hidden` bit(1) NOT NULL DEFAULT b'0' COMMENT '店铺删除(BOOL): 1-是; 0-否',
  `payment_type_id` int(10) UNSIGNED NOT NULL DEFAULT 1302 COMMENT '支付方式(ENUM): 1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付; ',
  `payment_time` timestamp(0) NOT NULL COMMENT '付款时间',
  `order_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 10 COMMENT '订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;',
  `order_is_review` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单审核(ENUM):0-未审核;1-审核通过;2-审核驳回;',
  `order_finance_review` bit(1) NOT NULL DEFAULT b'0' COMMENT '财务状态(BOOL):0-未审核;1-已审核',
  `order_is_paid` smallint(5) UNSIGNED NOT NULL DEFAULT 3010 COMMENT '付款状态(ENUM):3010-未付款;3011-付款待审核;3012-部分付款;3013-已付款',
  `order_is_out` smallint(5) UNSIGNED NOT NULL DEFAULT 3020 COMMENT '出库状态(ENUM):3020-未出库;3021-部分出库通过拆单解决这种问题;3022-已出库',
  `order_is_shipped` smallint(5) UNSIGNED NOT NULL DEFAULT 3030 COMMENT '发货状态(ENUM):3030-未发货;3032-已发货;3031-部分发货',
  `order_is_received` bit(1) NOT NULL DEFAULT b'0' COMMENT '收货状态(BOOL):0-未收货;1-已收货',
  `order_received_time` timestamp(0) NOT NULL COMMENT '订单签收时间',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `delivery_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送方式',
  `order_is_offline` bit(1) NOT NULL DEFAULT b'0' COMMENT '线下订单(BOOL):0-线上;1-线下',
  `order_express_print` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否打印(BOOL):0-未打印;1-已打印',
  `activity_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动编号(DOT)',
  `activity_type_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动类型(DOT)',
  `salesperson_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '销售员编号:用户编号',
  `order_is_sync` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否ERP同步(BOOL):0-未同步;1-已同步;2-同步失败;3-无需同步',
  `order_sync_msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '同步消息',
  `store_is_selfsupport` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否自营(ENUM): 1-自营;0-非自营',
  `store_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺',
  `order_erp_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ERP订单编号',
  `distributor_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分销商编号:用户编号',
  `order_is_cb` bit(1) NOT NULL DEFAULT b'0' COMMENT '跨境订单(BOOL):0-否; 1-是',
  `order_is_cb_sync` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否报关(BOOL):0-否; 1-是',
  `src_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源订单',
  `order_is_transfer` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否代发(BOOL):0-否; 1-是',
  `order_is_transfer_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '转单执行结果',
  `order_fx_is_settlemented` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '佣金是否发放(ENUM):0 -未发放;1 -已发放; 2-待发放; 3-停止发放',
  `order_fx_settlement_time` timestamp(0) NOT NULL COMMENT '佣金结算时间',
  `order_type` mediumint(8) UNSIGNED NOT NULL DEFAULT 3061 COMMENT '订单类型(ENUM)',
  `order_withdraw_confirm` bit(1) NOT NULL DEFAULT b'0' COMMENT '提现审核(BOOL):0-未审核; 1-已审核',
  `payment_form_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '支付方式(BOOL):1-先预约后支付;0-先支付后预约',
  `cart_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销',
  `card_kind_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品绑定卡片类型(ENUM): 1001-次卡商品; 1002-优惠券商品;1003-券码商品;',
  `company_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '公司编号',
  `pay_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '公司付款用户编号',
  `create_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '下单时间:检索使用',
  `update_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前状态的处理时间',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `order_state_id`(`order_state_id`) USING BTREE,
  INDEX `kind_id`(`kind_id`) USING BTREE,
  INDEX `buyer_user_id`(`user_id`, `order_state_id`, `kind_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE,
  INDEX `store_id_2`(`store_id`, `order_state_id`, `order_is_paid`, `update_time`) USING BTREE,
  INDEX `store_id_3`(`store_id`) USING BTREE,
  INDEX `src_order_id`(`src_order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_invoice
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_invoice`;
CREATE TABLE `trade_order_invoice`  (
  `order_invoice_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '发票编号',
  `order_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `store_id` int(11) NOT NULL COMMENT '店铺编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `invoice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票抬头',
  `invoice_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发票内容',
  `invoice_amount` decimal(16, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '开票金额',
  `invoice_company_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '纳税人识别号',
  `invoice_is_company` bit(1) NOT NULL DEFAULT b'1' COMMENT '公司开票(BOOL):0-个人;1-公司',
  `invoice_is_electronic` bit(1) NOT NULL DEFAULT b'1' COMMENT '电子发票(ENUM):0-纸质发票;1-电子发票',
  `invoice_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '发票类型(ENUM):1-普通发票;2-增值税专用发票',
  `invoice_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开票状态(ENUM): 0-未开票; 1-已开票; 2-已拒绝',
  `invoice_status_remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '开票状态备注',
  `invoice_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '电子发票图片',
  `invoice_datetime` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开票时间',
  `invoice_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位地址',
  `invoice_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '单位电话',
  `invoice_bankname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '开户银行',
  `invoice_bankaccount` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '银行账号',
  `invoice_contact_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人',
  `invoice_contact_area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票人地区',
  `invoice_contact_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收票详细地址',
  `user_intl` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '+86' COMMENT '国家编码',
  `user_mobile` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号码(mobile)',
  `user_email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱(email)',
  `order_is_paid` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否支付(BOOL): 0-未支付; 1-已支付;',
  `invoice_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `invoice_cancel` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否取消(BOOL): 0-未取消; 1-取消;',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_invoice_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '订单发票管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_item
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_item`;
CREATE TABLE `trade_order_item`  (
  `order_item_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `product_id` bigint(20) UNSIGNED NOT NULL COMMENT '产品编号',
  `product_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `item_id` bigint(20) UNSIGNED NOT NULL COMMENT '货品编号',
  `item_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `category_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类编号',
  `item_cost_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '成本价',
  `item_unit_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '商品价格单价',
  `item_policy_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '会员专属价',
  `item_unit_points` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '资源1单价',
  `item_unit_sp` decimal(10, 2) NOT NULL COMMENT '资源2单价',
  `order_item_sale_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品实际成交价单价',
  `order_item_quantity` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '商品数量',
  `order_item_inventory_lock` smallint(5) UNSIGNED NOT NULL DEFAULT 1002 COMMENT '库存锁定(ENUM):1001-下单锁定;1002-支付锁定;',
  `order_item_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品图片',
  `order_item_return_num` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退货数量',
  `order_item_return_subtotal` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款总额',
  `order_item_return_agree_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款金额:同意额度',
  `order_item_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品实际总金额: order_item_sale_price * order_item_quantity',
  `order_item_discount_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '优惠金额:只考虑单品的，订单及店铺总活动优惠不影响',
  `order_item_adjust_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '手工调整金额',
  `order_item_points_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '积分费用',
  `order_item_points_add` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '赠送积分',
  `order_item_payment_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '实付金额: order_item_payment_amount =  order_item_amount - order_item_discount_amount - order_item_adjust_fee - order_item_point_fee',
  `order_item_evaluation_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价状态(ENUM): 0-未评价;1-已评价;2-失效评价',
  `activity_type_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动类型(ENUM):0-默认;1101-加价购=搭配宝;1102-店铺满赠-小礼品;1103-限时折扣;1104-优惠套装;1105-店铺代金券coupon优惠券;1106-拼团;1107-满减送;1108-阶梯价;1109-积分换购',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '促销活动ID:与activity_type_id搭配使用, 团购ID/限时折扣ID/优惠套装ID/积分兑换编号',
  `activity_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '礼包活动对应兑换码code',
  `order_item_commission_rate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '分佣金比例百分比',
  `order_item_commission_fee` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '佣金',
  `order_item_commission_fee_refund` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款佣金',
  `policy_discountrate` decimal(6, 2) UNSIGNED NOT NULL DEFAULT 100.00 COMMENT '价格策略折扣率',
  `order_item_voucher` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '分配优惠券额度',
  `order_item_reduce` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '分配满减额度',
  `order_item_sale_person_discount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '分配销售员折扣额度',
  `order_item_note` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `order_item_file` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单附件',
  `order_item_confirm_file` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家附件',
  `order_item_confirm_status` bit(1) NOT NULL DEFAULT b'0' COMMENT '买家确认状态(BOOL):0-为确认;1-已确认',
  `order_item_saler_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '单品分销者编号',
  `item_src_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '分销商品编号',
  `order_item_supplier_sync` bit(1) NOT NULL DEFAULT b'0' COMMENT '拆单同步状态(BOOL):0-未同步;1-已同步',
  `src_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源订单',
  `order_item_return_agree_num` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '同意退货数量',
  `order_give_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '满返优惠券id',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `item_plus_discount_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'PLUS优惠金额',
  `item_level_discount_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '会员等级优惠金额',
  PRIMARY KEY (`order_item_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE COMMENT '(null)',
  INDEX `item_id`(`item_id`) USING BTREE,
  INDEX `buyer_id`(`user_id`) USING BTREE,
  INDEX `store_id`(`store_id`, `user_id`, `product_id`) USING BTREE,
  INDEX `store_id_2`(`store_id`, `order_item_evaluation_status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单商品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_logistics
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_logistics`;
CREATE TABLE `trade_order_logistics`  (
  `order_logistics_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单物流编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `stock_bill_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '出入库单据id=stock_bill_id',
  `order_tracking_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单物流单号AIRWAY BILL number',
  `logistics_explain` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '卖家备注发货备忘',
  `logistics_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发货时间配送时间',
  `logistics_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '对应快递公司',
  `express_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '快递名称',
  `express_id` mediumint(9) NOT NULL DEFAULT 0 COMMENT '快递编号',
  `ss_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '送货编号',
  `logistics_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '送货联系电话',
  `logistics_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '送货联系手机',
  `logistics_contacter` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '送货联系人',
  `logistics_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '送货联系地址',
  `logistics_postcode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  `logistics_enable` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否有效(BOOL):1-有效; 0-无效',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_logistics_id`) USING BTREE,
  UNIQUE INDEX `order_id_2`(`order_id`, `stock_bill_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `stock_bill_id`(`stock_bill_id`) USING BTREE,
  INDEX `logistics_id`(`logistics_id`) USING BTREE,
  INDEX `order_tracking_number`(`order_tracking_number`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单发货物流信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_return
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_return`;
CREATE TABLE `trade_order_return`  (
  `return_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '退单号',
  `service_type_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 2 COMMENT '服务类型(ENUM):1-退款;2-退货;3-换货;4-维修',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
  `return_refund_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款金额',
  `return_refund_point` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '积分部分',
  `store_id` int(10) UNSIGNED NOT NULL COMMENT '店铺编号',
  `buyer_user_id` int(10) UNSIGNED NOT NULL COMMENT '买家编号',
  `buyer_store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家是否有店铺',
  `return_add_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '添加时间',
  `return_reason_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退款理由编号',
  `return_buyer_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家退货手机号',
  `return_buyer_message` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '买家退货备注',
  `return_addr_contacter` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人',
  `return_tel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `return_addr` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货地址详情',
  `return_post_code` int(11) NOT NULL DEFAULT 0 COMMENT '邮编',
  `express_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流公司编号',
  `return_tracking_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '物流名称',
  `return_tracking_number` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '物流单号',
  `plantform_return_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 3180 COMMENT '申请状态平台(ENUM):3180-未申请;3181-待处理;3182-为已完成',
  `return_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 3105 COMMENT '卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成',
  `return_is_paid` bit(1) NOT NULL DEFAULT b'0' COMMENT '退款完成',
  `return_is_shipping_fee` bit(1) NOT NULL DEFAULT b'0' COMMENT '退货类型(BOOL): 0-退款单;1-退运费单',
  `return_shipping_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退运费额度',
  `return_flag` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '退货类型(ENUM): 0-不用退货;1-需要退货',
  `return_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '申请类型(ENUM): 1-退款申请; 2-退货申请; 3-虚拟退款  ',
  `return_order_lock` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '订单锁定类型(BOOL):1-不用锁定;2-需要锁定',
  `return_item_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 2030 COMMENT '物流状态(LIST):2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;',
  `return_store_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '商家处理时间',
  `return_store_message` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家备注',
  `return_commision_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退还佣金',
  `return_finish_time` timestamp(0) NOT NULL COMMENT '退款完成时间',
  `return_platform_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台留言',
  `return_is_settlemented` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单是否结算(BOOL): 0-未结算; 1-已结算',
  `return_settlement_time` timestamp(0) NOT NULL COMMENT '订单结算时间',
  `return_channel_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'money' COMMENT '退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信',
  `return_channel_flag` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '渠道是否退款(ENUM): 0-待退; 1-已退; 2-异常',
  `return_channel_time` timestamp(0) NOT NULL COMMENT '渠道退款时间',
  `return_channel_trans_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '渠道退款单号',
  `deposit_trade_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易号',
  `payment_channel_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付渠道',
  `trade_payment_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实付金额:在线支付金额',
  `return_contact_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `return_store_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核人员id',
  `return_withdraw_confirm` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现审核(BOOL):0-未审核; 1-已审核',
  `return_financial_confirm` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退款财务确认(BOOL):0-未确认; 1-已确认',
  `return_financial_confirm_time` timestamp(0) NOT NULL COMMENT '退款财务确认时间',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `return_purchase_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '取消退货备注',
  PRIMARY KEY (`return_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `buyer_user_id`(`buyer_user_id`) USING BTREE,
  INDEX `store_id`(`store_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '退款退货表-发货退货,卖家也可以决定不退货退款，买家申请退款不支持。卖家可以主动退款。' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_return_item
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_return_item`;
CREATE TABLE `trade_order_return_item`  (
  `order_return_item_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `return_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '退单号',
  `order_item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单项目编号',
  `return_item_num` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退货商品数量',
  `return_item_store_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家备注',
  `return_reason_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '退款理由',
  `return_item_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退货申请原因',
  `return_item_subtotal` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款总额',
  `return_item_commision_fee` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '退款佣金总额',
  `return_item_image` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退款凭据(DOT)',
  `return_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 3105 COMMENT '卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`order_return_item_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `return_id`(`return_id`) USING BTREE,
  INDEX `order_item_id`(`order_item_id`) USING BTREE,
  INDEX `return_state_id`(`return_state_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单退货详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_return_reason
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_return_reason`;
CREATE TABLE `trade_order_return_reason`  (
  `return_reason_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '售后编号',
  `return_reason_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '售后理由',
  `return_reason_sort` int(11) NOT NULL DEFAULT 225 COMMENT '售后排序',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`return_reason_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '退货原因表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_order_state_log
-- ----------------------------
DROP TABLE IF EXISTS `trade_order_state_log`;
CREATE TABLE `trade_order_state_log`  (
  `order_state_log_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '状态编号',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `order_state_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;',
  `order_state_pre_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;',
  `order_state_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作类别',
  `order_state_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '操作用户',
  `user_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作账号',
  `order_state_note` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作备注',
  `order_state_is_sync` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否同步(BOOL):0-否;1-是',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  `order_reduce_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '减少金额',
  PRIMARY KEY (`order_state_log_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单状态记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_user_cart
-- ----------------------------
DROP TABLE IF EXISTS `trade_user_cart`;
CREATE TABLE `trade_user_cart`  (
  `cart_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '买家编号',
  `store_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺编号',
  `chain_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '门店编号',
  `product_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品编号',
  `item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品编号',
  `cart_quantity` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '购买商品数量',
  `cart_type` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销',
  `activity_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动Id-加价购等等加入购物的需要提示',
  `activity_item_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '加入购物车所属活动Item编号',
  `cart_select` bit(1) NOT NULL DEFAULT b'1' COMMENT '选中状态(BOOL):0-未选;1-已选',
  `cart_ttl` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '有效时间戳',
  `cart_time` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间戳',
  `cart_file` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本',
  `subsite_id` int(10) UNSIGNED NOT NULL DEFAULT 1001 COMMENT '租户编号',
  PRIMARY KEY (`cart_id`) USING BTREE,
  UNIQUE INDEX `user_id_2`(`user_id`, `item_id`, `activity_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE COMMENT '(null)',
  INDEX `user_cart_ibfk_1`(`store_id`) USING BTREE,
  INDEX `user_cart_ibfk_3`(`item_id`) USING BTREE,
  INDEX `user_id_3`(`user_id`, `store_id`, `item_id`, `activity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '购物车表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
