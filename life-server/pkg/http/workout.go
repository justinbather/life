package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/service"
)

type WorkoutHandler struct {
	Service *service.WorkoutService
}

func (h *WorkoutHandler) WorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutType := params["type"]

	h.Service.Logger.Infof("Got workout handler request for type %s", workoutType)
	h.Service.Logger.Info(workoutType)
}

// func (h *WorkoutHandler) WorkoutHandler(w http.ResponseWriter, r *http.Request) {}
