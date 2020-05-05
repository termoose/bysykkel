package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockClient struct{}

func (m *mockClient) doAvailabilityRequest() (*http.Response, error) {
	response := `{
  "last_updated": 1540219230,
  "data": {
    "stations": [
      {
        "is_installed": 1,
        "is_renting": 1,
        "num_bikes_available": 7,
        "num_docks_available": 5,
        "last_reported": 1540219230,
        "is_returning": 1,
        "station_id": "175"
      },
      {
        "is_installed": 1,
        "is_renting": 1,
        "num_bikes_available": 4,
        "num_docks_available": 8,
        "last_reported": 1540219230,
        "is_returning": 1,
        "station_id": "47"
      }]}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(response)))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func (m *mockClient) doStationsRequest() (*http.Response, error) {
	response := `{
  "last_updated": 1553592653,
  "data": {
    "stations": [
      {  
        "station_id":"627",
        "name":"Skøyen Stasjon",
        "address":"Skøyen Stasjon",
        "lat":59.9226729,
        "lon":10.6788129,
        "capacity":20
      },
      {  
        "station_id":"623",
        "name":"7 Juni Plassen",
        "address":"7 Juni Plassen",
        "lat":59.9150596,
        "lon":10.7312715,
        "capacity":15
      }]}}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(response)))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestClient_doAvailabilityRequest(t *testing.T) {
	client := &mockClient{}
	data, err := GetStatusData(client)

	if err != nil {
		t.Error(err)
	}

	if data.LastUpdated != 1540219230 ||
		len(data.Data.Stations) != 2 {
		t.Errorf("Some parsing error")
	}
}

func TestClient_doStationsRequest(t *testing.T) {
	client := &mockClient{}
	data, err := GetStationData(client)

	if err != nil {
		t.Error(err)
	}

	if data.LastUpdated != 1553592653 ||
		len(data.Data.Stations) != 2 ||
		data.Data.Stations[0].Name != "Skøyen Stasjon" {
		t.Errorf("Some parsing error")
	}
}
