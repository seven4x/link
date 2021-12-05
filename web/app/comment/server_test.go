package comment

import (
	"github.com/Seven4X/link/web/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//https://echo.labstack.com/guide/testing
func TestServer_ListComment(t *testing.T) {
	e := app.SetupEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/:lid")
	context.SetParamNames("lid")
	context.SetParamValues("1")
	err := listComment(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	println(rec.Body.String())

}
