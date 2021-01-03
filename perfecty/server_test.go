package perfecty

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rwngallego/perfecty-push/perfecty/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerMonitor(t *testing.T) {
	req := httptest.NewRequest("GET", "/monitor", nil)
	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/monitor", handlers.Monitor)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
