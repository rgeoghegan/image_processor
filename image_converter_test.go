package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertBodyEqual(t *testing.T, response *http.Response, content string) {
	responseBody, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	assert.Equal(t, string(responseBody), content)
}

func assertBodyEqualBytes(t *testing.T, response *http.Response, content []byte) {
	responseBody, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	assert.Equal(t, responseBody, content)
}

// Reads a file you know is on disk, panics otherwise
func sloppilyGetFileContent(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return content
}

func TestImageConverterErrors(t *testing.T) {
	var handler http.HandlerFunc = imageConverter
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint only responds to POSTs
	res, err := http.Get(ts.URL)
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 405)
	assertBodyEqual(t, res, "Only accepts POST requests\n")

	// Endpoint requires png content-type
	body := strings.NewReader("[1, 2, 3]")
	res, err = http.Post(ts.URL, "application/json", body) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Only accepts image/png\n")

	// Endpoint requires a png
	body = strings.NewReader("I am obviously not a valid png")
	res, err = http.Post(ts.URL, "image/png", body) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Error converting image: Unsupported image format\n")
}

func TestImageConverterSuccess(t *testing.T) {
	var handler http.HandlerFunc = imageConverter
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint requires png content-type
	examplePng := sloppilyGetFileContent("example.png")

	res, err := http.Post(ts.URL, "image/png", bytes.NewBuffer(examplePng) )
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	exampleJpeg := sloppilyGetFileContent("example.jpeg")
	require.NoError(t, err)
	assertBodyEqualBytes(t, res, exampleJpeg)
}
