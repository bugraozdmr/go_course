package delivery

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type EventHandler struct {
	eventUseCase usecase.EventUseCase
}

type EventUseCase interface {
	CreateEventHandler(c echo.Context) error
	FindEventByIDHandler(c echo.Context) error
}

func (e *EventHandler) CreateEventHandler(c echo.Context) error {
	var event analytics.Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	ctx := c.Request().Context()
	if err := e.eventUseCase.CreateEvent(ctx, &event); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error in creating event",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "event created successfully",
	})
}

func (e *EventHandler) FindEventByIDHandler(c echo.Context) error {
	eventID := c.Param("id")
	event, err := e.eventUseCase.FindEventByID(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "event not found",
		})
	}
	return c.JSON(http.StatusOK, event)
}

func EventDelivery(eventUseCase usecase.EventUseCase) EventUseCase {
	return &EventHandler{
		eventUseCase: eventUseCase,
	}
}
