# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.36)
# Database: ferrytale
# Generation Time: 2017-07-22 15:12:11 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table docks
# ------------------------------------------------------------

DROP TABLE IF EXISTS `docks`;

CREATE TABLE `docks` (
  `id` varchar(191) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `longitude` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `latitude` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

LOCK TABLES `docks` WRITE;
/*!40000 ALTER TABLE `docks` DISABLE KEYS */;

INSERT INTO `docks` (`id`, `name`, `longitude`, `latitude`, `created_at`, `updated_at`)
VALUES
	('283c3e8d-372f-483d-9951-08a8d3d4b5b9','Centraal Station','','','2017-07-20 16:47:16',NULL),
	('3d141219-288a-4dbd-ac63-11ca62c60df9','Westerdoksdijk','','','0000-00-00 00:00:00',NULL),
	('6327dfb2-1b91-489a-bcd0-9f5f82950fe1','Distelweg','','','0000-00-00 00:00:00',NULL),
	('8eb3e3de-1e26-409a-b12a-08fea8c357dc','NDSM','','','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	('8fb6aada-a948-4d43-a0b0-84dfcb4027b2','Velsen Zuid','','','0000-00-00 00:00:00',NULL),
	('ad6e8694-7f8f-44eb-a41d-76bc7c2dd919','Assendelft','','','0000-00-00 00:00:00',NULL),
	('be013d60-9125-460f-b198-bf73e050a5cd','Velsen Noord','','','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	('e0f74832-49b4-4100-9319-669b692642be','Spaarndam','','','2017-07-20 16:47:18','0000-00-00 00:00:00'),
	('eaef9a9b-309b-408c-abbf-a22c0f312dcf','Buiksloterwegveer','','','0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `docks` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table ferries
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ferries`;

CREATE TABLE `ferries` (
  `id` varchar(191) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

LOCK TABLES `ferries` WRITE;
/*!40000 ALTER TABLE `ferries` DISABLE KEYS */;

INSERT INTO `ferries` (`id`, `name`, `created_at`, `updated_at`)
VALUES
	('5c6a48b6-96e1-437e-be61-c1466730d278','900','0000-00-00 00:00:00',NULL),
	('cbce4a2a-1210-4e26-b897-ce2b7df3eba4','901','0000-00-00 00:00:00',NULL);

/*!40000 ALTER TABLE `ferries` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table lines
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lines`;

CREATE TABLE `lines` (
  `id` varchar(191) NOT NULL DEFAULT '',
  `description` varchar(255) DEFAULT NULL,
  `ferry_id` varchar(191) NOT NULL DEFAULT '',
  `a_dock_id` varchar(191) DEFAULT NULL,
  `b_dock_id` varchar(191) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `lines` WRITE;
/*!40000 ALTER TABLE `lines` DISABLE KEYS */;

INSERT INTO `lines` (`id`, `description`, `ferry_id`, `a_dock_id`, `b_dock_id`, `created_at`, `updated_at`)
VALUES
	('fd9449c9-e1f6-4821-8ef8-08ea1dd96604','Westerdoksdijk to NDSM','5c6a48b6-96e1-437e-be61-c1466730d278','3d141219-288a-4dbd-ac63-11ca62c60df9','8eb3e3de-1e26-409a-b12a-08fea8c357dc','0000-00-00 00:00:00',NULL);

/*!40000 ALTER TABLE `lines` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table schedules
# ------------------------------------------------------------

DROP TABLE IF EXISTS `schedules`;

CREATE TABLE `schedules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `line_id` varchar(191) DEFAULT NULL,
  `day` varchar(100) DEFAULT NULL,
  `departure_at` time DEFAULT NULL,
  `arrival_at` time DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `schedules` WRITE;
/*!40000 ALTER TABLE `schedules` DISABLE KEYS */;

INSERT INTO `schedules` (`id`, `line_id`, `day`, `departure_at`, `arrival_at`)
VALUES
	(1,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','06:35:00','06:41:00'),
	(2,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','06:55:00','07:01:00'),
	(3,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','07:15:00','07:21:00'),
	(4,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','07:35:00','07:41:00'),
	(5,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','07:55:00','08:01:00'),
	(6,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','08:15:00','08:21:00'),
	(7,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','08:35:00','08:41:00'),
	(8,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','08:55:00','09:01:00'),
	(9,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','09:15:00','09:21:00'),
	(10,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','09:35:00','09:41:00'),
	(11,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','09:55:00','10:15:00'),
	(12,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','10:25:00','10:45:00'),
	(13,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','10:55:00','11:15:00'),
	(14,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','11:25:00','11:45:00'),
	(15,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','11:55:00','12:15:00'),
	(16,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','12:25:00','12:45:00'),
	(17,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','12:55:00','13:15:00'),
	(18,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','13:25:00','13:45:00'),
	(19,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','13:55:00','14:15:00'),
	(20,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','14:25:00','14:45:00'),
	(21,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','14:55:00','15:15:00'),
	(22,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','15:25:00','15:45:00'),
	(23,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','15:55:00','16:01:00'),
	(24,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','16:15:00','16:21:00'),
	(25,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','16:35:00','16:41:00'),
	(26,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','16:55:00','17:01:00'),
	(27,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','17:15:00','17:21:00'),
	(28,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','17:35:00','17:41:00'),
	(29,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','17:55:00','18:01:00'),
	(30,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','18:15:00','18:21:00'),
	(31,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','18:35:00','18:41:00'),
	(32,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','18:55:00','19:01:00'),
	(33,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','19:15:00','19:21:00'),
	(34,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','19:35:00','19:41:00'),
	(35,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','19:55:00','20:01:00'),
	(36,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','20:15:00','20:21:00'),
	(37,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','20:35:00','20:41:00'),
	(38,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','20:55:00','21:01:00'),
	(39,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','21:15:00','21:21:00'),
	(40,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','21:35:00','21:41:00'),
	(41,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','21:55:00','22:01:00'),
	(42,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','22:15:00','22:21:00'),
	(43,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','22:35:00','22:41:00'),
	(44,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','22:55:00','23:01:00'),
	(45,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','23:15:00','23:21:00'),
	(46,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','23:35:00','23:41:00'),
	(47,'fd9449c9-e1f6-4821-8ef8-08ea1dd96604','weekday','23:55:00','00:01:00');

/*!40000 ALTER TABLE `schedules` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
