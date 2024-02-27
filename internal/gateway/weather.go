package gateway

import "github.com/josimarz/fc-goexpert-weather-api/internal/temp"

type WeatherApi interface {
	Search(string) (*temp.Temp, error)
}
