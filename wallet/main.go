package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jeffprestes/test-cryptum-golang/model"
	"github.com/jeffprestes/test-cryptum-golang/util"
)

func main() {
	util.LoadEnv()
	apiError := util.NewAPIGetErrorEmpty()
	client := util.GetHttpClient()

	address := "0xcffad3200574698b78f32232aa9d63eabd290703"
	chain := "POLYGON"
	url := util.APIENDPOINT + "/wallet/" + address + "/info?protocol=" + chain
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		apiError.SetError(err)
		apiError.LogErrorWithAditionalMsg("Error generating request")
		return
	}
	util.AddCryptumRequestHeader(request)
	response, err := client.Do(request)
	if err != nil {
		apiError.SetError(err)
		apiError.LogErrorWithAditionalMsg("url:" + url)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err := fmt.Errorf("server returned an error. Status Code: %d - Status: %s", response.StatusCode, response.Status)
		apiError.SetError(err)
		apiError.LogError()
		return
	}
	wallet := model.WalletResponse{}
	err = json.NewDecoder(response.Body).Decode(&wallet)
	if err != nil {
		apiError.SetError(err)
		apiError.LogErrorWithAditionalMsg("Error reading the response from server")
		return
	}
	log.Printf("Server response is: %#v\n\n", wallet)
}
