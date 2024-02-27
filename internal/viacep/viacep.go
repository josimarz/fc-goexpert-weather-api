package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/josimarz/fc-goexpert-weather-api/internal/location"
)

var (
	errUnableToSearchCep = errors.New("can not find zipcode")
)

type Response struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"GIA"`
	DDD         string `json:"DDD"`
	SIAFI       string `json:"SIAFI"`
	Erro        string `json:"erro"`
}

type ViaCepApi struct{}

func (a *ViaCepApi) Search(cep string) (*location.Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := http.Get(url)
	if err != nil {
		return nil, errUnableToSearchCep
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errUnableToSearchCep
	}
	r := &Response{}
	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return nil, errUnableToSearchCep
	}
	if r.Erro == "true" {
		return nil, errUnableToSearchCep
	}
	return &location.Location{
		Region: location.Regions[r.UF],
		City:   r.Localidade,
	}, nil
}
