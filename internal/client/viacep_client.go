package client

import (
	"fmt"
	"go-weather/internal/model"

	"github.com/go-resty/resty/v2"
)

type ViaCEPClientImpl struct{}

type viaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

func (v *ViaCEPClientImpl) GetCityFromCEP(cep string) (string, error) {
	client := resty.New()
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp := viaCEPResponse{}
	res, err := client.R().SetResult(&resp).Get(url)

	if err != nil || res.StatusCode() != 200 || resp.Erro {
		return "", fmt.Errorf("not found")
	}

	return resp.Localidade, nil
}

func NewViaCEPClient() model.ViaCEPClient {
	return &ViaCEPClientImpl{}
}
