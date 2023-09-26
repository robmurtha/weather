package server

import (
	"context"
	"github.com/robmurtha/weather-service/service"
)

type Server struct {
	ctx            context.Context
	weatherService service.Weather
}

func New(ctx context.Context, weatherService service.Weather) *Server {
	return &Server{
		ctx:            ctx,
		weatherService: weatherService,
	}
}
