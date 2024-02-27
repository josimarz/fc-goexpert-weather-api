package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josimarz/fc-goexpert-weather-api/internal/location"
	"github.com/josimarz/fc-goexpert-weather-api/internal/temp"
)

type MockCepApi struct{}

func (m *MockCepApi) Search(cep string) (*location.Location, error) {
	return &location.Location{
		Region: "Santa Catarina",
		City:   "Blumenau",
	}, nil
}

type MockWeatherApi struct{}

func (m *MockWeatherApi) Search(location string) (*temp.Temp, error) {
	return &temp.Temp{
		C: 30.5,
		F: 86.9,
		K: 303.65,
	}, nil
}

func TestDefaultHandler(t *testing.T) {
	h := NewDefaultHandler(&MockCepApi{}, &MockWeatherApi{})
	t.Run("ServeHTTP", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/?cep=89023-600", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		expected := `{"temp_C":30.5,"temp_F":86.9,"temp_K":303.65}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
		req, err = http.NewRequest("GET", "/?cep=asdfsda", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
		}
		expected = "invalid zipcode"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
