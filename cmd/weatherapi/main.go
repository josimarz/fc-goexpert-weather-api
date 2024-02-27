package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/josimarz/fc-goexpert-weather-api/internal/handler"
	"github.com/josimarz/fc-goexpert-weather-api/internal/viacep"
	"github.com/josimarz/fc-goexpert-weather-api/internal/weather"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}
	cepApi := &viacep.ViaCepApi{}
	weatherApi := weather.NewWeatherApi(os.Getenv("API_KEY"))
	h := handler.NewDefaultHandler(cepApi, weatherApi)
	mux := http.NewServeMux()
	mux.Handle("/", h)
	http.ListenAndServe(":8080", mux)
}
