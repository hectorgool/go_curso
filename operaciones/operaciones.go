/*+
Una librería no tiene: package main
Una librería no tiene: func main()
para compilar una libreria:
go build
./pkg/linux_amd64/github.com/hectorgool/curso_go/operaciones.a
*/
package operaciones

import(
  "log"
)

//Los nombres de las funciones en las librerias debe empezar con una mayuscula, de otra forma no se pueden importar
func Suma(x, y int) int {
  return x + y
}

//Los nombres de las funciones en las librerias debe empezar con una mayuscula, de otra forma no se pueden importar
func Resta(x, y int) int {
  return x - y
}

//Los nombres de las funciones en las librerias debe empezar con una mayuscula, de otra forma no se pueden importar
func Multiplicacion(x, y int) int {
  return x * y
}

//Los nombres de las funciones en las librerias debe empezar con una mayuscula, de otra forma no se pueden importar
func Division(x, y int) (float32, error) {
  var err error
  if( y == 0 ){
    log.Fatal(err)
  }
  return x/y, nil
}
