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
// these tests require improvement in that the database should have been mocked
// currently they use the real database, but the `make test` command automatically
// sets up and tears down the test containers

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

func (ts *TestSuite) TestHandleGetTracks_Returns200Tracks() {
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

	t := models.SetupThreeTracks()

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), models.TrackResponse{Tracks: t, Message: "request successful"}, r)
		assert.Equal(ts.T(), len(t), len(r.Tracks))
	}
}

func (ts *TestSuite) TestHandleGetTracksBySignalId_Returns200AllTracksWithThatId() {
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

	t := models.SetupThreeTracks()

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), len(t), len(r.Tracks))
		assert.Equal(ts.T(), 55, r.Tracks[0].TrackId)
		assert.Equal(ts.T(), 3247, r.Tracks[1].TrackId)
	}
}

func (ts *TestSuite) TestHandleGetSignals_Returns200AllSignals() {
	req := httptest.NewRequest(http.MethodGet, "/signals", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handleGetSignals(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.SignalResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	t := models.SetupThreeTracksEachWithFiveSignals()

	var sigs []models.Signal

	for _, t := range t {
		sigs = append(sigs, t.SignalIds...)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), len(sigs), len(r.Signals))
	}
}
