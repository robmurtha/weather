package model

import "github.com/robmurtha/weather-service/open_weather"

type GetCurrentWeatherRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type GetCurrentWeatherResponse struct {
	CurrentConditions *open_weather.CurrentConditions `json:"currentConditions,omitempty"`
	Alerts            []string                        `json:"alerts,omitempty"`

	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Error      error  `json:"error,omitempty"`
}

func (req GetCurrentWeatherRequest) IsValid() bool {
	if req.Latitude == 0 || req.Longitude == 0 {
		return false
	}

	return true
}
