package data

import (
	"bysykkel/api"
	"strconv"
)

type Station struct {
	Id        int
	Name      string
	Address   string
	Lat       float32
	Lon       float32
	Capacity  int
	Available int
}

type Stations []Station

func GetLocation(id int) (float32, float32) {
	client := api.NewAPIClient()

	stations, stationErr := api.GetStationData(client)
	if stationErr != nil {
		return 0, 0
	}

	return api.GetLocation(stations, id)
}

func GetFrontendData() (Stations, error) {
	client := api.NewAPIClient()

	stations, stationErr := api.GetStationData(client)
	if stationErr != nil {
		return nil, stationErr
	}

	statuses, statusErr := api.GetStatusData(client)
	if statusErr != nil {
		return nil, statusErr
	}

	var result Stations

	for _, station := range stations.Data.Stations {
		id, _ := strconv.Atoi(station.Id)
		avail := api.GetAvailability(statuses, id)

		s := Station{
			Id:        id,
			Name:      station.Name,
			Address:   station.Address,
			Lat:       station.Latitude,
			Lon:       station.Longitude,
			Capacity:  station.Capacity,
			Available: avail.BikesAvailable,
		}

		result = append(result, s)
	}

	return result, nil
}