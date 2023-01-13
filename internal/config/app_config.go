package config

import "os"

type AppConfig struct {
	AwsRegion string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{AwsRegion: os.Getenv("aws_region")}
}
