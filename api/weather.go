package api

import (
	"net/http"
	"strconv"

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
	Latitude float64 `form:"-" json:"lat"`
	// in: query
	// required: true
	// example: 0
	Longitude float64 `form:"-" json:"lon"`
}

func getWeather(logic models.Logic) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			WriteJSONError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var params getWeatherParams

		var err error
		params.Latitude, err = strconv.ParseFloat(req.URL.Query().Get("lat"), 64)
		if err != nil {
			WriteJSONError(w, "no lat given", http.StatusBadRequest)
			return
		}

		params.Longitude, err = strconv.ParseFloat(req.URL.Query().Get("lon"), 64)
		if err != nil {
			WriteJSONError(w, "no lon given", http.StatusBadRequest)
			return
		}

		weather, err := logic.Weather(params.Latitude, params.Longitude)
		if err != nil {
			WriteJSONError(w, "Unknown error", http.StatusInternalServerError)
			return
		}

		WriteJSON(w, weather)
	}
}
