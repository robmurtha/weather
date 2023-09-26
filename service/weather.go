package service

import (
	"context"
	"github.com/robmurtha/weather-service/model"
	"github.com/robmurtha/weather-service/open_weather"
	"go.uber.org/zap"
)

type Weather interface {
	GetCurrent(context.Context, *model.GetCurrentWeatherRequest) (*model.GetCurrentWeatherResponse, error)
}

type weather struct {
	client open_weather.Client
	logger *zap.Logger
}

func New(openWeatherClient open_weather.Client, logger *zap.Logger) Weather {
	return &weather{
		client: openWeatherClient,
		logger: logger,
	}
}

func (w *weather) GetCurrent(ctx context.Context, request *model.GetCurrentWeatherRequest) (*model.GetCurrentWeatherResponse, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	currentConditions, resp, err := w.client.GetCurrentHourlyWeatherByCoord(
		ctx,
		open_weather.Coord{
			Lat: request.Latitude,
			Lon: request.Longitude,
		})

	result := model.GetCurrentWeatherResponse{
		CurrentConditions: currentConditions,
		Status:            resp.Status,
		StatusCode:        resp.StatusCode,
	}

	return &result, err
}
