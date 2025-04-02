package di

import (
	"VisiGo/internal/analytics/delivery"
	"VisiGo/internal/analytics/repository"
	"VisiGo/internal/analytics/usecase"
	"VisiGo/internal/config"
	db "VisiGo/internal/database"
	"VisiGo/internal/routes"
	"VisiGo/internal/server"
)

func Init(conf config.Config) *server.ServerStruct {
	// 1. Initialize the server
	server := server.NewHTTPServer()

	// 2. Initialize the database connection
	database := db.ConnectPGDB(conf)

	// 3. Initialize repositories for Visit, PageView, Event
	visitRepo := repository.NewVisitRepository(database)
	pageViewRepo := repository.NewPageViewRepository(database)
	eventRepo := repository.NewEventRepository(database)
	oneShotRepo := repository.NewOneShotRepository(database)

	// 4. Initialize the use cases for Visit, PageView, Event
	visitUseCase := usecase.VisitUseC(visitRepo)
	pageViewUseCase := usecase.PageViewUseC(pageViewRepo)
	eventUseCase := usecase.EventUseC(eventRepo)
	oneShotUseCase := usecase.OneShotUseC(oneShotRepo)

	// 5. Initialize the delivery layer (handlers) with the use cases
	visitDelivery := delivery.VisitDelivery(visitUseCase)
	pageViewDelivery := delivery.PageViewDelivery(pageViewUseCase)
	eventDelivery := delivery.EventDelivery(eventUseCase)
	oneShotDelivery := delivery.OneShotDelivery(oneShotUseCase)

	// 6. Initialize the routes and pass the server and delivery handlers
	visitRoutes := routes.NewVisitInit(server, visitDelivery)
	pageViewRoutes := routes.NewPageViewInit(server, pageViewDelivery)
	eventRoutes := routes.NewEventInit(server, eventDelivery)
	oneShotRoutes := routes.NewOneShotInit(server, oneShotDelivery)

	// 7. Set up the routes (Visit, PageView, Event)
	visitRoutes.VisitRoutes()
	pageViewRoutes.PageViewRoutes()
	eventRoutes.EventRoutes()
	oneShotRoutes.OneShotRoutes()

	// 8. Return the initialized server with routes setup
	return server
}
