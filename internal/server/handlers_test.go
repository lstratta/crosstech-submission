package server

import (
	"bytes"
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

// the suite also needs tests that try to break the application
// as well as edge case tests

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
		// assert.Equal(ts.T(), models.TrackResponse{Tracks: t, Message: "request successful"}, r)
		assert.Equal(ts.T(), len(t), len(r.Tracks))
	}
}

func (ts *TestSuite) TestHandleGetTrackByTrackId_Returns200AndSignal() {
	req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)
	c.SetPath("/tracks/:id")
	c.SetParamNames("id")
	c.SetParamValues("55")

	h := ts.srv.handleGetTrackByTrackId(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.TrackResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	t := models.SetupThreeTracksEachWithFiveSignals()

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), 1, len(r.Tracks))
		assert.Equal(ts.T(), t[0].SignalIds[0].SignalId, r.Tracks[0].SignalIds[0].SignalId)
	}
}

func (ts *TestSuite) TestHandleGetSignalBySignalId_Returns200AndSignal() {
	req := httptest.NewRequest(http.MethodGet, "/signals", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)
	c.SetPath("/signals/:id")
	c.SetParamNames("id")
	c.SetParamValues("453")

	h := ts.srv.handleGetSignalBySignalId(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.SignalResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), 2, len(r.Signals))
		assert.Equal(ts.T(), 453, r.Signals[0].SignalId)
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

func (ts *TestSuite) TestHandlePostTrack_Returns201CreatedRecord() {
	trk := &models.Track{
		TrackId: 92774,
		Source:  "Test Station 3",
		Target:  "Test Station 4",
		SignalIds: []models.Signal{
			{
				SignalId:   13393,
				SignalName: "SIG:WM791(CO)WEMBLEY CENTRAL STN",
				ELR:        "MFD1",
				Mileage:    8.3815,
			},
			{
				SignalId:   13399,
				SignalName: "SIG:WM1252(PL)WEMBLEY CENTRAL STN",
				ELR:        "XGF1",
				Mileage:    2.9309,
			},
		},
	}

	d, err := json.Marshal(trk)
	if err != nil {
		ts.T().Errorf("error marshalling data: %s", err)
	}
	rd := bytes.NewReader(d)

	req := httptest.NewRequest(http.MethodPost, "/tracks", rd)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handlePostTrack(c)

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
		assert.Equal(ts.T(), http.StatusCreated, res.Code)
		assert.Equal(ts.T(), 1, len(r.Tracks))
		assert.Equal(ts.T(), *trk, r.Tracks[0])
	}
}

func (ts *TestSuite) TestHandlePostSignal_Returns201CreatedRecord() {
	sig := &models.Signal{
		SignalId:   99999,
		SignalName: "SIG:WM791(CO)WEMBLEY CENTRAL STN",
		ELR:        "AAA3",
		Mileage:    1.2345,
	}

	d, err := json.Marshal(sig)
	if err != nil {
		ts.T().Errorf("error marshalling data: %s", err)
	}
	rd := bytes.NewReader(d)

	req := httptest.NewRequest(http.MethodPost, "/signals", rd)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handlePostSignal(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.SignalResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusCreated, res.Code)
		assert.Equal(ts.T(), 1, len(r.Signals))
		assert.Equal(ts.T(), *sig, r.Signals[0])
	}
}

func (ts *TestSuite) TestHandlePutSignal_Returns200AndCreatedRecord() {
	sig := &models.Signal{
		SignalId:   453,
		SignalName: "SIG:AW148(CO) ACTON WELLS JCN -- UPDATED",
		ELR:        "LPC5",
		Mileage:    3.1745,
	}

	d, err := json.Marshal(sig)
	if err != nil {
		ts.T().Errorf("error marshalling data: %s", err)
	}
	rd := bytes.NewReader(d)

	req := httptest.NewRequest(http.MethodPut, "/signals", rd)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handleUpdateSignal(c)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		ts.T().Errorf("error reading response body: %s", err)
	}

	r := models.SignalResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		ts.T().Errorf("error unmashalling json: %s", err)
	}

	if assert.NoError(ts.T(), h) {
		assert.Equal(ts.T(), http.StatusOK, res.Code)
		assert.Equal(ts.T(), sig.SignalName, r.Signals[0].SignalName)
	}
}

func (ts *TestSuite) TestHandlePutTrack_Returns200AndCreatedRecord() {
	trk := &models.Track{
		TrackId: 55,
		Source:  "Acton Central -- UPDATED",
		Target:  "Willesden Junction",
	}

	d, err := json.Marshal(trk)
	if err != nil {
		ts.T().Errorf("error marshalling data: %s", err)
	}
	rd := bytes.NewReader(d)

	req := httptest.NewRequest(http.MethodPut, "/tracks", rd)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)

	h := ts.srv.handleUpdateTrack(c)

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
		assert.Equal(ts.T(), trk.Source, r.Tracks[0].Source)
	}
}

func (ts *TestSuite) TestHandleDeleteTrack_Returns200AndMessage() {
	req := httptest.NewRequest(http.MethodDelete, "/tracks/", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)
	c.SetPath("/tracks/:id")
	c.SetParamNames("id")
	c.SetParamValues("55")

	h := ts.srv.handleDeleteTrackById(c)

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
		assert.Equal(ts.T(), "delete successful", r.Message)
		assert.Equal(ts.T(), "", r.Error)
	}
}

func (ts *TestSuite) TestHandleDeleteSignal_Returns200AndMessage() {
	req := httptest.NewRequest(http.MethodDelete, "/signals/", nil)
	res := httptest.NewRecorder()

	c := echo.New().NewContext(req, res)
	c.SetPath("/signals/:id")
	c.SetParamNames("id")
	c.SetParamValues("453")

	h := ts.srv.handleDeleteSignalById(c)

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
		assert.Equal(ts.T(), "delete successful", r.Message)
	}
}
