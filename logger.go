package main

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		slog.Info("response",
			"req_method", r.Method,
			"req_ip", r.RemoteAddr,
			"req_path", r.RequestURI,
			"size", r.ContentLength,
			"duration", time.Since(start).Milliseconds(),
		)
	}()
	l.handler.ServeHTTP(w, r)
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
