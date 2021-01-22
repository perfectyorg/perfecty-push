package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/rwngallego/perfecty-push/internal/application"
	"net/http"
)

type (
	PutPublicUsersResponse struct {
		UUID    uuid.UUID `json:"uuid"`
		Success bool      `json:"success"`
	}

	PublicHandler struct {
		registrationService *application.RegistrationService
		preferenceService   *application.PreferenceService
	}
)

func NewPublicHandlers(mux *httprouter.Router, rs *application.RegistrationService, ps *application.PreferenceService) {
	h := PublicHandler{
		registrationService: rs,
		preferenceService:   ps,
	}

	mux.PUT("/v1/public/users", h.PutUsers)
}

func (h *PublicHandler) PutUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		id  uuid.UUID
		err error
		b   []byte
	)
	if id, err = uuid.NewRandom(); err != nil {
		log.Error().Err(err).Msg("Cannot generate the UUID")
	}

	res := PutPublicUsersResponse{Success: true, UUID: id}
	b, err = json.Marshal(res)
	if err != nil {
		log.Error().Err(err).Msg("Cannot marshal response")
	}
	w.Write(b)
}
