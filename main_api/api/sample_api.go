package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Gps struct {
	GpsId int `json:"gps_id"`
	Latitude float32 `json:latitude`
	Longitude float32 `json:longitude`
}

type GpsData struct {
	GpsData []Gps `json:"gps_data"`
}

func FetchSampleApi() GpsData {
	resp, err := http.Get("http://flask_host:5000/api/get_sample")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data GpsData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	return data
}
