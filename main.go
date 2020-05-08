package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("app started")

	log.Fatal(http.ListenAndServe(":8080", NewRouter()))

}
