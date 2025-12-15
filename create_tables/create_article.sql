CREATE TABLE `article` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `titile` varchar(200) NOT NULL DEFAULT '',
    `status` tinyint(1) NOT NULL DEFAULT '0' ,
    `image_url` varchar(2000) NOT NULL DEFAULT '',
    `cate_id` int(11) NOT NULL DEFAULT '',
    `is_top` tinyint(1) NOT NULL DEFAULT '0',
    `views` int(11) NOT NULL DEFAULT '0',
    `desc` varchar(2000) NOT NULL DEFAULT '',
    `content` longtext,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
);