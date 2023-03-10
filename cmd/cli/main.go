package main

import (
	"github.com/jrolstad/weather-insights/internal/clients"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/orchestration"
	"github.com/jrolstad/weather-insights/internal/repositories"
	"time"
)

func main() {
	start := time.Now().UTC().AddDate(0, 0, -1)
	end := time.Now().UTC()

	appConfig := config.NewAppConfig()
	secretClient := clients.NewSecretClient(appConfig)
	client, err := clients.NewWeatherDataClient(appConfig, secretClient)
	if err != nil {
		logging.LogPanic(err)
	}
	repository := repositories.NewObservationRepository(appConfig)

	err = orchestration.GetWeatherData(start, end, client, repository)
	if err != nil {
		logging.LogPanic(err)
	}

	logging.LogInfo("GetWeatherData Complete", "start", start.String(), "end", end.String())
}
