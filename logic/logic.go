package logic

import "github.com/delabroj/weather/models"

type logic struct {
	openWeatherClient models.OpenWeatherClient
}

func NewLogic(openWeatherClient models.OpenWeatherClient) logic {
	return logic{
		openWeatherClient: openWeatherClient,
	}
}
