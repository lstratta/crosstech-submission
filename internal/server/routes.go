// simply holds all the routes for endpoints
package server

func (s *Server) routes() {
	r := s.router

	r.GET("/ping", s.handlePing)
	r.GET("/tracks", s.handleGetTracks)
	r.GET("/tracks/:id", s.handleGetTrackByTrackId)
	r.GET("/signals", s.handleGetSignals)
	r.GET("/signals/:id", s.handleGetSignalBySignalId)

	r.POST("/tracks", s.handlePostTrack)
	r.POST("/signals", s.handlePostSignal)

	r.PUT("/tracks", s.handleUpdateTrack)
	r.PUT("/signals", s.handleUpdateSignal)

	r.DELETE("/tracks/:id", s.handleDeleteTrackById)
	r.DELETE("/signals/:id", s.handleDeleteSignalById)
}
