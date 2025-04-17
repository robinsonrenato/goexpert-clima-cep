package service

import (
	"context"
	"go-weather/internal/model"
	"go-weather/internal/util"
)

type WeatherService interface {
	GetWeatherByCEP(ctx context.Context, cep string) (*WeatherResult, error)
}

// WeatherResult representa o resultado do serviço de clima.
type WeatherResult struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type weatherService struct {
	ViaCEP  model.ViaCEPClient
	Weather model.WeatherClient
}

func NewWeatherService(viaCEP model.ViaCEPClient, weather model.WeatherClient) WeatherService {
	return &weatherService{
		ViaCEP:  viaCEP,
		Weather: weather,
	}
}

func (s *weatherService) GetWeatherByCEP(ctx context.Context, cep string) (*WeatherResult, error) {
	city, err := s.ViaCEP.GetCityFromCEP(cep)
	if err != nil {
		return nil, err
	}

	tempC, err := s.Weather.GetTemperatureByCity(city)
	if err != nil {
		return nil, err
	}

	return &WeatherResult{
		City:  city,
		TempC: tempC, // Adiciona a temperatura em Celsius
		TempF: util.CelsiusToFahrenheit(tempC),
		TempK: tempC + 273.15,
	}, nil
}
