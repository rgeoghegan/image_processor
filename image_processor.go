package main

import (
	"log"
	"net/http"
)

func run() error {
	http.HandleFunc("/convert", imageConverter)
	http.HandleFunc("/resize", imageResize)
	http.HandleFunc("/compress", imageCompress)

	return http.ListenAndServe(":8080", nil)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
