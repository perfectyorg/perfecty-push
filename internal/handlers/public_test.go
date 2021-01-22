package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCanRegister(t *testing.T) {
	req := httptest.NewRequest("PUT", "/v1/public/users", nil)
	rr := httptest.NewRecorder()
	router := httprouter.New()

	NewPublicHandlers(router, nil, nil)
	router.ServeHTTP(rr, req)

	body := PutPublicUsersResponse{}
	err := json.NewDecoder(rr.Body).Decode(&body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, true, body.Success)
	assert.NotEqual(t, nil, body.UUID)
}
