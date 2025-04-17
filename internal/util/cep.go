package util

import (
	"errors"
	"regexp"
)

func IsValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func ValidateCEP(cep string) (string, error) {
	if !IsValidCEP(cep) {
		return "", errors.New("invalid CEP")
	}
	return cep, nil
}
