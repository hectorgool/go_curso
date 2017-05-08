/*
go test

PASS
ok      github.com/hectorgool/curso_go/operaciones      0.002s

*/
package operaciones

import "testing"

func TestSuma(t *testing.T) {
  var result int
  result = Suma(15,10)
  if result!= 25 {
    t.Error("Expected 25, got ", result)
  }
}

func TestResta(t *testing.T) {
  var result int
  result = Resta(15,10)

  if result!= 5 {
    t.Error("Expected 5, got ", result)
  }
}
