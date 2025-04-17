package controller

import (
	"encoding/json"
	"go-weather/internal/service"
	"go-weather/internal/util"
	"net/http"
)

type WeatherHandler struct {
	Service service.WeatherService
}

func NewWeatherHandler(service service.WeatherService) http.HandlerFunc {
	h := &WeatherHandler{Service: service}
	return h.Handle
}

func (h *WeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if !util.IsValidCEP(cep) {
		http.Error(w, `{"message":"invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := h.Service.GetWeatherByCEP(r.Context(), cep)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, `{"message":"can not find zipcode"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"message":"internal error"}`, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
