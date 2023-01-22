package main

import (
	"flag"
	"github.com/jrolstad/weather-insights/internal/clients"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/orchestration"
	"github.com/jrolstad/weather-insights/internal/repositories"
	"strings"
	"time"
)

func main() {
	action := flag.String("action", "", "Type of action to use.  Valid values are daylight and weather")
	flag.Parse()

	start := time.Now().UTC().AddDate(0, 0, -1)
	end := time.Now().UTC()

	appConfig := config.NewAppConfig()
	secretClient := clients.NewSecretClient(appConfig)
	daylightClient := clients.NewDaylightClient(appConfig)
	client, err := clients.NewWeatherDataClient(appConfig, secretClient)
	if err != nil {
		logging.LogPanic(err)
	}
	repository := repositories.NewObservationRepository(appConfig)
	daylightRepository := repositories.NewDaylightRepository(appConfig)

	if strings.EqualFold(*action, "daylight") {
		err = orchestration.GetDaylightData(end, client, daylightClient, daylightRepository)
		if err != nil {
			logging.LogPanic(err)
		}

		logging.LogInfo("GetDaylightData Complete", "date", end.String())
	}

	if strings.EqualFold(*action, "weather") {
		err = orchestration.GetWeatherData(start, end, client, repository)
		if err != nil {
			logging.LogPanic(err)
		}

		logging.LogInfo("GetWeatherData Complete", "start", start.String(), "end", end.String())
	}
}
