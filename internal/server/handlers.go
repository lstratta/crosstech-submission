// handlers.go holds all the handler functions that then interact with the database
// Usually there would be another layer that handles business logic
// but there wasn't a need in this insance
package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/models"
)

type response struct {
	Message string `json:"message"`
}

// GET methods

// Test function
func (s *Server) handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "Pong!",
	})
}

// Returns all tracks
// handles route both with and without query param
func (s *Server) handleGetTracks(c echo.Context) error {
	p := c.QueryParam("signal-id")
	if p != "" {
		id, err := strconv.Atoi(p)
		if err != nil {
			return trackResponse(c, http.StatusBadRequest, err, "cannot convert query param to integer", nil)
		}

		t, err := s.db.TracksBySignalId(id)
		if err != nil {
			return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks by signal id", nil)
		}

		return trackResponse(c, http.StatusOK, err, "request successful", t)

	}

	t, err := s.db.Tracks()
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks", nil)
	}

	return trackResponse(c, http.StatusOK, nil, "request successful", t)

}

// Returns all tracks with a specific ID
func (s *Server) handleGetTrackByTrackId(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return trackResponse(c, http.StatusBadRequest, err, "make sure param is an integer", nil)
	}

	t, err := s.db.TracksById(id)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks by id", nil)
	}

	if len(t) < 1 {
		return trackResponse(c, http.StatusOK, err, "request successful, no data found", t)
	}

	return trackResponse(c, http.StatusOK, err, "request successful", t)

}

// Returns a slice of signals
func (s *Server) handleGetSignals(c echo.Context) error {
	sig, err := s.db.Signals()
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error finding signals", nil)
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

// Returns all signals with a specific ID
func (s *Server) handleGetSignalBySignalId(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return trackResponse(c, http.StatusBadRequest, err, "make sure param is an integer", nil)
	}
	sig, err := s.db.SignalsById(id)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error finding signals by id", nil)
	}

	if len(sig) < 1 {
		return signalResponse(c, http.StatusOK, err, "request successful, but no data found", sig)
	}

	return signalResponse(c, http.StatusOK, err, "request successful", sig)

}

// POST methods

// Creates a new track
func (s *Server) handlePostTrack(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return trackResponse(c, http.StatusBadRequest, err, "error reading request body", nil)
	}

	t := models.Track{}
	if err = json.Unmarshal(b, &t); err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error unmarshalling body json", nil)

	}

	res, err := s.db.CreateTrackWithSignals(&t)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error creating track record", nil)

	}

	return trackResponse(c, http.StatusCreated, err, "successfully created", []models.Track{*res})

}

// Creates a new signal
func (s *Server) handlePostSignal(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return signalResponse(c, http.StatusBadRequest, err, "error reading request body", nil)
	}

	sig := models.Signal{}
	if err = json.Unmarshal(b, &sig); err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error unmarshalling json", nil)
	}

	res, err := s.db.CreateSignal(&sig)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error creating signal record", nil)
	}

	return signalResponse(c, http.StatusCreated, err, "successfully created", []models.Signal{*res})
}

// PUT methods

// Updates a signal
func (s *Server) handleUpdateSignal(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error reading request body", nil)
	}
	sig := models.Signal{}
	if err = json.Unmarshal(b, &sig); err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error unmarshalling json", nil)
	}

	res, err := s.db.UpdateSignal(&sig)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error creating signal record", nil)
	}

	return signalResponse(c, http.StatusOK, nil, "update successful", []models.Signal{*res})
}

// Updates a track
func (s *Server) handleUpdateTrack(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error reading request body", nil)
	}
	t := models.Track{}
	if err = json.Unmarshal(b, &t); err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error unmarshalling json", nil)
	}

	res, err := s.db.UpdateTrack(&t)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error creating signal record", nil)
	}

	return trackResponse(c, http.StatusOK, nil, "update successful", []models.Track{*res})
}

// DELETE methods

// Deletes a signal
func (s *Server) handleDeleteSignalById(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return signalResponse(c, http.StatusBadRequest, err, "error converting string to int", nil)
	}
	err = s.db.DeleteSignalById(id)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error deleting signal by id", nil)
	}

	return signalResponse(c, http.StatusOK, nil, "delete successful", nil)
}

// Deletes a track
func (s *Server) handleDeleteTrackById(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return trackResponse(c, http.StatusBadRequest, err, "error converting string to int", nil)
	}
	err = s.db.DeleteTrackById(id)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error deleting track by id", nil)
	}

	return trackResponse(c, http.StatusOK, nil, "delete successful", nil)
}

// wrapper for responding to signal requests
func signalResponse(c echo.Context, status int, err error, message string, res []models.Signal) error {
	if err != nil {
		return c.JSON(status, &models.SignalResponse{
			Signals: []models.Signal{},
			Message: message,
			Error:   err.Error(),
		})
	}

	return c.JSON(status, &models.SignalResponse{
		Message: message,
		Signals: res,
	})
}

// wrapper for responding to track requests
func trackResponse(c echo.Context, status int, err error, message string, res []models.Track) error {
	if err != nil {
		return c.JSON(status, &models.TrackResponse{
			Tracks:  []models.Track{},
			Message: message,
			Error:   err.Error(),
		})
	}

	return c.JSON(status, &models.TrackResponse{
		Message: message,
		Tracks:  res,
	})
}
