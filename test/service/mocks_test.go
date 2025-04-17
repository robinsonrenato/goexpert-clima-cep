package service_test

import (
	"context"
	"errors"
	"fmt"
	"go-weather/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockViaCEPClient simula o client do ViaCEP
type MockViaCEPClient struct {
	mock.Mock
}

func (m *MockViaCEPClient) GetCityFromCEP(cep string) (string, error) {
	args := m.Called(cep)
	fmt.Printf("GetCityFromCEP chamado com: %s\n", cep) // Adicione este log
	return args.String(0), args.Error(1)
}

// MockWeatherClient simula o client da WeatherAPI
type MockWeatherClient struct {
	mock.Mock
}

func (m *MockWeatherClient) GetTemperatureByCity(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

func TestWeatherService_InvalidCEP(t *testing.T) {
	viaCepMock := new(MockViaCEPClient)
	weatherMock := new(MockWeatherClient)

	// Configurar o mock para o CEP "123"
	viaCepMock.On("GetCityFromCEP", "123").Return("", errors.New("invalid CEP"))

	svc := service.NewWeatherService(viaCepMock, weatherMock)

	// Chamar o método que está sendo testado
	result, err := svc.GetWeatherByCEP(context.Background(), "123")

	// Verificar os resultados
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verificar se o mock foi chamado corretamente
	viaCepMock.AssertCalled(t, "GetCityFromCEP", "123")
	viaCepMock.AssertExpectations(t) // Verifica se todas as expectativas foram atendidas
}
