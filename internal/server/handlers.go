package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/models"
)

type response struct {
	Message string `json:"message"`
}

// GET method calls

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
			return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
				Tracks: []models.Track{},
				Errors: []models.Errors{
					{
						Error:   "internal server error",
						Message: "there was an error with the server",
						Detail:  err.Error(),
					},
				},
			})
		}

		return c.JSON(http.StatusOK, &models.TrackResponse{
			Tracks: t,
		})
	}

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

func (s *Server) handleGetTrackByTrackId(c echo.Context) error {
	id := c.Param("id")
	t, err := s.db.TracksById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
			Tracks: []models.Track{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	if len(t) < 1 {
		return c.JSON(http.StatusOK, &models.TrackResponse{
			Tracks:  t,
			Message: "request OK, no data found",
		})
	}

	return c.JSON(http.StatusOK, &models.TrackResponse{
		Tracks: t,
	})
}

func (s *Server) handleGetSignals(c echo.Context) error {
	sig, err := s.db.Signals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.SignalResponse{
			Signals: []models.Signal{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

func (s *Server) handleGetSignalBySignalId(c echo.Context) error {
	id := c.Param("id")
	sig, err := s.db.SignalsById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.SignalResponse{
			Signals: []models.Signal{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	if len(sig) < 1 {
		return c.JSON(http.StatusOK, &models.SignalResponse{
			Signals: sig,
			Message: "request OK, no data found",
		})
	}

	return c.JSON(http.StatusOK, &models.SignalResponse{
		Signals: sig,
	})
}

// POST method calls

func (s *Server) handlePostTrack(c echo.Context) error {

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
			Tracks: []models.Track{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	t := models.Track{}
	if err = json.Unmarshal(b, &t); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
			Tracks: []models.Track{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	res, err := s.db.CreateTrack(&t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.TrackResponse{
			Tracks: []models.Track{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	return c.JSON(http.StatusOK, &models.TrackResponse{
		Message: "POST successful",
		Tracks: []models.Track{
			{TrackId: res.TrackId,
				Source:    res.Source,
				Target:    res.Target,
				SignalIds: res.SignalIds},
		},
	})
}

func (s *Server) handPostSignal(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return signalResponse(c, err, "internal server error", nil)
	}

	sig := models.Signal{}
	if err = json.Unmarshal(b, &sig); err != nil {
		return signalResponse(c, err, "internal server error", nil)
	}

	res, err := s.db.CreateSignal(&sig)
	if err != nil {
		return signalResponse(c, err, "internal server error", res)
	}

	return signalResponse(c, err, "successfully created", res)
}

func signalResponse(c echo.Context, err error, message string, res *models.Signal) error {

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.SignalResponse{
			Signals: []models.Signal{},
			Errors: []models.Errors{
				{
					Error:   "internal server error",
					Message: "there was an error with the server",
					Detail:  err.Error(),
				},
			},
		})
	}

	c.JSON(http.StatusOK, &models.SignalResponse{
		Message: "POST successful",
		Signals: []models.Signal{
			{
				SignalId:   res.SignalId,
				SignalName: res.SignalName,
				ELR:        res.ELR,
				Mileage:    res.Mileage,
			},
		},
	})
}
