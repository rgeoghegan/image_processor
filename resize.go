package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h2non/bimg"
)

func imageResize(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accepts POST requests", 405)
		return
	}

	params := r.URL.Query()
	width, err := parseIntParam(params, "width")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	height, err := parseIntParam(params, "height")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	content, err := io.ReadAll(r.Body)
	if respondWithServerError(w, err) {
		return
	}

	image := bimg.NewImage(content)
	if image.Type() != "png" {
		http.Error(w, "Error converting image: Unsupported image format", 400)
		return
	}

	resized, err := image.Resize(width, height)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error resizing image: %s", err.Error()), 400)
		return
	}

	w.Write(resized)
	return
}
