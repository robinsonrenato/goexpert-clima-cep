package main

import (
	"go-weather/internal/client"
	"go-weather/internal/controller"
	"go-weather/internal/service"
	"log"
	"net/http"
)

// WeatherClient represents a client for weather API.

func main() {
	viaCepClient := client.NewViaCEPClient()
	weatherClient := client.NewWeatherClient()
	weatherService := service.NewWeatherService(viaCepClient, weatherClient)

	handler := controller.NewWeatherHandler(weatherService)

	http.HandleFunc("/weather", handler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

