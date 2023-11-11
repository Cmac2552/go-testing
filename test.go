package main

import (
	"fmt"
	"log"
	"net/http"
)

func main2() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("worked")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
