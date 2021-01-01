package perfecty

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Monitor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"running": true}`)
}

func StartServer() (err error) {
	router := httprouter.New()
	router.GET("/monitor", Monitor)

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	fmt.Printf("Listening on %s\n", address)

	err = http.ListenAndServe(address, router)
	return
}
