package logic

import (
	"fmt"

	"github.com/delabroj/weather"
)

func (l logic) Weather(lat, lon float64) (weather.BasicWeather, error) {
	ret, err := l.openWeatherClient.WeatherByCoordinates(lat, lon)
	if err != nil {
		return weather.BasicWeather{}, fmt.Errorf("Unable to retrieve weather from open weather client: %s", err)
	}

	return ret, nil
}
