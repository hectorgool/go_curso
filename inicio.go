package main

import (
  "fmt"
  "github.com/hectorgool/curso_go/operaciones"
)

func main() {
  var x, y int = 10, 5
  fmt.Println(operaciones.Suma(x, y))
  fmt.Println(operaciones.Resta(x, y))
}
