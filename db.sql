CREATE DATABASE biblioteca;
USE biblioteca;

CREATE TABLE libros (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  titulo VARCHAR(100) NOT NULL,
  descripcion VARCHAR(450) NOT NULL,
  autor VARCHAR(200) NOT NULL,
  editorial VARCHAR(200) NOT NULL,
  fecha_publicacion DATE NOT NULL
);
