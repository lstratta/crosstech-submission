package server

import (
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) middleware() {
	s.router.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: s.conf.AllowedOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
				AllowHeaders: []string{"Accept", "Authorisation", "Content-Type", "X-CSRF-Token"},
			},
		),
		middleware.Logger(),
		middleware.Recover(),
	)
}
