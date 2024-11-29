package handlers

import (
	"net/http"

	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type userHandler struct {
	service service.UserService
	logger  *prettylog.Logger
}

func NewUserHandler(service service.UserService, logger *prettylog.Logger) *userHandler {
	return &userHandler{service: service, logger: logger}
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *userHandler) Signup(w http.ResponseWriter, r *http.Request) {
	req, err := decode[UserRequest](r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 201, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	req, err := decode[UserRequest](r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByUsernameAndPass(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, "Username and/or password incorrect", http.StatusNotFound)
		return
	}

	err = encode(w, r, 201, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
