package util

import (
	"net/http"
	"os"
)

func AddCryptumRequestHeader(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-key", os.Getenv("API_KEY"))
}
