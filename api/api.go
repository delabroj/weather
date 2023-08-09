package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/delabroj/weather/models"
	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Config models.Config
	Logic  models.Logic
}

type Server struct {
	Mux *mux.Router
	cfg *ServerConfig
}

func NewServer(cfg *ServerConfig) *Server {
	srv := Server{
		Mux: mux.NewRouter(),
		cfg: cfg,
	}

	srv.addRoutes()

	return &srv
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	lrw := &loggingResponseWriter{w, http.StatusOK}
	requestURL := req.URL.String()

	srv.Mux.ServeHTTP(lrw, req)

	duration := time.Since(start)
	var durationString string
	durationus := float64(duration.Nanoseconds() / 1000)
	switch {
	case durationus < 1000:
		durationString = fmt.Sprintf("%.3gus", durationus)
	case durationus < 1000000:
		durationString = fmt.Sprintf("%.3gms", durationus/1000)
	default:
		durationString = fmt.Sprintf("%.3gs", durationus/1000000)
	}

	log.Printf(
		`HTTP: method=%s remote-addr="%s" status=%d duration=%v url=%s`,
		req.Method,
		req.RemoteAddr,
		lrw.statusCode,
		durationString,
		requestURL,
	)
}

// swagger:model APIError
type apiError struct {
	Error string `json:"error,omitempty"`
}

func WriteJSONError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	WriteJSON(w, apiError{err})
}

func WriteJSON(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	d, _ := json.Marshal(j)

	if _, err := w.Write(d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
