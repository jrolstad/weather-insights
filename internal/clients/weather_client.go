package clients

import (
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/models"
	"github.com/lrosenman/ambient"
	"time"
)

type WeatherDataClient interface {
	GetWeatherStations() ([]*models.WeatherStation, error)
	GetObservations(station *models.WeatherStation, start time.Time, end time.Time) ([]*models.Observation, error)
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
		queryLimit:        288, //This is the max allowed
	}, nil
}

type AmbientWeatherClient struct {
	authenticationKey ambient.Key
	queryLimit        int64
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

func (c *AmbientWeatherClient) GetObservations(station *models.WeatherStation, start time.Time, end time.Time) ([]*models.Observation, error) {
	data, err := ambient.DeviceMac(c.authenticationKey, station.MacAddress, end, c.queryLimit)
	if err != nil {
		return make([]*models.Observation, 0), err
	}

	mappedData := mapObservations(station, data.Record)
	return mappedData, nil
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

func mapObservations(source *models.WeatherStation, toMap []ambient.Record) []*models.Observation {
	result := make([]*models.Observation, 0)
	if toMap == nil {
		return result
	}

	for _, item := range toMap {
		mappedItem := mapObservation(source, item)
		result = append(result, mappedItem)
	}

	return result
}

func mapObservation(source *models.WeatherStation, toMap ambient.Record) *models.Observation {
	return &models.Observation{
		Station: source,
		At:      toMap.Date,
		Pressure: &models.PressureObservation{
			Barometer: toMap.Baromabsin,
		},
		Rain: &models.RainObservation{
			DailyRain:   toMap.Dailyrainin,
			HourlyRain:  toMap.Hourlyrainin,
			LastRain:    toMap.LastRain,
			WeeklyRain:  toMap.Weeklyrainin,
			YearlyRain:  toMap.Yearlyrainin,
			EventRain:   toMap.Eventrainin,
			MonthlyRain: toMap.Monthlyrainin,
			TotalRain:   toMap.Totalrainin,
		},
		Humidity: &models.HumidityObservation{
			Dewpoint:        toMap.Dewpoint,
			DewpointIndoor:  toMap.Dewpointin,
			Humidity:        toMap.Humidity,
			HumidityIndoors: toMap.Humidityin,
		},
		Temperature: &models.TemperatureObservation{
			TemperatureFeelsLike:         toMap.Feelslike,
			TemperatureFeelsLikeIndoors:  toMap.Feelslikein,
			TemperatureFahrenheit:        toMap.Tempf,
			TemperatureFahrenheitInddors: toMap.Tempinf,
		},
		Wind: &models.WindObservation{
			MaxDailyGust:                  toMap.Maxdailygust,
			WindDirection:                 toMap.Winddir,
			WindGust:                      toMap.Windgustmph,
			WindGustDirection:             toMap.Windgustdir,
			WindSpeed:                     toMap.Windspeedmph,
			WindDirectionAverage2Minutes:  toMap.Winddir_avg2m,
			WindSpeedAverage2Minutes:      toMap.Windspdmph_avg2m,
			WindDirectionAverage10Minutes: toMap.Winddir_avg10m,
			WindSpeedAverage10Minutes:     toMap.Windspdmph_avg10m,
		},
		AirQuality: &models.AirQualityObservation{
			Pm25:            toMap.Pm25,
			Pm25Daily:       toMap.Pm25_24h,
			Pm25Inddor:      toMap.Aqi_pm25_in,
			Pm25DailyIndoor: toMap.Aqi_pm25_in_24h,
		},
		Solar: &models.SolarObservation{
			SolarRadiation: toMap.Solarradiation,
			UvIndex:        toMap.Uv,
		},
		RawData: toMap,
	}
}
