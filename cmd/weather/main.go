package main

import (
	"log"
	"net/http"
	"time"

	"github.com/delabroj/weather/api"
	_ "github.com/delabroj/weather/docs" // imported to capture top-level OpenAPI metadata when `swagger generate spec` runs
	"github.com/delabroj/weather/logic"
	"github.com/delabroj/weather/models"
	"github.com/delabroj/weather/openweather"
)

func main() {
	cfg, err := models.NewConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	openWeatherClient := openweather.NewOpenWeatherClient(cfg.OpenWeatherAPIKey)
	logic := logic.NewLogic(openWeatherClient)

	serverConfig := api.ServerConfig{
		Logic: logic,
	}

	apiHandler := api.NewServer(&serverConfig)

	srv := http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      apiHandler,
	}

	log.Println(srv.ListenAndServe())

	log.Println("weather closing")
}
