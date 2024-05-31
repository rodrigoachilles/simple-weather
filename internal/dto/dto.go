package dto

import "fmt"

type OutputLocale struct {
	Localidade string `json:"localidade"`
}

type OutputWeather struct {
	Current CurrentWeather `json:"current"`
}

type CurrentWeather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

type OutputResult struct {
	Locale string  `json:"locale"`
	TempC  float64 `json:"temp_C"`
	TempF  float64 `json:"temp_F"`
	TempK  float64 `json:"temp_K"`
}

func (o OutputResult) String() string {
	return fmt.Sprintf("{ TempC: %f, TempF: %f, TempK: %f }", o.TempC, o.TempF, o.TempK)
}

type OutputError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
