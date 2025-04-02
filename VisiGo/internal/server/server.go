package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ServerStruct struct {
	Engine *echo.Echo
}

func (s *ServerStruct) StartServer() {
	s.Engine.Start("localhost:8080")
}

func NewHTTPServer() *ServerStruct {
	e := echo.New()

	// CORS middleware ekleniyor
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:80"}, // Frontend'iniz nerede çalışıyorsa o domaini ekleyin
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderContentType, "Authorization", "Referer"},
		AllowCredentials: true, // Kimlik doğrulama bilgilerini (cookies, headers) geçmesine izin ver
	}))

	// Burada handler ekleyebilirsiniz
	e.POST("/shot", func(c echo.Context) error {
		return c.String(200, `{"status": "success"}`)
	})

	return &ServerStruct{
		Engine: e,
	}
}
