SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`
(
    `id`           int(11)      NOT NULL AUTO_INCREMENT,
    `title`        varchar(255) NULL     DEFAULT NULL,
    `desc`         varchar(255) NULL     DEFAULT NULL,
    `content`      varchar(255) NULL     DEFAULT NULL,
    `position`     int(11)      NULL     DEFAULT NULL,
    `head_line`    varchar(113) NULL     DEFAULT NULL,
    `keywords`     varchar(255) NULL     DEFAULT NULL,
    `image_thumb`  varchar(255) NULL     DEFAULT NULL,
    `image_s`      varchar(255) NULL     DEFAULT NULL,
    `image_m`      varchar(255) NULL     DEFAULT NULL,
    `image_l`      varchar(255) NULL     DEFAULT NULL,
    `created_at`   datetime(0)  NULL     DEFAULT NULL,
    `updated_at`   datetime(0)  NULL     DEFAULT NULL,
    `published_at` datetime(0)  NULL     DEFAULT NULL,
    `is_deleted`  tinyint(1)   NOT NULL DEFAULT 0,
        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = Compact;


DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT,
    `keyword`     varchar(255) NULL     DEFAULT NULL,
    `cat_type`    varchar(255) NULL     DEFAULT NULL,
    `is_deleted` tinyint(1)   NOT NULL DEFAULT 0,
    primary key (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = Compact;

DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop`
(
    `id`           int(11)       NOT NULL AUTO_INCREMENT,
    `shopname`     varchar(255)  NOT NULL,
    `name`         varchar(255)  NULL DEFAULT NULL,
    `desc`         varchar(500)  NULL     DEFAULT NULL,
    `prefecture_n` varchar(50)   NULL     DEFAULT NULL,
    `city_n`       varchar(50)   NULL     DEFAULT NULL,
    `zip_cd`       char(7),
    `tel`          varchar(14)   NULL     DEFAULT NULL,
    `email`        varchar(50)   NULL     DEFAULT NULL,
    `house_number` varchar(255)  NULL     DEFAULT NULL,
    `map_level`    char(1)       NULL     DEFAULT NULL,
    `ad_id`        int           NOT NULL DEFAULT 0,
    `point`        int           NOT NULL DEFAULT 0,
    `lat`          Decimal(9, 6) NULL     DEFAULT NULL,
    `lon`          Decimal(9, 6) NULL     DEFAULT NULL,
    `url`          varchar(255)  NULL     DEFAULT NULL,
    `created_at`   datetime(0)   NULL     DEFAULT NULL,
    `updated_at`   datetime(0)   NULL     DEFAULT NULL,
    `deleted_note` TEXT          NULL     DEFAULT NULL,
    `is_deleted`  tinyint(1)    NOT NULL DEFAULT 0,
    primary key (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = Compact;



SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `m_cat`;
CREATE TABLE `m_cat`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT,
    `name`        varchar(255) NULL     DEFAULT NULL,
    `code`        varchar(255) NULL     DEFAULT NULL,
    `cat_id`      varchar(255) NULL     DEFAULT NULL,
    `type`        tinyint(1)   NOT NULL DEFAULT 0,
    `created_at`  datetime(0)  NULL     DEFAULT NULL,
    `updated_at`  datetime(0)  NULL     DEFAULT NULL,
    `is_deleted` tinyint(1)   NOT NULL DEFAULT 0,
    primary key (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = Compact;
