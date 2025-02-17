package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Routes() {
	s.router.GET("/", home)
}

type response struct {
	Message string `json:"message"`
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "Hello, from the router!",
	})
}
