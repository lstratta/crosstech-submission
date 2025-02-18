package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/models"
)

type response struct {
	Message string `json:"message"`
}

func (s *Server) handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "Pong!",
	})
}

func (s *Server) handleGetAllTracks(c echo.Context) error {
	t, err := s.db.Tracks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
			Tracks: []models.Track{},
		})
	}

	return c.JSON(http.StatusOK, &models.TrackResponse{
		Tracks: t,
	})
}

func (s *Server) handleGetAllSignals(c echo.Context) error {
	sig, err := s.db.Signals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.SignalResponse{
			Signals: []models.Signal{},
		})
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

func (s *Server) handleGetSignalsBySignalId(c echo.Context) error {
	id := c.Param("id")
	sig, err := s.db.SignalsById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.SignalResponse{
			Signals: []models.Signal{},
		})
	}

	if len(sig) < 1 {
		return c.JSON(http.StatusNotFound, &models.SignalResponse{
			Signals: sig,
		})
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

// func (s *Server) handleGetTracksBySignalId(c echo.Context) error {
// 	// id := c.Param("id")
// 	// t, err := s.db.TracksBySignalId(id)
// }
