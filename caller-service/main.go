package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	googleMetadataURL = "http://metadata.google.internal/computeMetadata/v1/instance/zone"
	pauseTime         = 3
)

type Model struct {
	SourcePodName  string `json:"SourcePodName"`
	SourceNodeZone string `json:"SourceNodeZone"`
	DestPodName    string `json:"DestPodName"`
	DestNodeZone   string `json:"DestNodeZone"`
}

func getNodeZone() string {
	req, err := http.NewRequest("GET", googleMetadataURL, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	req.Header.Add("Metadata-Flavor", "Google")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	zone_temp := strings.Split(string(resBody), "/")
	zone := zone_temp[len(zone_temp)-1]
	return zone
}

func caller() {

	podName := os.Getenv("PODNAME")

	zone := getNodeZone()
	if zone == "" {
		log.Fatal("Failed to get node zone")
	}

	params := url.Values{}
	params.Add("name", podName)
	params.Add("zone", zone)

	url := os.Getenv("ENDPOINT") + "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("accept", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil && response.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	var msg Model
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(resBody, &msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", msg)
}
func main() {
	log.Print("Starting Caller Serice")
	for {
		caller()
		time.Sleep(pauseTime * time.Second)
	}
}
