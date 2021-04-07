package openweather

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/delabroj/weather"
)

const baseURL = "https://api.openweathermap.org"

type openWeatherClient struct {
	apiKey string
	client *http.Client
}

func NewOpenWeatherClient(apiKey string) openWeatherClient {
	return openWeatherClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: time.Second * 60},
	}
}

type openWeatherAPIWeather struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

func (c openWeatherClient) WeatherByCoordinates(lat, lon float64) (weather.BasicWeather, error) {
	values := c.newURLValues()
	values.Set("lat", fmt.Sprint(lat))
	values.Set("lon", fmt.Sprint(lon))

	u := baseURL + "/data/2.5/weather?" + values.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return weather.BasicWeather{}, err
	}

	var response openWeatherAPIWeather
	if err = c.makeRequest(req, &response); err != nil {
		return weather.BasicWeather{}, err
	}

	ret := weather.BasicWeather{
		Weather: response.Weather[0].Main,
	}

	if response.Main.Temperature < 288.7 {
		ret.Temperature = "cold"
	} else if response.Main.Temperature < 299.8 {
		ret.Temperature = "moderate"
	} else {
		ret.Temperature = "hot"
	}

	return ret, nil
}

func (c openWeatherClient) newURLValues() url.Values {
	ret := url.Values{}
	ret.Set("appid", c.apiKey)
	return ret
}

func (c openWeatherClient) makeRequest(req *http.Request, dataObj interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		if ne, ok := err.(net.Error); ok {
			if ne.Timeout() || ne.Temporary() {
				return fmt.Errorf("Timeout or temporary network error: %s", ne)
			}
		}

		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response: %d - %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(dataObj); err != nil {
		return err
	}

	return nil
}
