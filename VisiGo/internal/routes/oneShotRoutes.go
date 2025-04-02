package routes

import (
	"VisiGo/internal/analytics/delivery"
	"VisiGo/internal/server"
)

type OneShotRoutes struct {
	Server         *server.ServerStruct
	OneShotUseCase delivery.OneShotUseCase
}

func (e *OneShotRoutes) OneShotRoutes() {
	e.Server.Engine.POST("/shot", e.OneShotUseCase.CreateOneShotHandler)
}

func NewOneShotInit(server *server.ServerStruct, OneShotUseCase delivery.OneShotUseCase) *OneShotRoutes {
	return &OneShotRoutes{
		Server:         server,
		OneShotUseCase: OneShotUseCase,
	}
}
