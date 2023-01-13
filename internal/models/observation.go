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
	Dewpoint       float64
	DewpointIndoor float64
	Humidity       int
	HumidityIndoor int
}

type TemperatureObservation struct {
	FeelsLike        float64
	FeelsLikeIndoor  float64
	Fahrenheit       float64
	FahrenheitIndoor float64
}

type AirQualityObservation struct {
	Pm25Indoor      int
	Pm25DailyIndoor int
	Pm25            float64
	Pm25Daily       float64
}

type SolarObservation struct {
	SolarRadiation float64
	UvIndex        float64
}

type WindObservation struct {
	Direction                 int
	Gust                      float64
	GustDirection             int
	Speed                     float64
	DirectionAverage2Minutes  int
	SpeedAverage2Minutes      float64
	DirectionAverage10Minutes int
	SpeedAverage10Minutes     float64
	MaxDailyGust              float64
}

type RainObservation struct {
	Daily      float64
	Event      float64
	Hourly     float64
	LastRainAt time.Time
	Monthly    float64
	Total      float64
	Yearly     float64
	Weekly     float64
}
