-- MySQL dump 10.13  Distrib 5.7.13, for Linux (x86_64)
--
-- Host: localhost    Database: banya
-- ------------------------------------------------------
-- Server version	5.7.13

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
-- Table structure for table `food`
--

DROP TABLE IF EXISTS `food`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `food` (
  `belyash` int(2) NOT NULL DEFAULT '5',
  `cheb` int(2) DEFAULT NULL,
  `borcht` int(2) DEFAULT NULL,
  `salyan` int(2) DEFAULT NULL,
  `fish_salyan` int(2) DEFAULT NULL,
  `akrosh` int(2) DEFAULT NULL,
  `harcho` int(2) DEFAULT NULL,
  `pea` int(2) DEFAULT NULL,
  `greek` int(2) DEFAULT NULL,
  `beet` int(2) DEFAULT NULL,
  `olivier` int(2) DEFAULT NULL,
  `cake` int(2) DEFAULT NULL,
  `pel` int(2) DEFAULT NULL,
  `vareniki` int(2) DEFAULT NULL,
  `salo` int(2) DEFAULT NULL,
  `shuba` int(2) DEFAULT NULL,
  `pastilla` int(2) DEFAULT NULL,
  `fish` int(2) DEFAULT NULL,
  `strog` int(2) DEFAULT NULL,
  `kebab` int(2) DEFAULT NULL,
  `veg_plate` int(2) DEFAULT NULL,
  `plov` int(2) DEFAULT NULL,
  `tabak` int(2) DEFAULT NULL,
  `beef` int(2) DEFAULT NULL,
  `kotlet` int(2) DEFAULT NULL,
  `saus` int(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `food`
--

LOCK TABLES `food` WRITE;
/*!40000 ALTER TABLE `food` DISABLE KEYS */;
/*!40000 ALTER TABLE `food` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `foodstuffs`
--

DROP TABLE IF EXISTS `foodstuffs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `foodstuffs` (
  `name` varchar(20) DEFAULT NULL,
  `price` int(2) DEFAULT NULL,
  `notes` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foodstuffs`
--

LOCK TABLES `foodstuffs` WRITE;
/*!40000 ALTER TABLE `foodstuffs` DISABLE KEYS */;
INSERT INTO `foodstuffs` VALUES ('belyash',5,'fried meat donut'),('cheb',5,'cheburek, georgian deep fried savoury pastry'),('borcht',6,'russian soup'),('salyan',6,'russian soup');
/*!40000 ALTER TABLE `foodstuffs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_id`
--

DROP TABLE IF EXISTS `invoice_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_id` (
  `banya` float(6,2) DEFAULT NULL,
  `food` float(6,2) DEFAULT NULL,
  `drink` float(6,2) DEFAULT NULL,
  `massage` float(6,2) DEFAULT NULL,
  `misc` float(6,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_id`
--

LOCK TABLES `invoice_id` WRITE;
/*!40000 ALTER TABLE `invoice_id` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoice_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit`
--

DROP TABLE IF EXISTS `visit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `visit` (
  `date` date DEFAULT NULL,
  `unique_id` int(10) DEFAULT NULL,
  `bracelet_num` int(2) DEFAULT NULL,
  `entry_time` time DEFAULT NULL,
  `exit_time` time DEFAULT NULL,
  `invoice_id` int(10) DEFAULT NULL,
  `total` float(6,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit`
--

LOCK TABLES `visit` WRITE;
/*!40000 ALTER TABLE `visit` DISABLE KEYS */;
/*!40000 ALTER TABLE `visit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-06-28 20:23:26
