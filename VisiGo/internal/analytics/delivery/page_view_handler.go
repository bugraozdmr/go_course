package delivery

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/analytics/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type PageViewHandler struct {
	pageViewUseCase usecase.PageViewUseCase
}

type PageViewUseCase interface {
	CreatePageViewHandler(c echo.Context) error
	FindPageViewByIDHandler(c echo.Context) error
}

func (p *PageViewHandler) CreatePageViewHandler(c echo.Context) error {
	var pageView analytics.PageView
	if err := c.Bind(&pageView); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	ctx := c.Request().Context()
	if err := p.pageViewUseCase.CreatePageView(ctx, &pageView); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error in creating page view",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "page view created successfully",
	})
}

func (p *PageViewHandler) FindPageViewByIDHandler(c echo.Context) error {
	pageViewID := c.Param("id")
	pageView, err := p.pageViewUseCase.FindPageViewByID(pageViewID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "page view not found",
		})
	}
	return c.JSON(http.StatusOK, pageView)
}

func PageViewDelivery(pageViewUseCase usecase.PageViewUseCase) PageViewUseCase {
	return &PageViewHandler{
		pageViewUseCase: pageViewUseCase,
	}
}
