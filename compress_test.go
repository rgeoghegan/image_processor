package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageCompressErrors(t *testing.T) {
	var handler http.HandlerFunc = imageCompress
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Endpoint only responds to POSTs
	res, err := http.Get(ts.URL)
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 405)
	assertBodyEqual(t, res, "Only accepts POST requests\n")

	exampleJpeg := sloppilyGetFileContent("example.jpeg")

	// Endpoint requires valid compression level parameters
	url := fmt.Sprintf("%s?level=banana", ts.URL)
	res, err = http.Post(url, "image/png", bytes.NewBuffer(exampleJpeg))
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 400)
	assertBodyEqual(t, res, "Invalid or missing level url parameter\n")
}

func TestImageCompressSuccess(t *testing.T) {
	var handler http.HandlerFunc = imageCompress
	ts := httptest.NewServer(handler)
	defer ts.Close()

	exampleJpeg := sloppilyGetFileContent("example.jpeg")
	url := fmt.Sprintf("%s?level=5", ts.URL)
	res, err := http.Post(url, "image/jpeg", bytes.NewBuffer(exampleJpeg))
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	exampleJpegCompressed := sloppilyGetFileContent("example.compressed.jpeg")
	assertBodyEqualBytes(t, res, exampleJpegCompressed)
}
