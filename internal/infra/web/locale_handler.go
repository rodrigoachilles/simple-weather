package web

import (
	"encoding/json"
	"fmt"
	"github.com/rodrigoachilles/simple-weather/internal/dto"
	"github.com/rodrigoachilles/simple-weather/internal/usecase"
	"github.com/rodrigoachilles/simple-weather/pkg/log"
	"net/http"
)

type LocaleHandler struct {
	localeFinder  usecase.Finder
	weatherFinder usecase.Finder
}

func NewLocaleHandler(localeFinder usecase.Finder, weatherFinder usecase.Finder) *LocaleHandler {
	return &LocaleHandler{
		localeFinder:  localeFinder,
		weatherFinder: weatherFinder,
	}
}

func (h *LocaleHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cep := r.PathValue("cep")
	if len(cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "invalid zipcode",
		})
		return
	}

	localeOutputRaw, err := h.localeFinder.Execute(cep)
	if err != nil {
		log.Logger.Error().Err(err).Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	localeOutput := localeOutputRaw.(*dto.LocaleOutput)
	if localeOutput.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusNotFound,
			Message:    "can not find zipcode",
		})
		return
	}

	weatherOutputRaw, err := h.weatherFinder.Execute(localeOutput.Localidade)
	if err != nil {
		log.Logger.Error().Err(err).Msg(err.Error())
		if err.Error() == "API key is invalid" || err.Error() == "API key is not provided" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
			})
			return
		}

		if err.Error() == "can not find zipcode" {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	weatherOutput := weatherOutputRaw.(*dto.WeatherOutput)

	w.WriteHeader(http.StatusOK)
	result := dto.ResultOutput{
		Locale: localeOutput.Localidade,
		TempC:  weatherOutput.Current.TempC,
		TempF:  weatherOutput.Current.TempF,
		TempK:  weatherOutput.Current.TempC + 273.15,
	}
	log.Logger.Info().Msg(fmt.Sprintf("%s", result))

	_ = json.NewEncoder(w).Encode(result)
}
