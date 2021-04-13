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

-- Volcando estructura para tabla articstyle.styles
CREATE TABLE IF NOT EXISTS `styles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `author` int(11) NOT NULL,
  `label` varchar(50) NOT NULL,
  `created` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4;

-- Volcando datos para la tabla articstyle.styles: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `styles` DISABLE KEYS */;
INSERT INTO `styles` (`id`, `title`, `author`, `label`, `created`) VALUES
	(1, 'Prueba', 2, '$2a$10$AdJfNmlaTxKaZDbu9HF1Iuraj5Saq7MoXWaRn2mi59a', '2021-04-13 09:08:34.1563952'),
	(2, 'Prueba', 2, '$2a$10$9Qk/ozNO5Hq2D79F74PSSufCNKfMuFiQLnMurtgIPCk', '2021-04-13 09:09:25.2094373'),
	(3, 'Prueba', 2, '$2a$10$Aqwd17VTcPY6U/ilmG8MMOpuwjoawVzcrgcAMLIKhUW', '2021-04-13 09:22:08.4961525'),
	(4, 'Prueba', 2, '$2a$10$4.Le0GvDxXnr808gnFCVyuTicOullrDbhy9RId01VFV', '2021-04-13 09:24:25.3238405'),
	(5, 'Prueba', 2, '$2a$10$bbW5RCOb.qCSWJhOKh.KL.L0lucphYNLS.abWM5DR/1', '2021-04-13 09:31:27.9671036'),
	(6, 'Prueba', 2, '$2a$10$Epj0cCCfijzmJhLYrLKX6.wWfop9mGoLqwbyayXhX.B', '2021-04-13 09:32:04.2941986'),
	(7, 'Prueba', 2, '$2a$10$4f20U6aPfNMU7TDYQSUbCeS6ZhuNj8npMIezkN.jjBU', '2021-04-13 09:32:07.1710944'),
	(8, 'Prueba', 2, '$2a$10$RPn8OHKABwxe6GZojxKwvOQfDtoUQYI16axf.DtWK7U', '2021-04-13 09:36:08.5364081'),
	(9, 'Prueba', 2, '$2a$10$zxMK8yKdezx563cqTWS29uiUTm3WdRoRq43DhTZKcxn', '2021-04-13 09:37:22.7281491'),
	(10, 'Prueba', 2, '$2a$10$MUGSyGzRuqOw.O1GwtEmMukox7XXxpCSx6iQeLHNfHy', '2021-04-13 09:38:24.1508619'),
	(11, 'Prueba', 2, '$2a$10$17eutrQSUw7DH2lTMZj.zuqZJSVvtEW.g6SAGTnRjJX', '2021-04-13 09:44:35.7289228'),
	(12, 'Prueba', 2, '$2a$10$Ejn/44OnUF3fbfv/JBkQROM.dOuXpQxrXCC.mLYkhOx', '2021-04-13 09:44:48.38253'),
	(13, 'Prueba', 2, '$2a$10$9kCTTUiFRrzHlDHcF3V3/.MbWJyS5vIdL6hxlukbi1s', '2021-04-13 10:00:28.3560676'),
	(14, 'Prueba', 2, '$2a$10$sjvse8vz2vZDj5RarGNxveUElqnCEEpW.tOBuMEIEi0', '2021-04-13 10:00:46.0384365'),
	(15, 'Prueba', 2, '$2a$10$O9YBhctlvDpNhuMcQuEsk.YZfIEFM.B586ioQRZM/j4', '2021-04-13 10:01:09.1602373'),
	(16, 'Prueba', 2, '$2a$10$I/zEtMBhAgEQmhu45luKneInwDZLMq7bH2e07Sdllnj', '2021-04-13 10:45:49.9821069'),
	(17, 'Prueba', 2, '$2a$10$NdrlOleLWBq0TQ4YoG5REuuVpeFiHhtre4InsXNCrSF', '2021-04-13 10:46:53.2330328'),
	(18, 'Prueba', 2, '$2a$10$WuTS3M8FBIPsoip0RgZCouhCYkjud4GUurbVdKfxrqF', '2021-04-13 10:47:07.6719479'),
	(19, 'Prueba', 2, '$2a$10$wvgto30VVAHrluLQqgLI4.NwvjKDs/vEVDyuzT2NgGk', '2021-04-13 11:21:17.5827693'),
	(20, 'Prueba', 2, '$2a$10$FEZSvuBdF8kJGt3EBxnzc..9uFl7BFx.9lNEBwLrChV', '2021-04-13 11:21:22.6971307'),
	(21, 'Prueba', 2, '$2a$10$zsYNzKUPY5uBzXKLgAE9C.4W8IN8fbIPdmvSCE2f1/H', '2021-04-13 11:22:24.8108512'),
	(22, 'Prueba', 2, '$2a$10$mpxZkAtfo36TUQps23TBKeqrGlpKcQTfWXZm2PL8OoF', '2021-04-13 11:22:28.2564468'),
	(23, 'Prueba', 2, '$2a$10$b5eAstnLC79SU4x8PO7iwu4rXxq.SKGVWriI3RHoTqA', '2021-04-13 11:22:48.0221214'),
	(24, 'Prueba', 2, '$2a$10$VTZfFF1hCDAS46cAUSCA7elxeuQIxCIkOCaPRLgzMEW', '2021-04-13 11:22:53.6501108'),
	(25, 'Prueba', 2, '$2a$10$LcTu5.6wktkzdQb/bmSV8e5FVrXyfZCFGEOzGU4/xv8', '2021-04-13 11:22:54.9166906'),
	(26, 'Prueba', 2, '$2a$10$KEpdI3C08SKdkFnmruaOy.5ojBvwYWKFvP1NEnix7ap', '2021-04-13 11:22:55.7426226'),
	(27, 'Prueba', 2, '$2a$10$L3a.Ygykb88shDybPfma2.ySUtIzBqCmH2XIYl6y/O4', '2021-04-13 11:22:56.763013'),
	(28, 'Prueba', 2, '$2a$10$g6iufpsRHEgBjjI5n.TgR.RDJ4m0vIWoUczTnX64J5i', '2021-04-13 11:22:58.1484838');
/*!40000 ALTER TABLE `styles` ENABLE KEYS */;

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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;

-- Volcando datos para la tabla articstyle.users: ~10 rows (aproximadamente)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password`, `email`, `name`, `last_name`, `date_registered`, `birthday`, `rank`, `img`) VALUES
	(1, 'Beniszkee', 'gilipollas', 'jose.mc.2004@gmail.com', 'Jose', 'Martinez', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}'),
	(2, 'Beniszke', '123456', 'jose.mc.2003@gmail.com', 'Jose', 'Martinez', '2021-04-11', '2021-04-11', 1, '{"avatar":"default", "banner": "default"}'),
	(3, 'Beniszkees', '123456', 'jose.mc.2002@gmail.com', 'Jose', 'Martinez', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}'),
	(4, 'Benisze', '123456', 'jose.mc.200@gmail.com', 'Jose', 'Martinez', '2021-04-11', '2021-04-11', 1, '{"avatar":"default", "banner": "default"}'),
	(5, 'Blazester', '$2a$08$MjxMZlZJmubu7UEFh.gD2uDzDjmzQHEUEo8rJtP3gxQynDKtRdKOe', 'blazester@gmail.com', 'Jose', 'Garcia', '2021-04-10', '2021-04-32', 2, '{"avatar":"default", "banner": "default"}'),
	(6, 'Blazestr', '$2a$08$XT.usZs8S9YEVXjiRHULLOCR6O62D2L2dTBDwYhaa/ax0xKEmtffG', 'blazestr@gmail.com', 'Jose', 'Garcia', '2021-04-11', '', 1, '{"images" : {"avatar": "default", "banner": "default"}}'),
	(7, 'Blazes', '$2a$08$uViDHJUYbjiAXLoJ6g39..Q03vC/mt2BWHXtNj7kQCEHkpe1LoI/K', 'blazest@gmail.com', 'Jose', 'Garcia', '2021-04-11', '', 1, '{"avatar":"default", "banner": "default"}'),
	(8, 's', '$2a$08$rnEu0tihXHQGLUa0Dx8hReSVx824ZObrZDsf3GMb6tscdRV9M17C2', 's', 's', 'sssss', '2021-04-11 15:18:35.0894088', '', 1, '{"avatar":"default", "banner": "default"}'),
	(11, 'ssSs', '$2a$08$5GoFKgF3uSWjHSyMfVjSjej2Po3UzaXTzoZmrXCXJwVKsSQs.FN.m', 'SSss', 's', 'sssss', '2021-04-11 15:54:23.0061824', '', 1, '{"avatar":"default", "banner": "default"}'),
	(15, 'ssssSss', '$2a$08$LOdlkWNDAXZ3ZYrqEEYGPO9om7bpjw5vg28ZkxkZjtlGc1dU3X5Ci', 'SsSssss', 's', 'sssss', '2021-04-13 10:10:22.4085445', '', 1, '{"avatar":"default", "banner": "default"}'),
	(17, 'sss!2sSss', '$2a$08$e4896bDtnJ0JKz5GyGWh3eWM0AaAeRINRBeNoM8yxv9UVlt610W1G', 'SsSss2sws', 's@', 'sssss', '2021-04-13 10:33:10.4598584', '', 1, '{"avatar":"default", "banner": "default"}'),
	(18, 'ssss1Sss', '$2a$08$HtjyQDzLcFLsVJiV4.aLiuMum3GfRx2eJGYSEgHVyfUVgWFqTOiRm', 'SsSss2sws@ee', 's d', 'sssss', '2021-04-13 10:52:26.3317315', '', 1, '{"avatar":"default", "banner": "default"}'),
	(19, 'ssss1Ssss', '$2a$08$Qsr4ECPszJEvTuXgoO17JOOVeSKbH/QKJERiAnOxirFXAuYY1aBTi', 'SsSss2sws@ee.com', 's d', 'sssss', '2021-04-13 10:53:24.0052913', '', 1, '{"avatar":"default", "banner": "default"}'),
	(20, 'ssss1Sssss', '$2a$08$mnkhlIdY7M7xfTtXm797nOiFjhdLglCDvvGcv/WV0s2HmZ78frVs6', 'SsSss2sws@eecom', 's d', 'sssss', '2021-04-13 10:53:30.3478562', '', 1, '{"avatar":"default", "banner": "default"}');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
