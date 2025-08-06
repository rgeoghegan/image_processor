package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h2non/bimg"
)

func imageCompress(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accepts POST requests", 405)
		return
	}

	params := r.URL.Query()
	level, err := parseIntParam(params, "level")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	content, err := io.ReadAll(r.Body)
	if respondWithServerError(w, err) {
		return
	}

	image := bimg.NewImage(content)
	compressed, err := image.Process(bimg.Options{
		Compression: level,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Error compressing image: %s", err.Error()), 400)
		return
	}

	w.Write(compressed)
	return
}
