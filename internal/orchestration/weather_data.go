package orchestration

import (
	"github.com/jrolstad/weather-insights/internal/clients"
	"github.com/jrolstad/weather-insights/internal/core"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/models"
	"github.com/jrolstad/weather-insights/internal/repositories"
	"time"
)

func GetWeatherData(start time.Time,
	end time.Time,
	weatherDataClient clients.WeatherDataClient,
	dataRepository repositories.ObservationRepository) error {

	stations, err := weatherDataClient.GetWeatherStations()
	if err != nil {
		return err
	}

	processingErrors := make([]error, 0)
	for _, item := range stations {
		waitForApiThrottling()
		err = getObservations(item, start, end, weatherDataClient, dataRepository)
		if err != nil {
			processingErrors = append(processingErrors, err)
		}
	}

	return core.ConsolidateErrors(processingErrors)
}

func getObservations(station *models.WeatherStation,
	start time.Time,
	end time.Time,
	weatherDataClient clients.WeatherDataClient,
	observationRepository repositories.ObservationRepository) error {

	logging.LogInfo("Obtaining observations from station", "mac", station.MacAddress, "name", station.Name)
	observations, err := weatherDataClient.GetObservations(station, start, end)
	if err != nil {
		return err
	}

	err = observationRepository.Save(observations)

	return err
}

func waitForApiThrottling() {
	time.Sleep(time.Second * 1)
}
