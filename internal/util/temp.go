package util

func CelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func CelsiusToKelvin(tempC float64) float64 {
	return tempC + 273.15
}
