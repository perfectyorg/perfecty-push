package internal

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/rwngallego/perfecty-push/internal/handlers"
	"net"
	"net/http"
	"time"
)

func StartServer() (err error) {
	mux := httprouter.New()

	//handlers
	mux.GET("/monitor", handlers.Monitor)
	mux.GET("/v1/public/users", handlers.PutPublicUsers)

	address := fmt.Sprintf("%s:%d", Config.Server.Host, Config.Server.Port)
	log.Info().Msg("Listening on " + address)

	server := http.Server{
		Addr:    address,
		Handler: logger(cors.Default().Handler(mux)),
	}

	if Config.Server.Ssl.Enabled == true {
		err = server.ListenAndServeTLS(Config.Server.Ssl.CertFile, Config.Server.Ssl.KeyFile)
	} else {
		err = server.ListenAndServe()
	}

	return
}

// Logging requests

const (
	TraceKey = "trace"
)

type (
	LoggerResponseWriter struct {
		http.ResponseWriter
		StatusCode int
	}
)

func NewLoggerResponseWriter(w http.ResponseWriter) (lrw *LoggerResponseWriter) {
	lrw = &LoggerResponseWriter{w, http.StatusOK}
	return
}

func (lrw *LoggerResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			start = time.Now()
			ua    = r.Header.Get("User-Agent")
			trace uuid.UUID
			ip    string
			err   error
		)

		if trace, err = uuid.NewRandom(); err != nil {
			log.Error().Err(err).Msg("Could not generate the trace")
		}

		if ip, _, err = net.SplitHostPort(r.RemoteAddr); err != nil {
			log.Error().Err(err).Str("Remote-Addr", r.RemoteAddr).Msg("Could not get the host part")
		}

		lrw := NewLoggerResponseWriter(w)
		ctx := context.WithValue(r.Context(), TraceKey, trace)
		h.ServeHTTP(lrw, r.WithContext(ctx))

		log.Info().
			Str("Trace", trace.String()).
			Str("Method", r.Method).
			Stringer("Url", r.URL).
			Dur("Duration", time.Since(start)).
			Str("Ip", ip).
			Str("User-Agent", ua).
			Int("Code", lrw.StatusCode).
			Msg("")
	})
}
