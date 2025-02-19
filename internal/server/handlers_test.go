package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/models"
	"github.com/stretchr/testify/assert"
)

// basic set of tests for all the handlers

func (ts *TestSuite) TestHandlePing_ReturnsPongAsMessageInJson() {

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handlePing(c)

	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := &response{}
	json.Unmarshal(reqBody, r)

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), &response{Message: "Pong!"}, r)
	}
}

func (ts *TestSuite) TestHandleGetTracks_ReturnsTracks() {
	req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handleGetTracks(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.TrackResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), models.TrackResponse{Tracks: models.SetupThreeTracks(), Message: "request successful"}, r)
	}
}

func (ts *TestSuite) TestHandleGetTracksBySignalId_ReturnsAllTracksWithThatId() {
	req := httptest.NewRequest(http.MethodGet, "/tracks?query-id=453", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handleGetTracks(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.TrackResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), 55, r.Tracks[0].TrackId)
		assert.Equal(ts.T(), 3247, r.Tracks[1].TrackId)
	}
}
