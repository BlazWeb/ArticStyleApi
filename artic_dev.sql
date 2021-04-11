-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         10.4.14-MariaDB - mariadb.org binary distribution
-- SO del servidor:              Win64
-- HeidiSQL Versión:             11.1.0.6116
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Volcando estructura de base de datos para articstyle
CREATE DATABASE IF NOT EXISTS `articstyle` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `articstyle`;

-- Volcando estructura para tabla articstyle.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(75) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `name` varchar(50) NOT NULL,
  `last_name` varchar(75) NOT NULL,
  `date_registered` varchar(50) NOT NULL DEFAULT '',
  `birthday` varchar(50) NOT NULL DEFAULT '',
  `rank` int(11) DEFAULT 1,
  `img` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '{"avatar":"default", "banner": "default"}' CHECK (json_valid(`img`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- Volcando datos para la tabla articstyle.users: ~7 rows (aproximadamente)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password`, `email`, `name`, `last_name`, `date_registered`, `birthday`, `rank`, `img`) VALUES
	(1, 'Beniszkee', 'gilipollas', 'jose.mc.2004@gmail.com', 'Jose', 'Martinez', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}'),
	(2, 'Beniszke', '123456', 'jose.mc.2003@gmail.com', 'Jose', 'Martinez', '2021-04-11', '2021-04-11', 1, '{"avatar":"default", "banner": "default"}'),
	(3, 'Beniszkees', '123456', 'jose.mc.2002@gmail.com', 'Jose', 'Martinez', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}'),
	(4, 'Benisze', '123456', 'jose.mc.200@gmail.com', 'Jose', 'Martinez', '2021-04-11', '2021-04-11', 1, '{"avatar":"default", "banner": "default"}'),
	(5, 'Blazester', '$2a$08$MjxMZlZJmubu7UEFh.gD2uDzDjmzQHEUEo8rJtP3gxQynDKtRdKOe', 'blazester@gmail.com', 'Jose', 'Garcia', '2021-04-10', '2021-04-32', 2, '{"avatar":"default", "banner": "default"}'),
	(6, 'Blazestr', '$2a$08$XT.usZs8S9YEVXjiRHULLOCR6O62D2L2dTBDwYhaa/ax0xKEmtffG', 'blazestr@gmail.com', 'Jose', 'Garcia', '2021-04-11', '', 1, '{"images" : {"avatar": "default", "banner": "default"}}'),
	(7, 'Blazes', '$2a$08$uViDHJUYbjiAXLoJ6g39..Q03vC/mt2BWHXtNj7kQCEHkpe1LoI/K', 'blazest@gmail.com', 'Jose', 'Garcia', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
