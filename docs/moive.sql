CREATE TABLE `sp_douban_movie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(130) DEFAULT '' COMMENT '标题',
  `subtitle` varchar(120) DEFAULT '' COMMENT '副标题',
  `other` varchar(120) DEFAULT '' COMMENT '其他',
  `desc` varchar(130) DEFAULT '' COMMENT '简述',
  `year` varchar(110) DEFAULT '' COMMENT '年份',
  `area` varchar(120) DEFAULT '' COMMENT '地区',
  `tag` varchar(120) DEFAULT '' COMMENT '标签',
  `star` varchar(110) DEFAULT '' COMMENT 'star',
  `comment` varchar(110) DEFAULT  '' COMMENT '评分',
  `quote` varchar(130) DEFAULT '' COMMENT '引用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='豆瓣电影Top250';
