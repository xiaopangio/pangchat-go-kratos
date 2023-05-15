/*
 Navicat Premium Data Transfer

 Source Server         : pangchat
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3307
 Source Schema         : pangchat

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 15/05/2023 15:07:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cities
-- ----------------------------
DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities`  (
  `city_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市id',
  `city_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市名称',
  `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市所属省份id',
  PRIMARY KEY (`city_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cities
-- ----------------------------
INSERT INTO `cities` VALUES ('110100', '北京', '110000');
INSERT INTO `cities` VALUES ('120100', '天津', '120000');
INSERT INTO `cities` VALUES ('130100', '石家庄', '130000');
INSERT INTO `cities` VALUES ('130200', '唐山', '130000');
INSERT INTO `cities` VALUES ('130300', '秦皇岛', '130000');
INSERT INTO `cities` VALUES ('130400', '邯郸', '130000');
INSERT INTO `cities` VALUES ('130500', '邢台', '130000');
INSERT INTO `cities` VALUES ('130600', '保定', '130000');
INSERT INTO `cities` VALUES ('130700', '张家口', '130000');
INSERT INTO `cities` VALUES ('130800', '承德', '130000');
INSERT INTO `cities` VALUES ('130900', '沧州', '130000');
INSERT INTO `cities` VALUES ('131000', '廊坊', '130000');
INSERT INTO `cities` VALUES ('131100', '衡水', '130000');
INSERT INTO `cities` VALUES ('140100', '太原', '140000');
INSERT INTO `cities` VALUES ('140200', '大同', '140000');
INSERT INTO `cities` VALUES ('140300', '阳泉', '140000');
INSERT INTO `cities` VALUES ('140400', '长治', '140000');
INSERT INTO `cities` VALUES ('140500', '晋城', '140000');
INSERT INTO `cities` VALUES ('140600', '朔州', '140000');
INSERT INTO `cities` VALUES ('140700', '晋中', '140000');
INSERT INTO `cities` VALUES ('140800', '运城', '140000');
INSERT INTO `cities` VALUES ('140900', '忻州', '140000');
INSERT INTO `cities` VALUES ('141000', '临汾', '140000');
INSERT INTO `cities` VALUES ('141100', '吕梁', '140000');
INSERT INTO `cities` VALUES ('150100', '呼和浩特', '150000');
INSERT INTO `cities` VALUES ('150200', '包头', '150000');
INSERT INTO `cities` VALUES ('150300', '乌海', '150000');
INSERT INTO `cities` VALUES ('150400', '赤峰', '150000');
INSERT INTO `cities` VALUES ('150500', '通辽', '150000');
INSERT INTO `cities` VALUES ('150600', '鄂尔多斯', '150000');
INSERT INTO `cities` VALUES ('150700', '呼伦贝尔', '150000');
INSERT INTO `cities` VALUES ('150800', '巴彦淖尔', '150000');
INSERT INTO `cities` VALUES ('150900', '乌兰察布', '150000');
INSERT INTO `cities` VALUES ('152200', '兴安盟', '150000');
INSERT INTO `cities` VALUES ('152500', '锡林郭勒盟', '150000');
INSERT INTO `cities` VALUES ('152900', '阿拉善盟', '150000');
INSERT INTO `cities` VALUES ('210100', '沈阳', '210000');
INSERT INTO `cities` VALUES ('210200', '大连', '210000');
INSERT INTO `cities` VALUES ('210300', '鞍山', '210000');
INSERT INTO `cities` VALUES ('210400', '抚顺', '210000');
INSERT INTO `cities` VALUES ('210500', '本溪', '210000');
INSERT INTO `cities` VALUES ('210600', '丹东', '210000');
INSERT INTO `cities` VALUES ('210700', '锦州', '210000');
INSERT INTO `cities` VALUES ('210800', '营口', '210000');
INSERT INTO `cities` VALUES ('210900', '阜新', '210000');
INSERT INTO `cities` VALUES ('211000', '辽阳', '210000');
INSERT INTO `cities` VALUES ('211100', '盘锦', '210000');
INSERT INTO `cities` VALUES ('211200', '铁岭', '210000');
INSERT INTO `cities` VALUES ('211300', '朝阳', '210000');
INSERT INTO `cities` VALUES ('211400', '葫芦岛', '210000');
INSERT INTO `cities` VALUES ('220100', '长春', '220000');
INSERT INTO `cities` VALUES ('220200', '吉林', '220000');
INSERT INTO `cities` VALUES ('220300', '四平', '220000');
INSERT INTO `cities` VALUES ('220400', '辽源', '220000');
INSERT INTO `cities` VALUES ('220500', '通化', '220000');
INSERT INTO `cities` VALUES ('220600', '白山', '220000');
INSERT INTO `cities` VALUES ('220700', '松原', '220000');
INSERT INTO `cities` VALUES ('220800', '白城', '220000');
INSERT INTO `cities` VALUES ('222400', '延边朝鲜族自治州', '220000');
INSERT INTO `cities` VALUES ('230100', '哈尔滨', '230000');
INSERT INTO `cities` VALUES ('230200', '齐齐哈尔', '230000');
INSERT INTO `cities` VALUES ('230300', '鸡西', '230000');
INSERT INTO `cities` VALUES ('230400', '鹤岗', '230000');
INSERT INTO `cities` VALUES ('230500', '双鸭山', '230000');
INSERT INTO `cities` VALUES ('230600', '大庆', '230000');
INSERT INTO `cities` VALUES ('230700', '伊春', '230000');
INSERT INTO `cities` VALUES ('230800', '佳木斯', '230000');
INSERT INTO `cities` VALUES ('230900', '七台河', '230000');
INSERT INTO `cities` VALUES ('231000', '牡丹江', '230000');
INSERT INTO `cities` VALUES ('231100', '黑河', '230000');
INSERT INTO `cities` VALUES ('231200', '绥化', '230000');
INSERT INTO `cities` VALUES ('232700', '大兴安岭地区', '230000');
INSERT INTO `cities` VALUES ('310100', '上海', '310000');
INSERT INTO `cities` VALUES ('320100', '南京', '320000');
INSERT INTO `cities` VALUES ('320200', '无锡', '320000');
INSERT INTO `cities` VALUES ('320300', '徐州', '320000');
INSERT INTO `cities` VALUES ('320400', '常州', '320000');
INSERT INTO `cities` VALUES ('320500', '苏州', '320000');
INSERT INTO `cities` VALUES ('320600', '南通', '320000');
INSERT INTO `cities` VALUES ('320700', '连云港', '320000');
INSERT INTO `cities` VALUES ('320800', '淮安', '320000');
INSERT INTO `cities` VALUES ('320900', '盐城', '320000');
INSERT INTO `cities` VALUES ('321000', '扬州', '320000');
INSERT INTO `cities` VALUES ('321100', '镇江', '320000');
INSERT INTO `cities` VALUES ('321200', '泰州', '320000');
INSERT INTO `cities` VALUES ('321300', '宿迁', '320000');
INSERT INTO `cities` VALUES ('330100', '杭州', '330000');
INSERT INTO `cities` VALUES ('330200', '宁波', '330000');
INSERT INTO `cities` VALUES ('330300', '温州', '330000');
INSERT INTO `cities` VALUES ('330400', '嘉兴', '330000');
INSERT INTO `cities` VALUES ('330500', '湖州', '330000');
INSERT INTO `cities` VALUES ('330600', '绍兴', '330000');
INSERT INTO `cities` VALUES ('330700', '金华', '330000');
INSERT INTO `cities` VALUES ('330800', '衢州', '330000');
INSERT INTO `cities` VALUES ('330900', '舟山', '330000');
INSERT INTO `cities` VALUES ('331000', '台州', '330000');
INSERT INTO `cities` VALUES ('331100', '丽水', '330000');
INSERT INTO `cities` VALUES ('340100', '合肥', '340000');
INSERT INTO `cities` VALUES ('340200', '芜湖', '340000');
INSERT INTO `cities` VALUES ('340300', '蚌埠', '340000');
INSERT INTO `cities` VALUES ('340400', '淮南', '340000');
INSERT INTO `cities` VALUES ('340500', '马鞍山', '340000');
INSERT INTO `cities` VALUES ('340600', '淮北', '340000');
INSERT INTO `cities` VALUES ('340700', '铜陵', '340000');
INSERT INTO `cities` VALUES ('340800', '安庆', '340000');
INSERT INTO `cities` VALUES ('341000', '黄山', '340000');
INSERT INTO `cities` VALUES ('341100', '滁州', '340000');
INSERT INTO `cities` VALUES ('341200', '阜阳', '340000');
INSERT INTO `cities` VALUES ('341300', '宿州', '340000');
INSERT INTO `cities` VALUES ('341400', '巢湖', '340000');
INSERT INTO `cities` VALUES ('341500', '六安', '340000');
INSERT INTO `cities` VALUES ('341600', '亳州', '340000');
INSERT INTO `cities` VALUES ('341700', '池州', '340000');
INSERT INTO `cities` VALUES ('341800', '宣城', '340000');
INSERT INTO `cities` VALUES ('350100', '福州', '350000');
INSERT INTO `cities` VALUES ('350200', '厦门', '350000');
INSERT INTO `cities` VALUES ('350300', '莆田', '350000');
INSERT INTO `cities` VALUES ('350400', '三明', '350000');
INSERT INTO `cities` VALUES ('350500', '泉州', '350000');
INSERT INTO `cities` VALUES ('350600', '漳州', '350000');
INSERT INTO `cities` VALUES ('350700', '南平', '350000');
INSERT INTO `cities` VALUES ('350800', '龙岩', '350000');
INSERT INTO `cities` VALUES ('350900', '宁德', '350000');
INSERT INTO `cities` VALUES ('360100', '南昌', '360000');
INSERT INTO `cities` VALUES ('360200', '景德镇', '360000');
INSERT INTO `cities` VALUES ('360300', '萍乡', '360000');
INSERT INTO `cities` VALUES ('360400', '九江', '360000');
INSERT INTO `cities` VALUES ('360500', '新余', '360000');
INSERT INTO `cities` VALUES ('360600', '鹰潭', '360000');
INSERT INTO `cities` VALUES ('360700', '赣州', '360000');
INSERT INTO `cities` VALUES ('360800', '吉安', '360000');
INSERT INTO `cities` VALUES ('360900', '宜春', '360000');
INSERT INTO `cities` VALUES ('361000', '抚州', '360000');
INSERT INTO `cities` VALUES ('361100', '上饶', '360000');
INSERT INTO `cities` VALUES ('370100', '济南', '370000');
INSERT INTO `cities` VALUES ('370200', '青岛', '370000');
INSERT INTO `cities` VALUES ('370300', '淄博', '370000');
INSERT INTO `cities` VALUES ('370400', '枣庄', '370000');
INSERT INTO `cities` VALUES ('370500', '东营', '370000');
INSERT INTO `cities` VALUES ('370600', '烟台', '370000');
INSERT INTO `cities` VALUES ('370700', '潍坊', '370000');
INSERT INTO `cities` VALUES ('370800', '济宁', '370000');
INSERT INTO `cities` VALUES ('370900', '泰安', '370000');
INSERT INTO `cities` VALUES ('371000', '威海', '370000');
INSERT INTO `cities` VALUES ('371100', '日照', '370000');
INSERT INTO `cities` VALUES ('371200', '莱芜', '370000');
INSERT INTO `cities` VALUES ('371300', '临沂', '370000');
INSERT INTO `cities` VALUES ('371400', '德州', '370000');
INSERT INTO `cities` VALUES ('371500', '聊城', '370000');
INSERT INTO `cities` VALUES ('371600', '滨州', '370000');
INSERT INTO `cities` VALUES ('371700', '荷泽', '370000');
INSERT INTO `cities` VALUES ('410100', '郑州', '410000');
INSERT INTO `cities` VALUES ('410200', '开封', '410000');
INSERT INTO `cities` VALUES ('410300', '洛阳', '410000');
INSERT INTO `cities` VALUES ('410400', '平顶山', '410000');
INSERT INTO `cities` VALUES ('410500', '安阳', '410000');
INSERT INTO `cities` VALUES ('410600', '鹤壁', '410000');
INSERT INTO `cities` VALUES ('410700', '新乡', '410000');
INSERT INTO `cities` VALUES ('410800', '焦作', '410000');
INSERT INTO `cities` VALUES ('410900', '濮阳', '410000');
INSERT INTO `cities` VALUES ('411000', '许昌', '410000');
INSERT INTO `cities` VALUES ('411100', '漯河', '410000');
INSERT INTO `cities` VALUES ('411200', '三门峡', '410000');
INSERT INTO `cities` VALUES ('411300', '南阳', '410000');
INSERT INTO `cities` VALUES ('411400', '商丘', '410000');
INSERT INTO `cities` VALUES ('411500', '信阳', '410000');
INSERT INTO `cities` VALUES ('411600', '周口', '410000');
INSERT INTO `cities` VALUES ('411700', '驻马店', '410000');
INSERT INTO `cities` VALUES ('420100', '武汉', '420000');
INSERT INTO `cities` VALUES ('420200', '黄石', '420000');
INSERT INTO `cities` VALUES ('420300', '十堰', '420000');
INSERT INTO `cities` VALUES ('420500', '宜昌', '420000');
INSERT INTO `cities` VALUES ('420600', '襄樊', '420000');
INSERT INTO `cities` VALUES ('420700', '鄂州', '420000');
INSERT INTO `cities` VALUES ('420800', '荆门', '420000');
INSERT INTO `cities` VALUES ('420900', '孝感', '420000');
INSERT INTO `cities` VALUES ('421000', '荆州', '420000');
INSERT INTO `cities` VALUES ('421100', '黄冈', '420000');
INSERT INTO `cities` VALUES ('421200', '咸宁', '420000');
INSERT INTO `cities` VALUES ('421300', '随州', '420000');
INSERT INTO `cities` VALUES ('422800', '恩施土家族苗族自治州', '420000');
INSERT INTO `cities` VALUES ('429000', '省直辖行政单位', '420000');
INSERT INTO `cities` VALUES ('430100', '长沙', '430000');
INSERT INTO `cities` VALUES ('430200', '株洲', '430000');
INSERT INTO `cities` VALUES ('430300', '湘潭', '430000');
INSERT INTO `cities` VALUES ('430400', '衡阳', '430000');
INSERT INTO `cities` VALUES ('430500', '邵阳', '430000');
INSERT INTO `cities` VALUES ('430600', '岳阳', '430000');
INSERT INTO `cities` VALUES ('430700', '常德', '430000');
INSERT INTO `cities` VALUES ('430800', '张家界', '430000');
INSERT INTO `cities` VALUES ('430900', '益阳', '430000');
INSERT INTO `cities` VALUES ('431000', '郴州', '430000');
INSERT INTO `cities` VALUES ('431100', '永州', '430000');
INSERT INTO `cities` VALUES ('431200', '怀化', '430000');
INSERT INTO `cities` VALUES ('431300', '娄底', '430000');
INSERT INTO `cities` VALUES ('433100', '湘西土家族苗族自治州', '430000');
INSERT INTO `cities` VALUES ('440100', '广州', '440000');
INSERT INTO `cities` VALUES ('440200', '韶关', '440000');
INSERT INTO `cities` VALUES ('440300', '深圳', '440000');
INSERT INTO `cities` VALUES ('440400', '珠海', '440000');
INSERT INTO `cities` VALUES ('440500', '汕头', '440000');
INSERT INTO `cities` VALUES ('440600', '佛山', '440000');
INSERT INTO `cities` VALUES ('440700', '江门', '440000');
INSERT INTO `cities` VALUES ('440800', '湛江', '440000');
INSERT INTO `cities` VALUES ('440900', '茂名', '440000');
INSERT INTO `cities` VALUES ('441200', '肇庆', '440000');
INSERT INTO `cities` VALUES ('441300', '惠州', '440000');
INSERT INTO `cities` VALUES ('441400', '梅州', '440000');
INSERT INTO `cities` VALUES ('441500', '汕尾', '440000');
INSERT INTO `cities` VALUES ('441600', '河源', '440000');
INSERT INTO `cities` VALUES ('441700', '阳江', '440000');
INSERT INTO `cities` VALUES ('441800', '清远', '440000');
INSERT INTO `cities` VALUES ('441900', '东莞', '440000');
INSERT INTO `cities` VALUES ('442000', '中山', '440000');
INSERT INTO `cities` VALUES ('445100', '潮州', '440000');
INSERT INTO `cities` VALUES ('445200', '揭阳', '440000');
INSERT INTO `cities` VALUES ('445300', '云浮', '440000');
INSERT INTO `cities` VALUES ('450100', '南宁', '450000');
INSERT INTO `cities` VALUES ('450200', '柳州', '450000');
INSERT INTO `cities` VALUES ('450300', '桂林', '450000');
INSERT INTO `cities` VALUES ('450400', '梧州', '450000');
INSERT INTO `cities` VALUES ('450500', '北海', '450000');
INSERT INTO `cities` VALUES ('450600', '防城港', '450000');
INSERT INTO `cities` VALUES ('450700', '钦州', '450000');
INSERT INTO `cities` VALUES ('450800', '贵港', '450000');
INSERT INTO `cities` VALUES ('450900', '玉林', '450000');
INSERT INTO `cities` VALUES ('451000', '百色', '450000');
INSERT INTO `cities` VALUES ('451100', '贺州', '450000');
INSERT INTO `cities` VALUES ('451200', '河池', '450000');
INSERT INTO `cities` VALUES ('451300', '来宾', '450000');
INSERT INTO `cities` VALUES ('451400', '崇左', '450000');
INSERT INTO `cities` VALUES ('460100', '海口', '460000');
INSERT INTO `cities` VALUES ('460200', '三亚', '460000');
INSERT INTO `cities` VALUES ('469000', '省直辖县级行政单位', '460000');
INSERT INTO `cities` VALUES ('500100', '重庆', '500000');
INSERT INTO `cities` VALUES ('500300', '', '500000');
INSERT INTO `cities` VALUES ('510100', '成都', '510000');
INSERT INTO `cities` VALUES ('510300', '自贡', '510000');
INSERT INTO `cities` VALUES ('510400', '攀枝花', '510000');
INSERT INTO `cities` VALUES ('510500', '泸州', '510000');
INSERT INTO `cities` VALUES ('510600', '德阳', '510000');
INSERT INTO `cities` VALUES ('510700', '绵阳', '510000');
INSERT INTO `cities` VALUES ('510800', '广元', '510000');
INSERT INTO `cities` VALUES ('510900', '遂宁', '510000');
INSERT INTO `cities` VALUES ('511000', '内江', '510000');
INSERT INTO `cities` VALUES ('511100', '乐山', '510000');
INSERT INTO `cities` VALUES ('511300', '南充', '510000');
INSERT INTO `cities` VALUES ('511400', '眉山', '510000');
INSERT INTO `cities` VALUES ('511500', '宜宾', '510000');
INSERT INTO `cities` VALUES ('511600', '广安', '510000');
INSERT INTO `cities` VALUES ('511700', '达州', '510000');
INSERT INTO `cities` VALUES ('511800', '雅安', '510000');
INSERT INTO `cities` VALUES ('511900', '巴中', '510000');
INSERT INTO `cities` VALUES ('512000', '资阳', '510000');
INSERT INTO `cities` VALUES ('513200', '阿坝藏族羌族自治州', '510000');
INSERT INTO `cities` VALUES ('513300', '甘孜藏族自治州', '510000');
INSERT INTO `cities` VALUES ('513400', '凉山彝族自治州', '510000');
INSERT INTO `cities` VALUES ('520100', '贵阳', '520000');
INSERT INTO `cities` VALUES ('520200', '六盘水', '520000');
INSERT INTO `cities` VALUES ('520300', '遵义', '520000');
INSERT INTO `cities` VALUES ('520400', '安顺', '520000');
INSERT INTO `cities` VALUES ('522200', '铜仁地区', '520000');
INSERT INTO `cities` VALUES ('522300', '黔西南布依族苗族自治州', '520000');
INSERT INTO `cities` VALUES ('522400', '毕节地区', '520000');
INSERT INTO `cities` VALUES ('522600', '黔东南苗族侗族自治州', '520000');
INSERT INTO `cities` VALUES ('522700', '黔南布依族苗族自治州', '520000');
INSERT INTO `cities` VALUES ('530100', '昆明', '530000');
INSERT INTO `cities` VALUES ('530300', '曲靖', '530000');
INSERT INTO `cities` VALUES ('530400', '玉溪', '530000');
INSERT INTO `cities` VALUES ('530500', '保山', '530000');
INSERT INTO `cities` VALUES ('530600', '昭通', '530000');
INSERT INTO `cities` VALUES ('530700', '丽江', '530000');
INSERT INTO `cities` VALUES ('530800', '思茅', '530000');
INSERT INTO `cities` VALUES ('530900', '临沧', '530000');
INSERT INTO `cities` VALUES ('532300', '楚雄彝族自治州', '530000');
INSERT INTO `cities` VALUES ('532500', '红河哈尼族彝族自治州', '530000');
INSERT INTO `cities` VALUES ('532600', '文山壮族苗族自治州', '530000');
INSERT INTO `cities` VALUES ('532800', '西双版纳傣族自治州', '530000');
INSERT INTO `cities` VALUES ('532900', '大理白族自治州', '530000');
INSERT INTO `cities` VALUES ('533100', '德宏傣族景颇族自治州', '530000');
INSERT INTO `cities` VALUES ('533300', '怒江傈僳族自治州', '530000');
INSERT INTO `cities` VALUES ('533400', '迪庆藏族自治州', '530000');
INSERT INTO `cities` VALUES ('540100', '拉萨', '540000');
INSERT INTO `cities` VALUES ('542100', '昌都地区', '540000');
INSERT INTO `cities` VALUES ('542200', '山南地区', '540000');
INSERT INTO `cities` VALUES ('542300', '日喀则地区', '540000');
INSERT INTO `cities` VALUES ('542400', '那曲地区', '540000');
INSERT INTO `cities` VALUES ('542500', '阿里地区', '540000');
INSERT INTO `cities` VALUES ('542600', '林芝地区', '540000');
INSERT INTO `cities` VALUES ('610100', '西安', '610000');
INSERT INTO `cities` VALUES ('610200', '铜川', '610000');
INSERT INTO `cities` VALUES ('610300', '宝鸡', '610000');
INSERT INTO `cities` VALUES ('610400', '咸阳', '610000');
INSERT INTO `cities` VALUES ('610500', '渭南', '610000');
INSERT INTO `cities` VALUES ('610600', '延安', '610000');
INSERT INTO `cities` VALUES ('610700', '汉中', '610000');
INSERT INTO `cities` VALUES ('610800', '榆林', '610000');
INSERT INTO `cities` VALUES ('610900', '安康', '610000');
INSERT INTO `cities` VALUES ('611000', '商洛', '610000');
INSERT INTO `cities` VALUES ('620100', '兰州', '620000');
INSERT INTO `cities` VALUES ('620200', '嘉峪关', '620000');
INSERT INTO `cities` VALUES ('620300', '金昌', '620000');
INSERT INTO `cities` VALUES ('620400', '白银', '620000');
INSERT INTO `cities` VALUES ('620500', '天水', '620000');
INSERT INTO `cities` VALUES ('620600', '武威', '620000');
INSERT INTO `cities` VALUES ('620700', '张掖', '620000');
INSERT INTO `cities` VALUES ('620800', '平凉', '620000');
INSERT INTO `cities` VALUES ('620900', '酒泉', '620000');
INSERT INTO `cities` VALUES ('621000', '庆阳', '620000');
INSERT INTO `cities` VALUES ('621100', '定西', '620000');
INSERT INTO `cities` VALUES ('621200', '陇南', '620000');
INSERT INTO `cities` VALUES ('622900', '临夏回族自治州', '620000');
INSERT INTO `cities` VALUES ('623000', '甘南藏族自治州', '620000');
INSERT INTO `cities` VALUES ('630100', '西宁', '630000');
INSERT INTO `cities` VALUES ('632100', '海东地区', '630000');
INSERT INTO `cities` VALUES ('632200', '海北藏族自治州', '630000');
INSERT INTO `cities` VALUES ('632300', '黄南藏族自治州', '630000');
INSERT INTO `cities` VALUES ('632500', '海南藏族自治州', '630000');
INSERT INTO `cities` VALUES ('632600', '果洛藏族自治州', '630000');
INSERT INTO `cities` VALUES ('632700', '玉树藏族自治州', '630000');
INSERT INTO `cities` VALUES ('632800', '海西蒙古族藏族自治州', '630000');
INSERT INTO `cities` VALUES ('640100', '银川', '640000');
INSERT INTO `cities` VALUES ('640200', '石嘴山', '640000');
INSERT INTO `cities` VALUES ('640300', '吴忠', '640000');
INSERT INTO `cities` VALUES ('640400', '固原', '640000');
INSERT INTO `cities` VALUES ('640500', '中卫', '640000');
INSERT INTO `cities` VALUES ('650100', '乌鲁木齐', '650000');
INSERT INTO `cities` VALUES ('650200', '克拉玛依', '650000');
INSERT INTO `cities` VALUES ('652100', '吐鲁番地区', '650000');
INSERT INTO `cities` VALUES ('652200', '哈密地区', '650000');
INSERT INTO `cities` VALUES ('652300', '昌吉回族自治州', '650000');
INSERT INTO `cities` VALUES ('652700', '博尔塔拉蒙古自治州', '650000');
INSERT INTO `cities` VALUES ('652800', '巴音郭楞蒙古自治州', '650000');
INSERT INTO `cities` VALUES ('652900', '阿克苏地区', '650000');
INSERT INTO `cities` VALUES ('653000', '克孜勒苏柯尔克孜自治州', '650000');
INSERT INTO `cities` VALUES ('653100', '喀什地区', '650000');
INSERT INTO `cities` VALUES ('653200', '和田地区', '650000');
INSERT INTO `cities` VALUES ('654000', '伊犁哈萨克自治州', '650000');
INSERT INTO `cities` VALUES ('654200', '塔城地区', '650000');
INSERT INTO `cities` VALUES ('654300', '阿勒泰地区', '650000');
INSERT INTO `cities` VALUES ('659000', '省直辖行政单位', '650000');

-- ----------------------------
-- Table structure for friend_groups
-- ----------------------------
DROP TABLE IF EXISTS `friend_groups`;
CREATE TABLE `friend_groups`  (
  `user_id` bigint NOT NULL COMMENT '用户id',
  `group_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '好友分组名称',
  PRIMARY KEY (`user_id`, `group_name`) USING BTREE,
  INDEX `group_name`(`group_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friend_groups
-- ----------------------------
INSERT INTO `friend_groups` VALUES (1656926837948813312, '好友分组1');
INSERT INTO `friend_groups` VALUES (1656945494364000256, '好友分组1');
INSERT INTO `friend_groups` VALUES (1656945606888787968, '好友分组2');

-- ----------------------------
-- Table structure for friend_requests
-- ----------------------------
DROP TABLE IF EXISTS `friend_requests`;
CREATE TABLE `friend_requests`  (
  `request_id` bigint NOT NULL COMMENT '唯一标识',
  `requester_id` bigint NULL DEFAULT NULL COMMENT '请求者uid',
  `receiver_id` bigint NULL DEFAULT NULL COMMENT '被申请者uid',
  `note_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `group_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '好友分组',
  `desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '申请描述',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求状态 // 0:未处理 1:已同意 2:已拒绝',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`request_id`) USING BTREE,
  CONSTRAINT `status` CHECK ((`status` = 0) or (`status` = 1) or (`status` = 2))
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friend_requests
-- ----------------------------
INSERT INTO `friend_requests` VALUES (1657359761496084480, 1656945606888787968, 1656926837948813312, 'account1', '好友分组2', '我是account3，我想加您为好友', '0', '2023-05-13 20:18:54', '2023-05-13 20:18:54');

-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends`  (
  `user_id` bigint NOT NULL COMMENT '好友1uid',
  `friend_id` bigint NOT NULL COMMENT '好友2uid',
  `note_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '好友1给好友2的备注',
  `group_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '好友2所处好友1的分组',
  `last_ack_msg_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '最后接受消息的id',
  `become_at` datetime NULL DEFAULT NULL COMMENT '成为好友的时间',
  PRIMARY KEY (`user_id`, `friend_id`) USING BTREE,
  INDEX `group_name`(`group_name` ASC) USING BTREE,
  INDEX `id2`(`friend_id` ASC) USING BTREE,
  CONSTRAINT `id1` FOREIGN KEY (`user_id`) REFERENCES `users` (`u_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `id2` FOREIGN KEY (`friend_id`) REFERENCES `users` (`u_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friends
-- ----------------------------

-- ----------------------------
-- Table structure for provinces
-- ----------------------------
DROP TABLE IF EXISTS `provinces`;
CREATE TABLE `provinces`  (
  `province_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省份id',
  `province_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省份名称',
  PRIMARY KEY (`province_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of provinces
-- ----------------------------
INSERT INTO `provinces` VALUES ('110000', '北京');
INSERT INTO `provinces` VALUES ('120000', '天津');
INSERT INTO `provinces` VALUES ('130000', '河北');
INSERT INTO `provinces` VALUES ('140000', '山西');
INSERT INTO `provinces` VALUES ('150000', '内蒙古');
INSERT INTO `provinces` VALUES ('210000', '辽宁');
INSERT INTO `provinces` VALUES ('220000', '吉林');
INSERT INTO `provinces` VALUES ('230000', '黑龙江');
INSERT INTO `provinces` VALUES ('310000', '上海');
INSERT INTO `provinces` VALUES ('320000', '江苏');
INSERT INTO `provinces` VALUES ('330000', '浙江');
INSERT INTO `provinces` VALUES ('340000', '安徽');
INSERT INTO `provinces` VALUES ('350000', '福建');
INSERT INTO `provinces` VALUES ('360000', '江西');
INSERT INTO `provinces` VALUES ('370000', '山东');
INSERT INTO `provinces` VALUES ('410000', '河南');
INSERT INTO `provinces` VALUES ('420000', '湖北');
INSERT INTO `provinces` VALUES ('430000', '湖南');
INSERT INTO `provinces` VALUES ('440000', '广东');
INSERT INTO `provinces` VALUES ('450000', '广西');
INSERT INTO `provinces` VALUES ('460000', '海南');
INSERT INTO `provinces` VALUES ('500000', '重庆');
INSERT INTO `provinces` VALUES ('510000', '四川');
INSERT INTO `provinces` VALUES ('520000', '贵州');
INSERT INTO `provinces` VALUES ('530000', '云南');
INSERT INTO `provinces` VALUES ('540000', '西藏');
INSERT INTO `provinces` VALUES ('610000', '陕西');
INSERT INTO `provinces` VALUES ('620000', '甘肃');
INSERT INTO `provinces` VALUES ('630000', '青海');
INSERT INTO `provinces` VALUES ('640000', '宁夏');
INSERT INTO `provinces` VALUES ('650000', '新疆');
INSERT INTO `provinces` VALUES ('710000', '台湾');
INSERT INTO `provinces` VALUES ('810000', '香港');
INSERT INTO `provinces` VALUES ('820000', '澳门');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `u_id` bigint NOT NULL COMMENT '用户唯一id',
  `account_id` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '对外标识账号id',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `personal_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个性签名',
  `city_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '所在城市id',
  `expire` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号名修改的冷却期',
  `status` int NULL DEFAULT NULL COMMENT '登录状态',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`u_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1656926837948813312, 'account1', '19944162167', '$2a$12$53.4IRd0p/oyYJiJUInzGOveulEmaCan8mqf3dJrJU7dSiH1Wm/5i', 'account1', '', '', '', 1, 'img/avatar16838817417827241960708995825.png');
INSERT INTO `users` VALUES (1656945494364000256, 'account2', '', '$2a$12$NnyBX3xJlffOBKJ5yV/BD.ostmynnGTJKG.vlnwoXhkj7aRz.kcu.', 'account2', '', '', '', 2, 'img/avatar1683881564176151025513319748.png');
INSERT INTO `users` VALUES (1656945606888787968, 'account3', '', '$2a$12$DNeSuWIVoB5j3rhOM09pJeSJudWBOO09.6G4HUih9IGp7VOvzX8hC', 'account3', '', '', '', 2, 'img/avatar_16839779347701683977934459875491677069555.png');
INSERT INTO `users` VALUES (1656945759372709888, 'account4', '', '$2a$12$0wvq.w6j2Gtm8ShoRRit4.bb8HHJn9dh29I67vUqli2PAz/oY9bY.', 'account4', '', '', '', 2, 'img/avatar16839651208429769365066165038.png');

SET FOREIGN_KEY_CHECKS = 1;
