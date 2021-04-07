package weather

type OpenWeatherClient interface {
	WeatherByCoordinates(lat, lon float64) (BasicWeather, error)
}
