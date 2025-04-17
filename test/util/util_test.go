package util_test

import (
	"go-weather/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCEP(t *testing.T) {
	valid := []string{"12345678", "01001000"}
	invalid := []string{"abc", "123", "123456789", "12-345-678"}

	for _, v := range valid {
		if !util.IsValidCEP(v) {
			t.Errorf("CEP %s deveria ser válido", v)
		}
	}

	for _, v := range invalid {
		if util.IsValidCEP(v) {
			t.Errorf("CEP %s deveria ser inválido", v)
		}
	}
}

func TestTemperatureConversions(t *testing.T) {
	tempC := 25.0
	expectedKelvin := 298.15

	resultKelvin := util.CelsiusToKelvin(tempC)

	// Use assert.InDelta para lidar com diferenças de precisão
	assert.InDelta(t, expectedKelvin, resultKelvin, 0.01, "A conversão para Kelvin deve ser precisa até duas casas decimais")
}
