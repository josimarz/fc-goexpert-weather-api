package gateway

import "github.com/josimarz/fc-goexpert-weather-api/internal/location"

type CepApi interface {
	Search(string) (*location.Location, error)
}
