package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// NewMonitorHandler Register the monitor handler
func NewMonitorHandler(mux *httprouter.Router) {
	mux.GET("/monitor", Monitor)
}

func Monitor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"running": true}`)
}
