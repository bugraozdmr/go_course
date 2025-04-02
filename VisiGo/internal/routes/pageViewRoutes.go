package routes

import (
	"VisiGo/internal/analytics/delivery"
	"VisiGo/internal/server"
)

type PageViewRoutes struct {
	Server          *server.ServerStruct
	PageViewUseCase delivery.PageViewUseCase
}

func (p *PageViewRoutes) PageViewRoutes() {
	p.Server.Engine.POST("/pageview", p.PageViewUseCase.CreatePageViewHandler)
	p.Server.Engine.GET("/pageview/:id", p.PageViewUseCase.FindPageViewByIDHandler)
}

func NewPageViewInit(server *server.ServerStruct, pageViewUseCase delivery.PageViewUseCase) *PageViewRoutes {
	return &PageViewRoutes{
		Server:          server,
		PageViewUseCase: pageViewUseCase,
	}
}
