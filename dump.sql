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
INSERT INTO `foodstuffs` VALUES ('belyash',5,'fried meat donut'),('cheb',5,'cheburek, georgian deep fried savoury pastry'),('borcht',6,'russian soup'),('salyan',6,'russian soup'),('fish_salyan',6,'russian soup'),('akrosh',6,'russian cold soup'),('harcho',6,'georgian soup'),('pea',6,'russian soup'),('greek',6,'greek salad'),('beet',6,'russian beet salad'),('olivier',6,'russian salad'),('cake',8,'dessert'),('pastilla',8,'dessert'),('pel',8,'pelmeni'),('vareniki',8,'potato or cherry perogies'),('salo',8,'russian cured pork'),('shuba',8,'russian herring salad'),('fish',12,'russian susdal baked fish'),('strog',12,'pork strogonoff'),('kebob',12,'lyolyakibabi, georgian lamb kebobs'),('veg_plate',12,'vegetarian plate'),('plov',12,'georgian lamb pilaf'),('tabak',15,'georgian, chicken tabaka'),('beef',15,'beef bourginion'),('kotlet',15,'russian chicken cutlets'),('saus',15,'sausages');
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

-- Dump completed on 2016-06-28 23:43:50
