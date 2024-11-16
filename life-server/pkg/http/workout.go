package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type WorkoutHandler struct {
	service service.WorkoutService
	logger  *prettylog.Logger
}

func NewWorkoutHandler(service service.WorkoutService, logger *prettylog.Logger) *WorkoutHandler {
	return &WorkoutHandler{service: service, logger: logger}
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	type createReq struct {
		Type string `json:"type"`
	}

	h.logger.Info("Got CreateWorkout request")

	req, err := decode[createReq](r)
	if err != nil {
		h.logger.Errorf("Error decoding createWorkoutRequest. Err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workout, err := h.service.CreateWorkout(r.Context(), req.Type)
	if err != nil {
		h.logger.Errorf("Error doing create workout request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 201, workout)
	if err != nil {
		h.logger.Errorf("Error encoding workout in WorkoutHandler.CreateWorkout: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *WorkoutHandler) GetWorkoutsByType(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Got GetWorkoutsByType Request")

	params := mux.Vars(r)
	workoutType := params["type"]

	if workoutType == "" {
		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write([]byte("Error: A workout type must be specified"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	workouts, err := h.service.GetWorkoutsByType(r.Context(), workoutType)
	if err != nil {
		h.logger.Errorf("Error getting workouts by type: %s. Err: %s", workoutType, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(workouts) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = encode(w, r, 200, workouts)
	if err != nil {
		h.logger.Errorf("Error encoding []workout in WorkoutHandler.GetWorkoutsByType: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
