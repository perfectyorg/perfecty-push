package perfecty

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerMonitor(t *testing.T) {
	req := httptest.NewRequest("GET", "/monitor", nil)
	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/monitor", Monitor)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
