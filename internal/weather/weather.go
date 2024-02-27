package weather

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/josimarz/fc-goexpert-weather-api/internal/temp"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Response struct {
	Current struct {
		C float64 `json:"temp_c"`
		F float64 `json:"temp_f"`
	} `json:"current"`
}

type WeatherApi struct {
	key    string
	client HTTPClient
}

func NewWeatherApi(key string) *WeatherApi {
	return &WeatherApi{
		key:    key,
		client: &http.Client{},
	}
}

func (a *WeatherApi) Search(location string) (*temp.Temp, error) {
	req, err := http.NewRequest(http.MethodGet, "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", a.key)
	q.Add("q", location)
	q.Add("aqi", "no")
	req.URL.RawQuery = q.Encode()
	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, a.parseError(res.Body)
	}
	r := &Response{}
	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return nil, err
	}
	return &temp.Temp{
		C: r.Current.C,
		F: r.Current.F,
		K: temp.CToK(r.Current.C),
	}, nil
}

func (a *WeatherApi) parseError(r io.Reader) error {
	res := &ErrorResponse{}
	if err := json.NewDecoder(r).Decode(res); err != nil {
		return err
	}
	return errors.New(res.Error.Message)
}
