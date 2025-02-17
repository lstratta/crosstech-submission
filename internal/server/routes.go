package server

func (s *Server) routes() {
	r := s.router

	r.GET("/ping", s.handlePing)
}
