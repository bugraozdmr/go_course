package delivery

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type VisitHandler struct {
	visitUseCase usecase.VisitUseCase
}

type VisitUseCase interface {
	CreateVisitHandler(c echo.Context) error
	FindVisitByIDHandler(c echo.Context) error
}

func (v *VisitHandler) CreateVisitHandler(c echo.Context) error {
	var visit analytics.Visit
	if err := c.Bind(&visit); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	ctx := c.Request().Context()
	if err := v.visitUseCase.CreateVisit(ctx, &visit); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error in creating visit",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "visit created successfully",
	})
}

func (v *VisitHandler) FindVisitByIDHandler(c echo.Context) error {
	visitID := c.Param("id")
	visit, err := v.visitUseCase.FindVisitByID(visitID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "visit not found",
		})
	}
	return c.JSON(http.StatusOK, visit)
}

func VisitDelivery(visitUseCase usecase.VisitUseCase) VisitUseCase {
	return &VisitHandler{
		visitUseCase: visitUseCase,
	}
}
