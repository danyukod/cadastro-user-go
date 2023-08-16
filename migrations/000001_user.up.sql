CREATE TABLE IF NOT EXISTS `cadastro_user` (
    `id` varchar(36) NOT NULL PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `birth_date` datetime NOT NULL,
    `document` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `modified_at` datetime
    )ENGINE=InnoDB DEFAULT CHARSET=UTF8;