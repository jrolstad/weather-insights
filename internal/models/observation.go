package models

import (
	"time"
)

type Observation struct {
	At          time.Time
	Station     *WeatherStation
	Rain        *RainObservation
	Wind        *WindObservation
	Solar       *SolarObservation
	AirQuality  *AirQualityObservation
	Temperature *TemperatureObservation
	Humidity    *HumidityObservation
	Pressure    *PressureObservation
	RawData     interface{}
}

type PressureObservation struct {
	Barometer float64
}

type HumidityObservation struct {
	Dewpoint        float64
	DewpointIndoor  float64
	Humidity        int
	HumidityIndoors int
}

type TemperatureObservation struct {
	TemperatureFeelsLike         float64
	TemperatureFeelsLikeIndoors  float64
	TemperatureFahrenheit        float64
	TemperatureFahrenheitInddors float64
}

type AirQualityObservation struct {
	Pm25Inddor      int
	Pm25DailyIndoor int
	Pm25            float64
	Pm25Daily       float64
}

type SolarObservation struct {
	SolarRadiation float64
	UvIndex        float64
}

type WindObservation struct {
	WindDirection                 int
	WindGust                      float64
	WindGustDirection             int
	WindSpeed                     float64
	WindDirectionAverage2Minutes  int
	WindSpeedAverage2Minutes      float64
	WindDirectionAverage10Minutes int
	WindSpeedAverage10Minutes     float64
	MaxDailyGust                  float64
}

type RainObservation struct {
	DailyRain    float64
	EventRain    float64
	HourlyRain   float64
	LastRain     time.Time
	MonthlyRain  float64
	TotalRain    float64
	YearlyRain   float64
	WeeklyRain   float64
	MaxDailyGust float64
}
