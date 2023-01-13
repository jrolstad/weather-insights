package config

import "os"

type AppConfig struct {
	AwsRegion          string
	ApplicationKeyName string
	ApiKeyName         string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AwsRegion:          os.Getenv("aws_region"),
		ApplicationKeyName: os.Getenv("weather_application_key_name"),
		ApiKeyName:         os.Getenv("weather_api_key_name"),
	}
}
