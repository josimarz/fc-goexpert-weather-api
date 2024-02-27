package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/josimarz/fc-goexpert-weather-api/internal/gateway"
)

type Response struct {
	C float64 `json:"temp_C"`
	F float64 `json:"temp_F"`
	K float64 `json:"temp_K"`
}

type DefaultHandler struct {
	cepApi     gateway.CepApi
	weatherApi gateway.WeatherApi
}

func NewDefaultHandler(cepApi gateway.CepApi, weatherApi gateway.WeatherApi) *DefaultHandler {
	return &DefaultHandler{cepApi, weatherApi}
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	match, err := regexp.MatchString("\\b\\d{5}-?\\d{3}\\b", cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !match {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}
	loc, err := h.cepApi.Search(cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	search := fmt.Sprintf("%s, %s", loc.City, loc.Region)
	temp, err := h.weatherApi.Search(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := &Response{
		C: temp.C,
		F: temp.F,
		K: temp.K,
	}
	body, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
