package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h2non/bimg"
)

func imageConverter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accepts POST requests", 405)
		return
	}

	contentType, ok := r.Header["Content-Type"]
	if ok && len(contentType) > 0 {
		// If no header, let's go ahead and assume image/png, because
		// it makes testing with curl a bit easier
		if contentType[0] != "image/png" {
			http.Error(w, "Only accepts image/png", 400)
			return
		}
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

	asJpeg, err := image.Convert(bimg.JPEG)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting image: %s", err.Error()), 400)
		return
	}

	w.Write(asJpeg)
	return
}
