package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerMonitorReturnsOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/monitor", nil)
	rr := httptest.NewRecorder()

	router := httprouter.New()
	NewMonitorHandler(router)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
