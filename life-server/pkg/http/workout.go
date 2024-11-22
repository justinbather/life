package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

// Learn more about doing custom errors like this, where to put them etc
var ERR_USER_NOT_FOUND = errors.New("User not found")

type WorkoutHandler struct {
	service service.WorkoutService
	logger  *prettylog.Logger
}

func NewWorkoutHandler(service service.WorkoutService, logger *prettylog.Logger) *WorkoutHandler {
	return &WorkoutHandler{service: service, logger: logger}
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Got CreateWorkout request")

	req, err := decode[model.Workout](r)
	if err != nil {
		h.logger.Errorf("Error decoding createWorkoutRequest. Err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	workout, err := h.service.CreateWorkout(r.Context(), req)
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

func (h *WorkoutHandler) GetAllWorkouts(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	workouts, err := h.service.GetAllWorkouts(r.Context())
	if err != nil {
		h.logger.Errorf("Error getting all workouts Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(workouts) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = encode(w, r, 200, workouts)
	if err != nil {
		h.logger.Errorf("Error encoding []workout in WorkoutHandler.GetAllWorkouts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
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

func validate(w model.Workout) error {
	const REQUIRED = "%s is a required field."
	if w.Type == "" {
		return fmt.Errorf(REQUIRED, "Type")
	}

	return nil
}

func getUser(r *http.Request) (string, error) {
	params := mux.Vars(r)
	user := params["user"]
	if user == "" {
		return "", ERR_USER_NOT_FOUND
	}

	return user, nil
}
