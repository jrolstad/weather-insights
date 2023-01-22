package clients

import (
	"encoding/json"
	"fmt"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/models"
	"io/ioutil"
	"net/http"
	"time"
)

type DaylightClient interface {
	GetTimes(date time.Time, latitude float64, longitude float64) (*models.DaylightTimes, error)
}

func NewDaylightClient(config *config.AppConfig) DaylightClient {
	return &SunriseSunsetApiDaylightClient{
		baseUri: config.SunriseSunsetBaseUri,
	}
}

type SunriseSunsetApiDaylightClient struct {
	baseUri string
}

func (c *SunriseSunsetApiDaylightClient) GetTimes(date time.Time, latitude float64, longitude float64) (*models.DaylightTimes, error) {
	uri := fmt.Sprintf("%s/json?lat=%v&lng=%v&date=%s", c.baseUri, latitude, longitude, date.Format("2006-01-02"))

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &SunriseSunsetApiResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	result := c.mapApiResponse(date, latitude, longitude, data)
	return result, nil

}

func (c *SunriseSunsetApiDaylightClient) mapApiResponse(measuredDate time.Time,
	latitudue float64,
	longitude float64,
	data *SunriseSunsetApiResponse) *models.DaylightTimes {
	if data == nil || data.Results == nil {
		return nil
	}

	return &models.DaylightTimes{
		Date:                      measuredDate,
		Latitude:                  latitudue,
		Longitude:                 longitude,
		Sunrise:                   data.Results.Sunrise,
		Sunset:                    data.Results.Sunset,
		SolarNoon:                 data.Results.SolarNoon,
		DayLength:                 data.Results.DayLength,
		CivilTwilightBegin:        data.Results.CivilTwilightBegin,
		CivilTwilightEnd:          data.Results.CivilTwilightEnd,
		NauticalTwilightBegin:     data.Results.NauticalTwilightBegin,
		NauticalTwilightEnd:       data.Results.NauticalTwilightEnd,
		AstronomicalTwilightBegin: data.Results.AstronomicalTwilightBegin,
		AstronomicalTwilightEnd:   data.Results.AstronomicalTwilightEnd,
	}
}

type SunriseSunsetApiResponse struct {
	Results *SunriseSunsetApiData `json:"results"`
	Status  string                `json:"status"`
}
type SunriseSunsetApiData struct {
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	SolarNoon                 string `json:"solar_noon"`
	DayLength                 string `json:"day_length"`
	CivilTwilightBegin        string `json:"civil_twilight_begin"`
	CivilTwilightEnd          string `json:"civil_twilight_end"`
	NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       string `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
}
