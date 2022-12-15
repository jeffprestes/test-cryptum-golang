package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jeffprestes/test-cryptum-golang/util"
)

func init() {
	util.LoadEnv()
}

func main() {
	client := util.GetHttpClient()

	url := util.APIENDPOINT + "/prices/ETH"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error generating request", err.Error())
	}
	util.AddCryptumRequestHeader(request)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error calling server", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalln("Server returned an error", response.StatusCode, response.Status)
	}
	buffer, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error reading the response from server:", err.Error())
	}
	log.Println("Server response is:", string(buffer))
}
