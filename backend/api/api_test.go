package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	api := &API{
		db: nil,
	}
	exp := `{"status":"OK"}` + "\n"
	// '\n' inside ` ` acts as '\\n'

	// Assertions
	if assert.NoError(t, api.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, exp, rec.Body.String())
	}
}
