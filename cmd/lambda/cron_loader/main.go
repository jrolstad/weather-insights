package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jrolstad/weather-insights/internal/clients"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/core"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/orchestration"
	"github.com/jrolstad/weather-insights/internal/repositories"
	"time"
)

var (
	appConfig             *config.AppConfig
	secretClient          clients.SecretClient
	weatherClient         clients.WeatherDataClient
	daylightClient        clients.DaylightClient
	observationRepository repositories.ObservationRepository
	daylightRepository    repositories.DaylightRepository
)

func init() {
	var err error
	appConfig = config.NewAppConfig()

	secretClient = clients.NewSecretClient(appConfig)
	weatherClient, err = clients.NewWeatherDataClient(appConfig, secretClient)
	if err != nil {
		logging.LogError(err)
	}
	daylightClient = clients.NewDaylightClient(appConfig)

	observationRepository = repositories.NewObservationRepository(appConfig)
	daylightRepository = repositories.NewDaylightRepository(appConfig)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	start := time.Now().UTC().AddDate(0, 0, -1)
	end := time.Now().UTC()

	processingErrors := make([]error, 0)

	logging.LogInfo("Retrieving weather data", "start", start, "end", end)
	weatherErr := orchestration.GetWeatherData(start, end, weatherClient, observationRepository)
	if weatherErr != nil {
		processingErrors = append(processingErrors, weatherErr)
	}
	logging.LogInfo("Weather data retrieval complete", "error", weatherErr)

	logging.LogInfo("Retrieving daylight data", "start", start)
	daylightErr := orchestration.GetDaylightData(end, weatherClient, daylightClient, daylightRepository)
	if daylightErr != nil {
		processingErrors = append(processingErrors, daylightErr)
	}
	logging.LogInfo("Daylight data retrieval complete", "error", daylightErr)

	return core.ConsolidateErrors(processingErrors)
}
