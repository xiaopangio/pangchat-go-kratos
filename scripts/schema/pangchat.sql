/*
 Navicat Premium Data Transfer

 Source Server         : cloud-pangchat
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 43.143.137.232:3307
 Source Schema         : pangchat

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 11/06/2023 19:13:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cities
-- ----------------------------
DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities`
(
    `city_id`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市id',
    `city_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市名称',
    `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市所属省份id',
    PRIMARY KEY (`city_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cities
-- ----------------------------
INSERT INTO `cities`
VALUES ('110100', '北京', '110000');
INSERT INTO `cities`
VALUES ('120100', '天津', '120000');
INSERT INTO `cities`
VALUES ('130100', '石家庄', '130000');
INSERT INTO `cities`
VALUES ('130200', '唐山', '130000');
INSERT INTO `cities`
VALUES ('130300', '秦皇岛', '130000');
INSERT INTO `cities`
VALUES ('130400', '邯郸', '130000');
INSERT INTO `cities`
VALUES ('130500', '邢台', '130000');
INSERT INTO `cities`
VALUES ('130600', '保定', '130000');
INSERT INTO `cities`
VALUES ('130700', '张家口', '130000');
INSERT INTO `cities`
VALUES ('130800', '承德', '130000');
INSERT INTO `cities`
VALUES ('130900', '沧州', '130000');
INSERT INTO `cities`
VALUES ('131000', '廊坊', '130000');
INSERT INTO `cities`
VALUES ('131100', '衡水', '130000');
INSERT INTO `cities`
VALUES ('140100', '太原', '140000');
INSERT INTO `cities`
VALUES ('140200', '大同', '140000');
INSERT INTO `cities`
VALUES ('140300', '阳泉', '140000');
INSERT INTO `cities`
VALUES ('140400', '长治', '140000');
INSERT INTO `cities`
VALUES ('140500', '晋城', '140000');
INSERT INTO `cities`
VALUES ('140600', '朔州', '140000');
INSERT INTO `cities`
VALUES ('140700', '晋中', '140000');
INSERT INTO `cities`
VALUES ('140800', '运城', '140000');
INSERT INTO `cities`
VALUES ('140900', '忻州', '140000');
INSERT INTO `cities`
VALUES ('141000', '临汾', '140000');
INSERT INTO `cities`
VALUES ('141100', '吕梁', '140000');
INSERT INTO `cities`
VALUES ('150100', '呼和浩特', '150000');
INSERT INTO `cities`
VALUES ('150200', '包头', '150000');
INSERT INTO `cities`
VALUES ('150300', '乌海', '150000');
INSERT INTO `cities`
VALUES ('150400', '赤峰', '150000');
INSERT INTO `cities`
VALUES ('150500', '通辽', '150000');
INSERT INTO `cities`
VALUES ('150600', '鄂尔多斯', '150000');
INSERT INTO `cities`
VALUES ('150700', '呼伦贝尔', '150000');
INSERT INTO `cities`
VALUES ('150800', '巴彦淖尔', '150000');
INSERT INTO `cities`
VALUES ('150900', '乌兰察布', '150000');
INSERT INTO `cities`
VALUES ('152200', '兴安盟', '150000');
INSERT INTO `cities`
VALUES ('152500', '锡林郭勒盟', '150000');
INSERT INTO `cities`
VALUES ('152900', '阿拉善盟', '150000');
INSERT INTO `cities`
VALUES ('210100', '沈阳', '210000');
INSERT INTO `cities`
VALUES ('210200', '大连', '210000');
INSERT INTO `cities`
VALUES ('210300', '鞍山', '210000');
INSERT INTO `cities`
VALUES ('210400', '抚顺', '210000');
INSERT INTO `cities`
VALUES ('210500', '本溪', '210000');
INSERT INTO `cities`
VALUES ('210600', '丹东', '210000');
INSERT INTO `cities`
VALUES ('210700', '锦州', '210000');
INSERT INTO `cities`
VALUES ('210800', '营口', '210000');
INSERT INTO `cities`
VALUES ('210900', '阜新', '210000');
INSERT INTO `cities`
VALUES ('211000', '辽阳', '210000');
INSERT INTO `cities`
VALUES ('211100', '盘锦', '210000');
INSERT INTO `cities`
VALUES ('211200', '铁岭', '210000');
INSERT INTO `cities`
VALUES ('211300', '朝阳', '210000');
INSERT INTO `cities`
VALUES ('211400', '葫芦岛', '210000');
INSERT INTO `cities`
VALUES ('220100', '长春', '220000');
INSERT INTO `cities`
VALUES ('220200', '吉林', '220000');
INSERT INTO `cities`
VALUES ('220300', '四平', '220000');
INSERT INTO `cities`
VALUES ('220400', '辽源', '220000');
INSERT INTO `cities`
VALUES ('220500', '通化', '220000');
INSERT INTO `cities`
VALUES ('220600', '白山', '220000');
INSERT INTO `cities`
VALUES ('220700', '松原', '220000');
INSERT INTO `cities`
VALUES ('220800', '白城', '220000');
INSERT INTO `cities`
VALUES ('222400', '延边朝鲜族自治州', '220000');
INSERT INTO `cities`
VALUES ('230100', '哈尔滨', '230000');
INSERT INTO `cities`
VALUES ('230200', '齐齐哈尔', '230000');
INSERT INTO `cities`
VALUES ('230300', '鸡西', '230000');
INSERT INTO `cities`
VALUES ('230400', '鹤岗', '230000');
INSERT INTO `cities`
VALUES ('230500', '双鸭山', '230000');
INSERT INTO `cities`
VALUES ('230600', '大庆', '230000');
INSERT INTO `cities`
VALUES ('230700', '伊春', '230000');
INSERT INTO `cities`
VALUES ('230800', '佳木斯', '230000');
INSERT INTO `cities`
VALUES ('230900', '七台河', '230000');
INSERT INTO `cities`
VALUES ('231000', '牡丹江', '230000');
INSERT INTO `cities`
VALUES ('231100', '黑河', '230000');
INSERT INTO `cities`
VALUES ('231200', '绥化', '230000');
INSERT INTO `cities`
VALUES ('232700', '大兴安岭地区', '230000');
INSERT INTO `cities`
VALUES ('310100', '上海', '310000');
INSERT INTO `cities`
VALUES ('320100', '南京', '320000');
INSERT INTO `cities`
VALUES ('320200', '无锡', '320000');
INSERT INTO `cities`
VALUES ('320300', '徐州', '320000');
INSERT INTO `cities`
VALUES ('320400', '常州', '320000');
INSERT INTO `cities`
VALUES ('320500', '苏州', '320000');
INSERT INTO `cities`
VALUES ('320600', '南通', '320000');
INSERT INTO `cities`
VALUES ('320700', '连云港', '320000');
INSERT INTO `cities`
VALUES ('320800', '淮安', '320000');
INSERT INTO `cities`
VALUES ('320900', '盐城', '320000');
INSERT INTO `cities`
VALUES ('321000', '扬州', '320000');
INSERT INTO `cities`
VALUES ('321100', '镇江', '320000');
INSERT INTO `cities`
VALUES ('321200', '泰州', '320000');
INSERT INTO `cities`
VALUES ('321300', '宿迁', '320000');
INSERT INTO `cities`
VALUES ('330100', '杭州', '330000');
INSERT INTO `cities`
VALUES ('330200', '宁波', '330000');
INSERT INTO `cities`
VALUES ('330300', '温州', '330000');
INSERT INTO `cities`
VALUES ('330400', '嘉兴', '330000');
INSERT INTO `cities`
VALUES ('330500', '湖州', '330000');
INSERT INTO `cities`
VALUES ('330600', '绍兴', '330000');
INSERT INTO `cities`
VALUES ('330700', '金华', '330000');
INSERT INTO `cities`
VALUES ('330800', '衢州', '330000');
INSERT INTO `cities`
VALUES ('330900', '舟山', '330000');
INSERT INTO `cities`
VALUES ('331000', '台州', '330000');
INSERT INTO `cities`
VALUES ('331100', '丽水', '330000');
INSERT INTO `cities`
VALUES ('340100', '合肥', '340000');
INSERT INTO `cities`
VALUES ('340200', '芜湖', '340000');
INSERT INTO `cities`
VALUES ('340300', '蚌埠', '340000');
INSERT INTO `cities`
VALUES ('340400', '淮南', '340000');
INSERT INTO `cities`
VALUES ('340500', '马鞍山', '340000');
INSERT INTO `cities`
VALUES ('340600', '淮北', '340000');
INSERT INTO `cities`
VALUES ('340700', '铜陵', '340000');
INSERT INTO `cities`
VALUES ('340800', '安庆', '340000');
INSERT INTO `cities`
VALUES ('341000', '黄山', '340000');
INSERT INTO `cities`
VALUES ('341100', '滁州', '340000');
INSERT INTO `cities`
VALUES ('341200', '阜阳', '340000');
INSERT INTO `cities`
VALUES ('341300', '宿州', '340000');
INSERT INTO `cities`
VALUES ('341400', '巢湖', '340000');
INSERT INTO `cities`
VALUES ('341500', '六安', '340000');
INSERT INTO `cities`
VALUES ('341600', '亳州', '340000');
INSERT INTO `cities`
VALUES ('341700', '池州', '340000');
INSERT INTO `cities`
VALUES ('341800', '宣城', '340000');
INSERT INTO `cities`
VALUES ('350100', '福州', '350000');
INSERT INTO `cities`
VALUES ('350200', '厦门', '350000');
INSERT INTO `cities`
VALUES ('350300', '莆田', '350000');
INSERT INTO `cities`
VALUES ('350400', '三明', '350000');
INSERT INTO `cities`
VALUES ('350500', '泉州', '350000');
INSERT INTO `cities`
VALUES ('350600', '漳州', '350000');
INSERT INTO `cities`
VALUES ('350700', '南平', '350000');
INSERT INTO `cities`
VALUES ('350800', '龙岩', '350000');
INSERT INTO `cities`
VALUES ('350900', '宁德', '350000');
INSERT INTO `cities`
VALUES ('360100', '南昌', '360000');
INSERT INTO `cities`
VALUES ('360200', '景德镇', '360000');
INSERT INTO `cities`
VALUES ('360300', '萍乡', '360000');
INSERT INTO `cities`
VALUES ('360400', '九江', '360000');
INSERT INTO `cities`
VALUES ('360500', '新余', '360000');
INSERT INTO `cities`
VALUES ('360600', '鹰潭', '360000');
INSERT INTO `cities`
VALUES ('360700', '赣州', '360000');
INSERT INTO `cities`
VALUES ('360800', '吉安', '360000');
INSERT INTO `cities`
VALUES ('360900', '宜春', '360000');
INSERT INTO `cities`
VALUES ('361000', '抚州', '360000');
INSERT INTO `cities`
VALUES ('361100', '上饶', '360000');
INSERT INTO `cities`
VALUES ('370100', '济南', '370000');
INSERT INTO `cities`
VALUES ('370200', '青岛', '370000');
INSERT INTO `cities`
VALUES ('370300', '淄博', '370000');
INSERT INTO `cities`
VALUES ('370400', '枣庄', '370000');
INSERT INTO `cities`
VALUES ('370500', '东营', '370000');
INSERT INTO `cities`
VALUES ('370600', '烟台', '370000');
INSERT INTO `cities`
VALUES ('370700', '潍坊', '370000');
INSERT INTO `cities`
VALUES ('370800', '济宁', '370000');
INSERT INTO `cities`
VALUES ('370900', '泰安', '370000');
INSERT INTO `cities`
VALUES ('371000', '威海', '370000');
INSERT INTO `cities`
VALUES ('371100', '日照', '370000');
INSERT INTO `cities`
VALUES ('371200', '莱芜', '370000');
INSERT INTO `cities`
VALUES ('371300', '临沂', '370000');
INSERT INTO `cities`
VALUES ('371400', '德州', '370000');
INSERT INTO `cities`
VALUES ('371500', '聊城', '370000');
INSERT INTO `cities`
VALUES ('371600', '滨州', '370000');
INSERT INTO `cities`
VALUES ('371700', '荷泽', '370000');
INSERT INTO `cities`
VALUES ('410100', '郑州', '410000');
INSERT INTO `cities`
VALUES ('410200', '开封', '410000');
INSERT INTO `cities`
VALUES ('410300', '洛阳', '410000');
INSERT INTO `cities`
VALUES ('410400', '平顶山', '410000');
INSERT INTO `cities`
VALUES ('410500', '安阳', '410000');
INSERT INTO `cities`
VALUES ('410600', '鹤壁', '410000');
INSERT INTO `cities`
VALUES ('410700', '新乡', '410000');
INSERT INTO `cities`
VALUES ('410800', '焦作', '410000');
INSERT INTO `cities`
VALUES ('410900', '濮阳', '410000');
INSERT INTO `cities`
VALUES ('411000', '许昌', '410000');
INSERT INTO `cities`
VALUES ('411100', '漯河', '410000');
INSERT INTO `cities`
VALUES ('411200', '三门峡', '410000');
INSERT INTO `cities`
VALUES ('411300', '南阳', '410000');
INSERT INTO `cities`
VALUES ('411400', '商丘', '410000');
INSERT INTO `cities`
VALUES ('411500', '信阳', '410000');
INSERT INTO `cities`
VALUES ('411600', '周口', '410000');
INSERT INTO `cities`
VALUES ('411700', '驻马店', '410000');
INSERT INTO `cities`
VALUES ('420100', '武汉', '420000');
INSERT INTO `cities`
VALUES ('420200', '黄石', '420000');
INSERT INTO `cities`
VALUES ('420300', '十堰', '420000');
INSERT INTO `cities`
VALUES ('420500', '宜昌', '420000');
INSERT INTO `cities`
VALUES ('420600', '襄樊', '420000');
INSERT INTO `cities`
VALUES ('420700', '鄂州', '420000');
INSERT INTO `cities`
VALUES ('420800', '荆门', '420000');
INSERT INTO `cities`
VALUES ('420900', '孝感', '420000');
INSERT INTO `cities`
VALUES ('421000', '荆州', '420000');
INSERT INTO `cities`
VALUES ('421100', '黄冈', '420000');
INSERT INTO `cities`
VALUES ('421200', '咸宁', '420000');
INSERT INTO `cities`
VALUES ('421300', '随州', '420000');
INSERT INTO `cities`
VALUES ('422800', '恩施土家族苗族自治州', '420000');
INSERT INTO `cities`
VALUES ('429000', '省直辖行政单位', '420000');
INSERT INTO `cities`
VALUES ('430100', '长沙', '430000');
INSERT INTO `cities`
VALUES ('430200', '株洲', '430000');
INSERT INTO `cities`
VALUES ('430300', '湘潭', '430000');
INSERT INTO `cities`
VALUES ('430400', '衡阳', '430000');
INSERT INTO `cities`
VALUES ('430500', '邵阳', '430000');
INSERT INTO `cities`
VALUES ('430600', '岳阳', '430000');
INSERT INTO `cities`
VALUES ('430700', '常德', '430000');
INSERT INTO `cities`
VALUES ('430800', '张家界', '430000');
INSERT INTO `cities`
VALUES ('430900', '益阳', '430000');
INSERT INTO `cities`
VALUES ('431000', '郴州', '430000');
INSERT INTO `cities`
VALUES ('431100', '永州', '430000');
INSERT INTO `cities`
VALUES ('431200', '怀化', '430000');
INSERT INTO `cities`
VALUES ('431300', '娄底', '430000');
INSERT INTO `cities`
VALUES ('433100', '湘西土家族苗族自治州', '430000');
INSERT INTO `cities`
VALUES ('440100', '广州', '440000');
INSERT INTO `cities`
VALUES ('440200', '韶关', '440000');
INSERT INTO `cities`
VALUES ('440300', '深圳', '440000');
INSERT INTO `cities`
VALUES ('440400', '珠海', '440000');
INSERT INTO `cities`
VALUES ('440500', '汕头', '440000');
INSERT INTO `cities`
VALUES ('440600', '佛山', '440000');
INSERT INTO `cities`
VALUES ('440700', '江门', '440000');
INSERT INTO `cities`
VALUES ('440800', '湛江', '440000');
INSERT INTO `cities`
VALUES ('440900', '茂名', '440000');
INSERT INTO `cities`
VALUES ('441200', '肇庆', '440000');
INSERT INTO `cities`
VALUES ('441300', '惠州', '440000');
INSERT INTO `cities`
VALUES ('441400', '梅州', '440000');
INSERT INTO `cities`
VALUES ('441500', '汕尾', '440000');
INSERT INTO `cities`
VALUES ('441600', '河源', '440000');
INSERT INTO `cities`
VALUES ('441700', '阳江', '440000');
INSERT INTO `cities`
VALUES ('441800', '清远', '440000');
INSERT INTO `cities`
VALUES ('441900', '东莞', '440000');
INSERT INTO `cities`
VALUES ('442000', '中山', '440000');
INSERT INTO `cities`
VALUES ('445100', '潮州', '440000');
INSERT INTO `cities`
VALUES ('445200', '揭阳', '440000');
INSERT INTO `cities`
VALUES ('445300', '云浮', '440000');
INSERT INTO `cities`
VALUES ('450100', '南宁', '450000');
INSERT INTO `cities`
VALUES ('450200', '柳州', '450000');
INSERT INTO `cities`
VALUES ('450300', '桂林', '450000');
INSERT INTO `cities`
VALUES ('450400', '梧州', '450000');
INSERT INTO `cities`
VALUES ('450500', '北海', '450000');
INSERT INTO `cities`
VALUES ('450600', '防城港', '450000');
INSERT INTO `cities`
VALUES ('450700', '钦州', '450000');
INSERT INTO `cities`
VALUES ('450800', '贵港', '450000');
INSERT INTO `cities`
VALUES ('450900', '玉林', '450000');
INSERT INTO `cities`
VALUES ('451000', '百色', '450000');
INSERT INTO `cities`
VALUES ('451100', '贺州', '450000');
INSERT INTO `cities`
VALUES ('451200', '河池', '450000');
INSERT INTO `cities`
VALUES ('451300', '来宾', '450000');
INSERT INTO `cities`
VALUES ('451400', '崇左', '450000');
INSERT INTO `cities`
VALUES ('460100', '海口', '460000');
INSERT INTO `cities`
VALUES ('460200', '三亚', '460000');
INSERT INTO `cities`
VALUES ('469000', '省直辖县级行政单位', '460000');
INSERT INTO `cities`
VALUES ('500100', '重庆', '500000');
INSERT INTO `cities`
VALUES ('500300', '', '500000');
INSERT INTO `cities`
VALUES ('510100', '成都', '510000');
INSERT INTO `cities`
VALUES ('510300', '自贡', '510000');
INSERT INTO `cities`
VALUES ('510400', '攀枝花', '510000');
INSERT INTO `cities`
VALUES ('510500', '泸州', '510000');
INSERT INTO `cities`
VALUES ('510600', '德阳', '510000');
INSERT INTO `cities`
VALUES ('510700', '绵阳', '510000');
INSERT INTO `cities`
VALUES ('510800', '广元', '510000');
INSERT INTO `cities`
VALUES ('510900', '遂宁', '510000');
INSERT INTO `cities`
VALUES ('511000', '内江', '510000');
INSERT INTO `cities`
VALUES ('511100', '乐山', '510000');
INSERT INTO `cities`
VALUES ('511300', '南充', '510000');
INSERT INTO `cities`
VALUES ('511400', '眉山', '510000');
INSERT INTO `cities`
VALUES ('511500', '宜宾', '510000');
INSERT INTO `cities`
VALUES ('511600', '广安', '510000');
INSERT INTO `cities`
VALUES ('511700', '达州', '510000');
INSERT INTO `cities`
VALUES ('511800', '雅安', '510000');
INSERT INTO `cities`
VALUES ('511900', '巴中', '510000');
INSERT INTO `cities`
VALUES ('512000', '资阳', '510000');
INSERT INTO `cities`
VALUES ('513200', '阿坝藏族羌族自治州', '510000');
INSERT INTO `cities`
VALUES ('513300', '甘孜藏族自治州', '510000');
INSERT INTO `cities`
VALUES ('513400', '凉山彝族自治州', '510000');
INSERT INTO `cities`
VALUES ('520100', '贵阳', '520000');
INSERT INTO `cities`
VALUES ('520200', '六盘水', '520000');
INSERT INTO `cities`
VALUES ('520300', '遵义', '520000');
INSERT INTO `cities`
VALUES ('520400', '安顺', '520000');
INSERT INTO `cities`
VALUES ('522200', '铜仁地区', '520000');
INSERT INTO `cities`
VALUES ('522300', '黔西南布依族苗族自治州', '520000');
INSERT INTO `cities`
VALUES ('522400', '毕节地区', '520000');
INSERT INTO `cities`
VALUES ('522600', '黔东南苗族侗族自治州', '520000');
INSERT INTO `cities`
VALUES ('522700', '黔南布依族苗族自治州', '520000');
INSERT INTO `cities`
VALUES ('530100', '昆明', '530000');
INSERT INTO `cities`
VALUES ('530300', '曲靖', '530000');
INSERT INTO `cities`
VALUES ('530400', '玉溪', '530000');
INSERT INTO `cities`
VALUES ('530500', '保山', '530000');
INSERT INTO `cities`
VALUES ('530600', '昭通', '530000');
INSERT INTO `cities`
VALUES ('530700', '丽江', '530000');
INSERT INTO `cities`
VALUES ('530800', '思茅', '530000');
INSERT INTO `cities`
VALUES ('530900', '临沧', '530000');
INSERT INTO `cities`
VALUES ('532300', '楚雄彝族自治州', '530000');
INSERT INTO `cities`
VALUES ('532500', '红河哈尼族彝族自治州', '530000');
INSERT INTO `cities`
VALUES ('532600', '文山壮族苗族自治州', '530000');
INSERT INTO `cities`
VALUES ('532800', '西双版纳傣族自治州', '530000');
INSERT INTO `cities`
VALUES ('532900', '大理白族自治州', '530000');
INSERT INTO `cities`
VALUES ('533100', '德宏傣族景颇族自治州', '530000');
INSERT INTO `cities`
VALUES ('533300', '怒江傈僳族自治州', '530000');
INSERT INTO `cities`
VALUES ('533400', '迪庆藏族自治州', '530000');
INSERT INTO `cities`
VALUES ('540100', '拉萨', '540000');
INSERT INTO `cities`
VALUES ('542100', '昌都地区', '540000');
INSERT INTO `cities`
VALUES ('542200', '山南地区', '540000');
INSERT INTO `cities`
VALUES ('542300', '日喀则地区', '540000');
INSERT INTO `cities`
VALUES ('542400', '那曲地区', '540000');
INSERT INTO `cities`
VALUES ('542500', '阿里地区', '540000');
INSERT INTO `cities`
VALUES ('542600', '林芝地区', '540000');
INSERT INTO `cities`
VALUES ('610100', '西安', '610000');
INSERT INTO `cities`
VALUES ('610200', '铜川', '610000');
INSERT INTO `cities`
VALUES ('610300', '宝鸡', '610000');
INSERT INTO `cities`
VALUES ('610400', '咸阳', '610000');
INSERT INTO `cities`
VALUES ('610500', '渭南', '610000');
INSERT INTO `cities`
VALUES ('610600', '延安', '610000');
INSERT INTO `cities`
VALUES ('610700', '汉中', '610000');
INSERT INTO `cities`
VALUES ('610800', '榆林', '610000');
INSERT INTO `cities`
VALUES ('610900', '安康', '610000');
INSERT INTO `cities`
VALUES ('611000', '商洛', '610000');
INSERT INTO `cities`
VALUES ('620100', '兰州', '620000');
INSERT INTO `cities`
VALUES ('620200', '嘉峪关', '620000');
INSERT INTO `cities`
VALUES ('620300', '金昌', '620000');
INSERT INTO `cities`
VALUES ('620400', '白银', '620000');
INSERT INTO `cities`
VALUES ('620500', '天水', '620000');
INSERT INTO `cities`
VALUES ('620600', '武威', '620000');
INSERT INTO `cities`
VALUES ('620700', '张掖', '620000');
INSERT INTO `cities`
VALUES ('620800', '平凉', '620000');
INSERT INTO `cities`
VALUES ('620900', '酒泉', '620000');
INSERT INTO `cities`
VALUES ('621000', '庆阳', '620000');
INSERT INTO `cities`
VALUES ('621100', '定西', '620000');
INSERT INTO `cities`
VALUES ('621200', '陇南', '620000');
INSERT INTO `cities`
VALUES ('622900', '临夏回族自治州', '620000');
INSERT INTO `cities`
VALUES ('623000', '甘南藏族自治州', '620000');
INSERT INTO `cities`
VALUES ('630100', '西宁', '630000');
INSERT INTO `cities`
VALUES ('632100', '海东地区', '630000');
INSERT INTO `cities`
VALUES ('632200', '海北藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('632300', '黄南藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('632500', '海南藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('632600', '果洛藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('632700', '玉树藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('632800', '海西蒙古族藏族自治州', '630000');
INSERT INTO `cities`
VALUES ('640100', '银川', '640000');
INSERT INTO `cities`
VALUES ('640200', '石嘴山', '640000');
INSERT INTO `cities`
VALUES ('640300', '吴忠', '640000');
INSERT INTO `cities`
VALUES ('640400', '固原', '640000');
INSERT INTO `cities`
VALUES ('640500', '中卫', '640000');
INSERT INTO `cities`
VALUES ('650100', '乌鲁木齐', '650000');
INSERT INTO `cities`
VALUES ('650200', '克拉玛依', '650000');
INSERT INTO `cities`
VALUES ('652100', '吐鲁番地区', '650000');
INSERT INTO `cities`
VALUES ('652200', '哈密地区', '650000');
INSERT INTO `cities`
VALUES ('652300', '昌吉回族自治州', '650000');
INSERT INTO `cities`
VALUES ('652700', '博尔塔拉蒙古自治州', '650000');
INSERT INTO `cities`
VALUES ('652800', '巴音郭楞蒙古自治州', '650000');
INSERT INTO `cities`
VALUES ('652900', '阿克苏地区', '650000');
INSERT INTO `cities`
VALUES ('653000', '克孜勒苏柯尔克孜自治州', '650000');
INSERT INTO `cities`
VALUES ('653100', '喀什地区', '650000');
INSERT INTO `cities`
VALUES ('653200', '和田地区', '650000');
INSERT INTO `cities`
VALUES ('654000', '伊犁哈萨克自治州', '650000');
INSERT INTO `cities`
VALUES ('654200', '塔城地区', '650000');
INSERT INTO `cities`
VALUES ('654300', '阿勒泰地区', '650000');
INSERT INTO `cities`
VALUES ('659000', '省直辖行政单位', '650000');

-- ----------------------------
-- Table structure for emojis
-- ----------------------------
DROP TABLE IF EXISTS `emojis`;
CREATE TABLE `emojis`
(
    `e_id`      bigint                                                        NOT NULL AUTO_INCREMENT,
    `e_name`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `e_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`e_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 86
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of emojis
-- ----------------------------
INSERT INTO `emojis`
VALUES (1, 'grinning', '😁');
INSERT INTO `emojis`
VALUES (2, 'heart eyes', '😍');
INSERT INTO `emojis`
VALUES (3, 'thumbs up', '👍');
INSERT INTO `emojis`
VALUES (4, 'crying', '😢');
INSERT INTO `emojis`
VALUES (5, 'laughing', '😆');
INSERT INTO `emojis`
VALUES (6, 'angry', '😠');
INSERT INTO `emojis`
VALUES (7, 'smiley', '🙂');
INSERT INTO `emojis`
VALUES (8, 'wink', '😉');
INSERT INTO `emojis`
VALUES (9, 'poo', '💩');
INSERT INTO `emojis`
VALUES (10, 'unicorn', '🦄');
INSERT INTO `emojis`
VALUES (11, 'ghost', '👻');
INSERT INTO `emojis`
VALUES (12, 'alien', '👽');
INSERT INTO `emojis`
VALUES (13, 'monkey', '🐒');
INSERT INTO `emojis`
VALUES (14, 'chicken', '🐔');
INSERT INTO `emojis`
VALUES (15, 'pig', '🐷');
INSERT INTO `emojis`
VALUES (16, 'dog', '🐶');
INSERT INTO `emojis`
VALUES (17, 'cat', '🐱');
INSERT INTO `emojis`
VALUES (18, 'koala', '🐨');
INSERT INTO `emojis`
VALUES (19, 'panda', '🐼');
INSERT INTO `emojis`
VALUES (20, 'snake', '🐍');
INSERT INTO `emojis`
VALUES (21, 'spider', '🕷️');
INSERT INTO `emojis`
VALUES (22, 'scorpion', '🦂');
INSERT INTO `emojis`
VALUES (23, 'fish', '🐟');
INSERT INTO `emojis`
VALUES (24, 'octopus', '🐙');
INSERT INTO `emojis`
VALUES (25, 'dolphin', '🐬');
INSERT INTO `emojis`
VALUES (26, 'turtle', '🐢');
INSERT INTO `emojis`
VALUES (27, 'bird', '🐦');
INSERT INTO `emojis`
VALUES (28, 'penguin', '🐧');
INSERT INTO `emojis`
VALUES (29, 'elephant', '🐘');
INSERT INTO `emojis`
VALUES (30, 'giraffe', '🦒');
INSERT INTO `emojis`
VALUES (31, 'lion', '🦁');
INSERT INTO `emojis`
VALUES (32, 'tiger', '🐯');
INSERT INTO `emojis`
VALUES (33, 'bear', '🐻');
INSERT INTO `emojis`
VALUES (34, 'zebra', '🦓');
INSERT INTO `emojis`
VALUES (35, 'horse', '🐴');
INSERT INTO `emojis`
VALUES (36, 'beetle', '🐞');
INSERT INTO `emojis`
VALUES (37, 'snail', '🐌');
INSERT INTO `emojis`
VALUES (38, 'flower', '🌸');
INSERT INTO `emojis`
VALUES (39, 'plant', '🌿');
INSERT INTO `emojis`
VALUES (40, 'tree', '🌳');
INSERT INTO `emojis`
VALUES (41, 'sun', '☀️');
INSERT INTO `emojis`
VALUES (42, 'moon', '🌙');
INSERT INTO `emojis`
VALUES (43, 'star', '⭐');
INSERT INTO `emojis`
VALUES (44, 'cloud', '☁️');
INSERT INTO `emojis`
VALUES (45, 'rain', '🌧️');
INSERT INTO `emojis`
VALUES (46, 'thunderstorm', '⛈️');
INSERT INTO `emojis`
VALUES (47, 'snowflake', '❄️');
INSERT INTO `emojis`
VALUES (48, 'fire', '🔥');
INSERT INTO `emojis`
VALUES (49, 'bomb', '💣');
INSERT INTO `emojis`
VALUES (50, 'gun', '🔫');
INSERT INTO `emojis`
VALUES (51, 'rocket', '🚀');
INSERT INTO `emojis`
VALUES (52, 'airplane', '✈️');
INSERT INTO `emojis`
VALUES (53, 'car', '🚗');
INSERT INTO `emojis`
VALUES (54, 'truck', '🚚');
INSERT INTO `emojis`
VALUES (55, 'bus', '🚌');
INSERT INTO `emojis`
VALUES (56, 'train', '🚆');
INSERT INTO `emojis`
VALUES (57, 'bike', '🚲');
INSERT INTO `emojis`
VALUES (58, 'boat', '⛵');
INSERT INTO `emojis`
VALUES (59, 'ship', '🚢');
INSERT INTO `emojis`
VALUES (60, 'house', '🏠');
INSERT INTO `emojis`
VALUES (61, 'building', '🏢');
INSERT INTO `emojis`
VALUES (62, 'hospital', '🏥');
INSERT INTO `emojis`
VALUES (63, 'school', '🏫');
INSERT INTO `emojis`
VALUES (64, 'hotel', '🏨');
INSERT INTO `emojis`
VALUES (65, 'church', '⛪');
INSERT INTO `emojis`
VALUES (66, 'mosque', '🕌');
INSERT INTO `emojis`
VALUES (67, 'synagogue', '🕍');
INSERT INTO `emojis`
VALUES (68, 'tent', '⛺');
INSERT INTO `emojis`
VALUES (69, 'umbrella', '☂️');
INSERT INTO `emojis`
VALUES (70, 'book', '📖');
INSERT INTO `emojis`
VALUES (71, 'newspaper', '📰');
INSERT INTO `emojis`
VALUES (72, 'magazine', '📰');
INSERT INTO `emojis`
VALUES (73, 'phone', '📱');
INSERT INTO `emojis`
VALUES (74, 'computer', '💻');
INSERT INTO `emojis`
VALUES (75, 'television', '📺');
INSERT INTO `emojis`
VALUES (76, 'camera', '📷');
INSERT INTO `emojis`
VALUES (77, 'video camera', '📹');
INSERT INTO `emojis`
VALUES (78, 'movie camera', '🎥');
INSERT INTO `emojis`
VALUES (79, 'microphone', '🎤');
INSERT INTO `emojis`
VALUES (80, 'guitar', '🎸');
INSERT INTO `emojis`
VALUES (81, 'trumpet', '🎺');
INSERT INTO `emojis`
VALUES (82, 'violin', '🎻');
INSERT INTO `emojis`
VALUES (83, 'saxophone', '🎷');
INSERT INTO `emojis`
VALUES (84, 'piano', '🎹');
INSERT INTO `emojis`
VALUES (85, 'drum', '🥁');

-- ----------------------------
-- Table structure for friend_groups
-- ----------------------------
DROP TABLE IF EXISTS `friend_groups`;
CREATE TABLE `friend_groups`
(
    `user_id`    bigint                                                       NOT NULL COMMENT '用户id',
    `group_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '好友分组名称',
    PRIMARY KEY (`user_id`, `group_name`) USING BTREE,
    INDEX `group_name` (`group_name` ASC) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of friend_groups
-- ----------------------------
INSERT INTO `friend_groups`
VALUES (1658763585775472640, '同学');
INSERT INTO `friend_groups`
VALUES (1656926837948813312, '好友分组1');
INSERT INTO `friend_groups`
VALUES (1656945494364000256, '好友分组1');
INSERT INTO `friend_groups`
VALUES (1656945606888787968, '好友分组2');
INSERT INTO `friend_groups`
VALUES (1656926837948813312, '我的好友');
INSERT INTO `friend_groups`
VALUES (1656945759372709888, '我的好友');
INSERT INTO `friend_groups`
VALUES (1658763585775472640, '我的好友');

-- ----------------------------
-- Table structure for friend_requests
-- ----------------------------
DROP TABLE IF EXISTS `friend_requests`;
CREATE TABLE `friend_requests`
(
    `request_id`   bigint                                                       NOT NULL COMMENT '唯一标识',
    `requester_id` bigint                                                       NOT NULL COMMENT '请求者uid',
    `receiver_id`  bigint                                                       NOT NULL COMMENT '被申请者uid',
    `note_name`    varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
    `group_name`   varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '好友分组',
    `desc`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL     DEFAULT NULL COMMENT '申请描述',
    `status`       varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '0' COMMENT '请求状态 // 0:未处理 1:已同意 2:已拒绝',
    `create_at`    datetime                                                     NULL     DEFAULT NULL COMMENT '创建时间',
    `update_at`    datetime                                                     NULL     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`request_id`) USING BTREE,
    CONSTRAINT `status` CHECK ((`status` = 0) or (`status` = 1) or (`status` = 2))
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of friend_requests
-- ----------------------------
INSERT INTO `friend_requests`
VALUES (1659475293473083392, 1658763585775472640, 1656926837948813312, 'account1', '我的好友',
        '我是阿达的三大，我想加您为好友', '1', '2023-05-19 16:25:16', '2023-05-19 16:25:16');
INSERT INTO `friend_requests`
VALUES (1659917991258624000, 1658763585775472640, 1656945494364000256, 'account2', '我的好友',
        '我是阿达的三大，我想加您为好友', '1', '2023-05-20 21:44:24', '2023-05-20 21:44:24');
INSERT INTO `friend_requests`
VALUES (1663490420371361792, 1658763585775472640, 1656945606888787968, 'account3', '我的好友',
        '我是阿达的三大，我想加您为好友', '1', '2023-05-30 18:19:57', '2023-05-30 18:19:57');
INSERT INTO `friend_requests`
VALUES (1663498041161682944, 1658763585775472640, 1656945759372709888, 'account4', '我的好友',
        '我是阿达的三大，我想加您为好友', '1', '2023-05-30 18:50:14', '2023-05-30 18:50:14');

-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends`
(
    `user_id`    bigint                                                       NOT NULL COMMENT '好友1uid',
    `friend_id`  bigint                                                       NOT NULL COMMENT '好友2uid',
    `note_name`  varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '好友1给好友2的备注',
    `group_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '好友2所处好友1的分组',
    `become_at`  datetime                                                     NOT NULL COMMENT '成为好友的时间',
    `is_deleted` int                                                          NULL DEFAULT 0,
    PRIMARY KEY (`user_id`, `friend_id`) USING BTREE,
    INDEX `group_name` (`group_name` ASC) USING BTREE,
    INDEX `id2` (`friend_id` ASC) USING BTREE,
    CONSTRAINT `id1` FOREIGN KEY (`user_id`) REFERENCES `users` (`u_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `id2` FOREIGN KEY (`friend_id`) REFERENCES `users` (`u_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of friends
-- ----------------------------
INSERT INTO `friends`
VALUES (1656926837948813312, 1658763585775472640, '阿达的三大', '好友分组1', '2023-05-19 16:25:48', 0);
INSERT INTO `friends`
VALUES (1656945494364000256, 1658763585775472640, '阿达的三大', '好友分组1', '2023-05-20 21:50:28', 0);
INSERT INTO `friends`
VALUES (1656945606888787968, 1658763585775472640, '阿达的三大', '好友分组2', '2023-05-30 18:27:49', 0);
INSERT INTO `friends`
VALUES (1656945759372709888, 1658763585775472640, '阿达的三大', '我的好友', '2023-05-30 18:56:01', 0);
INSERT INTO `friends`
VALUES (1658763585775472640, 1656926837948813312, 'account1', '我的好友', '2023-05-19 16:25:48', 0);
INSERT INTO `friends`
VALUES (1658763585775472640, 1656945494364000256, 'account2', '我的好友', '2023-05-20 21:50:28', 0);
INSERT INTO `friends`
VALUES (1658763585775472640, 1656945606888787968, 'account3', '我的好友', '2023-05-30 18:27:49', 0);
INSERT INTO `friends`
VALUES (1658763585775472640, 1656945759372709888, 'account4', '我的好友', '2023-05-30 18:56:01', 0);

-- ----------------------------
-- Table structure for group_admins
-- ----------------------------
DROP TABLE IF EXISTS `group_admins`;
CREATE TABLE `group_admins`
(
    `group_id`         varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '群id',
    `administrator_id` bigint                                                       NOT NULL COMMENT '管理员id',
    `is_deleted`       int                                                          NULL DEFAULT NULL COMMENT '是否删除 0：未删除，1已删除',
    PRIMARY KEY (`group_id`, `administrator_id`) USING BTREE,
    INDEX `administrator_id` (`administrator_id` ASC) USING BTREE,
    CONSTRAINT `group_admins_ibfk_2` FOREIGN KEY (`administrator_id`) REFERENCES `users` (`u_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `group_admins_ibfk_3` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_admins
-- ----------------------------

-- ----------------------------
-- Table structure for group_members
-- ----------------------------
DROP TABLE IF EXISTS `group_members`;
CREATE TABLE `group_members`
(
    `member_id`        bigint                                                       NOT NULL COMMENT '群成员id',
    `group_id`         varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '群聊id',
    `group_note_name`  varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '群备注昵称',
    `member_note_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '成员备注昵称',
    `become_at`        datetime                                                     NOT NULL COMMENT '成为成员的时间',
    `is_deleted`       int                                                          NOT NULL DEFAULT 0 COMMENT '是否删除 0未删除，1已删除',
    `role`             int                                                          NULL     DEFAULT NULL COMMENT '权限，0代表成员，1代表admin，2代表leader',
    PRIMARY KEY (`member_id`, `group_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_members
-- ----------------------------

-- ----------------------------
-- Table structure for group_messages
-- ----------------------------
DROP TABLE IF EXISTS `group_messages`;
CREATE TABLE `group_messages`
(
    `message_id` bigint                                                       NOT NULL COMMENT '消息id',
    `group_id`   varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '群聊id',
    PRIMARY KEY (`message_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_messages
-- ----------------------------

-- ----------------------------
-- Table structure for group_requests
-- ----------------------------
DROP TABLE IF EXISTS `group_requests`;
CREATE TABLE `group_requests`
(
    `request_id`   bigint                                                        NOT NULL COMMENT '群申请id',
    `requester_id` bigint                                                        NOT NULL COMMENT '申请者id',
    `group_id`     varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '群聊id',
    `desc`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL     DEFAULT NULL COMMENT '申请说明',
    `status`       varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci   NOT NULL DEFAULT '0' COMMENT '请求状态 // 0:未处理 1:已同意 2:已拒绝',
    `create_at`    datetime                                                      NULL     DEFAULT NULL COMMENT '请求创建时间',
    `update_at`    datetime                                                      NULL     DEFAULT NULL COMMENT '请求处理时间',
    PRIMARY KEY (`request_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_requests
-- ----------------------------

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups`
(
    `group_id`        varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '群唯一标识、群号：10-13位',
    `group_name`      varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '群名称',
    `group_avatar`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL     DEFAULT NULL COMMENT '群头像地址',
    `gourp_desc`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL     DEFAULT NULL COMMENT '群介绍',
    `create_at`       datetime                                                      NOT NULL COMMENT '群创建时间',
    `group_leader_id` bigint                                                        NOT NULL COMMENT '群主id',
    `is_deleted`      int                                                           NOT NULL DEFAULT 0 COMMENT '是否删除，0表示未删除，1表示已删除',
    PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of groups
-- ----------------------------

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`
(
    `message_id`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '唯一消息标识',
    `type`        int                                                           NOT NULL COMMENT '消息类型：1为文字，2为图片，3为视频，4为文件',
    `content`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NOT NULL COMMENT '消息体',
    `sender_id`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '发送者',
    `receiver_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '接受者',
    `send_at`     varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '发送时间',
    PRIMARY KEY (`message_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of messages
-- ----------------------------
INSERT INTO `messages`
VALUES ('1663107218041475072', 1, '123456', '1658763585775472640', '1656926837948813312', '1970-01-01 08:00:00');
INSERT INTO `messages`
VALUES ('1663110152275890176', 1, '123456', '1658763585775472640', '1656926837948813312', '2023-5-29 17:8:54.49');
INSERT INTO `messages`
VALUES ('1663117989479714816', 1, '你好', '1658763585775472640', '1656926837948813312', '2023-5-29 17:40:2.584');
INSERT INTO `messages`
VALUES ('1663118724187557888', 1, '12345', '1658763585775472640', '1656926837948813312', '2023-5-29 17:42:57.753');
INSERT INTO `messages`
VALUES ('1663118916383150080', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-29 17:43:43.577');
INSERT INTO `messages`
VALUES ('1663119080141361152', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-29 17:44:22.620');
INSERT INTO `messages`
VALUES ('1663119254959951872', 1, '1234', '1658763585775472640', '1656926837948813312', '2023-5-29 17:45:4.300');
INSERT INTO `messages`
VALUES ('1663120397194760192', 1, '1234', '1658763585775472640', '1656926837948813312', '2023-5-29 17:49:36.629');
INSERT INTO `messages`
VALUES ('1663122663146524672', 1, 'asd', '1658763585775472640', '1656926837948813312', '2023-5-29 17:58:36.874');
INSERT INTO `messages`
VALUES ('1663122810089771008', 1, 'asd', '1658763585775472640', '1656926837948813312', '2023-5-29 17:59:11.908');
INSERT INTO `messages`
VALUES ('1663122911461904384', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-29 17:59:36.77');
INSERT INTO `messages`
VALUES ('1663124347377684480', 1, 'asda', '1658763585775472640', '1656926837948813312', '2023-5-29 18:5:18.425');
INSERT INTO `messages`
VALUES ('1663124507361021952', 1, '1234', '1658763585775472640', '1656926837948813312', '2023-5-29 18:5:56.570');
INSERT INTO `messages`
VALUES ('1663130399250845696', 1, 'nihao', '1658763585775472640', '1656926837948813312', '2023-5-29 18:29:21.306');
INSERT INTO `messages`
VALUES ('1663130922863562752', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-29 18:31:26.143');
INSERT INTO `messages`
VALUES ('1663132129032146944', 1, '123456', '1658763585775472640', '1656926837948813312', '2023-5-29 18:36:13.716');
INSERT INTO `messages`
VALUES ('1663132712497582080', 1, 'asd', '1658763585775472640', '1656926837948813312', '2023-5-29 18:38:32.827');
INSERT INTO `messages`
VALUES ('1663133623483633664', 1, '789a49sdas', '1658763585775472640', '1656926837948813312', '2023-5-29 18:42:10.21');
INSERT INTO `messages`
VALUES ('1663134647862693888', 1, '789798', '1658763585775472640', '1656926837948813312', '2023-5-29 18:46:14.252');
INSERT INTO `messages`
VALUES ('1663454663459999744', 1, '你好', '1658763585775472640', '1656926837948813312', '2023-5-30 15:57:51.913');
INSERT INTO `messages`
VALUES ('1663454772071501824', 1, '你好', '1658763585775472640', '1656945494364000256', '2023-5-30 15:58:17.813');
INSERT INTO `messages`
VALUES ('1663460983726673920', 1, '12345', '1658763585775472640', '1656926837948813312', '2023-5-30 16:22:58.785');
INSERT INTO `messages`
VALUES ('1663461087728635904', 1, '你好', '1658763585775472640', '1656945494364000256', '2023-5-30 16:23:23.582');
INSERT INTO `messages`
VALUES ('1663499876383592448', 1, 'abc', '1656945759372709888', '1658763585775472640', '2023-5-30 18:57:31.517');
INSERT INTO `messages`
VALUES ('1663500000828592128', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-30 18:58:1.187');
INSERT INTO `messages`
VALUES ('1663500175848509440', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-30 18:58:42.916');
INSERT INTO `messages`
VALUES ('1663500250721030144', 1, 'asd', '1656945759372709888', '1658763585775472640', '2023-5-30 18:59:0.767');
INSERT INTO `messages`
VALUES ('1663500331960504320', 1, '123', '1658763585775472640', '1656945606888787968', '2023-5-30 18:59:20.136');
INSERT INTO `messages`
VALUES ('1663500403871846400', 1, '456', '1658763585775472640', '1656945494364000256', '2023-5-30 18:59:37.282');
INSERT INTO `messages`
VALUES ('1663500698471370752', 1, '123', '1658763585775472640', '1656945494364000256', '2023-5-30 19:0:47.519');
INSERT INTO `messages`
VALUES ('1663500932282847232', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-30 19:1:43.263');
INSERT INTO `messages`
VALUES ('1663501613790138368', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-30 19:4:25.748');
INSERT INTO `messages`
VALUES ('1663501717817266176', 1, '123', '1658763585775472640', '1656926837948813312', '2023-5-30 19:4:50.551');
INSERT INTO `messages`
VALUES ('1663501882162679808', 1, '你好', '1658763585775472640', '1656926837948813312', '2023-5-30 19:5:29.733');
INSERT INTO `messages`
VALUES ('1663502390696873984', 1, '1234', '1658763585775472640', '1656926837948813312', '2023-5-30 19:7:30.976');
INSERT INTO `messages`
VALUES ('1663503083017080832', 1, '1234', '1658763585775472640', '1656926837948813312', '2023-5-30 19:10:16.39');
INSERT INTO `messages`
VALUES ('1663529601042747392', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-30 20:55:38.429');
INSERT INTO `messages`
VALUES ('1663529804827201536', 1, 'nihao', '1656945759372709888', '1658763585775472640', '2023-5-30 20:56:27.15');
INSERT INTO `messages`
VALUES ('1663530482135994368', 1, '你好', '1656945759372709888', '1658763585775472640', '2023-5-30 20:59:8.498');
INSERT INTO `messages`
VALUES ('1663531128528572416', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-30 21:1:42.608');
INSERT INTO `messages`
VALUES ('1663531182211469312', 1, '456', '1656945759372709888', '1658763585775472640', '2023-5-30 21:1:55.409');
INSERT INTO `messages`
VALUES ('1663531262163292160', 1, '阿萨德', '1656945759372709888', '1658763585775472640', '2023-5-30 21:2:14.471');
INSERT INTO `messages`
VALUES ('1663531506062069760', 1, '阿萨德', '1656945759372709888', '1658763585775472640', '2023-5-30 21:3:12.621');
INSERT INTO `messages`
VALUES ('1663532473117577216', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-30 21:7:3.184');
INSERT INTO `messages`
VALUES ('1663533453418696704', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-30 21:10:56.908');
INSERT INTO `messages`
VALUES ('1663536148502286336', 1, '123456', '1658763585775472640', '1656945759372709888', '2023-5-30 21:21:39.463');
INSERT INTO `messages`
VALUES ('1663536484755443712', 1, '8885\n<br>', '1658763585775472640', '1656945759372709888', '2023-5-30 21:22:59.634');
INSERT INTO `messages`
VALUES ('1663538055832670208', 1, '789', '1658763585775472640', '1656945759372709888', '2023-5-30 21:29:14.207');
INSERT INTO `messages`
VALUES ('1663538320476475392', 1, '没好事', '1658763585775472640', '1656945759372709888', '2023-5-30 21:30:17.304');
INSERT INTO `messages`
VALUES ('1663538590384132096', 1, '你好啊', '1656945759372709888', '1658763585775472640', '2023-5-30 21:31:21.652');
INSERT INTO `messages`
VALUES ('1663540599862595584', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-30 21:39:20.751');
INSERT INTO `messages`
VALUES ('1663541480553189376', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-30 21:42:50.724');
INSERT INTO `messages`
VALUES ('1663541687399485440', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-30 21:43:40.42');
INSERT INTO `messages`
VALUES ('1663541906472177664', 1, '你好啊', '1656945759372709888', '1658763585775472640', '2023-5-30 21:44:32.273');
INSERT INTO `messages`
VALUES ('1663723249017556992', 1, 'ad', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:7.702');
INSERT INTO `messages`
VALUES ('1663723293741420544', 1, 'asd阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:18.368');
INSERT INTO `messages`
VALUES ('1663723306487910400', 1, '1231312', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:21.406');
INSERT INTO `messages`
VALUES ('1663723312045363200', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:22.732');
INSERT INTO `messages`
VALUES ('1663723317531512832', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:24.39');
INSERT INTO `messages`
VALUES ('1663723324225622016', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:25.635');
INSERT INTO `messages`
VALUES ('1663723452575518720', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:56.236');
INSERT INTO `messages`
VALUES ('1663723457948422144', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:45:57.517');
INSERT INTO `messages`
VALUES ('1663723469939937280', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:0.376');
INSERT INTO `messages`
VALUES ('1663723475606441984', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:1.725');
INSERT INTO `messages`
VALUES ('1663723481449107456', 1, '阿萨德瓦我的', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:3.120');
INSERT INTO `messages`
VALUES ('1663723485219786752', 1, '瓦打我', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:4.20');
INSERT INTO `messages`
VALUES ('1663723488151605248', 1, '阿迪王', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:4.719');
INSERT INTO `messages`
VALUES ('1663723494283677696', 1, 'ad a', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:6.180');
INSERT INTO `messages`
VALUES ('1663723530061090816', 1, 'l', '1658763585775472640', '1656945759372709888', '2023-5-31 9:46:14.709');
INSERT INTO `messages`
VALUES ('1663724131578810368', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:48:38.124');
INSERT INTO `messages`
VALUES ('1663724208393293824', 1, '123123121111111111111111', '1658763585775472640', '1656945759372709888',
        '2023-5-31 9:48:56.438');
INSERT INTO `messages`
VALUES ('1663724227036975104', 1, '不不不', '1658763585775472640', '1656945759372709888', '2023-5-31 9:49:0.883');
INSERT INTO `messages`
VALUES ('1663724451147026432', 1, 'sad', '1658763585775472640', '1656945759372709888', '2023-5-31 9:49:54.315');
INSERT INTO `messages`
VALUES ('1663724467848744960', 1, 'sad', '1658763585775472640', '1656945759372709888', '2023-5-31 9:49:58.297');
INSERT INTO `messages`
VALUES ('1663724476870692864', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 9:50:0.444');
INSERT INTO `messages`
VALUES ('1663724622585008128', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 9:50:35.188');
INSERT INTO `messages`
VALUES ('1663724730600919040', 1, '你哈 啊', '1658763585775472640', '1656945759372709888', '2023-5-31 9:51:0.937');
INSERT INTO `messages`
VALUES ('1663836150952890368', 1, '你好啊', '1656945759372709888', '1658763585775472640', '2023-5-31 17:13:45.591');
INSERT INTO `messages`
VALUES ('1663836398270025728', 1, '你还是', '1658763585775472640', '1656945759372709888', '2023-5-31 17:14:44.587');
INSERT INTO `messages`
VALUES ('1663837193912717312', 1, '你好', '1656945759372709888', '1658763585775472640', '2023-5-31 17:17:54.284');
INSERT INTO `messages`
VALUES ('1663838499243036672', 1, '12345', '1658763585775472640', '1656945759372709888', '2023-5-31 17:23:5.498');
INSERT INTO `messages`
VALUES ('1663839390780100608', 1, '你哈', '1658763585775472640', '1656945759372709888', '2023-5-31 17:26:38.56');
INSERT INTO `messages`
VALUES ('1663839478273282048', 1, '你好', '1658763585775472640', '1656945759372709888', '2023-5-31 17:26:58.918');
INSERT INTO `messages`
VALUES ('1663841813317816320', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-31 17:36:15.634');
INSERT INTO `messages`
VALUES ('1663842164041322496', 1, '1234', '1656945759372709888', '1658763585775472640', '2023-5-31 17:37:39.256');
INSERT INTO `messages`
VALUES ('1663842892260577280', 1, '你好', '1658763585775472640', '1656945759372709888', '2023-5-31 17:40:32.876');
INSERT INTO `messages`
VALUES ('1663845067028172800', 1, '你哈', '1658763585775472640', '1656945759372709888', '2023-5-31 17:49:11.381');
INSERT INTO `messages`
VALUES ('1663845182879043584', 1, '你好', '1656945759372709888', '1658763585775472640', '2023-5-31 17:49:39.2');
INSERT INTO `messages`
VALUES ('1663845488174043136', 1, '你好啊', '1658763585775472640', '1656945759372709888', '2023-5-31 17:50:51.790');
INSERT INTO `messages`
VALUES ('1663845609825636352', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 17:51:20.793');
INSERT INTO `messages`
VALUES ('1663846717163835392', 1, '啊飒飒的', '1658763585775472640', '1656945759372709888', '2023-5-31 17:55:44.803');
INSERT INTO `messages`
VALUES ('1663846776303521792', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 17:55:58.905');
INSERT INTO `messages`
VALUES ('1663847691722952704', 1, '你好', '1658763585775472640', '1656945759372709888', '2023-5-31 17:59:37.156');
INSERT INTO `messages`
VALUES ('1663847763701403648', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 17:59:54.319');
INSERT INTO `messages`
VALUES ('1663847898657329152', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 18:0:26.495');
INSERT INTO `messages`
VALUES ('1663848129872531456', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 18:1:21.618');
INSERT INTO `messages`
VALUES ('1663848144179302400', 1, '捱三顶五', '1658763585775472640', '1656945759372709888', '2023-5-31 18:1:25.31');
INSERT INTO `messages`
VALUES ('1663848651463593984', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 18:3:25.977');
INSERT INTO `messages`
VALUES ('1663848922629541888', 1, '你好', '1658763585775472640', '1656945759372709888', '2023-5-31 18:4:30.629');
INSERT INTO `messages`
VALUES ('1663849280206540800', 1, '123', '1656945759372709888', '1658763585775472640', '2023-5-31 18:5:55.881');
INSERT INTO `messages`
VALUES ('1663849354605105152', 1, '你好', '1656945759372709888', '1658763585775472640', '2023-5-31 18:6:13.617');
INSERT INTO `messages`
VALUES ('1663849520070397952', 1, '哇呵呵', '1658763585775472640', '1656945759372709888', '2023-5-31 18:6:53.69');
INSERT INTO `messages`
VALUES ('1663851107291500544', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 18:13:11.491');
INSERT INTO `messages`
VALUES ('1663851145266728960', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 18:13:20.547');
INSERT INTO `messages`
VALUES ('1663851206092525568', 1, '1234', '1656945759372709888', '1658763585775472640', '2023-5-31 18:13:35.49');
INSERT INTO `messages`
VALUES ('1663851576055304192', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 18:15:3.254');
INSERT INTO `messages`
VALUES ('1663851664081162240', 1, '阿萨德阿萨德', '1658763585775472640', '1656945759372709888',
        '2023-5-31 18:15:24.241');
INSERT INTO `messages`
VALUES ('1663852052104613888', 1, '1234', '1658763585775472640', '1656945759372709888', '2023-5-31 18:16:56.752');
INSERT INTO `messages`
VALUES ('1663852881805053952', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 18:20:14.569');
INSERT INTO `messages`
VALUES ('1663852897596608512', 1, '1564564', '1658763585775472640', '1656945759372709888', '2023-5-31 18:20:18.335');
INSERT INTO `messages`
VALUES ('1663852915552423936', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-5-31 18:20:22.616');
INSERT INTO `messages`
VALUES ('1663853166652821504', 1, '捱三顶五', '1658763585775472640', '1656945759372709888', '2023-5-31 18:21:22.480');
INSERT INTO `messages`
VALUES ('1663853197191548928', 1, '156456', '1658763585775472640', '1656945759372709888', '2023-5-31 18:21:29.764');
INSERT INTO `messages`
VALUES ('1663853514071216128', 1, '阿萨德我', '1658763585775472640', '1656945759372709888', '2023-5-31 18:22:45.313');
INSERT INTO `messages`
VALUES ('1663853528524787712', 1, '888', '1658763585775472640', '1656945759372709888', '2023-5-31 18:22:48.759');
INSERT INTO `messages`
VALUES ('1663855105755058176', 1, '对方', '1658763585775472640', '1656945759372709888', '2023-5-31 18:29:4.799');
INSERT INTO `messages`
VALUES ('1663855158058029056', 1, '123', '1658763585775472640', '1656945759372709888', '2023-5-31 18:29:17.271');
INSERT INTO `messages`
VALUES ('1663856078716145664', 1, '1', '1658763585775472640', '1656945759372709888', '2023-5-31 18:32:56.772');
INSERT INTO `messages`
VALUES ('1663856122701811712', 1, '2', '1656945759372709888', '1658763585775472640', '2023-5-31 18:33:7.259');
INSERT INTO `messages`
VALUES ('1663856140271751168', 1, '3', '1658763585775472640', '1656945759372709888', '2023-5-31 18:33:11.449');
INSERT INTO `messages`
VALUES ('1663856451572994048', 1, '12345', '1658763585775472640', '1656945759372709888', '2023-5-31 18:34:25.668');
INSERT INTO `messages`
VALUES ('1663856477900640256', 1, '1', '1658763585775472640', '1656945759372709888', '2023-5-31 18:34:31.945');
INSERT INTO `messages`
VALUES ('1663856493654446080', 1, '2', '1656945759372709888', '1658763585775472640', '2023-5-31 18:34:35.702');
INSERT INTO `messages`
VALUES ('1663856506585485312', 1, '3', '1658763585775472640', '1656945759372709888', '2023-5-31 18:34:38.785');
INSERT INTO `messages`
VALUES ('1663856528064516096', 1, '4', '1656945759372709888', '1658763585775472640', '2023-5-31 18:34:43.906');
INSERT INTO `messages`
VALUES ('1663856556870995968', 1, '5', '1658763585775472640', '1656945759372709888', '2023-5-31 18:34:50.774');
INSERT INTO `messages`
VALUES ('1663856737112821760', 1, '6', '1656945759372709888', '1658763585775472640', '2023-5-31 18:35:33.744');
INSERT INTO `messages`
VALUES ('1663857272494755840', 1, '123456', '1658763585775472640', '1656945759372709888', '2023-5-31 18:37:41.390');
INSERT INTO `messages`
VALUES ('1663857762762756096', 1, '7', '1658763585775472640', '1656945759372709888', '2023-5-31 18:39:38.280');
INSERT INTO `messages`
VALUES ('1663857774712328192', 1, '8', '1658763585775472640', '1656945759372709888', '2023-5-31 18:39:41.129');
INSERT INTO `messages`
VALUES ('1663857793838354432', 1, '9', '1658763585775472640', '1656945759372709888', '2023-5-31 18:39:45.689');
INSERT INTO `messages`
VALUES ('1663858439601786880', 1, '10', '1656945759372709888', '1658763585775472640', '2023-5-31 18:42:19.650');
INSERT INTO `messages`
VALUES ('1663858855605440512', 1, '11', '1658763585775472640', '1656945759372709888', '2023-5-31 18:43:58.835');
INSERT INTO `messages`
VALUES ('1663860565652541440', 1, '12', '1656945759372709888', '1658763585775472640', '2023-5-31 18:50:46.541');
INSERT INTO `messages`
VALUES ('1663860690667966464', 1, '13', '1658763585775472640', '1656945759372709888', '2023-5-31 18:51:16.346');
INSERT INTO `messages`
VALUES ('1663860763422363648', 1, '14', '1658763585775472640', '1656945759372709888', '2023-5-31 18:51:33.692');
INSERT INTO `messages`
VALUES ('1663860778001764352', 1, '15', '1658763585775472640', '1656945759372709888', '2023-5-31 18:51:37.169');
INSERT INTO `messages`
VALUES ('1663860807701630976', 1, '16', '1658763585775472640', '1656945759372709888', '2023-5-31 18:51:44.251');
INSERT INTO `messages`
VALUES ('1663861101042864128', 1, '17', '1658763585775472640', '1656945759372709888', '2023-5-31 18:52:54.188');
INSERT INTO `messages`
VALUES ('1663861233238937600', 1, '18', '1658763585775472640', '1656945759372709888', '2023-5-31 18:53:25.707');
INSERT INTO `messages`
VALUES ('1663861314742652928', 1, '19', '1658763585775472640', '1656945759372709888', '2023-5-31 18:53:45.139');
INSERT INTO `messages`
VALUES ('1663861579382263808', 1, '20', '1658763585775472640', '1656945759372709888', '2023-5-31 18:54:48.234');
INSERT INTO `messages`
VALUES ('1663861627520290816', 1, '21', '1658763585775472640', '1656945759372709888', '2023-5-31 18:54:59.711');
INSERT INTO `messages`
VALUES ('1663861656859447296', 1, '22', '1656945759372709888', '1658763585775472640', '2023-5-31 18:55:6.705');
INSERT INTO `messages`
VALUES ('1663861970866016256', 1, '23', '1656945759372709888', '1658763585775472640', '2023-5-31 18:56:21.569');
INSERT INTO `messages`
VALUES ('1663862014746824704', 1, '24', '1658763585775472640', '1656945759372709888', '2023-5-31 18:56:32.33');
INSERT INTO `messages`
VALUES ('1663862148360572928', 1, '25', '1656945759372709888', '1658763585775472640', '2023-5-31 18:57:3.889');
INSERT INTO `messages`
VALUES ('1663862177502597120', 1, '26', '1658763585775472640', '1656945759372709888', '2023-5-31 18:57:10.837');
INSERT INTO `messages`
VALUES ('1663862861497110528', 1, '27', '1658763585775472640', '1656945759372709888', '2023-5-31 18:59:53.912');
INSERT INTO `messages`
VALUES ('1663863082851504128', 1, '28', '1658763585775472640', '1656945759372709888', '2023-5-31 19:0:46.689');
INSERT INTO `messages`
VALUES ('1663863138447003648', 1, '29', '1656945759372709888', '1658763585775472640', '2023-5-31 19:0:59.943');
INSERT INTO `messages`
VALUES ('1663863331879915520', 1, '30', '1658763585775472640', '1656945759372709888', '2023-5-31 19:1:46.60');
INSERT INTO `messages`
VALUES ('1663863522645250048', 1, '31', '1656945759372709888', '1658763585775472640', '2023-5-31 19:2:31.543');
INSERT INTO `messages`
VALUES ('1663863657236271104', 1, '32', '1658763585775472640', '1656945759372709888', '2023-5-31 19:3:3.633');
INSERT INTO `messages`
VALUES ('1663864111479394304', 1, '33', '1656945759372709888', '1658763585775472640', '2023-5-31 19:4:51.932');
INSERT INTO `messages`
VALUES ('1663864154261295104', 1, '34', '1658763585775472640', '1656945759372709888', '2023-5-31 19:5:2.132');
INSERT INTO `messages`
VALUES ('1663864250134695936', 1, '35', '1656945759372709888', '1658763585775472640', '2023-5-31 19:5:24.991');
INSERT INTO `messages`
VALUES ('1663864404510248960', 1, '36', '1658763585775472640', '1656945759372709888', '2023-5-31 19:6:1.793');
INSERT INTO `messages`
VALUES ('1663864453935927296', 1, '37', '1656945759372709888', '1658763585775472640', '2023-5-31 19:6:13.580');
INSERT INTO `messages`
VALUES ('1664138198801977344', 1, '123', '1658763585775472640', '1656945759372709888', '2023-6-1 13:13:59.440');
INSERT INTO `messages`
VALUES ('1664138256477851648', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 13:14:13.193');
INSERT INTO `messages`
VALUES ('1664138437248159744', 1, '2', '1658763585775472640', '1656945759372709888', '2023-6-1 13:14:56.293');
INSERT INTO `messages`
VALUES ('1664138608623226880', 1, '123', '1658763585775472640', '1656945759372709888', '2023-6-1 13:15:37.151');
INSERT INTO `messages`
VALUES ('1664138965206175744', 1, '1564', '1658763585775472640', '1656945759372709888', '2023-6-1 13:17:2.166');
INSERT INTO `messages`
VALUES ('1664175134396649472', 1, '1234', '1658763585775472640', '1656945759372709888', '2023-6-1 15:40:45.574');
INSERT INTO `messages`
VALUES ('1664175357143552000', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-1 15:41:38.679');
INSERT INTO `messages`
VALUES ('1664175369906819072', 1, 'sad', '1658763585775472640', '1656945759372709888', '2023-6-1 15:41:41.724');
INSERT INTO `messages`
VALUES ('1664175605643481088', 1, '456', '1658763585775472640', '1656945759372709888', '2023-6-1 15:42:37.928');
INSERT INTO `messages`
VALUES ('1664175650082131968', 1, '789', '1658763585775472640', '1656945759372709888', '2023-6-1 15:42:48.524');
INSERT INTO `messages`
VALUES ('1664175847386386432', 1, '444', '1658763585775472640', '1656945759372709888', '2023-6-1 15:43:35.565');
INSERT INTO `messages`
VALUES ('1664175872220860416', 1, '879', '1658763585775472640', '1656945759372709888', '2023-6-1 15:43:41.486');
INSERT INTO `messages`
VALUES ('1664175921130639360', 1, '44444', '1658763585775472640', '1656945759372709888', '2023-6-1 15:43:53.147');
INSERT INTO `messages`
VALUES ('1664176115175919616', 1, '789', '1658763585775472640', '1656945759372709888', '2023-6-1 15:44:39.411');
INSERT INTO `messages`
VALUES ('1664176258734362624', 1, '999', '1658763585775472640', '1656945759372709888', '2023-6-1 15:45:13.635');
INSERT INTO `messages`
VALUES ('1664176283216515072', 1, '1000', '1658763585775472640', '1656945759372709888', '2023-6-1 15:45:19.475');
INSERT INTO `messages`
VALUES ('1664176645050732544', 1, '123', '1658763585775472640', '1656945759372709888', '2023-6-1 15:46:45.740');
INSERT INTO `messages`
VALUES ('1664176971900260352', 1, '1896', '1658763585775472640', '1656945759372709888', '2023-6-1 15:48:3.668');
INSERT INTO `messages`
VALUES ('1664177002820669440', 1, '7777', '1658763585775472640', '1656945759372709888', '2023-6-1 15:48:11.41');
INSERT INTO `messages`
VALUES ('1664179234744045568', 1, '7789', '1658763585775472640', '1656945759372709888', '2023-6-1 15:57:3.172');
INSERT INTO `messages`
VALUES ('1664181146910461952', 1, 'abc', '1658763585775472640', '1656945759372709888', '2023-6-1 16:4:39.67');
INSERT INTO `messages`
VALUES ('1664181522627825664', 1, '456', '1658763585775472640', '1656945759372709888', '2023-6-1 16:6:8.646');
INSERT INTO `messages`
VALUES ('1664183433770831872', 1, '888', '1658763585775472640', '1656945759372709888', '2023-6-1 16:13:44.299');
INSERT INTO `messages`
VALUES ('1664184021774503936', 1, '9', '1658763585775472640', '1656945759372709888', '2023-6-1 16:16:4.487');
INSERT INTO `messages`
VALUES ('1664184142067142656', 1, '9991656456', '1658763585775472640', '1656945759372709888', '2023-6-1 16:16:33.169');
INSERT INTO `messages`
VALUES ('1664184310069989376', 1, '你好', '1658763585775472640', '1656945759372709888', '2023-6-1 16:17:13.225');
INSERT INTO `messages`
VALUES ('1664185916593278976', 1, '89', '1658763585775472640', '1656945759372709888', '2023-6-1 16:23:36.250');
INSERT INTO `messages`
VALUES ('1664186222974603264', 1, '999', '1658763585775472640', '1656945759372709888', '2023-6-1 16:24:49.298');
INSERT INTO `messages`
VALUES ('1664186996643336192', 1, '489', '1658763585775472640', '1656945759372709888', '2023-6-1 16:27:53.754');
INSERT INTO `messages`
VALUES ('1664187257046700032', 1, 'sad', '1658763585775472640', '1656945759372709888', '2023-6-1 16:28:55.839');
INSERT INTO `messages`
VALUES ('1664187563163783168', 1, '88', '1658763585775472640', '1656945759372709888', '2023-6-1 16:30:8.819');
INSERT INTO `messages`
VALUES ('1664188621680283648', 1, '99', '1658763585775472640', '1656945759372709888', '2023-6-1 16:34:21.193');
INSERT INTO `messages`
VALUES ('1664189270639775744', 1, '96314', '1658763585775472640', '1656945759372709888', '2023-6-1 16:36:55.917');
INSERT INTO `messages`
VALUES ('1664190657029214208', 1, '999', '1658763585775472640', '1656945759372709888', '2023-6-1 16:42:26.458');
INSERT INTO `messages`
VALUES ('1664191639121301504', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:20.604');
INSERT INTO `messages`
VALUES ('1664191643714064384', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:21.703');
INSERT INTO `messages`
VALUES ('1664191648583651328', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:22.864');
INSERT INTO `messages`
VALUES ('1664191653717479424', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:24.87');
INSERT INTO `messages`
VALUES ('1664191658410905600', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:25.207');
INSERT INTO `messages`
VALUES ('1664191673665589248', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:28.844');
INSERT INTO `messages`
VALUES ('1664191682251329536', 1, '4564', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:30.891');
INSERT INTO `messages`
VALUES ('1664191691294248960', 1, '认同', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:33.47');
INSERT INTO `messages`
VALUES ('1664191703273181184', 1, '省份', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:35.903');
INSERT INTO `messages`
VALUES ('1664191710780985344', 1, '省份', '1658763585775472640', '1656945759372709888', '2023-6-1 16:46:37.693');
INSERT INTO `messages`
VALUES ('1664192338836066304', 1, '456', '1658763585775472640', '1656945759372709888', '2023-6-1 16:49:7.432');
INSERT INTO `messages`
VALUES ('1664192483585691648', 1, '9', '1658763585775472640', '1656945759372709888', '2023-6-1 16:49:41.944');
INSERT INTO `messages`
VALUES ('1664193127650430976', 1, '456', '1658763585775472640', '1656945759372709888', '2023-6-1 16:52:15.496');
INSERT INTO `messages`
VALUES ('1664193244050755584', 1, '789', '1658763585775472640', '1656945759372709888', '2023-6-1 16:52:43.253');
INSERT INTO `messages`
VALUES ('1664193284987162624', 1, '9', '1658763585775472640', '1656945759372709888', '2023-6-1 16:52:53.12');
INSERT INTO `messages`
VALUES ('1664193642669019136', 1, '1', '1658763585775472640', '1656945759372709888', '2023-6-1 16:54:18.290');
INSERT INTO `messages`
VALUES ('1664194534688428032', 1, '9', '1658763585775472640', '1656945759372709888', '2023-6-1 16:57:50.963');
INSERT INTO `messages`
VALUES ('1664194752695767040', 1, '6', '1658763585775472640', '1656945759372709888', '2023-6-1 16:58:42.941');
INSERT INTO `messages`
VALUES ('1664195347481628672', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-1 17:1:4.747');
INSERT INTO `messages`
VALUES ('1664195472094400512', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-1 17:1:34.459');
INSERT INTO `messages`
VALUES ('1664195500494032896', 1, 'www', '1658763585775472640', '1656945759372709888', '2023-6-1 17:1:41.231');
INSERT INTO `messages`
VALUES ('1664196331033333760', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-1 17:4:59.247');
INSERT INTO `messages`
VALUES ('1664196364889755648', 1, '5', '1658763585775472640', '1656945759372709888', '2023-6-1 17:5:7.319');
INSERT INTO `messages`
VALUES ('1664196382958817280', 1, '6', '1658763585775472640', '1656945759372709888', '2023-6-1 17:5:11.626');
INSERT INTO `messages`
VALUES ('1664197045038092288', 1, '7', '1658763585775472640', '1656945759372709888', '2023-6-1 17:7:49.476');
INSERT INTO `messages`
VALUES ('1664197210809569280', 1, '8', '1658763585775472640', '1656945759372709888', '2023-6-1 17:8:29.2');
INSERT INTO `messages`
VALUES ('1664197219634384896', 1, '9', '1658763585775472640', '1656945759372709888', '2023-6-1 17:8:31.105');
INSERT INTO `messages`
VALUES ('1664541669946560512', 1, '456', '1658763585775472640', '1656945759372709888', '2023-6-2 15:57:14.456');
INSERT INTO `messages`
VALUES ('1664541758823862272', 1, '888', '1658763585775472640', '1656945759372709888', '2023-6-2 15:57:35.650');
INSERT INTO `messages`
VALUES ('1664542584329998336', 1, '你好啊', '1658763585775472640', '1656945759372709888', '2023-6-2 16:0:52.463');
INSERT INTO `messages`
VALUES ('1664542633969586176', 1, '123', '1658763585775472640', '1656945759372709888', '2023-6-2 16:1:4.298');
INSERT INTO `messages`
VALUES ('1664543055669104640', 1, '对对对', '1658763585775472640', '1656945759372709888', '2023-6-2 16:2:44.841');
INSERT INTO `messages`
VALUES ('1664543402760343552', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-2 16:4:7.595');
INSERT INTO `messages`
VALUES ('1664544249728733184', 1, '你好啊', '1658763585775472640', '1656945759372709888', '2023-6-2 16:7:29.527');
INSERT INTO `messages`
VALUES ('1664544694534672384', 1, 'sad', '1658763585775472640', '1656945759372709888', '2023-6-2 16:9:15.577');
INSERT INTO `messages`
VALUES ('1664545197360418816', 1, '88', '1658763585775472640', '1656945759372709888', '2023-6-2 16:11:15.458');
INSERT INTO `messages`
VALUES ('1664545381280649216', 1, '爱仕达无多 啊', '1658763585775472640', '1656945759372709888',
        '2023-6-2 16:11:59.310');
INSERT INTO `messages`
VALUES ('1664545591016820736', 1, '是', '1658763585775472640', '1656945759372709888', '2023-6-2 16:12:49.316');
INSERT INTO `messages`
VALUES ('1664545771766157312', 1, '阿斯达', '1658763585775472640', '1656945759372709888', '2023-6-2 16:13:32.410');
INSERT INTO `messages`
VALUES ('1664545944835723264', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-2 16:14:13.672');
INSERT INTO `messages`
VALUES ('1664546202697338880', 1, '996', '1658763585775472640', '1656945759372709888', '2023-6-2 16:15:15.152');
INSERT INTO `messages`
VALUES ('1664546366690430976', 1, '是', '1658763585775472640', '1656945759372709888', '2023-6-2 16:15:54.250');
INSERT INTO `messages`
VALUES ('1664546623708991488', 1, '你哈', '1658763585775472640', '1656945759372709888', '2023-6-2 16:16:55.527');
INSERT INTO `messages`
VALUES ('1664546764964761600', 1, '撒', '1658763585775472640', '1656945759372709888', '2023-6-2 16:17:29.206');
INSERT INTO `messages`
VALUES ('1664546861773492224', 1, '阿萨德', '1658763585775472640', '1656945759372709888', '2023-6-2 16:17:52.287');
INSERT INTO `messages`
VALUES ('1664546932862750720', 1, '是', '1658763585775472640', '1656945759372709888', '2023-6-2 16:18:9.236');
INSERT INTO `messages`
VALUES ('1664546980656844800', 1, '999', '1656945759372709888', '1658763585775472640', '2023-6-2 16:18:20.631');
INSERT INTO `messages`
VALUES ('1664550983671222272', 1, '你好', '1656945606888787968', '1658763585775472640', '2023-6-2 16:34:15.24');
INSERT INTO `messages`
VALUES ('1664551929667129344', 1, '8888', '1656926837948813312', '1658763585775472640', '2023-6-2 16:38:0.565');
INSERT INTO `messages`
VALUES ('1664552043034972160', 1, '1011', '1656926837948813312', '1658763585775472640', '2023-6-2 16:38:27.596');
INSERT INTO `messages`
VALUES ('1664552058885246976', 1, '捱三顶五啊', '1656926837948813312', '1658763585775472640', '2023-6-2 16:38:31.375');
INSERT INTO `messages`
VALUES ('1664552070616715264', 1, '挖到我', '1656926837948813312', '1658763585775472640', '2023-6-2 16:38:34.173');
INSERT INTO `messages`
VALUES ('1664553660882882560', 1, '8888', '1656945494364000256', '1658763585775472640', '2023-6-2 16:44:53.321');
INSERT INTO `messages`
VALUES ('1664554037804011520', 1, '9999', '1656945494364000256', '1658763585775472640', '2023-6-2 16:46:23.185');
INSERT INTO `messages`
VALUES ('1664804941958483968', 1, '你好', '1656945494364000256', '1658763585775472640', '2023-6-3 9:23:23.396');
INSERT INTO `messages`
VALUES ('1664805146376278016', 1, '1', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:12.135');
INSERT INTO `messages`
VALUES ('1664805151438802944', 1, '2', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:13.342');
INSERT INTO `messages`
VALUES ('1664805156635545600', 1, '3', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:14.581');
INSERT INTO `messages`
VALUES ('1664805161106673664', 1, '4', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:15.647');
INSERT INTO `messages`
VALUES ('1664805168765472768', 1, '5', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:17.472');
INSERT INTO `messages`
VALUES ('1664805174390034432', 1, '6', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:18.814');
INSERT INTO `messages`
VALUES ('1664805182090776576', 1, '7', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:20.651');
INSERT INTO `messages`
VALUES ('1664805185886621696', 1, '8', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:21.555');
INSERT INTO `messages`
VALUES ('1664805191909642240', 1, '9', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:22.992');
INSERT INTO `messages`
VALUES ('1664805204765184000', 1, '10', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:26.57');
INSERT INTO `messages`
VALUES ('1664805211635453952', 1, '11', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:27.695');
INSERT INTO `messages`
VALUES ('1664805218367311872', 1, '12', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:29.300');
INSERT INTO `messages`
VALUES ('1664805228223926272', 1, '13', '1656945494364000256', '1658763585775472640', '2023-6-3 9:24:31.650');
INSERT INTO `messages`
VALUES ('1664806386627776512', 1, '1', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:7.834');
INSERT INTO `messages`
VALUES ('1664806395117047808', 1, '2', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:9.858');
INSERT INTO `messages`
VALUES ('1664806401630801920', 1, '3', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:11.412');
INSERT INTO `messages`
VALUES ('1664806429283848192', 1, '4', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:18.5');
INSERT INTO `messages`
VALUES ('1664806436728737792', 1, '5', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:19.780');
INSERT INTO `messages`
VALUES ('1664806443754196992', 1, '6', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:21.455');
INSERT INTO `messages`
VALUES ('1664806450611884032', 1, '7', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:23.90');
INSERT INTO `messages`
VALUES ('1664806458258100224', 1, '8', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:24.913');
INSERT INTO `messages`
VALUES ('1664806463723278336', 1, '9', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:26.215');
INSERT INTO `messages`
VALUES ('1664806471197528064', 1, '10', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:27.997');
INSERT INTO `messages`
VALUES ('1664806478348816384', 1, '11', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:29.702');
INSERT INTO `messages`
VALUES ('1664806487521759232', 1, '12', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:31.889');
INSERT INTO `messages`
VALUES ('1664806525169831936', 1, '13', '1656926837948813312', '1658763585775472640', '2023-6-3 9:29:40.866');
INSERT INTO `messages`
VALUES ('1664807236804808704', 1, '1', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:30.531');
INSERT INTO `messages`
VALUES ('1664807240546127872', 1, '2', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:31.424');
INSERT INTO `messages`
VALUES ('1664807245092753408', 1, '3', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:32.508');
INSERT INTO `messages`
VALUES ('1664807248964096000', 1, '4', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:33.431');
INSERT INTO `messages`
VALUES ('1664807255918252032', 1, '5', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:35.89');
INSERT INTO `messages`
VALUES ('1664807259613433856', 1, '6', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:35.970');
INSERT INTO `messages`
VALUES ('1664807266685030400', 1, '7', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:37.656');
INSERT INTO `messages`
VALUES ('1664807270321491968', 1, '8', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:38.523');
INSERT INTO `messages`
VALUES ('1664807275035889664', 1, '9', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:39.647');
INSERT INTO `messages`
VALUES ('1664807282115874816', 1, '10', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:41.335');
INSERT INTO `messages`
VALUES ('1664807286515699712', 1, '11', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:42.384');
INSERT INTO `messages`
VALUES ('1664807291825688576', 1, '12', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:43.650');
INSERT INTO `messages`
VALUES ('1664807304731561984', 1, '13', '1656945606888787968', '1658763585775472640', '2023-6-3 9:32:46.727');
INSERT INTO `messages`
VALUES ('1664807673457020928', 1, '1', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:14.638');
INSERT INTO `messages`
VALUES ('1664807680264376320', 1, '2', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:16.262');
INSERT INTO `messages`
VALUES ('1664807689143717888', 1, '3', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:18.378');
INSERT INTO `messages`
VALUES ('1664807693426102272', 1, '4', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:19.399');
INSERT INTO `messages`
VALUES ('1664807700799688704', 1, '5', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:21.158');
INSERT INTO `messages`
VALUES ('1664807705988042752', 1, '6', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:22.394');
INSERT INTO `messages`
VALUES ('1664807719418204160', 1, '7', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:25.597');
INSERT INTO `messages`
VALUES ('1664807722828173312', 1, '8', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:26.409');
INSERT INTO `messages`
VALUES ('1664807728037498880', 1, '9', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:27.651');
INSERT INTO `messages`
VALUES ('1664807732621873152', 1, '10', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:28.745');
INSERT INTO `messages`
VALUES ('1664807736665182208', 1, '11', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:29.708');
INSERT INTO `messages`
VALUES ('1664807740502970368', 1, '12', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:30.623');
INSERT INTO `messages`
VALUES ('1664807750686740480', 1, '13', '1656945759372709888', '1658763585775472640', '2023-6-3 9:34:33.51');
INSERT INTO `messages`
VALUES ('1664952045049745408', 1, '1', '1658763585775472640', '1656945606888787968', '2023-6-3 19:7:55.507');
INSERT INTO `messages`
VALUES ('1664952057649434624', 1, '2', '1658763585775472640', '1656945606888787968', '2023-6-3 19:7:58.513');
INSERT INTO `messages`
VALUES ('1664952064918163456', 1, '3', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:0.245');
INSERT INTO `messages`
VALUES ('1664952069850664960', 1, '4', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:1.422');
INSERT INTO `messages`
VALUES ('1664952074275655680', 1, '5', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:2.477');
INSERT INTO `messages`
VALUES ('1664952080588083200', 1, '6', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:3.982');
INSERT INTO `messages`
VALUES ('1664952086602715136', 1, '7', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:5.416');
INSERT INTO `messages`
VALUES ('1664952089974935552', 1, '8', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:6.220');
INSERT INTO `messages`
VALUES ('1664952095905681408', 1, '9', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:7.633');
INSERT INTO `messages`
VALUES ('1664952103677726720', 1, '10', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:9.486');
INSERT INTO `messages`
VALUES ('1664952108803166208', 1, '11', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:10.708');
INSERT INTO `messages`
VALUES ('1664952113421094912', 1, '12', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:11.810');
INSERT INTO `messages`
VALUES ('1664952121524490240', 1, '13', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:13.742');
INSERT INTO `messages`
VALUES ('1664952127748837376', 1, '14', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:15.226');
INSERT INTO `messages`
VALUES ('1664952134833016832', 1, '15', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:16.915');
INSERT INTO `messages`
VALUES ('1664952146556096512', 1, '16', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:19.710');
INSERT INTO `messages`
VALUES ('1664952153325703168', 1, '17', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:21.323');
INSERT INTO `messages`
VALUES ('1664952160019812352', 1, '18', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:22.919');
INSERT INTO `messages`
VALUES ('1664952170241331200', 1, '19', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:25.356');
INSERT INTO `messages`
VALUES ('1664952181570146304', 1, '20', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:28.58');
INSERT INTO `messages`
VALUES ('1664952188125843456', 1, '21', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:29.621');
INSERT INTO `messages`
VALUES ('1664952193003819008', 1, '22', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:30.784');
INSERT INTO `messages`
VALUES ('1664952199739871232', 1, '23', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:32.390');
INSERT INTO `messages`
VALUES ('1664952207138623488', 1, '24', '1658763585775472640', '1656945606888787968', '2023-6-3 19:8:34.154');
INSERT INTO `messages`
VALUES ('1664955474501439488', 1, '你好啊', '1656945606888787968', '1658763585775472640', '2023-6-3 19:21:33.152');
INSERT INTO `messages`
VALUES ('1667850512713650176', 1, '6666', '1658763585775472640', '1656945606888787968', '2023-6-11 19:5:25.83');

-- ----------------------------
-- Table structure for post_comments
-- ----------------------------
DROP TABLE IF EXISTS `post_comments`;
CREATE TABLE `post_comments`
(
    `comment_id`        bigint                                                        NOT NULL COMMENT '评论id',
    `post_id`           bigint                                                        NOT NULL COMMENT '推文id',
    `commenter_id`      bigint                                                        NOT NULL COMMENT '评论者id',
    `content`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
    `comment_time`      datetime                                                      NOT NULL COMMENT '评论时间',
    `parent_comment_id` bigint                                                        NULL DEFAULT NULL COMMENT '父评论id',
    PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_comments
-- ----------------------------

-- ----------------------------
-- Table structure for post_likes
-- ----------------------------
DROP TABLE IF EXISTS `post_likes`;
CREATE TABLE `post_likes`
(
    `like_id`   bigint   NOT NULL COMMENT '点赞id',
    `post_id`   bigint   NOT NULL COMMENT '推文id',
    `liker_id`  bigint   NOT NULL COMMENT '点赞者id',
    `like_time` datetime NOT NULL COMMENT '点赞时间',
    PRIMARY KEY (`like_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_likes
-- ----------------------------

-- ----------------------------
-- Table structure for post_medias
-- ----------------------------
DROP TABLE IF EXISTS `post_medias`;
CREATE TABLE `post_medias`
(
    `post_id`   bigint                                                        NOT NULL COMMENT '推文id',
    `media_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '媒体内容url',
    PRIMARY KEY (`post_id`, `media_url`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_medias
-- ----------------------------

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`
(
    `post_id`      bigint                                                NOT NULL COMMENT '推文ID',
    `publisher_id` bigint                                                NOT NULL COMMENT '发布者ID',
    `content`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '内容',
    `publish_time` datetime                                              NULL DEFAULT NULL COMMENT '发布时间',
    PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of posts
-- ----------------------------

-- ----------------------------
-- Table structure for provinces
-- ----------------------------
DROP TABLE IF EXISTS `provinces`;
CREATE TABLE `provinces`
(
    `province_id`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省份id',
    `province_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省份名称',
    PRIMARY KEY (`province_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of provinces
-- ----------------------------
INSERT INTO `provinces`
VALUES ('110000', '北京');
INSERT INTO `provinces`
VALUES ('120000', '天津');
INSERT INTO `provinces`
VALUES ('130000', '河北');
INSERT INTO `provinces`
VALUES ('140000', '山西');
INSERT INTO `provinces`
VALUES ('150000', '内蒙古');
INSERT INTO `provinces`
VALUES ('210000', '辽宁');
INSERT INTO `provinces`
VALUES ('220000', '吉林');
INSERT INTO `provinces`
VALUES ('230000', '黑龙江');
INSERT INTO `provinces`
VALUES ('310000', '上海');
INSERT INTO `provinces`
VALUES ('320000', '江苏');
INSERT INTO `provinces`
VALUES ('330000', '浙江');
INSERT INTO `provinces`
VALUES ('340000', '安徽');
INSERT INTO `provinces`
VALUES ('350000', '福建');
INSERT INTO `provinces`
VALUES ('360000', '江西');
INSERT INTO `provinces`
VALUES ('370000', '山东');
INSERT INTO `provinces`
VALUES ('410000', '河南');
INSERT INTO `provinces`
VALUES ('420000', '湖北');
INSERT INTO `provinces`
VALUES ('430000', '湖南');
INSERT INTO `provinces`
VALUES ('440000', '广东');
INSERT INTO `provinces`
VALUES ('450000', '广西');
INSERT INTO `provinces`
VALUES ('460000', '海南');
INSERT INTO `provinces`
VALUES ('500000', '重庆');
INSERT INTO `provinces`
VALUES ('510000', '四川');
INSERT INTO `provinces`
VALUES ('520000', '贵州');
INSERT INTO `provinces`
VALUES ('530000', '云南');
INSERT INTO `provinces`
VALUES ('540000', '西藏');
INSERT INTO `provinces`
VALUES ('610000', '陕西');
INSERT INTO `provinces`
VALUES ('620000', '甘肃');
INSERT INTO `provinces`
VALUES ('630000', '青海');
INSERT INTO `provinces`
VALUES ('640000', '宁夏');
INSERT INTO `provinces`
VALUES ('650000', '新疆');
INSERT INTO `provinces`
VALUES ('710000', '台湾');
INSERT INTO `provinces`
VALUES ('810000', '香港');
INSERT INTO `provinces`
VALUES ('820000', '澳门');

-- ----------------------------
-- Table structure for song_lrcs
-- ----------------------------
DROP TABLE IF EXISTS `song_lrcs`;
CREATE TABLE `song_lrcs`
(
    `lrc_song_name`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '歌曲名称',
    `lrc_singer_name`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '歌手名称',
    `lrc_song_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci         NULL COMMENT 'lrc歌词内容',
    PRIMARY KEY (`lrc_song_name`, `lrc_singer_name`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of song_lrcs
-- ----------------------------
INSERT INTO `song_lrcs`
VALUES ('Hug me (抱我)', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[00:01.54]Baby just hug me hug me\r\n[00:04.17]Baby just hug me hug me\r\n[00:07.50]每一次心动的原因\r\n[00:10.44]全部都是因为你\r\n[00:17.16]用一首歌的时间\r\n[00:18.66]在海边荡个秋千\r\n[00:20.31]巴厘岛的咖啡店\r\n[00:23.16]这空气突然变甜\r\n[00:24.72]是听到你的名字\r\n[00:26.23]不经意出现在耳边\r\n[00:29.34]想去的星球\r\n[00:33.21]要多久能带你走\r\n[00:35.37]想抱紧之后\r\n[00:39.30]就永远不再放手\r\n[00:41.01]Baby just hug me hug me\r\n[00:43.86]Baby just hug me hug me\r\n[00:46.71]每一次心动的原因\r\n[00:49.80]全部都是因为你\r\n[00:52.95]简单的爱你爱你\r\n[00:55.80]想和你在一起一起\r\n[00:59.70]不要只是\r\n[01:02.82]短暂的拥抱\r\n[01:08.70]你属于天意\r\n[01:10.05]那种命中注定\r\n[01:11.79]我还是无法\r\n[01:13.11]直视你的眼睛\r\n[01:15.33]在拥抱过后\r\n[01:16.44]画下了彩虹\r\n[01:18.39]被时间偷走\r\n[01:19.47]没做完的梦\r\n[01:20.85]想去的星球\r\n[01:24.84]还来不及带你走\r\n[01:27.06]想抱紧之后\r\n[01:30.69]就永远不再放手\r\n[01:32.57]Baby just hug me hug me\r\n[01:35.52]Baby just hug me hug me\r\n[01:38.37]每一次心动的原因\r\n[01:41.46]全部都是因为你\r\n[01:44.55]就这样爱你爱你\r\n[01:47.46]只要能在一起一起\r\n[01:51.18]不会只是\r\n[01:54.42]短暂的拥抱\r\n[01:59.34]Could u just\r\n[02:13.92]想去的星球\r\n[02:17.25]还来不及带你走\r\n[02:19.47]想抱紧之后\r\n[02:23.25]就永远不再放手');
INSERT INTO `song_lrcs`
VALUES ('It\'s You', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[00:09.31]Love rush\r\n[00:11.33]Drunk hush\r\n[00:13.31]Obsession\r\n[00:14.31]Pure\r\n[00:15.14]I know it\'s you\r\n[00:17.31]皇室里沉睡的带刺玫瑰\r\n[00:22.36]流浪着曾盛开在哪\r\n[00:25.29]街边的紫蔷薇野蛮生长\r\n[00:29.93]含苞待放\r\n[00:31.55]Ah ah\r\n[00:32.51]I\'m like oh\r\n[00:36.60]We don\'t waste time\r\n[00:38.31]We don\'t waste time\r\n[00:40.04]让我们相拥 oh\r\n[00:44.25]We don\'t waste time\r\n[00:46.17]We don\'t waste time\r\n[00:48.45]我想在你耳边 对你唱\r\n[00:52.62]When you say my name in the darkness\r\n[00:56.47]住进了深海 你被永远珍藏\r\n[01:03.80]It\'s feeling like you\r\n[01:11.13]It\'s got me feeling like oh\r\n[01:20.21]No no no no\r\n[01:20.81]Every time I see your face\r\n[01:22.62]Every time I catch your eyes\r\n[01:24.78]Every time I feel your body\r\n[01:26.30]It rushes right over me I know\r\n[01:32.21]It\'s you\r\n[01:34.23]It\'s you\r\n[01:36.60]深夜里仰望的孤独月光\r\n[01:41.47]凄美着黑色的眼珠\r\n[01:44.60]迷路的小女孩卷紧身躯\r\n[01:49.35]请别害怕\r\n[01:50.87]Ah ah\r\n[01:51.68]I\'m like oh\r\n[01:55.68]We don\'t waste time\r\n[01:57.60]We don\'t waste time\r\n[01:59.14]让我们解脱 oh\r\n[02:03.56]We don\'t waste time\r\n[02:05.58]We don\'t waste time\r\n[02:07.76]我想在你耳边 对你唱\r\n[02:11.77]When you say my name in the darkness\r\n[02:15.66]你不需要承受这所有的一切\r\n[02:23.13]It\'s feeling like you\r\n[02:30.46]It\'s got me feeling like oh\r\n[02:39.52]No no no no\r\n[02:40.18]Every time I see your face\r\n[02:42.19]Every time I catch your eyes\r\n[02:44.05]Every time I feel your body\r\n[02:45.71]It rushes right over me I know\r\n[02:51.49]It\'s you\r\n[02:53.40]It\'s you\r\n[02:56.21]Ayayiyayiya…\r\n[03:23.31]We don\'t waste time\r\n[03:24.98]We don\'t waste time\r\n[03:35.94]Love rush\r\n[03:37.71]Drunk hush\r\n[03:39.53]Obsession\r\n[03:40.49]Pure\r\n[03:41.37]I know it\'s you');
INSERT INTO `song_lrcs`
VALUES ('nobody cares', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:立落]\r\n[00:02.90]I\'ve been sleeping all day with you.\r\n[00:07.77]You said  you care, like you do care.\r\n[00:13.66]Tell me.\r\n[00:14.97]I\'ve been spending all day with you.\r\n[00:19.94]You said you care, like you do care.\r\n[00:26.07]Tell me.\r\n[00:27.07]Sitting in the sofa, played me like a little doll. So fun. Have fun.\r\n[00:38.00]Okay.\r\n[00:39.19]Then you turned the lights off; put me in the background.\r\n[00:44.27]Alright, fine.\r\n[00:47.01]I\'m so fine.\r\n[00:49.40]Okay.\r\n[00:50.49]Don\'t you say no to someone\r\n[00:54.18]I don\'t wanna let you \r\n[00:56.34]Don\'t you stay known in no time\r\n[01:00.12]I don\'t wanna \r\n[01:02.40]Don\'t you say no to someone\r\n[01:06.12]I don\'t wanna let you \r\n[01:08.33]Don\'t you stay known in no time\r\n[01:12.11]I don\'t wanna \r\n[01:13.49]Don\'t you stay with me.\r\n[01:19.53]Don\'t you stay with me.\r\n[01:25.55]Don\'t you stay with me.\r\n[01:31.59]Don\'t you stay with me.\r\n[01:39.23]I\'ve been sleeping all day with you.\r\n[01:43.94]You said you care, like you do care.\r\n[01:51.03]I\'ve been spending all day with you.\r\n[01:56.00]You said you care, like you do care.\r\n[02:04.58]Don\'t you…');
INSERT INTO `song_lrcs`
VALUES ('Pull Up', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[00:01.37]Whoa oh oh oh oh\r\n[00:05.97]Whoa oh oh oh oh\r\n[00:10.67]Whoa oh oh oh oh\r\n[00:15.51]Whoa oh oh oh oh\r\n[00:18.39]Didn\'t think I\'d be here\r\n[00:20.67]Lost my way but can you help me\r\n[00:23.01]Wasn\'t always so clear down this road\r\n[00:25.99]Now I\'m trying to break free\r\n[00:27.83]I been losing my my mind\r\n[00:29.95]I been drifting down down down\r\n[00:32.48]I can\'t sleep at night\r\n[00:34.75]Been going around round round\r\n[00:37.64]Ay I was so lost confused\r\n[00:39.79]Didn\'t know why who how to pull through\r\n[00:42.60]In a world I thought I knew though was too hard\r\n[00:46.26]But I was a fool\r\n[00:47.53]I\'ve been going nowhere\r\n[00:49.70]Need to know if you care\r\n[00:52.16]Are you gonna save me\r\n[00:54.33]Cos I been acting crazy\r\n[00:56.66]Pull up\r\n[00:58.88]There\'s nothing to gain and I\'m feeling the pain\r\n[01:01.26]Pull up\r\n[01:03.69]Don\'t say anymore cos I\'m losing control\r\n[01:05.92]Pull up\r\n[01:08.19]Cos I\'m running low\r\n[01:10.71]Pull up\r\n[01:12.95]But I can\'t let go\r\n[01:15.18]Baby telling me all the time\r\n[01:17.60]Got me saying what\'s on my mind\r\n[01:19.78]I can\'t tell what\'s up or down\r\n[01:22.15]I don\'t know my left from right\r\n[01:24.33]Think I\'ll start it again\r\n[01:25.63]Again and again\r\n[01:26.82]Second chances don\'t come easily\r\n[01:28.84]I\'m not gonna pretend\r\n[01:31.21]Never thought that things would be the same\r\n[01:33.74]Will it ever end\r\n[01:36.59]Holding out for something new\r\n[01:38.46]I\'m making a stand and I\'m gonna make it through\r\n[01:43.97]I\'ve been going nowhere\r\n[01:46.25]Need to know if you care\r\n[01:48.52]Are you gonna save me\r\n[01:50.80]Cos I been acting crazy\r\n[01:53.07]Pull up\r\n[01:55.36]There\'s nothing to gain and I\'m feeling the pain\r\n[01:57.70]Pull up\r\n[02:00.12]Don\'t say anymore cos I\'m losing control\r\n[02:02.34]Pull up\r\n[02:04.67]Cos I\'m running low\r\n[02:07.15]Pull up\r\n[02:09.53]But I can\'t let go\r\n[02:12.30]Need a little honesty\r\n[02:14.37]Tell me what you need from me\r\n[02:16.75]Show me that you\'re ready\r\n[02:18.72]Cos I\'m ready too\r\n[02:21.44]I don\'t care I\'m in control\r\n[02:23.86]Gotta let go oh ohhhh\r\n[02:26.24]Pull up\r\n[02:28.32]There\'s nothing to gain and I\'m feeling the pain\r\n[02:30.89]Pull up\r\n[02:33.21]Don\'t say anymore cos I\'m losing control\r\n[02:35.34]Pull up\r\n[02:37.32]Cos I\'m running low\r\n[02:40.07]Pull up\r\n[02:42.51]But I can\'t let go\r\n[02:44.78]Pull up\r\n[02:47.12]There\'s nothing to gain and I\'m feeling the pain\r\n[02:49.45]Pull up\r\n[02:51.83]Don\'t say anymore cos I\'m losing control\r\n[02:54.16]Pull up\r\n[02:56.48]Cos I\'m running low\r\n[02:58.81]Pull up\r\n[03:01.25]But I can\'t let go\r\n');
INSERT INTO `song_lrcs`
VALUES ('情人', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n\r\n[00:01.722]\r\n[00:12.049]眼色 是幻觉\r\n[00:14.505]泳池边你的身影勾成线\r\n[00:17.416]温热 蔓延\r\n[00:20.795]多少个午夜\r\n[00:22.428]肆无忌惮\r\n[00:23.645]醉梦酣欢\r\n[00:25.095]无意追逐\r\n[00:26.629]无法止步\r\n[00:28.627]热度 包围了我\r\n[00:32.910]All I wanna do is fool around\r\n[00:35.594]我的心在小鹿乱撞\r\n[00:38.276]从日落到清晨的月光\r\n[00:41.160]抱你到天亮\r\n[00:43.369]你轻轻一个吻\r\n[00:45.985]我疯狂体会\r\n[00:48.769]气氛开始升温\r\n[00:51.552]危险又迷人\r\n[00:53.813]I really wanna dance tonight\r\n[00:57.463]Feel a little bit dangerous\r\n[01:00.196]少了些安全感\r\n[01:03.329]做我的情人\r\n[01:04.896]I know you want it\r\n[01:12.757]掉落 人间 你像丘比特赐予我的首选\r\n[01:18.747]靠在 枕边\r\n[01:22.983]ah 光绕过你天使般的脸\r\n[01:25.750]ah 这感觉实在太危险\r\n[01:28.466]能否再对我温柔一点点\r\n[01:31.231]不忍心再带你去冒险\r\n[01:34.287]All I wanna do is fool around\r\n[01:36.984]我的心在小鹿乱撞\r\n[01:39.802]从日落到清晨的月光\r\n[01:42.585]抱你到天亮\r\n[01:44.871]你轻轻一个吻\r\n[01:47.419]我疯狂体会\r\n[01:50.070]气氛开始升温\r\n[01:52.818]危险又迷人\r\n[01:55.319]I really wanna dance tonight\r\n[01:58.902]Feel a little bit dangerous\r\n[02:01.669]少了些安全感\r\n[02:04.452]做我的情人\r\n[02:06.115]I know you want it\r\n[02:19.517]怪这感觉 狂热\r\n[02:21.200]灯光 晃了\r\n[02:22.383]音乐 放着\r\n[02:23.414]感受体温 上升\r\n[02:25.230]妆 花了\r\n[02:26.514]你 晃着\r\n[02:27.863]I know u really wanna\r\n[02:29.263]You know u really wanna\r\n[02:32.196]你轻轻一个吻\r\n[02:34.764]我疯狂体会\r\n[02:37.498]气氛开始升温\r\n[02:40.313]危险又迷人\r\n[02:42.513]I know you wanna dance tonight\r\n[02:46.296]Feel a little bit dangerous\r\n[02:49.180]少了些安全感\r\n[02:51.964]做我的情人\r\n[02:53.664]I know you want it\r\n[03:06.236]Be my lover');
INSERT INTO `song_lrcs`
VALUES ('感受她', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:立落]\r\n[00:13.80]感受她\r\n[00:17.03]透过窗\r\n[00:20.50]侵蚀这痛快\r\n[00:27.17]昨夜的酒\r\n[00:30.58]还让我晃\r\n[00:34.13]没理由重来\r\n[00:40.81]一种祸 一种惑  一种美妙的错\r\n[00:43.94]像一场火  点燃我 又烧光了我\r\n[00:47.38]无处躲 无处落 无所谓地难过\r\n[00:54.65]她的祸 她的惑 她是美妙的错\r\n[00:57.85]她的呼吸穿过鼓膜将我抚摸\r\n[01:01.17]贪婪地 吞噬了 交换过的轮廓\r\n[01:07.97]旧的心脏填满新的痒\r\n[01:11.24]限量发放一些奖赏\r\n[01:14.40]她逃离现场后\r\n[01:16.47]循环播放\r\n[01:17.96]一切情绪都很小心\r\n[01:21.36]旧的心脏爬满新的伤\r\n[01:24.85]解渴的药先干为敬\r\n[01:28.24]我配合退场 \r\n[01:29.73]表现得淡然\r\n[01:31.86]一切痕迹本该小心\r\n[02:03.81]房间空空落落 又影影绰绰\r\n[02:10.23]脚步昏昏沉沉 又心心念念\r\n[02:17.23]画面虚虚实实 又历历昭昭\r\n[02:23.89]就像是缠缠绕绕 又明明了了\r\n[02:30.33]旧的心脏填满新的痒\r\n[02:33.47]限量发放一些奖赏\r\n[02:36.87]她逃离现场后\r\n[02:38.98]循环播放\r\n[02:40.36]一切情绪都很小心\r\n[02:43.81]旧的心脏爬满新的伤\r\n[02:47.29]解渴的药先干为敬\r\n[02:50.53]我配合退场\r\n[02:51.86]表现得淡然\r\n[02:54.05]一切痕迹本该小心\r\n[02:58.55]感受她…\r\n[03:21.87]然后\r\n[03:23.23]戒了她');
INSERT INTO `song_lrcs`
VALUES ('标签', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n\r\n[00:09.54]太多的标签浮现 不知疲惫\r\n[00:11.92]挡了谁的路 换上新的梯队\r\n[00:14.24]目光放远些 把握每次机会\r\n[00:16.60]撕掉了标签作品做得金贵\r\n[00:19.33]Title\r\n[00:20.13]给我太多框架\r\n[00:21.65]Title\r\n[00:22.54]刻板印象放下\r\n[00:24.01]Title\r\n[00:24.69]一本正经说瞎话\r\n[00:26.37]Title\r\n[00:27.21]太多太多\r\n[00:28.70]Try\r\n[00:29.57]厚积薄发累积经验\r\n[00:31.00]Try\r\n[00:31.86]高级玩家从不检点\r\n[00:33.44]Try\r\n[00:34.22]小娄娄都绕在身边\r\n[00:35.72]牌打的太快分不清现实或老千\r\n[00:38.16]年轻气盛磕破了头吃过苦\r\n[00:40.50]花花世界看表里不一的主\r\n[00:42.81]分得清楚天时地利人和物\r\n[00:45.15]这一路\r\n[00:45.99]从不关心赢或输\r\n[00:47.46]Just Run\r\n[00:48.18]当妄想变成执念\r\n[00:49.86]Just Run\r\n[00:50.50]把油门踩到顶点\r\n[00:52.20]Just Run\r\n[00:52.80]又一次打破极限\r\n[00:54.26]台词简单却经典\r\n[00:56.65]循环这旋律一个礼拜\r\n[00:58.98]江湖的规矩置身事外\r\n[01:01.30]丢几个音符记在脑海\r\n[01:03.62]画我的拼图不被教坏\r\n[01:05.84]循环这旋律一个礼拜\r\n[01:08.21]江湖的规矩置身事外\r\n[01:10.47]丢几个音符记在脑海\r\n[01:12.77]You know I ain\'t worried about nothing about nothing\r\n\r\n[01:21.53]Why you always worried about something about something\r\n\r\n[01:30.46]贴好标签继续燥\r\n[01:32.70]乖乖收掉你的怯 I play it like a game\r\n\r\n[01:37.09]乖乖收掉你的怯 I play it like a game\r\n[01:41.09]\r\n[01:55.62]切换风格\r\n[01:56.54]保持诙谐\r\n[01:57.94]不同场合\r\n[01:58.83]不同嘴脸\r\n[02:00.10]“公众人物不是圣人就是玩笑”\r\n[02:02.35]“偶像流量连走路都不许崴脚”\r\n[02:04.48]欢迎光顾 人心危险\r\n[02:06.79]寻求帮助 套路堆叠\r\n[02:09.16]成功要付出的功课\r\n\r\n[02:11.78]励志要把规则打破\r\n[02:13.62]\r\n[02:14.73]怪我自命清高\r\n[02:16.78]懒得满足你的喜好\r\n\r\n[02:19.38]何必与狼共舞\r\n[02:21.49]天生要做高级动物\r\n[02:23.27]黑的酸的全都已经变得无所谓\r\n[02:25.29]如果嫉妒真的使你快乐\r\n[02:26.78]帮我“组个队”\r\n[02:27.65]Like plato 其他的没有\r\n[02:29.44]Try to play me like a play-doh（流过血和泪）\r\n[02:32.08]懒惰成性的人在往下坠\r\n[02:34.39]我把标签贴好突出重围\r\n[02:36.77]老人都说年轻别怕吃亏\r\n[02:39.13]经历了这过程就有回馈\r\n[02:41.52]Title\r\n[02:42.06]循环这旋律一个礼拜\r\n[02:44.39]江湖的规矩置身事外\r\n[02:46.77]丢几个音符记在脑海\r\n[02:49.09]画我的拼图不被教坏\r\n[02:51.33]循环这旋律一个礼拜\r\n[02:53.67]江湖的规矩置身事外\r\n[02:55.94]丢几个音符记在脑海\r\n[02:58.18]You know I ain\'t worried about nothing about nothing\r\n\r\n[03:07.11]Why you always worried about something about something\r\n[03:14.44]\r\n[03:15.91]贴好标签继续燥\r\n[03:18.15]乖乖收掉你的怯 I play it like a game\r\n[03:22.52]乖乖收掉你的怯 I play it like a game\r\n[03:26.49]I make it looks easy\r\n\r\n[03:30.55]Do u know what it takes\r\n\r\n[03:35.16]You don\'t wanna know');
INSERT INTO `song_lrcs`
VALUES ('欲', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:立落]\r\n[00:15.11]欲\r\n[00:19.39]沾满了血的印\r\n[00:21.80]人面兽心的你\r\n[00:24.32]欲情故纵的烂\r\n[00:27.16]游戏\r\n[00:29.39]那邪恶的信条\r\n[00:31.79]满足人心的迷\r\n[00:34.28]贪恋谁的双眸\r\n[00:37.08]Baby\r\n[00:55.15]欲\r\n[01:19.75]脑袋灌满酒精\r\n[01:21.83]排队的试验品\r\n[01:24.31]小心地追求名\r\n[01:27.20]和利\r\n[01:30.13]绑住灵魂\r\n[01:31.82]把爱变成禁忌\r\n[01:34.44]当欲望涌上来\r\n[01:37.20]窒息\r\n[01:44.62]欲');
INSERT INTO `song_lrcs`
VALUES ('爱与痛', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[00:18.49]他不停地追寻\r\n[00:21.12]没终点的终点\r\n[00:25.82]想多耀眼\r\n[00:29.64]当伤口愈合\r\n[00:32.28]再次挂上笑脸\r\n[00:36.65]迎接明天\r\n[00:41.15]一生以歌\r\n[00:43.55]去替代泪水\r\n[00:49.51]继续飞\r\n[00:52.08]无脚的鸟儿\r\n[00:55.03]继续飞\r\n[01:00.71]继续飞\r\n[01:12.57]在夜空中留下深情送给星星\r\n[01:23.34]看日出日落中追逐的这光景\r\n[01:35.01]如果你\r\n[01:37.51]太累\r\n[01:40.26]赠玫瑰\r\n[01:42.94]给你安慰\r\n[01:45.74]这万物\r\n[01:48.86]爱痛\r\n[01:51.14]为你唱\r\n[01:57.60]若你\r\n[02:00.16]太累\r\n[02:02.56]别迷惘\r\n[02:05.36]别后退\r\n[02:08.51]万物\r\n[02:11.40]爱痛\r\n[02:13.59]我为你唱\r\n[02:20.69]Love&Pain\r\n');
INSERT INTO `song_lrcs`
VALUES ('现象', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n\r\n[00:13.98]\r\n[00:14.70]It\'s time for show\r\n[00:15.74]你的喜好 我照单全收\r\n[00:18.16]被困的兽\r\n[00:19.18]各种规则 不必再纵容\r\n[00:21.78]I\'m so sorry about the struggle\r\n[00:23.49]又一次掀起躁动\r\n[00:26.04]我将冰冷的人心都变沸腾\r\n[00:29.46]太多的麻烦 来不及玩转\r\n[00:31.27]都烟消云散\r\n[00:32.94]We gon\' party we gon\' party 肆无忌惮\r\n[00:36.36]迷的梦魇 爬进双眼 流出黑的泪\r\n[00:39.82]喂养我的心脏\r\n[00:42.73]来\r\n[00:44.93]芸芸众生来\r\n[00:48.35]唤千秋万代\r\n[00:51.73]为何你不明白\r\n[00:53.92]\r\n[00:55.02]I\'m a phenomenon\r\n\r\n[01:02.13]It\'s a miracle\r\n\r\n[01:08.86]I\'m a phenomenon\r\n\r\n[01:22.63]It\'s only just begun\r\n[01:25.20]谁不在 仰望含血的丰碑\r\n[01:28.61]谁不在 变成提线的傀儡\r\n[01:32.05]Go get it go get it go with what you want\r\n[01:35.41]Go get it go get it 别再做哑巴\r\n[01:39.25]\r\n[01:39.78]Finally I\'m here\r\n[01:41.08]腐蚀着神经在流血\r\n[01:43.21]众生在感谢 笑着埋葬 笑着埋葬\r\n[01:46.18]把那些假意都咽进了肚子\r\n[01:47.96]笑我疯却都来模仿这路子\r\n[01:49.93]\r\n[01:50.90]我成为现象\r\n[01:53.05]来\r\n[01:55.22]芸芸众生来\r\n[01:58.53]唤千秋万代\r\n[02:01.88]为何你不明白\r\n[02:04.27]\r\n[02:05.28]I\'m a phenomenon\r\n\r\n[02:12.56]It\'s a miracle\r\n\r\n[02:19.08]I\'m a phenomenon\r\n[02:22.95]\r\n[02:33.73]万物现象\r\n[02:38.14]爱生恨终究归于遗忘\r\n[02:40.41]众生的象\r\n[02:44.88]善恶不朽成了灰\r\n[02:46.49]I\'m a phenomenon\r\n\r\n[02:53.53]It\'s a miracle\r\n\r\n[03:00.14]I\'m a phenomenon\r\n\r\n[03:27.72]It\'s only just begun');
INSERT INTO `song_lrcs`
VALUES ('蒙着眼', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:WannaOne_]\r\n\r\n[00:27.82]太多的fake friends\r\n[00:29.42]歧视链严重的圈子\r\n[00:30.94]大多数笑我偏执\r\n[00:32.61]谁懂得疯子的坚持\r\n[00:34.28]得得到得不到得不到\r\n[00:35.76]得不到得不到得得到\r\n[00:37.43]谣言在黑色里游走\r\n[00:38.67]这邪念多可笑多可笑\r\n[00:40.30]把冷眼都搅碎一口口吞下去\r\n[00:43.51]慢悠悠游过去这趟浑水\r\n[00:46.15]当你怕 又是为何\r\n[00:47.86]谁又空有一张躯壳\r\n[00:49.51]不愿做任人摆布的棋子\r\n[00:51.46]可谁比谁可怜 No pretending\r\n[00:55.61]早习惯过无人问津一旁的冷眼\r\n[01:01.88]只要剩最后一束光就不惜一切\r\n[01:08.34]不痛不痒不变\r\n[01:15.82]我相信真相会来 别蒙着眼\r\n[01:21.17]蒙着眼\r\n[01:24.57]蒙着眼\r\n[01:27.62]蒙着眼\r\n[01:43.94]消息来的越来越坏\r\n[01:45.56]哪有感同身受的道理\r\n[01:47.15]太多脏水泼满身\r\n[01:48.67]看着背叛转为动机\r\n[01:50.19]那就跟风耍一下\r\n[01:52.13]装聋还作哑\r\n[01:53.90]这水到底有多深\r\n[01:55.21]谁又管你真或假\r\n[01:56.83]擅读道德条文 疯了心\r\n[01:58.93]偷渡着败德\r\n[02:00.04]好让黑被技术型漂白神圣了邪恶\r\n[02:03.42]污名 别人付出的努力 有谁\r\n[02:05.42]能被资格给裁定\r\n[02:06.58]怪我不认被看不顺\r\n[02:08.26]杀红了眼 I\'m not pretending\r\n[02:12.53]早习惯过无人问津一旁的冷眼\r\n[02:18.81]只要剩最后一束光就不惜一切\r\n[02:25.67]不痛不痒不变\r\n[02:32.61]我相信真相会来 别蒙着眼\r\n[02:38.19]蒙着眼\r\n[02:41.37]蒙着眼\r\n[02:44.62]蒙着眼\r\n[02:49.32]是自觉还是自欺\r\n[02:52.57]争口气还是空气\r\n[02:55.61]不停压抑着呼吸\r\n[02:57.84]所有梦境终将破碎人们泪流成河\r\n[03:01.99]让你我都无惧存在\r\n[03:05.33]不回头只能负载\r\n[03:10.13]恶魔把我带入地狱\r\n[03:11.87]回不去只许混乱蒙着眼\r\n[03:16.76]蒙着眼\r\n[03:19.23]只许混乱蒙着眼\r\n[03:28.97]Made with pride\r\n[03:32.41]I can\'t find\r\n[03:35.69]恶魔把我带入地狱\r\n[03:37.48]回不去只许混乱蒙着眼');
INSERT INTO `song_lrcs`
VALUES ('迷', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:WannaOne_]\r\n\r\n[00:16.349]游荡在梦魇的海\r\n[00:19.469]奢求一滴泪\r\n[00:22.626]让我在长夜等待\r\n[00:26.126]听你诉苦水\r\n[00:29.138]始终没能逃离你的高墙快把我包围\r\n[00:36.209]像只猫\r\n[00:37.769]在迷宫中\r\n[00:38.962]无处可逃\r\n[00:41.431]暧昧的讯号吊人的胃\r\n[00:46.328]做人质或是偷心的贼\r\n[00:49.494]影子被吻后\r\n[00:51.971]就会枯萎\r\n[00:53.130]反复品尝这记号\r\n[00:57.018]这脆弱的爱\r\n[00:58.813]就快枯萎\r\n[01:00.041]我选择投降\r\n[01:01.991]任它抚慰\r\n[01:09.776]你留下迷任我去猜\r\n[01:18.714]恰到好处的欺瞒\r\n[01:24.683]无法逃离你的爱\r\n[01:32.433]缘分终是不对\r\n[01:36.261]我却执迷不悔\r\n[01:40.947]原谅你和你的双唇还有未尝过的吻\r\n[01:47.238]来时和消失一样快那梦境是否发生\r\n[01:54.176]一边回味\r\n[01:56.010]一边自卑\r\n[01:57.302]无力判断是和非\r\n[02:00.255]玫瑰凋落\r\n[02:03.418]结局了了\r\n[02:06.787]你留下迷任我去猜\r\n[02:16.234]恰到好处的欺瞒\r\n[02:22.749]无法逃离你的爱\r\n[02:30.031]缘分终是不对\r\n[02:34.143]何必执迷不悔\r\n[02:37.091]I\'m dying dying without you\r\n[02:42.973]Can\'t you see me turning blue\r\n[02:46.081]Dying dying without you.\r\n[02:52.863]Lonely lonely without you\r\n[02:55.852]Baby if you only knew\r\n[02:59.272]Trying Trying without you.\r\n[03:03.933]你的迷留给谁解开\r\n[03:24.086]你的迷留给谁解开\r\n[03:35.756]你的迷留给');
INSERT INTO `song_lrcs`
VALUES ('默片', '蔡徐坤',
        '[00:00.00]欢迎来访百视音乐网www.44h4.com\r\n[by:WannaOne_]\r\n\r\n[00:14.278] 那张古早唱片\r\n[00:17.119] 曾刻下你的容颜\r\n[00:19.582] 唱针划过弧线\r\n[00:23.250] 寄存了多少光年\r\n[00:26.935] 我无力分辨\r\n[00:28.599] 穿梭平行时间线\r\n[00:30.766] 为何不能再次遇见\r\n[00:32.975] 只是你的脸\r\n[00:34.998] 仿佛还清晰可见\r\n[00:37.885] 却模糊成遥远默片\r\n[00:39.753] 刻在百年之前未解的爱\r\n[00:47.270] 我仍守着花海等你归来\r\n[00:53.587] 夜幕降临燃烧我所有孤寂\r\n[00:58.403] 最荒唐的人最清醒\r\n[01:00.545] 当时穿越我们的心电感应\r\n[01:04.466] 也许早已失灵\r\n[01:07.033] 也许早已经失灵\r\n[01:26.020] 一寸密纹一寸情\r\n[01:27.550] 百年踪迹百年心\r\n[01:36.802] 灵魂坠入万丈深渊\r\n[01:41.412] 甘之若饴困在人间\r\n[01:44.358] 遗落了缱绻\r\n[01:48.279] 今生还在念\r\n[01:49.905] 错失的曾经\r\n[01:53.063] 时光交替你的泪水已成冰\r\n[01:56.032] 唯一的宿命\r\n[01:59.999] 蜿蜒世纪来回找寻\r\n[02:03.448] 最后背影\r\n[02:06.345] 刻在百年之前未解的爱\r\n[02:11.939] 我仍守着花海等你归来\r\n[02:18.108] 夜幕降临燃烧我所有孤寂\r\n[02:22.201] 最荒唐的人最清醒\r\n[02:25.313] 当时穿越我们的心电感应\r\n[02:29.584] 也许早已失灵\r\n[02:32.144] 也许早已经失灵\r\n[02:50.914] 一寸密纹一寸情\r\n[02:52.601] 百年踪迹百年心');

-- ----------------------------
-- Table structure for songs
-- ----------------------------
DROP TABLE IF EXISTS `songs`;
CREATE TABLE `songs`
(
    `song_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    PRIMARY KEY (`song_name`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of songs
-- ----------------------------
INSERT INTO `songs`
VALUES ('Hug me (抱我) - 蔡徐坤');
INSERT INTO `songs`
VALUES ('It\'s You - 蔡徐坤');
INSERT INTO `songs`
VALUES ('nobody cares - 蔡徐坤');
INSERT INTO `songs`
VALUES ('Pull Up - 蔡徐坤');
INSERT INTO `songs`
VALUES ('情人 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('感受她 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('标签 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('欲 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('爱与痛 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('现象 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('蒙着眼 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('迷 - 蔡徐坤');
INSERT INTO `songs`
VALUES ('默片 - 蔡徐坤');

-- ----------------------------
-- Table structure for tool_options
-- ----------------------------
DROP TABLE IF EXISTS `tool_options`;
CREATE TABLE `tool_options`
(
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`name`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_options
-- ----------------------------
INSERT INTO `tool_options`
VALUES ('个人名片', 'card');
INSERT INTO `tool_options`
VALUES ('位置', 'location');
INSERT INTO `tool_options`
VALUES ('我的收藏', 'collection');
INSERT INTO `tool_options`
VALUES ('拍摄', 'take-photo');
INSERT INTO `tool_options`
VALUES ('文件', 'file');
INSERT INTO `tool_options`
VALUES ('相册', 'photo');
INSERT INTO `tool_options`
VALUES ('红包', 'red-packet');
INSERT INTO `tool_options`
VALUES ('视频通话', 'video');
INSERT INTO `tool_options`
VALUES ('语音输入', 'voice2');
INSERT INTO `tool_options`
VALUES ('转账', 'transfer');
INSERT INTO `tool_options`
VALUES ('音乐', 'music');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `u_id`          bigint                                                        NOT NULL COMMENT '用户唯一id',
    `account_id`    varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '对外标识账号id',
    `phone`         varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NULL DEFAULT NULL COMMENT '手机号',
    `password`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `nick_name`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '昵称',
    `personal_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个性签名',
    `city_id`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '所在城市id',
    `expire`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号名修改的冷却期',
    `status`        int                                                           NULL DEFAULT NULL COMMENT '登录状态',
    `avatar_url`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
    PRIMARY KEY (`u_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users`
VALUES (1656926837948813312, 'account1', '19944162167', '$2a$12$53.4IRd0p/oyYJiJUInzGOveulEmaCan8mqf3dJrJU7dSiH1Wm/5i',
        'account1', '', '', '', 2, 'img/avatar16838817417827241960708995825.png');
INSERT INTO `users`
VALUES (1656945494364000256, 'account2', '', '$2a$12$NnyBX3xJlffOBKJ5yV/BD.ostmynnGTJKG.vlnwoXhkj7aRz.kcu.', 'account2',
        '', '', '', 2, 'img/avatar_168459101028516845910107667614276700704523.png');
INSERT INTO `users`
VALUES (1656945606888787968, 'account3', '', '$2a$12$DNeSuWIVoB5j3rhOM09pJeSJudWBOO09.6G4HUih9IGp7VOvzX8hC', 'account3',
        '', '', '', 2, 'img/avatar_16839779347701683977934459875491677069555.png');
INSERT INTO `users`
VALUES (1656945759372709888, 'account4', '', '$2a$12$0wvq.w6j2Gtm8ShoRRit4.bb8HHJn9dh29I67vUqli2PAz/oY9bY.', 'account4',
        '', '', '', 1, 'img/avatar16839651208429769365066165038.png');
INSERT INTO `users`
VALUES (1658763585775472640, 'account123', '', '$2a$12$tWAHKnAMeSKB3uEHRM/bd.p2YDK.WaPFVKl8Hm7FYFX1ujb6Cytfm',
        '阿达的三大', '', '', '', 1, 'img/avatar_16843150285221684315031767746191820431259.png');

SET FOREIGN_KEY_CHECKS = 1;
