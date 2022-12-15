package util

import (
	"net/http"
	"time"
)

func GetHttpClient() (client http.Client) {
	client = http.Client{
		Timeout: 30 * time.Second,
	}
	return
}
