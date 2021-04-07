package api

import (
	"net/http"
)

func (srv Server) addRoutes() {
	logic := srv.cfg.Logic

	srv.Mux.Handle("/weather", getWeather(logic))

	srv.Mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		WriteJSONError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
}
