package server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/lstratta/crosstech-submission/internal/config"
)

func (s *Server) middleware(c config.Config) {
	s.router.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: c.AllowedOrigins,
			},
		),
	)
}
