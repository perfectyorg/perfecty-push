package internal

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"os"
	"time"
)

// Logging

func LoadLogging() (err error) {
	var (
		host  string
		level zerolog.Level
	)
	host, _ = os.Hostname()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if Config.Logging.Level != "" {
		if level, err = zerolog.ParseLevel(Config.Logging.Level); err != nil {
			return
		}
		zerolog.SetGlobalLevel(level)
	}

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("host", host).
		Logger()
	if Config.Logging.Pretty == true {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Logger = logger

	return
}

// Logging middleware

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

func LoggerMiddleware(h http.Handler) http.Handler {
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
