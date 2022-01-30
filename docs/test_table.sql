-- t_book
CREATE TABLE `t_book` (
    `id` tinyint(11) NOT NULL AUTO_INCREMENT,
    `isbn` varchar(30) NOT NULL DEFAULT '' COMMENT 'ISBN',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '书名',
    `price` int(11) NOT NULL DEFAULT '0' COMMENT '定价，单位：分',
    `author` varchar(20) NOT NULL DEFAULT '' COMMENT '作者',
    `edition` varchar(20) NOT NULL DEFAULT '' COMMENT '版次',
    `press` varchar(20) NOT NULL DEFAULT '' COMMENT '出版社',
    `address` varchar(30) NOT NULL DEFAULT '' COMMENT '社址',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', 
    PRIMARY KEY(`id`),
    KEY `idx_isbn`(`isbn`),
    KEY `idx_name`(`name`)
) ENGINE=InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '书';

-- test data
INSERT INTO `t_book` (isbn, name, price, author, edition, press, address) 
VALUES 
("978-7-5063-6869-1", "《天幕红尘》", 4800, "豆豆", "2013 年 6 月第 1 版", "作家出版社", "北京农展馆南里 10 号"),
("978-7-5063-6869-2", "《天幕红尘》", 4800, "豆豆", "2013 年 6 月第 1 版", "作家出版社", "北京农展馆南里 10 号"),
("978-7-5063-6869-3", "《天幕红尘》", 4800, "豆豆", "2013 年 6 月第 1 版", "作家出版社", "北京农展馆南里 10 号"),
("978-7-5063-6869-4", "《天幕红尘》", 4800, "豆豆", "2013 年 6 月第 1 版", "作家出版社", "北京农展馆南里 10 号"),
("978-7-5063-6869-5", "《天幕红尘》", 4800, "豆豆", "2013 年 6 月第 1 版", "作家出版社", "北京农展馆南里 10 号");
