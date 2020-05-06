package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const stationPath = "https://gbfs.urbansharing.com/oslobysykkel.no/station_information.json"
const availabilityPath = "https://gbfs.urbansharing.com/oslobysykkel.no/station_status.json"

type ClientInterface interface {
	doAvailabilityRequest() (*http.Response, error)
	doStationsRequest()     (*http.Response, error)
}

type Client struct {
	Client *http.Client
}

func NewAPIClient() *Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		Client: client,
	}
}

func GetStationData(client ClientInterface) (*stationsData, error) {
	response, err := client.doStationsRequest()

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	data := &stationsData{}

	err = decoder.Decode(&data)

	if err != nil {
		return nil, fmt.Errorf("parsing: %v", err)
	}

	return data, nil
}

func GetStatusData(client ClientInterface) (*statusData, error) {
	response, err := client.doAvailabilityRequest()

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	data := &statusData{}

	err = decoder.Decode(&data)

	if err != nil {
		return nil, fmt.Errorf("parsing: %v", err)
	}

	return data, nil
}

func GetLocation(stations *stationsData, id int) (float32, float32) {
	for _, elem := range stations.Data.Stations {
		elemId, _ := strconv.Atoi(elem.Id)

		if elemId == id {
			fmt.Printf("Found lat long %d %d\n", elem.Latitude, elem.Longitude)
			return elem.Latitude, elem.Longitude
		}
	}

	return 0, 0
}

func GetAvailability(statuses *statusData, id int) *AvailabilityData {
	for _, elem := range statuses.Data.Stations {
		elemId, _ := strconv.Atoi(elem.Id)

		if elemId == id {
			return &elem
		}
	}

	return nil
}

func (c *Client) doAvailabilityRequest() (*http.Response, error) {
	return c.doRequest(availabilityPath)
}

func (c *Client) doStationsRequest() (*http.Response, error) {
	return c.doRequest(stationPath)
}

func (c *Client) doRequest(path string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Add("Client-Identifier", "termoose-intervju")

	response, err := c.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("doRequest: %v", err)
	}

	return response, nil
}