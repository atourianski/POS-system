-- MySQL dump 10.13  Distrib 5.7.16, for Linux (x86_64)
--
-- Host: localhost    Database: banya
-- ------------------------------------------------------
-- Server version	5.7.16

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
-- Table structure for table `drinks`
--

DROP TABLE IF EXISTS `drinks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `drinks` (
  `name` varchar(20) DEFAULT NULL,
  `price` int(2) DEFAULT NULL,
  `notes` varchar(100) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `drinks`
--

LOCK TABLES `drinks` WRITE;
/*!40000 ALTER TABLE `drinks` DISABLE KEYS */;
INSERT INTO `drinks` VALUES ('bb',5,'bottled beer',1),('draft',6,'Millst or Czechvar',1),('Moscow Martini',9,'Vodka, homemade pickle brine',1),('The Julius',9,'Vodka, pickle juice, clamato,worcestershire, lemon juice, celery salt',1),('Red Orachard',9,'Vodka, cherry jam, amaretto, lemon juice',1),('Honey Bison',9,'Zubrowka, honey, lemon juice',1),('Blushing Anastasia',11,'Vodka, lemon juice, peach schnapps, cognac, cranberry juice, honey',1),('Bloody Anastasia',11,'Vodka, fresh beetroot juice, sweet syrup, lemon juice',1),('kvas',2,'fermented bread drink',1);
/*!40000 ALTER TABLE `drinks` ENABLE KEYS */;
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
  `notes` varchar(50) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foodstuffs`
--

LOCK TABLES `foodstuffs` WRITE;
/*!40000 ALTER TABLE `foodstuffs` DISABLE KEYS */;
INSERT INTO `foodstuffs` VALUES ('belyash',5,'fried meat donut',NULL),('cheb',5,'cheburek, georgian deep fried savoury pastry',1),('borcht',6,'russian soup',1),('salyan',6,'russian soup',NULL),('fish_salyan',6,'russian soup',NULL),('akrosh',6,'russian cold soup',NULL),('harcho',6,'georgian soup',NULL),('pea',6,'russian soup',NULL),('greek',6,'greek salad',1),('beet',6,'russian beet salad',NULL),('olivier',6,'russian salad',NULL),('cake',8,'dessert',1),('pastilla',8,'dessert',NULL),('pel',8,'pelmeni',NULL),('vareniki',8,'potato or cherry perogies',NULL),('salo',8,'russian cured pork',1),('shuba',8,'russian herring salad',1),('fish',12,'russian susdal baked fish',NULL),('strog',12,'pork strogonoff',NULL),('kebob',12,'lyolyakibabi, georgian lamb kebobs',NULL),('veg_plate',12,'vegetarian plate',NULL),('plov',12,'georgian lamb pilaf',NULL),('tabak',15,'georgian, chicken tabaka',NULL),('beef',15,'beef bourginion',1),('kotlet',15,'russian chicken cutlets',NULL),('saus',15,'sausages',NULL);
/*!40000 ALTER TABLE `foodstuffs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoices`
--

DROP TABLE IF EXISTS `invoices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoices` (
  `invoice_id` int(5) NOT NULL AUTO_INCREMENT,
  `bracelet_id` int(2) DEFAULT NULL,
  `banya` float(6,2) DEFAULT NULL,
  `food` float(6,2) DEFAULT NULL,
  `drink` float(6,2) DEFAULT NULL,
  `misc` float(6,2) DEFAULT NULL,
  PRIMARY KEY (`invoice_id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoices`
--

LOCK TABLES `invoices` WRITE;
/*!40000 ALTER TABLE `invoices` DISABLE KEYS */;
INSERT INTO `invoices` VALUES (27,51,35.00,0.00,0.00,3.00),(28,51,35.00,0.00,5.00,0.00),(29,44,35.00,16.00,0.00,0.00),(30,45,35.00,8.00,55.00,3.00),(31,1,35.00,16.00,15.00,9.00),(32,2,35.00,8.00,15.00,3.00),(33,3,35.00,0.00,0.00,0.00),(34,4,35.00,24.00,10.00,3.00),(35,4,35.00,0.00,0.00,0.00),(36,34,35.00,8.00,5.00,3.00),(37,37,35.00,8.00,55.00,18.00),(38,39,35.00,0.00,0.00,0.00),(39,21,35.00,8.00,5.00,3.00),(40,1,35.00,0.00,0.00,0.00),(41,2,35.00,0.00,0.00,0.00),(42,3,35.00,0.00,0.00,0.00),(43,3,35.00,0.00,0.00,0.00),(44,7,35.00,47.00,9.00,0.00),(45,8,35.00,0.00,0.00,0.00),(46,9,35.00,0.00,0.00,0.00);
/*!40000 ALTER TABLE `invoices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit`
--

DROP TABLE IF EXISTS `visit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `visit` (
  `date` date DEFAULT NULL,
  `bracelet_num` int(2) DEFAULT NULL,
  `entry_time` time DEFAULT NULL,
  `exit_time` time DEFAULT NULL,
  `invoice_id` int(10) DEFAULT NULL,
  `total` float(6,2) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit`
--

LOCK TABLES `visit` WRITE;
/*!40000 ALTER TABLE `visit` DISABLE KEYS */;
INSERT INTO `visit` VALUES ('2016-10-15',51,'23:37:16','23:41:10',27,NULL,0),('2016-10-15',51,'23:46:35','23:49:31',28,NULL,0),('2016-10-15',44,'00:01:42','00:50:44',29,NULL,0),('2016-10-15',45,'00:48:18','00:52:50',30,NULL,0),('2016-10-15',1,'01:02:53','01:45:22',31,NULL,0),('2016-10-15',2,'01:02:57','01:45:36',32,NULL,0),('2016-10-15',3,'01:03:04','01:45:46',33,NULL,0),('2016-10-15',4,'01:45:53','01:46:11',34,NULL,0),('2016-10-15',4,'01:46:19','01:47:13',35,NULL,0),('2016-10-15',34,'02:51:37','02:52:27',36,NULL,0),('2016-10-15',37,'02:54:11','02:56:00',37,NULL,0),('2016-10-15',39,'02:54:16','13:03:12',38,NULL,0),('2016-10-16',21,'12:55:58','14:26:56',39,NULL,0),('2016-10-16',1,'13:07:41',NULL,40,NULL,1),('2016-10-16',2,'13:07:45',NULL,41,NULL,1),('2016-10-16',3,'13:07:50','14:00:34',42,NULL,0),('2016-10-16',3,'14:36:48',NULL,43,NULL,1),('2016-10-16',7,'19:31:08',NULL,44,NULL,1),('2016-10-16',8,'19:49:12',NULL,45,NULL,1),('2016-10-16',9,'19:51:18','00:36:03',46,NULL,0);
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

-- Dump completed on 2016-10-17  1:19:08
