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
		Station:                       source,
		At:                            toMap.Date,
		Barometer:                     toMap.Baromabsin,
		DailyRain:                     toMap.Dailyrainin,
		Dewpoint:                      toMap.Dewpoint,
		DewpointIndoor:                toMap.Dewpointin,
		EventRain:                     toMap.Eventrainin,
		TemperatureFeelsLike:          toMap.Feelslike,
		TemperatureFeelsLikeIndoors:   toMap.Feelslikein,
		HourlyRain:                    toMap.Hourlyrainin,
		Humidity:                      toMap.Humidity,
		HumidityIndoors:               toMap.Humidityin,
		LastRain:                      toMap.LastRain,
		MaxDailyGust:                  toMap.Maxdailygust,
		Pm25:                          toMap.Pm25,
		Pm25Daily:                     toMap.Pm25_24h,
		MonthlyRain:                   toMap.Monthlyrainin,
		SolarRadiation:                toMap.Solarradiation,
		TemperatureFahrenheit:         toMap.Tempf,
		TemperatureFahrenheitInddors:  toMap.Tempinf,
		TotalRain:                     toMap.Totalrainin,
		UvIndex:                       toMap.Uv,
		WeeklyRain:                    toMap.Weeklyrainin,
		WindDirection:                 toMap.Winddir,
		WindGust:                      toMap.Windgustmph,
		WindGustDirection:             toMap.Windgustdir,
		WindSpeed:                     toMap.Windspeedmph,
		WindDirectionAverage2Minutes:  toMap.Winddir_avg2m,
		WindSpeedAverage2Minutes:      toMap.Windspdmph_avg2m,
		WindDirectionAverage10Minutes: toMap.Winddir_avg10m,
		WindSpeedAverage10Minutes:     toMap.Windspdmph_avg10m,
		YearlyRain:                    toMap.Yearlyrainin,
		Pm25Inddor:                    toMap.Aqi_pm25_in,
		Pm25DailyIndoor:               toMap.Aqi_pm25_in_24h,
		RawData:                       toMap,
	}
}
