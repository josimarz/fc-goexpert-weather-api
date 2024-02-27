package viacep

import (
	"testing"

	"github.com/josimarz/fc-goexpert-weather-api/internal/location"
)

func TestViaCepApi(t *testing.T) {
	api := &ViaCepApi{}
	t.Run("Search", func(t *testing.T) {
		cep := "89023600"
		loc, err := api.Search(cep)
		want := &location.Location{
			Region: "Santa Catarina",
			City:   "Blumenau",
		}
		if err != nil || !want.Equals(loc) {
			t.Errorf("Search(%v) = (%v, %v), want (%v, %v)", cep, loc, err, want, nil)
		}
		cep = "00000000"
		loc, err = api.Search(cep)
		if err != errUnableToSearchCep || loc != nil {
			t.Errorf("Search(%v) = (%v, %v), want (%v, %v)", cep, loc, err, nil, errUnableToSearchCep)
		}
	})
}
