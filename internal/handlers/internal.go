package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rwngallego/perfecty-push/internal/application"
	"net/http"
)

type (
	InternalHandler struct {
		scheduleService *application.ScheduleService
	}
)

// NewInternalHandlers Register the internal handlers
func NewInternalHandlers(mux *httprouter.Router, scheduleService *application.ScheduleService) {
	h := InternalHandler{scheduleService: scheduleService}
	mux.Handle("GET", "/v1/notifications", h.getNotifications)
}

func (h *InternalHandler) getNotifications(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
