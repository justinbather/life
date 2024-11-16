package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/http"
	"github.com/justinbather/life/life-server/pkg/service"
)

func main() {
	fmt.Println("Life Server Running on 8080")

	service := service.NewWorkoutService()

	h := handler.WorkoutHandler{Service: service}

	r := mux.NewRouter()
	r.HandleFunc("/workouts/{type}", h.WorkoutsHandler).Methods(http.MethodPost, http.MethodGet)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
