package models

// swagger:model BasicWeather
type BasicWeather struct {
	Temperature string
	Weather     string
}

type WeatherLogic interface {
	Weather(lat, lon float64) (BasicWeather, error)
}
