package main

import (
	"fmt"
	"net/http"
)

func imageConverter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
	return
}
