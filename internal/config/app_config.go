package config

import "os"

type AppConfig struct {
	AwsRegion             string
	ApplicationKeyName    string
	ApiKeyName            string
	ObservationBucketName string
	SunriseSunsetBaseUri  string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AwsRegion:             os.Getenv("aws_region"),
		ApplicationKeyName:    os.Getenv("weather_application_key_name"),
		ApiKeyName:            os.Getenv("weather_api_key_name"),
		ObservationBucketName: os.Getenv("weather_observation_bucket_name"),
		SunriseSunsetBaseUri:  os.Getenv("weather_sunrise_sunset_base_uri"),
	}
}
