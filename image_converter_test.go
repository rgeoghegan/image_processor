package main

import (
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

func TestImageConverterErrors(t *testing.T) {
	var handler http.HandlerFunc = imageConverter
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint only responds to gets
	res, err := http.Get(ts.URL)
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 405)
	assertBodyEqual(t, res, "Only accepts POST requests")

	// Endpoint requires png content-type
	body := strings.NewReader("[1, 2, 3]")
	res, err = http.Post(ts.URL, "application/json", body) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Only accepts image/png")


	// Endpoint requires a png
	body = strings.NewReader("I am obviously not a valid png")
	res, err = http.Post(ts.URL, "image/png", body) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Error converting image: you suck")
}

func TestImageConverterSuccess(t *testing.T) {
	var handler http.HandlerFunc = imageConverter
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint requires png content-type
	examplePng, err := os.Open("example.png")
	require.NoError(t, err)
	defer examplePng.Close()

	res, err := http.Post(ts.URL, "image/png", examplePng) 
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	exampleJpegFile, err := os.Open("example.jpeg")
	require.NoError(t, err)
	defer exampleJpegFile.Close()

	exampleJpeg, err := io.ReadAll(exampleJpegFile)
	require.NoError(t, err)
	assertBodyEqualBytes(t, res, exampleJpeg)
}
