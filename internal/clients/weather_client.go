package clients

import (
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/models"
	"github.com/lrosenman/ambient"
	"time"
)

type WeatherDataClient interface {
	GetWeatherStations() ([]*models.WeatherStation, error)
	GetObservations(stationId string, start time.Time, end time.Time) ([]*models.Observation, error)
}

func NewWeatherDataClient(appConfig *config.AppConfig, secretClient SecretClient) (WeatherDataClient, error) {
	applicationKey, err := secretClient.GetSecret(appConfig.ApplicationKeyName)
	if err != nil {
		return nil, err
	}
	apiKey, err := secretClient.GetSecret(appConfig.ApiKeyName)
	if err != nil {
		return nil, err
	}

	return &AmbientWeatherClient{
		authenticationKey: ambient.NewKey(applicationKey, apiKey),
	}, nil
}

type AmbientWeatherClient struct {
	authenticationKey ambient.Key
}

func (c *AmbientWeatherClient) GetWeatherStations() ([]*models.WeatherStation, error) {
	result := make([]*models.WeatherStation, 0)

	data, err := ambient.Device(c.authenticationKey)
	if err != nil {
		return result, err
	}

	mappedData := mapDevices(data.DeviceRecord)
	return mappedData, nil
}

func (c *AmbientWeatherClient) GetObservations(stationId string, start time.Time, end time.Time) ([]*models.Observation, error) {
	return make([]*models.Observation, 0), nil
}

func mapDevices(toMap []ambient.DeviceRecord) []*models.WeatherStation {
	result := make([]*models.WeatherStation, 0)
	if toMap == nil {
		return result
	}

	for _, item := range toMap {
		mappedItem := mapDevice(item)
		result = append(result, mappedItem)
	}

	return result
}

func mapDevice(toMap ambient.DeviceRecord) *models.WeatherStation {
	return &models.WeatherStation{
		MacAddress: toMap.Macaddress,
		Name:       toMap.Info.Name,
		Location:   toMap.Info.Location,
	}
}
