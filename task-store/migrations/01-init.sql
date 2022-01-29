
-- +migrate Up

DROP TABLE IF EXISTS `stores`;
CREATE TABLE `stores` (
  `store_id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `store_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`store_id`)
);


INSERT INTO `stores` (`store_name`)
VALUES ('STORE_1'),
('STORE_2');

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `stock` int DEFAULT NULL,
  `store_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `products` (`product_name`, `stock`, `store_id`)
VALUES ('EVM-SAMPLE-PRODUCT-1', 100, 1);

DROP TABLE IF EXISTS `buyers`;
CREATE TABLE `buyers` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` CHAR(136) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `buyers` (`name`)
VALUES ('BAMBANG'),
('DARMO');

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `buyer_id` int NOT NULL,
    `quantity` VARCHAR(255) NOT NULL,
    `store_id` int NOT NULL,

    PRIMARY KEY (`id`)
);

INSERT INTO `orders` (`buyer_id`, `quantity`, `store_id`)
VALUES
(1, 5, 1),
(2, 10, 1);

DROP TABLE IF EXISTS `order_details`;
CREATE TABLE `order_details` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `order_id` int NOT NULL,
    `product_id` int NOT NULL,
    `qty` int NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `order_details` (`order_id`, `product_id`, `qty`)
VALUES
(1, 1, 20),
(2, 1, 5);