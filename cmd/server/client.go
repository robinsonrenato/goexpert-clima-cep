package main

import "net/http"

// WeatherClient represents a client for weather API.
type WeatherClient struct {
    httpClient *http.Client
}

// NewWeatherClient creates a new instance of WeatherClient.
func NewWeatherClient() *WeatherClient {
    return &WeatherClient{
        httpClient: &http.Client{},
    }
}
