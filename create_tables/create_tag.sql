CREATE TABLE `tag` (
    `id` int(11) unsigned NOT NULL DEFAULT AUTO_INCREMENT,
    `ident` varchar(20) NOT NULL DEFAULT '',
    `title` varchar(100) NOT NULL DEFAULT '',
    `content` longtext,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)    
);