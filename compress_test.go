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

func TestImageCompressSuccess(t *testing.T) {
	var handler http.HandlerFunc = imageCompress
	ts := httptest.NewServer(handler)
	defer ts.Close()

	examplePng := sloppilyGetFileContent("example.jpeg")
	url := fmt.Sprintf("%s?level=5", ts.URL)
	res, err := http.Post(url, "image/jpeg", bytes.NewBuffer(examplePng))
	require.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	exampleJpegCompressed := sloppilyGetFileContent("example.compressed.jpeg")
	assertBodyEqualBytes(t, res, exampleJpegCompressed)
}
