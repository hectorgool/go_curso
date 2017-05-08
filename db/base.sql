DROP database curso_go;
CREATE database curso_go;
use curso_go;

DROP TABLES IF EXISTS Tours;
CREATE TABLE tours (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  description varchar(255),
  price decimal(6,2) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO tours (id, title, description, price)
VALUES ('1', 'titulo 1', 'descripcion 1', '33.33');

INSERT INTO tours (id, title, description, price)
VALUES ('2', 'titulo 2', 'descripcion 2', '444.44');
