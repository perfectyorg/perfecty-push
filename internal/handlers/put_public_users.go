package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"net/http"
)

type PutPublicUsersResponse struct {
	UUID    uuid.UUID `json:"uuid"`
	Success bool      `json:"success"`
}

func PutPublicUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
