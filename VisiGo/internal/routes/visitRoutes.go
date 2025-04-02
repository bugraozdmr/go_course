package routes

import (
	"VisiGo/internal/analytics/delivery"
	"VisiGo/internal/server"
)

type VisitRoutes struct {
	Server       *server.ServerStruct
	VisitUseCase delivery.VisitUseCase
}

func (v *VisitRoutes) VisitRoutes() {
	v.Server.Engine.POST("/visit", v.VisitUseCase.CreateVisitHandler)
	v.Server.Engine.GET("/visit/:id", v.VisitUseCase.FindVisitByIDHandler)
}

func NewVisitInit(server *server.ServerStruct, visitUseCase delivery.VisitUseCase) *VisitRoutes {
	return &VisitRoutes{
		Server:       server,
		VisitUseCase: visitUseCase,
	}
}
