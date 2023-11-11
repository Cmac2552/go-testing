package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main4() {

	db, err := sql.Open("sqlite", "./DB_1.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}

func w() {
	fmt.Println("test")
}
