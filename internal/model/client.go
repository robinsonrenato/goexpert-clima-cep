package model

type ViaCEPClient interface {
	GetCityFromCEP(cep string) (string, error)
}

type WeatherClient interface {
	GetTemperatureByCity(city string) (float64, error)
}
