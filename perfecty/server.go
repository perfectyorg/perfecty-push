package perfecty

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/rwngallego/perfecty-push/perfecty/handlers"
	"net"
	"net/http"
	"time"
)

func startServer() (err error) {
	router := httprouter.New()

	//routes
	router.GET("/monitor", handlers.Monitor)

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Info().Msg("Listening on " + address)

	server := http.Server{
		Addr:    address,
		Handler: &Logger{router},
	}

	if config.Server.Ssl.Enabled == true {
		err = server.ListenAndServeTLS(config.Server.Ssl.CertFile, config.Server.Ssl.KeyFile)
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
	Logger struct {
		handler http.Handler
	}

	LoggerResponseWriter struct {
		http.ResponseWriter
		StatusCode int
	}
)

func NewLoggerRW(w http.ResponseWriter) (loggerRw *LoggerResponseWriter) {
	loggerRw = &LoggerResponseWriter{w, http.StatusOK}
	return
}

func (loggerRw *LoggerResponseWriter) WriteHeader(code int) {
	loggerRw.StatusCode = code
	loggerRw.ResponseWriter.WriteHeader(code)
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	loggerRw := NewLoggerRW(w)
	ctx := context.WithValue(r.Context(), TraceKey, trace)
	l.handler.ServeHTTP(loggerRw, r.WithContext(ctx))

	log.Info().
		Str("Trace", trace.String()).
		Str("Method", r.Method).
		Stringer("Url", r.URL).
		Dur("Duration", time.Since(start)).
		Str("Ip", ip).
		Str("User-Agent", ua).
		Int("Code", loggerRw.StatusCode).
		Msg("")
}
