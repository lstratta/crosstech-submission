package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

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
