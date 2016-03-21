package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"util"
)

func Send(msg util.Message, host string, port int, endpoint string) {
	url := "http://" + host + ":" + strconv.Itoa(port) + endpoint

	data, marshalErr := json.Marshal(msg)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	request, reqError := http.NewRequest("POST", url, bytes.NewReader(data))
	if reqError != nil {
		log.Fatal(reqError)
	}

	client := new(http.Client)
	_, clientErr := client.Do(request)
	if clientErr != nil {
		log.Println(clientErr)
	}

}
