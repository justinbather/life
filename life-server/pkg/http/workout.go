package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/service"
)

type WorkoutHandler struct {
	service service.WorkoutService
}

func NewWorkoutHandler(service service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{service: service}
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	type createReq struct {
		Type string `json:"type"`
	}

	req, err := decode[createReq](r)
	if err != nil {
		fmt.Printf("Error decoding createWorkoutRequest. Err: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workout, err := h.service.CreateWorkout(r.Context(), req.Type)
	if err != nil {
		fmt.Printf("Error doing create workout request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 201, workout)
	if err != nil {
		fmt.Printf("Error encoding workout in WorkoutHandler.CreateWorkout: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *WorkoutHandler) GetWorkoutsByType(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("Error getting workouts by type: %s. Err: %s", workoutType, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(workouts) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = encode(w, r, 200, workouts)
	if err != nil {
		fmt.Printf("Error encoding []workout in WorkoutHandler.GetWorkoutsByType: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
