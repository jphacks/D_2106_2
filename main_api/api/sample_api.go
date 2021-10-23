package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SampleBody struct {
	Message string `json:"message"`
}

func FetchSampleApi() string {
	resp, err := http.Get("http://localhost:5000/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data SampleBody

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	return data.Message
}
