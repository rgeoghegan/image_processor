package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/h2non/bimg"
)

func parseIntParam(urlValues url.Values, name string) (int, error) {
	values, ok := urlValues[name]
	if !ok {
		return 0, fmt.Errorf("Invalid or missing %s url parameter", name)
	}
	if len(values) == 0 {
		return 0, fmt.Errorf("Invalid or missing %s url parameter", name)
	}

	asInt, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, fmt.Errorf("Invalid or missing %s url parameter", name)
	}
	return asInt, nil
}

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
		http.Error(w, fmt.Sprintf("Error converting image: %s", err.Error()), 400)
		return
	}

	w.Write(resized)
	return
}
