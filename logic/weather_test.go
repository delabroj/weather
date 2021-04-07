package logic

import (
	"errors"
	"testing"

	"github.com/delabroj/weather"
	"github.com/delabroj/weather/mocks"
	"github.com/stretchr/testify/require"
)

func TestWeather(t *testing.T) {
	latOK := 12.2
	lonOK := 21.1

	weatherOK := weather.BasicWeather{
		Temperature: "cold",
		Weather:     "Cloud",
	}

	cases := []struct {
		name string

		retErr error

		expData weather.BasicWeather
		expErr  error
	}{
		{
			name:   "openWeatherClient.Weather error",
			retErr: errors.New("openWeatherClient.Weather error"),
			expErr: errors.New("Unable to retrieve weather from open weather client: openWeatherClient.Weather error"),
		},
		{
			name:    "ok",
			expData: weatherOK,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockOpenWeather := mocks.NewMockOpenWeatherClient()

			mockOpenWeather.WeatherByCoordinatesFunc = func(lat, lon float64) (weather.BasicWeather, error) {
				require.Equal(t, latOK, lat)
				require.Equal(t, lonOK, lon)
				return weatherOK, tc.retErr
			}

			data, err := logic{mockOpenWeather}.Weather(latOK, lonOK)
			require.Equal(t, tc.expErr, err)
			require.Equal(t, tc.expData, data)
		})
	}
}
