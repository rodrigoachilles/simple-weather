package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rodrigoachilles/simple-weather/internal/dto"
	"github.com/rodrigoachilles/simple-weather/pkg/log"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	urlWeatherApi = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"
	keyWeatherApi = "KEY_WEATHER_API"
)

type WeatherFinder struct {
	httpClient *http.Client
}

func NewWeatherFinder(httpClient *http.Client) *WeatherFinder {
	return &WeatherFinder{httpClient: httpClient}
}

func (w *WeatherFinder) Execute(query string) (interface{}, error) {
	key := os.Getenv(keyWeatherApi)
	if key == "" {
		return nil, errors.New("API key is not provided")
	}
	requestURL := fmt.Sprintf(urlWeatherApi, key, url.QueryEscape(query))

	log.Logger.Debug().Msg(fmt.Sprintf("Calling api url: %s", requestURL))

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")

	res, err := w.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()

	log.Logger.Debug().Msg(fmt.Sprintf("Response body: %s", string(body)))

	if res.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("API key is invalid")
	}

	if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New("can not find zipcode")
	}

	var output dto.WeatherOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
