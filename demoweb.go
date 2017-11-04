package main

import (
	"fmt"
	"net/http"
	"os"
	// "database/sql"
	// _ "github.com/lib/pq"
)

// const (	PORT = "5000")

func main() {

	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// log.Fatal(err)
	// }

	http.HandleFunc("/", hello)
	fmt.Println("listening... on port: " + os.Getenv("PORT"))
	// fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	// err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world! making it better..")
}
