package openweather

import (
	"testing"

	"github.com/delabroj/weather"
	"github.com/stretchr/testify/require"
)

func TestWeatherByCoordinates(t *testing.T) {
	cfg, err := weather.NewConfigFromEnv()
	require.NoError(t, err)

	cases := []struct {
		name     string
		lat, lon float64
		expData  bool
		expErr   bool
	}{
		{
			name:   "invalid latitude",
			lat:    1000,
			expErr: true,
		},
		{
			name:    "ok",
			lat:     45,
			lon:     45,
			expData: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewOpenWeatherClient(cfg.OpenWeatherAPIKey)

			data, err := client.WeatherByCoordinates(tc.lat, tc.lon)
			require.Equal(t, tc.expErr, err != nil)
			require.Equal(t, tc.expData, data != weather.BasicWeather{})
		})
	}
}
