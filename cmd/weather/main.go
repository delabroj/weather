package main

import (
	"log"
	"net/http"
	"time"

	"github.com/delabroj/weather"
	"github.com/delabroj/weather/api"
	"github.com/delabroj/weather/logic"
	"github.com/delabroj/weather/openweather"
)

func main() {
	cfg, err := weather.NewConfigFromEnv()
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
