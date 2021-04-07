package logic

import "github.com/delabroj/weather"

type logic struct {
	openWeatherClient weather.OpenWeatherClient
}

func NewLogic(openWeatherClient weather.OpenWeatherClient) logic {
	return logic{
		openWeatherClient: openWeatherClient,
	}
}
