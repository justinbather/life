package handlers

import (
	"net/http"

	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type userHandler struct {
	userService service.UserService
	authService service.AuthService
	logger      *prettylog.Logger
}

func NewUserHandler(userService service.UserService, authService service.AuthService, logger *prettylog.Logger) *userHandler {
	return &userHandler{userService: userService, authService: authService, logger: logger}
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

func (h *userHandler) Signup(w http.ResponseWriter, r *http.Request) {
	req, err := decode[UserRequest](r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req.Username, req.Password)
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

	user, err := h.userService.GetUserByUsernameAndPass(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, "Username and/or password incorrect", http.StatusNotFound)
		return
	}

	signed, expires, err := h.authService.CreateToken(user.Id)
	if err != nil {
		http.Error(w, "TODO: Error signing key", http.StatusInternalServerError)
		return
	}

	response := UserResponse{Token: signed, Expires: expires}

	err = encode(w, r, 201, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
