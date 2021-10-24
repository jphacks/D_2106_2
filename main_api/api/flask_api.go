package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jphacks/D_2106_2/config"
	"github.com/jphacks/D_2106_2/domain"
)

type Gps struct {
	GpsId     int     `json:"gps_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Cluster struct {
	ClusterId      int     `json:"cluster_id"`
	GpsIdBelongsTo []int   `json:"gps_id_belongs_to"`
	MeanLatitude   float64 `json:"mean_latitude"`
	MeanLongitude  float64 `json:"mean_longitude"`
}

type GpsData struct {
	GpsData []Gps `json:"gps_data"`
}

type ClusterData struct {
	ClusterData []Cluster `json:"cluster_data"`
}

func GetSampleApi() GpsData {
	hostname, _ := config.GetDataApiHostname()
	resp, err := http.Get("http://" + hostname + ":5000/api/get_sample")
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

func GetCheckClusteringApi() (*ClusterData, error) {
	data := GetSampleApi()
	responseData, err := GetClusteringApi(data)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func GetClusteringApi(data GpsData) (*ClusterData, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		"POST",
		"http://flask_host:5000/api/clustering",
		bytes.NewBuffer([]byte(string(jsonData))),
	)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseData *ClusterData

	if err := json.Unmarshal(body, &responseData); err != nil {
		log.Fatal(err)
	}
	return responseData, nil
}

func Coordinates2GpsData(coordinates []domain.Coordinate) GpsData {
	var dataList []Gps
	for _, coordinate := range coordinates {
		gps := Gps{
			GpsId:     coordinate.Id,
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		}
		dataList = append(dataList, gps)
	}
	data := GpsData{
		GpsData: dataList,
	}
	return data
}
