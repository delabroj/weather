package api

import (
	"net/http"
	"strconv"

	"github.com/delabroj/weather/models"
)

func getWeather(logic models.Logic) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			WriteJSONError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		lat, err := strconv.ParseFloat(req.URL.Query().Get("lat"), 64)
		if lat == 0 || err != nil {
			WriteJSONError(w, "no lat given", http.StatusBadRequest)
			return
		}

		lon, err := strconv.ParseFloat(req.URL.Query().Get("lon"), 64)
		if lat == 0 || err != nil {
			WriteJSONError(w, "no lon given", http.StatusBadRequest)
			return
		}

		weather, err := logic.Weather(lat, lon)
		if err != nil {
			WriteJSONError(w, "Unknown error", http.StatusInternalServerError)
			return
		}

		WriteJSON(w, weather)
	}
}
