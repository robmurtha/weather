package open_weather

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"path"
)

// Client provides access to the OpenWeather REST API.
// See https://openweathermap.org/api
type Client interface {
	GetCurrentHourlyWeatherByCoord(context.Context, Coord) (*CurrentConditions, *http.Response, error)
}

type client struct {
	httpClient *http.Client
	endpoint   string
	appID      string

	logger *zap.Logger
}

// New accepts an http client, OpenWeather endpoint and api key and returns a client ready to use.
func New(httpClient *http.Client, endPoint, appID string, logger *zap.Logger) Client {
	return &client{
		httpClient: httpClient,
		endpoint:   endPoint,
		appID:      appID,
		logger:     logger,
	}
}

// GetCurrentHourlyWeatherByCoord returns the current hourly weather conditions for the given location.
// Coord.Lat must be a valid latitude in the range of (-90:90)
// Coord.Lon must be a valid longitude in the range of (-180:180)
func (c *client) GetCurrentHourlyWeatherByCoord(ctx context.Context, coord Coord) (*CurrentConditions, *http.Response, error) {
	var err error
	var result CurrentConditions

	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s",
		path.Join(c.endpoint, "data/2.5/weather"),
		coord.Lat,
		coord.Lon,
		c.appID)

	body := &bytes.Buffer{}
	req, rerr := http.NewRequestWithContext(ctx, http.MethodGet, url, body)
	if rerr != nil {
		err = fmt.Errorf("request error: %w", rerr)
		c.logger.Panic("request error",
			zap.Float64("lat", coord.Lat),
			zap.Float64("lon", coord.Lon),
			zap.String("url", url),
			zap.Error(rerr))
	}

	resp, cerr := c.httpClient.Do(req)
	if cerr != nil {
		err = fmt.Errorf("unable to fetch weather data: %w", err)

		c.logger.Error("unable to fetch weather data",
			zap.Float64("lat", coord.Lat),
			zap.Float64("lon", coord.Lon),
			zap.String("url", url),
			zap.Any("resp", resp),
			zap.Error(cerr),
		)
	}

	if merr := json.Unmarshal(body.Bytes(), &result); merr != nil {
		c.logger.Error("unable tp process OpenWeather API response body",
			zap.Float64("lat", coord.Lat),
			zap.Float64("lon", coord.Lon),
			zap.String("url", url),
			zap.String("body", body.String()),
			zap.Error(merr),
		)
		return nil, resp, errors.New("unable tp process OpenWeather API response body")
	}

	return &result, resp, nil
}
