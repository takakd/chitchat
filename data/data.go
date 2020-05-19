package data

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq" // here
	"crypto/rand"
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

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7atch/goouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}