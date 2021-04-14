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

-- Volcando estructura para tabla articstyle.followers
CREATE TABLE IF NOT EXISTS `followers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL COMMENT 'La id de la persona que siguen',
  `follower_id` int(11) NOT NULL COMMENT 'La id de la persona que ha seguido a alguien',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='Esta tabla, incluira todos los seguidores de la gente\r\n';

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla articstyle.posts
CREATE TABLE IF NOT EXISTS `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `profile` int(11) NOT NULL COMMENT 'Perfil donde se encuentra el post',
  `author` int(11) NOT NULL COMMENT 'El creador del post',
  `content` varchar(255) NOT NULL COMMENT 'Contenido del post',
  `likes` int(11) NOT NULL DEFAULT 0 COMMENT 'La cantidad de likes del post',
  `time` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='Esta tabla almacena todos los post de los perfiles de los usuarios';

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla articstyle.styles
CREATE TABLE IF NOT EXISTS `styles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author` int(11) NOT NULL COMMENT 'El creador del estilo',
  `title` varchar(255) NOT NULL COMMENT 'Titulo del estilo',
  `likes` int(11) DEFAULT NULL COMMENT 'El numero de likes del estilo',
  `download` int(11) DEFAULT NULL COMMENT 'Las veces que han descargado el estilo',
  `created` date NOT NULL COMMENT 'La fecha en la que se registrado el estilo',
  `label` varchar(255) NOT NULL COMMENT 'La id establecida al repositorio de imagenes del estilo',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='Esta tabla almacenara todos los estilos creados en ArticStyle';

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla articstyle.style_save
CREATE TABLE IF NOT EXISTS `style_save` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `style_id` int(11) NOT NULL COMMENT 'Esta columna, representa la id del estilo guardado por el usuario',
  `user_id` int(11) NOT NULL COMMENT 'Usuario que ha guradado el estilo',
  `love` tinyint(1) DEFAULT 0 COMMENT 'Este valor, demuestra si este estilo guardado esta en favoritos',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='Esta tabla, guardara todos los estilos guardados por un usuario';

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla articstyle.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT 'Nombre de usuario',
  `name` varchar(50) DEFAULT NULL COMMENT 'Nombre real del usuario',
  `last_name` varchar(50) DEFAULT NULL COMMENT 'Apellidos reales del usuario',
  `date_registered` date DEFAULT NULL COMMENT 'Fecha en la que se registro el usuario',
  `birthday` varchar(50) DEFAULT NULL COMMENT 'Fecha de nacimiento del usuario',
  `password` varchar(255) NOT NULL COMMENT 'Contraseña del usuario',
  `email` varchar(255) NOT NULL COMMENT 'Correo electronico del usuario',
  `rank` int(11) DEFAULT NULL COMMENT 'Rango en el que se encuentra el usuario',
  `img` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Este JSON tendra los datos las imagenes de Avatar y las imagenes de banner' CHECK (json_valid(`img`)),
  `last_ip` varchar(50) DEFAULT NULL COMMENT 'Esta es la ultima ip del usuario conectado',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='Tabla de usuarios de ArticStyle, esta almanezara todos los datos de los usuarios de ticho proyecto';

-- La exportación de datos fue deseleccionada.

-- Volcando estructura para tabla articstyle.users_bans
CREATE TABLE IF NOT EXISTS `users_bans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT 'Esta columna, incluira la id del usuario baneado',
  `last_ip` varchar(50) DEFAULT NULL COMMENT 'Esta columna, incluira la ip del usuario baneado',
  `mac` varchar(255) DEFAULT NULL COMMENT 'Incluira la mac del dispositivo baneado',
  `reason` longtext DEFAULT NULL COMMENT 'Incluira la razon del baneo',
  `time` date NOT NULL COMMENT 'La fecha que fue baneado',
  `duration` date NOT NULL COMMENT 'Duración del baneo',
  `status` bit(1) NOT NULL DEFAULT b'0' COMMENT 'Este indicara el estado del baneo, si esta cumplido sera 0 si esta en curso sera 1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Esta tabla incluira los baneos de un usuario';

-- La exportación de datos fue deseleccionada.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
