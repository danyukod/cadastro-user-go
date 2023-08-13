CREATE TABLE IF NOT EXISTS `pix_key` (
    `id` varchar(36) NOT NULL PRIMARY KEY,
    `user_name` varchar(255) NOT NULL,
    `user_last_name` varchar(255) NOT NULL,
    `birthdate` datetime NOT NULL,
    `document` datetime NOT NULL,
    `created_at` datetime NOT NULL,
    `modified_at` datetime
    )ENGINE=InnoDB DEFAULT CHARSET=UTF8;