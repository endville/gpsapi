CREATE DATABASE  IF NOT EXISTS `endville_gps` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `endville_gps`;
-- MySQL dump 10.13  Distrib 5.6.24, for osx10.8 (x86_64)
--
-- Host: 190.168.1.138    Database: endville_gps
-- ------------------------------------------------------
-- Server version	5.6.24-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL,
  `group_name` varchar(32) NOT NULL,
  `password` char(32) NOT NULL,
  `group_profile_id` int(11) unsigned NOT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_group_name` (`group_name`) USING BTREE,
  KEY `index_parent_id` (`parent_id`),
  KEY `index_profile_id` (`group_profile_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `group_profile`
--

DROP TABLE IF EXISTS `group_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_real_name` varchar(32) DEFAULT NULL,
  `contact_name` varchar(16) DEFAULT NULL,
  `contact_phone` varchar(16) DEFAULT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `group_roles`
--

DROP TABLE IF EXISTS `group_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(11) unsigned NOT NULL,
  `role_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `level` tinyint(4) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `log_by` varchar(32) NOT NULL,
  `log_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_datetime` (`log_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `terminal_sn` varchar(32) NOT NULL,
  `remote_addr` varchar(32) NOT NULL,
  `message_type` varchar(4) NOT NULL,
  `message_body` varchar(255) NOT NULL,
  `send_by` varchar(32) NOT NULL,
  `send_on` datetime NOT NULL,
  `feed_back` varchar(255) DEFAULT NULL,
  `feed_back_on` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `TerminalSN` (`terminal_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mileage`
--

DROP TABLE IF EXISTS `mileage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mileage` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `terminal_id` int(11) unsigned NOT NULL DEFAULT '0',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `group_id` int(11) unsigned NOT NULL DEFAULT '0',
  `mileage` int(11) unsigned NOT NULL DEFAULT '0',
  `record_on` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `index_record_on` (`record_on`),
  KEY `index_terminal_id` (`terminal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `right`
--

DROP TABLE IF EXISTS `right`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `right` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `right_name` varchar(32) NOT NULL,
  `category` varchar(32) NOT NULL DEFAULT '',
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL,
  `role_name` varchar(32) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role_rights`
--

DROP TABLE IF EXISTS `role_rights`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_rights` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL,
  `right_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `terminal`
--

DROP TABLE IF EXISTS `terminal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `terminal` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `terminal_sn` varchar(32) NOT NULL,
  `password` varchar(32) DEFAULT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `group_id` int(11) unsigned NOT NULL,
  `terminal_profile_id` int(11) unsigned NOT NULL,
  `terminal_carrier_id` int(11) unsigned NOT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  `online_on` datetime DEFAULT NULL,
  `offline_on` datetime DEFAULT NULL,
  `motion_state` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知,1:运动,2:静止',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_terminal_sn` (`terminal_sn`) USING BTREE,
  KEY `index_user_id` (`user_id`) USING BTREE,
  KEY `index_group_id` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `terminal_carrier`
--

DROP TABLE IF EXISTS `terminal_carrier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `terminal_carrier` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `license_plate_number` varchar(32) DEFAULT NULL,
  `vehicle_identification_number` char(7) DEFAULT NULL,
  `carrier_type` varchar(16) DEFAULT NULL,
  `brand` varchar(16) DEFAULT NULL,
  `color` varchar(16) DEFAULT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `terminal_profile`
--

DROP TABLE IF EXISTS `terminal_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `terminal_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `terminal_sn` varchar(32) NOT NULL,
  `tmsisdn` char(11) NOT NULL,
  `pmsisdn` char(11) NOT NULL,
  `imsi` char(15) NOT NULL,
  `imei` char(15) NOT NULL,
  `product_code` varchar(16) NOT NULL DEFAULT '',
  `is_activated` tinyint(4) NOT NULL,
  `mileage` int(11) unsigned NOT NULL DEFAULT '0',
  `activate_on` datetime DEFAULT NULL,
  `expire_on` datetime DEFAULT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_terminal_sn` (`terminal_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL,
  `password` char(32) NOT NULL,
  `user_profile_id` int(11) unsigned NOT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_nam` (`user_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_profile`
--

DROP TABLE IF EXISTS `user_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` char(18) DEFAULT NULL,
  `real_name` varchar(8) DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL,
  `address` varchar(64) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `phone` char(11) DEFAULT NULL,
  `sim_number` char(20) DEFAULT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `warning`
--

DROP TABLE IF EXISTS `warning`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warning` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `terminal_sn` varchar(32) NOT NULL,
  `terminal_id` int(10) unsigned NOT NULL DEFAULT '0',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `group_id` int(11) unsigned NOT NULL DEFAULT '0',
  `longitude` double(10,6) NOT NULL,
  `latitude` double(10,6) NOT NULL,
  `speed` double(6,3) unsigned NOT NULL COMMENT '公里/小时',
  `direction` double(6,3) unsigned NOT NULL COMMENT '方向',
  `status` tinyint(4) NOT NULL,
  `cell_id` varchar(32) NOT NULL,
  `voltage` double(6,3) unsigned DEFAULT NULL COMMENT '电压',
  `temperature` int(10) NOT NULL COMMENT '温度',
  `type` tinyint(4) NOT NULL,
  `state` tinyint(4) NOT NULL,
  `flag` tinyint(4) NOT NULL,
  `create_on` datetime NOT NULL,
  `modify_on` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `TerminalSN` (`terminal_sn`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-11-03 14:43:38
