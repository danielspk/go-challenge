CREATE TABLE IF NOT EXISTS `offices`
(
    `id`         INT          NOT NULL AUTO_INCREMENT,
    `address`    VARCHAR(160) NOT NULL,
    `latitude`   FLOAT(10, 6) NOT NULL,
    `longitude`  FLOAT(10, 6) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `point` (`latitude` ASC, `longitude` ASC)
) ENGINE = `InnoDB`
  DEFAULT charset = `utf8`;
