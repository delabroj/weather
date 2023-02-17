package models

import "github.com/joeshaw/envdecode"

type Config struct {
	OpenWeatherAPIKey string `env:"OPEN_WEATHER_MAP_API_KEY"`
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envdecode.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
