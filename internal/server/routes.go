package server

func (s *Server) routes() {
	r := s.router

	r.GET("/ping", s.handlePing)
	r.GET("/tracks", s.handleGetAllTracks)
	r.GET("/signals", s.handleGetAllSignals)
	r.GET("/signals/:id", s.handleGetSignalsBySignalId)
}
