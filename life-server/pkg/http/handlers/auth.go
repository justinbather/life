package handlers

import (
	"net/http"

	"github.com/justinbather/prettylog"
)

type authHandler struct {
	logger *prettylog.Logger
}

func NewAuthHandler(logger *prettylog.Logger) *authHandler {
	return &authHandler{logger}
}

func (h *authHandler) GetToken(w http.ResponseWriter, r *http.Request) {
}
