package api

import (
	"net/http"

	"github.com/delabroj/weather/models"
)

// swagger:route GET /weather Weather GetWeather
//
// Get weather
//
// Get weather by position
//
// responses:
//	200: BasicWeather
//	default: APIError

// swagger:parameters GetWeather
type getWeatherParams struct {
	// in: query
	// required: true
	// example: 0
	Latitude float64 `form:"lat" json:"lat"`
	// in: query
	// required: true
	// example: 0
	Longitude float64 `form:"lon" json:"lon"`
}

func getWeather(logic models.Logic) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var params getWeatherParams

		err := bindQuery(req, &params)
		if err != nil {
			WriteJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		weather, err := logic.Weather(params.Latitude, params.Longitude)
		if err != nil {
			WriteJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		WriteJSON(w, weather)
	}
}
