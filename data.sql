/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 5.7.24 : Database - cart
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`cart` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `cart`;

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `total` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `order` */

insert  into `order`(`id`,`total`) values 
(1,200000);

/*Table structure for table `order_detail` */

DROP TABLE IF EXISTS `order_detail`;

CREATE TABLE `order_detail` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `order_id` int(10) NOT NULL,
  `item_id` int(10) NOT NULL,
  `item_name` varchar(100) NOT NULL,
  `item_price` double NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_order` (`order_id`),
  CONSTRAINT `fk_order` FOREIGN KEY (`order_id`) REFERENCES `order` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `order_detail` */

insert  into `order_detail`(`id`,`order_id`,`item_id`,`item_name`,`item_price`) values 
(1,1,1,'SEPATU',50000),
(2,1,2,'KAOS',50000),
(3,1,3,'TOPI',100000);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
