package service_test

import (
	"context"
	"errors"
	"go-weather/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherService_Success(t *testing.T) {
	viaCepMock := new(MockViaCEPClient)
	weatherMock := new(MockWeatherClient)

	// Configurar o mock para o CEP "01001000"
	viaCepMock.On("GetCityFromCEP", "01001000").Return("São Paulo", nil)
	weatherMock.On("GetTemperatureByCity", "São Paulo").Return(25.0, nil)

	svc := service.NewWeatherService(viaCepMock, weatherMock)

	// Chamar o método que está sendo testado
	result, err := svc.GetWeatherByCEP(context.Background(), "01001000")

	// Verificar os resultados
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verificar o resultado esperado
	expectedResult := &service.WeatherResult{
		City:  "São Paulo",
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.15,
	}
	assert.Equal(t, expectedResult.City, result.City)
	assert.Equal(t, expectedResult.TempC, result.TempC)
	assert.Equal(t, expectedResult.TempF, result.TempF)
	assert.InDelta(t, expectedResult.TempK, result.TempK, 0.01, "TempK deve ser aproximadamente 298.15")

	// Verificar se os mocks foram chamados corretamente
	viaCepMock.AssertCalled(t, "GetCityFromCEP", "01001000")
	weatherMock.AssertCalled(t, "GetTemperatureByCity", "São Paulo")
	viaCepMock.AssertExpectations(t)
	weatherMock.AssertExpectations(t)
}

func TestWeatherService_CEPNotFound(t *testing.T) {
	viaCepMock := new(MockViaCEPClient)
	weatherMock := new(MockWeatherClient)

	viaCepMock.On("GetCityFromCEP", "99999999").Return("", errors.New("not found"))

	svc := service.NewWeatherService(viaCepMock, weatherMock)

	result, err := svc.GetWeatherByCEP(context.Background(), "99999999")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestWeatherService_WeatherAPIError(t *testing.T) {
	viaCepMock := new(MockViaCEPClient)
	weatherMock := new(MockWeatherClient)

	viaCepMock.On("GetCityFromCEP", "01001000").Return("São Paulo", nil)
	weatherMock.On("GetTemperatureByCity", "São Paulo").Return(0.0, errors.New("weather API error"))

	svc := service.NewWeatherService(viaCepMock, weatherMock)

	result, err := svc.GetWeatherByCEP(context.Background(), "01001000")

	assert.Error(t, err)
	assert.Nil(t, result)
}
