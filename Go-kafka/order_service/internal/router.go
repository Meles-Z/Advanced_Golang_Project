package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/meles-z/go-kafa/order_service/internal/app/handlers"
)

func Router(e *echo.Echo, s *Server) {
	orderGroup := e.Group("/order")
	orderGroup.POST("/create", handers.)
}
