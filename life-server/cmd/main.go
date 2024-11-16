package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	handler "github.com/justinbather/life/life-server/pkg/http"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"

	"github.com/jackc/pgx/v5"
)

func main() {

	db, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Printf("Error making connection to postgres db on app startup. Err: %s\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	logger := prettylog.New()

	repository := repository.NewWorkoutRepository(db, logger)
	service := service.NewWorkoutService(repository, logger)
	handler := handler.NewWorkoutHandler(service, logger)

	r := mux.NewRouter()
	r.HandleFunc("/workouts/{type}", handler.GetWorkoutsByType).Methods(http.MethodGet)
	r.HandleFunc("/workouts", handler.CreateWorkout).Methods(http.MethodPost)

	fmt.Println("Life Server Running on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
