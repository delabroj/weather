package logic

import (
	"fmt"

	"github.com/delabroj/weather/models"
)

func (l logic) Weather(lat, lon float64) (models.BasicWeather, error) {
	ret, err := l.openWeatherClient.WeatherByCoordinates(lat, lon)
	if err != nil {
		return models.BasicWeather{}, fmt.Errorf("Unable to retrieve weather from open weather client: %s", err)
	}

	return ret, nil
}
