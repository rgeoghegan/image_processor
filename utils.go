package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func respondWithServerError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	http.Error(w, err.Error(), 500)
	return true
}

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
