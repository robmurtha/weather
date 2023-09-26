package main

import (
	"context"
	"github.com/robmurtha/weather-service/open_weather"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/robmurtha/weather-service/server"
	"github.com/robmurtha/weather-service/service"
)

const (
	openWeatherKey      = "acafdf86e4ee8458b55f8beee8575104"
	openWeatherEndpoint = "https://api.openweathermap.org/data/2.5/weather"

	defaultLogLevel = "info"
)

func main() {
	var key, endPoint string
	logLevelString := defaultLogLevel

	if key = os.Getenv("WEATHER_SERVICE_OPEN_WEATHER_KEY"); key == "" {
		key = openWeatherKey
	}

	if key = os.Getenv("WEATHER_SERVICE_OPEN_WEATHER_ENDPOINT"); endPoint == "" {
		endPoint = openWeatherEndpoint
	}

	if l := os.Getenv("WEATHER_SERVICE_LOG_LEVEL"); l != "" {
		logLevelString = l
	}

	logLevel, lerr := zap.ParseAtomicLevel(logLevelString)
	if lerr != nil {
		panic(lerr)
	}

	logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), logLevel))
	defer logger.Sync()

	openWeatherClient := open_weather.New(http.DefaultClient, endPoint, key, logger)
	svr := server.New(context.Background(), service.New(openWeatherClient, logger))
	http.HandleFunc("/weather/current/hourly", svr.HandleGetCurrentWeather)
	address := ":8080"

	logger.Info("starting server", zap.String("address", address))
	err := http.ListenAndServe(address, nil)

	if err != nil {
		panic(err)
	}
}
