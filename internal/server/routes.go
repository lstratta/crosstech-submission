package server

func (s *Server) routes() {
	r := s.router

	r.GET("/ping", s.handlePing)
	r.GET("/tracks", s.handleGetTracks)
	r.GET("/tracks/:id", s.handleGetTrackByTrackId)
	r.GET("/signals", s.handleGetSignals)
	r.GET("/signals/:id", s.handleGetSignalBySignalId)

	r.POST("/tracks", s.handlePostTrack)
	r.POST("/signals", s.handPostSignal)

	r.PUT("/tracks", s.handleUpdateTrack)
	r.PUT("/signals", s.handleUpdateSignal)

	r.DELETE("/tracks/:id", s.handleDeleteTrack)
	r.DELETE("/signals/:id", s.handleDeleteSignal)
}
