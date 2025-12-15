CREATE TABLE `article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT 'article id'
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'tag id'
    PRIMARY KEY (`id`)
);