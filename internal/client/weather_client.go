package client

import (
	"errors"
	"go-weather/internal/model"
	"go-weather/internal/util"
	"net/http"
)

// WeatherClient representa o cliente para a API de clima.
type WeatherClient struct {
	httpClient *http.Client
}

// NewWeatherClient cria uma nova instância de WeatherClient.
func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		httpClient: &http.Client{},
	}
}

// GetTemperatureByCity simula a busca de temperatura por cidade.
func (wc *WeatherClient) GetTemperatureByCity(city string) (float64, error) {
	// Simulação de implementação. Substitua pela lógica real de chamada à API.
	if city == "" {
		return 0, errors.New("cidade inválida")
	}
	return 25.0, nil // Retorna uma temperatura fixa como exemplo.
}

type WeatherService struct {
	ViaCEP  model.ViaCEPClient
	Weather model.WeatherClient
}

func NewWeatherService(viacep model.ViaCEPClient, weather model.WeatherClient) *WeatherService {
	return &WeatherService{
		ViaCEP:  viacep,
		Weather: weather,
	}
}

func (s *WeatherService) GetWeatherByCEP(cep string) (*model.TemperatureResponse, error) {
	city, err := s.ViaCEP.GetCityFromCEP(cep)
	if err != nil {
		return nil, err
	}

	tempC, err := s.Weather.GetTemperatureByCity(city)
	if err != nil {
		return nil, err
	}

	return &model.TemperatureResponse{
		TempC: tempC,
		TempF: util.CelsiusToFahrenheit(tempC),
		TempK: util.CelsiusToKelvin(tempC),
		City:  city,
	}, nil
}
