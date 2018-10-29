package app

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type requestContextKey int

const (
	// LoggerKey is the key at which to store the logger in request context
	LoggerKey requestContextKey = iota
)

// Middleware is a factory for creating a middleware
func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Level = logrus.InfoLevel

		ctx := context.WithValue(r.Context(), LoggerKey, logger)

		next(w, r.WithContext(ctx), ps)

		logger.Infof("Completed request\n")
	}
}
