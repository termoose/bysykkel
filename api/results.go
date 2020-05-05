package api

type stationData struct {
	Id        string  `json:"station_id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
	Capacity  int     `json:"capacity"`
}

type availabilityData struct {
	Id             string `json:"station_id"`
	IsInstalled    int    `json:"is_installed"`
	IsRenting      int    `json:"is_renting"`
	BikesAvailable int    `json:"num_bikes_available"`
	DocksAvailable int    `json:"num_docks_available"`
	LastReported   int    `json:"last_reported"`
	IsReturning    int    `json:"is_returning"`
}

type stationList struct {
	Stations []stationData `json:"stations"`
}

type statusList struct {
	Stations []availabilityData `json:"stations"`
}

type statusData struct {
	LastUpdated int        `json:"last_updated"`
	Data        statusList `json:"data"`
}

type stationsData struct {
	LastUpdated int         `json:"last_updated"`
	Data        stationList `json:"data"`
}
