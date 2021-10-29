package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Municipalities struct {
	City string `json:"city"`
	// City_Kana  string `json:"city_kana"`
	Prefecture string `json:"prefecture"`
}
type Location struct {
	Location []Municipalities `json:"location"`
}

type Response struct {
	Response Location `json:"response"`
}

func GetMunicipalitiesByGeoLocation(latitude float64, longtitude float64) (string, string, error) {
	url := "http://geoapi.heartrails.com/api/json"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return "", "", err
	}

	params := request.URL.Query()
	params.Add("method", "searchByGeoLocation")
	params.Add("x", strconv.FormatFloat(longtitude, 'f', -1, 64))
	params.Add("y", strconv.FormatFloat(latitude, 'f', -1, 64))

	request.URL.RawQuery = params.Encode()

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err)
		return "", "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		return "", "", err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Print(err)
		return "", "", err
	}

	if len(res.Response.Location) == 0 {
		return "", "", errors.New("coudn't obtain municipalities")
	}
	return res.Response.Location[0].City, res.Response.Location[0].Prefecture, nil
}
