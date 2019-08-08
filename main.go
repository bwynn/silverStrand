package main

import (
	"log"
	"fmt"
	"net/http"
)
func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World!")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
