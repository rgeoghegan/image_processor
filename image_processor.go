package main

import (
	"log"
	"net/http"
)

func run() error {
	http.HandleFunc("/convert", imageConverter)

	return http.ListenAndServe(":8080", nil)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
