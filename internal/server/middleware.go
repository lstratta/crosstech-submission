package server

import (
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) middleware() {
	s.router.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: s.conf.AllowedOrigins,
			},
		),
	)
}
