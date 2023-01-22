package orchestration

import (
	"github.com/jrolstad/weather-insights/internal/clients"
	"github.com/jrolstad/weather-insights/internal/core"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/models"
	"github.com/jrolstad/weather-insights/internal/repositories"
	"time"
)

func GetDaylightData(date time.Time,
	weatherDataClient clients.WeatherDataClient,
	daylightClient clients.DaylightClient,
	dataRepository repositories.DaylightRepository) error {

	stations, err := weatherDataClient.GetWeatherStations()
	if err != nil {
		return err
	}

	processingErrors := make([]error, 0)
	for _, station := range stations {
		logging.LogInfo("Processing daylight times for weather station", "macaddress", station.MacAddress, "name", station.Name)
		err = saveDaylightDataForStation(date, station, daylightClient, dataRepository)
		if err != nil {
			processingErrors = append(processingErrors, err)
		}
	}

	return core.ConsolidateErrors(processingErrors)
}

func saveDaylightDataForStation(date time.Time,
	station *models.WeatherStation,
	daylightClient clients.DaylightClient,
	dataRepository repositories.DaylightRepository) error {

	if !hasDefinedLocation(station) {
		logging.LogInfo("Station does not have a location", "name", station.Name, "latitude", station.Latitude, "longitude", station.Longitude)
		return nil
	}

	daylightData, err := daylightClient.GetTimes(date, station.Latitude, station.Longitude)
	if err != nil {
		return err
	}

	daylightData.LocationIdentifier = station.MacAddress

	err = dataRepository.Save(daylightData)
	return err
}

func hasDefinedLocation(station *models.WeatherStation) bool {
	if station == nil {
		return false
	}

	if station.Latitude == 0 && station.Longitude == 0 {
		return false
	}

	return true
}
