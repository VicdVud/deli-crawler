CREATE DATABASE IF NOT EXISTS `deli` DEFAULT CHARACTER SET utf8mb4;

USE `deli`;

CREATE TABLE IF NOT EXISTS `attendanceback` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '姓名',
  `department` varchar(15) NOT NULL DEFAULT '' COMMENT '部门',
  `year` int(4) NOT NULL DEFAULT '0' COMMENT '年份',
  `month` int(4) NOT NULL DEFAULT '0' COMMENT '月份',
  `day` int(4) NOT NULL DEFAULT '0' COMMENT '日期',
  `week` varchar(15) NOT NULL DEFAULT '' COMMENT '星期',
  `date_type` varchar(15) NOT NULL DEFAULT '' COMMENT '日期类型',
  `clock_in` varchar(15) NOT NULL DEFAULT '' COMMENT '签到时间',
  `clock_out` varchar(15) NOT NULL DEFAULT '' COMMENT '签退时间',
  `duration` float(5, 3) unsigned NOT NULL DEFAULT '0' COMMENT '工作时长',
  `late` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '迟到时间（分钟）',
  `leave_early` int(4) unsigned NOT NULL DEFAULT '0' COMMENT '早退时间（分钟）',
  `absent` int(2) unsigned NOT NULL DEFAULT '0' COMMENT '旷工时间（小时）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考勤信息表';
