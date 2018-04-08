package app

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type requestContextKey int

const (
	// LoggerKey is the key at which to store the logger in request context
	LoggerKey requestContextKey = iota
)

// Middleware is a factory for creating a middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Level = logrus.InfoLevel

		ctx := context.WithValue(r.Context(), LoggerKey, logger)

		next.ServeHTTP(w, r.WithContext(ctx))

		logger.Infof("Completed request\n")
	})
}
