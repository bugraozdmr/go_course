package delivery

import (
	"VisiGo/internal/analytics/repository"
	"VisiGo/internal/analytics/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type OneShotHandler struct {
	OneShotUseCase usecase.OneShotUseCase
}

type OneShotUseCase interface {
	CreateOneShotHandler(c echo.Context) error
}

func (e *OneShotHandler) CreateOneShotHandler(c echo.Context) error {
	//! this violates the clean architecture !
	var OneShot repository.OneShot
	if err := c.Bind(&OneShot); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	ctx := c.Request().Context()
	if err := e.OneShotUseCase.CreateAll(ctx, &OneShot); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Created successfully",
	})
}

func OneShotDelivery(oneShotUseCase usecase.OneShotUseCase) OneShotUseCase {
	return &OneShotHandler{
		OneShotUseCase: oneShotUseCase,
	}
}