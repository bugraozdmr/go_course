package routes

import (
	"VisiGo/internal/analytics/delivery"
	"VisiGo/internal/server"
)

type EventRoutes struct {
	Server       *server.ServerStruct
	EventUseCase delivery.EventUseCase
}

func (e *EventRoutes) EventRoutes() {
	e.Server.Engine.POST("/event", e.EventUseCase.CreateEventHandler)
	e.Server.Engine.GET("/event/:id", e.EventUseCase.FindEventByIDHandler)
}

func NewEventInit(server *server.ServerStruct, eventUseCase delivery.EventUseCase) *EventRoutes {
	return &EventRoutes{
		Server:       server,
		EventUseCase: eventUseCase,
	}
}
