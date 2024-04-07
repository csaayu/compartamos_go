# Go

Crud rest con mvc:

1.- Crear la base de datos: "compartamos" o ejecutar el siguiente script:




--
-- Base de datos: `compartamos`
--
CREATE DATABASE compartamos;

use compartamos;
-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `dni` bigint(20) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_spanish2_ci DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8_spanish2_ci DEFAULT NULL,
  `fecha_nacimiento` datetime(3) DEFAULT NULL,
  `sexo` varchar(10) COLLATE utf8_spanish2_ci DEFAULT NULL,
  `ciudad` varchar(255) COLLATE utf8_spanish2_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` (`id`, `dni`, `name`, `last_name`, `fecha_nacimiento`, `sexo`, `ciudad`) VALUES
(15, 76467728, 'Cristhian', 'Saavedra', '1993-04-06 23:02:00.000', 'Masculino', 'Iquitos');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;


