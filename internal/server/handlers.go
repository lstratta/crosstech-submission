package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Message string `json:"message"`
}

func (s *Server) handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "Pong!",
	})
}
