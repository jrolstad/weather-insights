package orchestration

import (
	"github.com/jrolstad/weather-insights/internal/clients"
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
		logging.LogError(err)
	}

	for _, item := range stations {
		waitForApiThrottling()
		getObservations(item, start, end, weatherDataClient, dataRepository)
	}

	return nil
}

func getObservations(station *models.WeatherStation,
	start time.Time,
	end time.Time,
	weatherDataClient clients.WeatherDataClient,
	observationRepository repositories.ObservationRepository) error {

	logging.LogInfo("Obtaining observations from station", "mac", station.MacAddress, "name", station.Name)
	observations, err := weatherDataClient.GetObservations(station.MacAddress, start, end)
	if err != nil {
		return err
	}

	err = observationRepository.Save(observations)

	return err
}

func waitForApiThrottling() {
	time.Sleep(time.Second * 1)
}
