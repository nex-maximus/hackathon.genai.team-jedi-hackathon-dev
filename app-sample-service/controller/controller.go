/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 **********************************************************************/

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type SampleController struct {
	lc     logger.LoggingClient
	sample string
}

type PayloadData struct {
	Inference []InferenceStruct `json:"inference"`
}

type InferenceStruct struct {
	VideoTimestamp string `json:"videoTimestamp"`
	Label          string `json:"label"`
	Accuracy       string `json:"accuracy"`
}

func New(lc logger.LoggingClient, sample string) *SampleController {
	return &SampleController{
		lc:     lc,
		sample: sample,
	}
}

func (p *SampleController) RegisterRoutes(service interfaces.ApplicationService) error {
	if err := service.AddRoute("/api/v1/data", p.postData, http.MethodPost); err != nil {
		return fmt.Errorf("could not register routes: %s", err.Error())
	}

	p.lc.Info("Routes added...")
	return nil
}

func (p *SampleController) postData(writer http.ResponseWriter, request *http.Request) {

	sounds := os.Getenv("SOUNDS")
	soundsSlice := strings.Split(sounds, ",")

	var data PayloadData
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//Connect to influxDB
	token := os.Getenv("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	url := "http://influxdb:8086"

	// Create a new InfluxDB client
	client := influxdb2.NewClient(url, token)
	org := "AiCSD"
	bucket := "sound"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// Check if detected sound is in the list
	for _, val := range data.Inference {
		found := stringInSlice(val.Label, soundsSlice)

		// Insert record in influxDB
		if found {
			tags := map[string]string{"data": "sound"}
			fields := map[string]interface{}{"videoTimestamp": val.VideoTimestamp, "label": val.Label, "accuracy": val.Accuracy}
			point := influxdb2.NewPoint("data", tags, fields, time.Now())
			if err := writeAPI.WritePoint(context.Background(), point); err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Printf("Received JSON data: %+v\n", data)
	writer.WriteHeader(200)

}

func stringInSlice(target string, slice []string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
