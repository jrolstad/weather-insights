package clients

import (
	"github.com/jrolstad/weather-insights/internal/models"
	"time"
)

type WeatherDataClient interface {
	GetWeatherStations() ([]*models.WeatherStation, error)
	GetObservations(stationId string, start time.Time, end time.Time) ([]*models.Observation, error)
}

func NewWeatherDataClient(secretClient SecretClient) WeatherDataClient {
	return nil
}

type AmbientWeatherClient struct {
}

func (c *AmbientWeatherClient) GetWeatherStations() ([]*models.WeatherStation, error) {

}

func (c *AmbientWeatherClient) GetObservations(stationId string, start time.Time, end time.Time) ([]*models.Observation, error) {

}
