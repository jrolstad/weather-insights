package models

import (
	"time"
)

type Observation struct {
	At                            time.Time
	Station                       *WeatherStation
	RawData                       interface{}
	Barometer                     float64
	DailyRain                     float64
	Dewpoint                      float64
	DewpointIndoor                float64
	EventRain                     float64
	TemperatureFeelsLike          float64
	TemperatureFeelsLikeIndoors   float64
	HourlyRain                    float64
	Humidity                      int
	HumidityIndoors               int
	LastRain                      time.Time
	MaxDailyGust                  float64
	Pm25                          float64
	Pm25Daily                     float64
	MonthlyRain                   float64
	SolarRadiation                float64
	TemperatureFahrenheit         float64
	TemperatureFahrenheitInddors  float64
	TotalRain                     float64
	UvIndex                       float64
	WeeklyRain                    float64
	WindDirection                 int
	WindGust                      float64
	WindGustDirection             int
	WindSpeed                     float64
	WindDirectionAverage2Minutes  int
	WindSpeedAverage2Minutes      float64
	WindDirectionAverage10Minutes int
	WindSpeedAverage10Minutes     float64
	YearlyRain                    float64
	Pm25Inddor                    int
	Pm25DailyIndoor               int
}
