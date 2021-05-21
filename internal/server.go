package internal

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/perfectyorg/perfecty-push/internal/application"
	"github.com/perfectyorg/perfecty-push/internal/handlers"
	"net/http"
)

func StartServer(rs *application.RegistrationService, ps *application.PreferenceService, ss *application.ScheduleService) (err error) {
	mux := httprouter.New()

	//handlers
	handlers.NewMonitorHandler(mux)
	handlers.NewPublicHandlers(mux, rs, ps)
	handlers.NewInternalHandlers(mux, ss)

	address := fmt.Sprintf("%s:%d", Config.Server.Host, Config.Server.Port)
	log.Info().Msg("Listening on " + address)

	server := http.Server{
		Addr:    address,
		Handler: LoggerMiddleware(cors.Default().Handler(mux)),
	}

	if Config.Server.Ssl.Enabled == true {
		err = server.ListenAndServeTLS(Config.Server.Ssl.CertFile, Config.Server.Ssl.KeyFile)
	} else {
		err = server.ListenAndServe()
	}

	return
}
