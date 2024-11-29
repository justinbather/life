package handlers

import (
	"net/http"

	"github.com/justinbather/prettylog"
)

type healthHandler struct {
	logger *prettylog.Logger
}

func NewHealthHandler(logger *prettylog.Logger) *healthHandler {
	return &healthHandler{logger}
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("Health Check: Service Healthy")

	w.WriteHeader(http.StatusOK)
}
