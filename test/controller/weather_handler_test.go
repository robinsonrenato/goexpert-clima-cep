package controller_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-weather/internal/controller"
	"go-weather/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWeatherService is a mock implementation of the WeatherService interface.
type MockWeatherService struct{}

// GetWeatherByCEP is a mock method that satisfies the WeatherService interface.
func (m *MockWeatherService) GetWeatherByCEP(ctx context.Context, cep string) (*service.WeatherResult, error) {
	if cep == "01001000" {
		return &service.WeatherResult{
			City:  "São Paulo",
			TempC: 25.0,
			TempF: 77.0,
			TempK: 298.15,
		}, nil
	}
	return nil, service.ErrInvalidCEP
}

// MockViaCEPClient é um mock da interface ViaCEPClient.
type MockViaCEPClient struct {
	mock.Mock
}

// GetCityFromCEP é o método mockado que implementa a interface ViaCEPClient.
func (m *MockViaCEPClient) GetCityFromCEP(cep string) (string, error) {
	args := m.Called(cep)
	return args.String(0), args.Error(1)
}

// MockWeatherClient é um mock da interface WeatherClient.
type MockWeatherClient struct {
	mock.Mock
}

// GetTemperatureByCity é o método mockado que implementa a interface WeatherClient.
func (m *MockWeatherClient) GetTemperatureByCity(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

func TestWeatherHandler_InvalidCEP(t *testing.T) {
	// Mock do serviço
	mockService := new(MockWeatherService)
	handler := controller.NewWeatherHandler(mockService) // Instanciar o handler corretamente

	req := httptest.NewRequest("GET", "/weather?cep=123", nil)
	w := httptest.NewRecorder()

	// Chamar o handler com ResponseWriter e Request
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Esperado status 422, obtido %d", w.Code)
	}
}

func TestWeatherHandler_ValidCEP(t *testing.T) {
	// Mock do serviço
	mockService := new(MockWeatherService)
	handler := controller.NewWeatherHandler(mockService) // Instanciar o handler corretamente

	req := httptest.NewRequest("GET", "/weather?cep=01001000", nil)
	w := httptest.NewRecorder()

	// Chamar o handler com ResponseWriter e Request
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperado status 200, obtido %d", w.Code)
	}
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
