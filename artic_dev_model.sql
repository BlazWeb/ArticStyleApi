-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         8.0.23 - MySQL Community Server - GPL
-- SO del servidor:              Win64
-- HeidiSQL Versión:             11.2.0.6213
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Volcando estructura de base de datos para articstyle
CREATE DATABASE IF NOT EXISTS `articstyle` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `articstyle`;

-- Volcando estructura para tabla articstyle.followers
CREATE TABLE IF NOT EXISTS `followers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL COMMENT 'La id de la persona que siguen',
  `follower_id` int NOT NULL COMMENT 'La id de la persona que ha seguido a alguien',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Esta tabla, incluira todos los seguidores de la gente\r\n';

-- Volcando datos para la tabla articstyle.followers: ~0 rows (aproximadamente)
DELETE FROM `followers`;
/*!40000 ALTER TABLE `followers` DISABLE KEYS */;
/*!40000 ALTER TABLE `followers` ENABLE KEYS */;

-- Volcando estructura para tabla articstyle.post
CREATE TABLE IF NOT EXISTS `post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `profile` int NOT NULL COMMENT 'Perfil donde se encuentra el post',
  `author` int NOT NULL COMMENT 'El creador del post',
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Contenido del post',
  `likes` int NOT NULL DEFAULT '0' COMMENT 'La cantidad de likes del post',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Esta tabla almacena todos los post de los perfiles de los usuarios';

-- Volcando datos para la tabla articstyle.post: ~0 rows (aproximadamente)
DELETE FROM `post`;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
/*!40000 ALTER TABLE `post` ENABLE KEYS */;

-- Volcando estructura para tabla articstyle.style_publish
CREATE TABLE IF NOT EXISTS `style_publish` (
  `id` int NOT NULL,
  `author` int NOT NULL COMMENT 'El creador del estilo',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Titulo del estilo',
  `likes` int DEFAULT NULL COMMENT 'El numero de likes del estilo',
  `download` int DEFAULT NULL COMMENT 'Las veces que han descargado el estilo',
  `registrado` date NOT NULL COMMENT 'La fecha en la que se registrado el estilo',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'La id establecida al repositorio de imagenes del estilo',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Esta tabla almacenara todos los estilos creados en ArticStyle';

-- Volcando datos para la tabla articstyle.style_publish: ~0 rows (aproximadamente)
DELETE FROM `style_publish`;
/*!40000 ALTER TABLE `style_publish` DISABLE KEYS */;
/*!40000 ALTER TABLE `style_publish` ENABLE KEYS */;

-- Volcando estructura para tabla articstyle.style_save
CREATE TABLE IF NOT EXISTS `style_save` (
  `id` int NOT NULL AUTO_INCREMENT,
  `style_id` int NOT NULL COMMENT 'Esta columna, representa la id del estilo guardado por el usuario',
  `user_id` int NOT NULL COMMENT 'Usuario que ha guradado el estilo',
  `love` bit(1) NOT NULL DEFAULT b'0' COMMENT 'Este valor, demuestra si este estilo guardado esta en favoritos',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Esta tabla, guardara todos los estilos guardados por un usuario';

-- Volcando datos para la tabla articstyle.style_save: ~0 rows (aproximadamente)
DELETE FROM `style_save`;
/*!40000 ALTER TABLE `style_save` DISABLE KEYS */;
/*!40000 ALTER TABLE `style_save` ENABLE KEYS */;

-- Volcando estructura para tabla articstyle.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Nombre de usuario',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Nombre real del usuario',
  `last_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Apellidos reales del usuario',
  `date_register` date DEFAULT NULL COMMENT 'Fecha en la que se registro el usuario',
  `birthday` date DEFAULT NULL COMMENT 'Fecha de nacimiento del usuario',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Contraseña del usuario',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Correo electronico del usuario',
  `rank` int DEFAULT NULL COMMENT 'Rango en el que se encuentra el usuario',
  `img` json DEFAULT NULL COMMENT 'Este JSON tendra los datos las imagenes de Avatar y las imagenes de banner',
  `last_ip` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Esta es la ultima ip del usuario conectado',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Tabla de usuarios de ArticStyle, esta almanezara todos los datos de los usuarios de ticho proyecto';

-- Volcando datos para la tabla articstyle.users: ~0 rows (aproximadamente)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- Volcando estructura para tabla articstyle.users_bans
CREATE TABLE IF NOT EXISTS `users_bans` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT 'Esta columna, incluira la id del usuario baneado',
  `last_ip` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Esta columna, incluira la ip del usuario baneado',
  `mac` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Incluira la mac del dispositivo baneado',
  `reason` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Incluira la razon del baneo',
  `time` date NOT NULL COMMENT 'La fecha que fue baneado',
  `duration` date NOT NULL COMMENT 'Duración del baneo',
  `status` bit(1) NOT NULL DEFAULT b'0' COMMENT 'Este indicara el estado del baneo, si esta cumplido sera 0 si esta en curso sera 1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Esta tabla incluira los baneos de un usuario';

-- Volcando datos para la tabla articstyle.users_bans: ~0 rows (aproximadamente)
DELETE FROM `users_bans`;
/*!40000 ALTER TABLE `users_bans` DISABLE KEYS */;
/*!40000 ALTER TABLE `users_bans` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
