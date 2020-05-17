package data

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq" // here
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5433", "user1", "abc", "chitchat")
	//Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	Db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
	return
}
