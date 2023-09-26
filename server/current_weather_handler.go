package server

import (
	"encoding/json"
	"fmt"
	"github.com/robmurtha/weather-service/model"
	"net/http"
	"strconv"
)

func (s *Server) HandleGetCurrentWeather(res http.ResponseWriter, req *http.Request) {
	var err error
	var lat, lon string

	var in model.GetCurrentWeatherRequest
	var out *model.GetCurrentWeatherResponse

	lat = req.URL.Query().Get("lat")
	lon = req.URL.Query().Get("lon")

	if in.Latitude, err = strconv.ParseFloat(lat, 64); err != nil {
		http.Error(res, fmt.Sprintf("invalid latitude: "+err.Error()), http.StatusBadRequest)
		return
	}

	if in.Longitude, err = strconv.ParseFloat(lon, 64); err != nil {
		http.Error(res, fmt.Sprintf("invalid longitude: "+err.Error()), http.StatusBadRequest)
		return
	}

	out, err = s.weatherService.GetCurrent(req.Context(), &in)
	if err != nil || out.StatusCode == 0 {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(out)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(out.StatusCode)
	res.Write(body)
}
