package main

import (
  sql "github.com/hectorgool/curso_go/db/sql"
  "fmt"
)

func main(){
  if result, err := sql.GetTours(); err != nil {
          fmt.Printf("Error: %s\n", err)
  } else {
          fmt.Printf("MySQL result: %s\n", result)
  }
}
