package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Seven4X/link/web/library/api"
	setup "github.com/Seven4X/link/web/library/echo"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	topicJson = `{"name":"t","refTopicId":12,"position":5}`
)

// todo mockdb
func TestCreateTopic(t *testing.T) {
	// Setup
	e := setup.NewEcho()

	req := httptest.NewRequest(http.MethodPost, "/topic", strings.NewReader(topicJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, newTopic(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		suc := api.SuccResp{}
		json.Unmarshal(rec.Body.Bytes(), &suc)
		println(rec.Body.String())
		assert.Equal(t, suc.Ok, 0)
	}
}
