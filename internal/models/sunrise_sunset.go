package models

import "time"

type DaylightTimes struct {
	Date                      time.Time
	LocationIdentifier        string
	Latitude                  float64
	Longitude                 float64
	Sunrise                   string
	Sunset                    string
	SolarNoon                 string
	DayLength                 string
	CivilTwilightBegin        string
	CivilTwilightEnd          string
	NauticalTwilightBegin     string
	NauticalTwilightEnd       string
	AstronomicalTwilightBegin string
	AstronomicalTwilightEnd   string
}
