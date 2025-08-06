package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h2non/bimg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageResizeErrors(t *testing.T) {
	var handler http.HandlerFunc = imageResize
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint only responds to POSTs
	res, err := http.Get(ts.URL)
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 405)
	assertBodyEqual(t, res, "Only accepts POST requests\n")

	// Get example png
	examplePng := sloppilyGetFileContent("example.png")

	// Endpoint requires valid size infomation
	url := fmt.Sprintf("%s?width=banana&height=42", ts.URL)
	res, err = http.Post(url, "image/png", bytes.NewBuffer(examplePng))
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Invalid or missing width url parameter\n")

	// Endpoint requires a valid file
	body := strings.NewReader("I am obviously not a valid png")
	url = fmt.Sprintf("%s?width=100&height=100", ts.URL)
	res, err = http.Post(url, "image/png", body) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Error converting image: Unsupported image format\n")
}

func TestImageResizeSuccess(t *testing.T) {
	var handler http.HandlerFunc = imageResize
	ts := httptest.NewServer(handler)
	defer ts.Close()

	examplePng := sloppilyGetFileContent("example.png")
	url := fmt.Sprintf("%s?width=100&height=100", ts.URL)
	res, err := http.Post(url, "image/png", bytes.NewBuffer(examplePng))
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	responseBody, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	image := bimg.NewImage(responseBody)

	size, err := image.Size()
	require.NoError(t, err)
	assert.Equal(t, 100, size.Height)
	assert.Equal(t, 100, size.Width)
}
