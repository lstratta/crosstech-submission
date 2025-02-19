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

func (s *Server) handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "Pong!",
	})
}

// handles route both with and without query param
func (s *Server) handleGetTracks(c echo.Context) error {
	sigId := c.QueryParam("signal-id")
	if sigId != "" {
		t, err := s.db.TracksBySignalId(sigId)
		if err != nil {
			return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks by signal id", nil)
		}

		return trackResponse(c, http.StatusOK, err, "request successful", t)

	}

	t, err := s.db.Tracks()
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks", nil)
	}

	return c.JSON(http.StatusOK, &models.TrackResponse{
		Tracks: t,
	})
}

func (s *Server) handleGetTrackByTrackId(c echo.Context) error {
	id := c.Param("id")
	t, err := s.db.TracksById(id)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error finding tracks by id", nil)
	}

	if len(t) < 1 {
		return trackResponse(c, http.StatusOK, err, "request successful, no data found", t)
	}

	return trackResponse(c, http.StatusOK, err, "request successful", t)

}

func (s *Server) handleGetSignals(c echo.Context) error {
	sig, err := s.db.Signals()
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error finding signals", nil)
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

func (s *Server) handleGetSignalBySignalId(c echo.Context) error {
	id := c.Param("id")
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

func (s *Server) handlePostTrack(c echo.Context) error {

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error reading request body", nil)
	}

	t := models.Track{}
	if err = json.Unmarshal(b, &t); err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error unmarshalling body json", nil)

	}

	res, err := s.db.CreateTrack(&t)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error creating track record", nil)

	}

	return trackResponse(c, http.StatusCreated, err, "successfully created", []models.Track{*res})

}

func (s *Server) handPostSignal(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error reading request body", nil)
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

func (s *Server) handleDeleteSignal(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error converting string to int", nil)
	}
	err = s.db.DeleteSignalById(id)
	if err != nil {
		return signalResponse(c, http.StatusInternalServerError, err, "error deleting signal by id", nil)
	}

	return signalResponse(c, http.StatusNoContent, nil, "delete successful", nil)
}

func (s *Server) handleDeleteTrack(c echo.Context) error {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error converting string to int", nil)
	}
	err = s.db.DeleteTrackById(id)
	if err != nil {
		return trackResponse(c, http.StatusInternalServerError, err, "error deleting track by id", nil)
	}

	return trackResponse(c, http.StatusNoContent, nil, "delete successful", nil)
}

func signalResponse(c echo.Context, status int, err error, message string, res []models.Signal) error {

	if err != nil {
		return c.JSON(status, &models.SignalResponse{
			Signals: []models.Signal{},
			Errors: []models.Errors{
				{
					Message: message,
					Detail:  err.Error(),
				},
			},
		})
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Message: message,
		Signals: res,
	})
}

func trackResponse(c echo.Context, status int, err error, message string, res []models.Track) error {

	if err != nil {
		return c.JSON(status, &models.TrackResponse{
			Tracks: []models.Track{},
			Errors: []models.Errors{
				{
					Message: message,
					Detail:  err.Error(),
				},
			},
		})
	}

	return c.JSON(http.StatusOK, &models.TrackResponse{
		Message: message,
		Tracks:  res,
	})
}
