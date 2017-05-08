/*
go get github.com/go-sql-driver/mysql
go build
*/

package sql

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "os"
  "encoding/json"
)

type Tour struct {
  id int
  title string
  description string
  price float32
}


var db *sql.DB

func init() {
  var err error
  db, err = sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWD")+"@/"+os.Getenv("MYSQL_DATABASE"))
  if err != nil {
    log.Fatal(err)
  }
}

func GetTours() (string,error) {
  rows, err := db.Query("SELECT * FROM tours")
  if err != nil {
    if err == sql.ErrNoRows {
      fmt.Println("No Records Found")
      return "", err
    }
    log.Fatal(err)
  }
  defer rows.Close()

  var Tours []*Tour

  for rows.Next() {
    tour := &Tour{}
    err := rows.Scan(&tour.id, &tour.title, &tour.description, &tour.price)
    if err != nil {
      log.Fatal(err)
    }
    Tours = append(Tours, tour)
  }

  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  for _, pr := range Tours {
    fmt.Printf("%s, %s, $%.2f\n", pr.title, pr.description, pr.price)
  }

  jsonTours, err := json.Marshal(Tours)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Println(jsonTours)

	return string(jsonTours), nil

}
