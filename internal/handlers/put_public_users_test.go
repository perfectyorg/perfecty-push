package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutPublicUsers(t *testing.T) {
	req := httptest.NewRequest("PUT", "/v1/public/users", nil)
	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/v1/public/users", PutPublicUsers)
	router.ServeHTTP(rr, req)

	body := PutPublicUsersResponse{}
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, true, body.Success)
	assert.NotEqual(t, nil, body.UUID)
}
