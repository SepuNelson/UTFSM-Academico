-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 14-11-2023 a las 10:36:57
-- Versión del servidor: 10.4.28-MariaDB
-- Versión de PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `prestigefoods`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `almuerzo`
--

CREATE TABLE `almuerzo` (
  `id` int(11) NOT NULL,
  `idReceta` int(11) NOT NULL,
  `idUsuario` int(11) NOT NULL,
  `favorito` tinyint(1) NOT NULL,
  `calificacionUsuario` int(1) NOT NULL,
  `reseñaUsuario` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `almuerzo`
--

INSERT INTO `almuerzo` (`id`, `idReceta`, `idUsuario`, `favorito`, `calificacionUsuario`, `reseñaUsuario`) VALUES
(1, 11, 2, 1, 3, 'Muy rico el Fetuccini!'),
(2, 24, 2, 0, 3, 'Le falta limon'),
(3, 17, 2, 1, 5, ''),
(6, 16, 4, 1, 5, 'Perfección'),
(7, 11, 1, 1, 4, 'Si'),
(8, 14, 1, 1, 4, ''),
(9, 22, 1, 1, 7, ''),
(10, 16, 1, 0, 1, ''),
(11, 20, 3, 0, 1, '');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `ingredientes`
--

CREATE TABLE `ingredientes` (
  `idIngrediente` int(11) NOT NULL,
  `ingrediente` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `ingredientes`
--

INSERT INTO `ingredientes` (`idIngrediente`, `ingrediente`) VALUES
(1, 'arroz'),
(2, 'pescado'),
(3, 'fideos'),
(4, 'pollo'),
(5, 'lechuga'),
(6, 'tomate'),
(7, 'cilantro'),
(8, 'crema de champiñones'),
(9, 'cebolla'),
(10, 'palta');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `ingrrec`
--

CREATE TABLE `ingrrec` (
  `id` int(11) NOT NULL,
  `idReceta` int(11) NOT NULL,
  `idIngrediente` int(11) NOT NULL,
  `cantidad` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `recetas`
--

CREATE TABLE `recetas` (
  `idReceta` int(11) NOT NULL,
  `nombre` varchar(25) NOT NULL,
  `foto` varchar(255) NOT NULL,
  `idPlato` int(1) NOT NULL,
  `tiempoPrep` int(3) NOT NULL,
  `instrucciones` varchar(255) NOT NULL,
  `apto para diabéticos` tinyint(1) NOT NULL,
  `sin gluten` tinyint(1) NOT NULL,
  `apto para intolerantes a la lactosa` tinyint(1) NOT NULL,
  `apto para veganos` tinyint(1) NOT NULL,
  `calificacion` decimal(1,1) NOT NULL,
  `reseña` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `recetas`
--

INSERT INTO `recetas` (`idReceta`, `nombre`, `foto`, `idPlato`, `tiempoPrep`, `instrucciones`, `apto para diabéticos`, `sin gluten`, `apto para intolerantes a la lactosa`, `apto para veganos`, `calificacion`, `reseña`) VALUES
(1, 'Arroz con Pollo', 'https://www.frutamare.com/wp-content/uploads/2021/04/arroz-con-pollo-al-curry.jpg', 3, 20, 'Saltee cebolla, ajo y pollo, agregue arroz, caldo y guisantes, cocinando hasta que el arroz esté listo.', 0, 0, 0, 0, 0.0, ''),
(11, 'Fetuccini', 'https://www.gourmet.cl/wp-content/uploads/2016/09/FETUCCINI-ALFREDO.jpg', 3, 15, 'Cocine la pasta hasta que esté al dente, mientras saltea champiñones con crema y mezcla con la pasta cocida, sirviéndolas con queso parmesano.', 1, 0, 0, 0, 0.0, ''),
(12, 'Lentejas', 'https://recetinas.com/wp-content/uploads/2020/01/lentejas-con-chorizo.jpg', 3, 30, 'Saltee cebolla y zanahoria, añada lentejas, tomates y caldo, cocinándolas hasta que estén tiernas.', 1, 0, 0, 1, 0.0, ''),
(14, 'Pasta Primavera', 'https://www.comedera.com/wp-content/uploads/2022/04/Pastas-primavera-shutterstock_1689065998.jpg', 1, 40, 'Cocine la pasta de acuerdo a las instrucciones del paquete. En una sartén, saltee vegetales variados con aceite de oliva y ajo. Mezcle con la pasta cocida y agregue queso parmesano.', 0, 0, 1, 1, 0.0, ''),
(15, 'Ensalada Griega', 'https://www.comedera.com/wp-content/uploads/2018/05/ensalada-griega.jpg', 2, 3, 'Combine pepinos, tomates, cebolla roja, aceitunas, queso feta y aderezo de aceite de oliva y jugo de limón. Mezcle bien y sirva frío.', 1, 1, 1, 1, 0.0, ''),
(16, 'Tarta de Manzana', 'https://images.hola.com/imagenes/cocina/recetas/20221128221837/tarta-manzana-mermelada/1-170-865/tarta-manzan-arguinano-t.jpg', 1, 50, 'Prepare la masa, corte las manzanas y hornee hasta que esté dorada. Sirva con helado o crema.', 0, 0, 0, 0, 0.0, ''),
(17, 'Sopa de Champiñones', 'https://mandolina.co/wp-content/uploads/2020/11/sopa-de-champinones-destacada.jpg', 2, 35, 'Saltee los champiñones y cebollas, agregue caldo y hierbas, cocine a fuego lento. Licúe y sirva con crema.', 1, 1, 1, 0, 0.0, ''),
(18, 'Pescado a la Parrilla', 'https://s1.eestatic.com/2022/01/27/actualidad/645695545_221414690_1706x960.jpg', 3, 30, 'Marine el pescado con hierbas y limón, luego a la parrilla hasta que esté cocido. Sirva con una guarnición de su elección.', 0, 0, 0, 1, 0.0, ''),
(19, 'Tacos de Camarones', 'https://assets.unileversolutions.com/recipes-v2/148465.jpg', 1, 25, 'Marine los camarones con especias, cocínelos y colóquelos en tortillas con aderezos y verduras. Sirva caliente.', 0, 0, 0, 0, 0.0, ''),
(20, 'Ensalada de Quinoa', 'https://www.goya.com/media/4003/quinoa-salad.jpg?quality=80', 2, 20, 'Cocine la quinoa, mezcle con vegetales picados, aderezo de limón y hierbas. Refrigere y sirva frío.', 1, 1, 1, 1, 0.0, ''),
(21, 'Batido de Frutas', 'https://s2.abcstatics.com/media/bienestar/2020/07/04/batidos-saludables-kdhH--1248x698@abc.jpeg', 3, 10, 'Mezcle frutas variadas con yogurt o leche, añada miel al gusto y licúe. Sirva frío.', 1, 1, 1, 1, 0.0, ''),
(22, 'Hamburguesa Vegetariana', 'https://thancguide.org/wp-content/uploads/2022/08/iStock-1310168994-scaled.jpg', 1, 30, 'Prepare una mezcla de vegetales y legumbres, forme hamburguesas y cocine a la parrilla. Sirva en pan integral.', 0, 1, 1, 1, 0.0, ''),
(23, 'Arroz con Leche', 'https://blog.renaware.com/wp-content/uploads/2023/03/arroz-con-leche-1.jpg', 2, 45, 'Cocine arroz en leche con canela y azúcar hasta que espese. Refrigere y sirva frío con canela espolvoreada.', 0, 0, 0, 0, 0.0, ''),
(24, 'Ceviche de Pescado', 'https://www.laylita.com/recetas/wp-content/uploads/2013/08/1-Cebiche-de-pescado.jpg', 3, 20, 'Marine el pescado en jugo de limón, agregue cebolla, tomate, cilantro y sirva frío.', 0, 0, 0, 0, 0.0, '');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tipoplato`
--

CREATE TABLE `tipoplato` (
  `idPlato` int(1) NOT NULL,
  `tipoPlato` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `tipoplato`
--

INSERT INTO `tipoplato` (`idPlato`, `tipoPlato`) VALUES
(1, 'Entrada'),
(2, 'Postre'),
(3, 'Plato de Fondo');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuario`
--

CREATE TABLE `usuario` (
  `idUsuario` int(11) NOT NULL,
  `nombre` varchar(25) NOT NULL,
  `mail` varchar(256) NOT NULL,
  `password` varchar(12) NOT NULL,
  `almuerzos` int(2) NOT NULL,
  `lastLog` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `usuario`
--

INSERT INTO `usuario` (`idUsuario`, `nombre`, `mail`, `password`, `almuerzos`, `lastLog`) VALUES
(1, 'Amadeus', 'a.mz@gm.de', 'requiem', 0, '2023-11-14 06:33:51'),
(2, 'Elon', 'kawai@usm.cl', 'twitter', 2, '2023-11-14 06:18:32'),
(3, 'gato', 's@g.com', 'asd', 0, '2023-11-14 06:36:00'),
(4, 'L', 'l.@gmail.com', 'kira', 4, '2023-11-10 01:13:39');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `almuerzo`
--
ALTER TABLE `almuerzo`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idReceta` (`idReceta`,`idUsuario`),
  ADD KEY `idUsuario` (`idUsuario`);

--
-- Indices de la tabla `ingredientes`
--
ALTER TABLE `ingredientes`
  ADD PRIMARY KEY (`idIngrediente`);

--
-- Indices de la tabla `ingrrec`
--
ALTER TABLE `ingrrec`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idReceta` (`idReceta`,`idIngrediente`),
  ADD KEY `idIngrediente` (`idIngrediente`);

--
-- Indices de la tabla `recetas`
--
ALTER TABLE `recetas`
  ADD PRIMARY KEY (`idReceta`),
  ADD KEY `idPlato` (`idPlato`);

--
-- Indices de la tabla `tipoplato`
--
ALTER TABLE `tipoplato`
  ADD PRIMARY KEY (`idPlato`);

--
-- Indices de la tabla `usuario`
--
ALTER TABLE `usuario`
  ADD PRIMARY KEY (`idUsuario`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `almuerzo`
--
ALTER TABLE `almuerzo`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT de la tabla `ingredientes`
--
ALTER TABLE `ingredientes`
  MODIFY `idIngrediente` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT de la tabla `ingrrec`
--
ALTER TABLE `ingrrec`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `recetas`
--
ALTER TABLE `recetas`
  MODIFY `idReceta` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT de la tabla `tipoplato`
--
ALTER TABLE `tipoplato`
  MODIFY `idPlato` int(1) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT de la tabla `usuario`
--
ALTER TABLE `usuario`
  MODIFY `idUsuario` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=194827666;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `almuerzo`
--
ALTER TABLE `almuerzo`
  ADD CONSTRAINT `almuerzo_ibfk_1` FOREIGN KEY (`idUsuario`) REFERENCES `usuario` (`idUsuario`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `almuerzo_ibfk_2` FOREIGN KEY (`idReceta`) REFERENCES `recetas` (`idReceta`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `ingrrec`
--
ALTER TABLE `ingrrec`
  ADD CONSTRAINT `ingrrec_ibfk_1` FOREIGN KEY (`idIngrediente`) REFERENCES `ingredientes` (`idIngrediente`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ingrrec_ibfk_2` FOREIGN KEY (`idIngrediente`) REFERENCES `ingredientes` (`idIngrediente`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ingrrec_ibfk_3` FOREIGN KEY (`idReceta`) REFERENCES `recetas` (`idReceta`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `recetas`
--
ALTER TABLE `recetas`
  ADD CONSTRAINT `recetas_ibfk_1` FOREIGN KEY (`idPlato`) REFERENCES `tipoplato` (`idPlato`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
