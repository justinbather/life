package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	handler "github.com/justinbather/life/life-server/pkg/http"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"

	"github.com/jackc/pgx/v5"
)

func main() {
	url := os.Getenv("DATABASE_URL")

	fmt.Println("url: ", url)

	m, err := migrate.New(
		"file://./db/migrations",
		url)
	if err != nil {
		fmt.Printf("Error building migration connection: %s", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("Error running migration: %s\n", err)
		os.Exit(1)
	}

	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Error making connection to postgres db on app startup. Err: %s\n", err)
		os.Exit(1)
	}

	defer db.Close(context.Background())

	logger := prettylog.New()

	wRepository := repository.NewWorkoutRepository(db, logger)
	wService := service.NewWorkoutService(wRepository, logger)
	wHandler := handler.NewWorkoutHandler(wService, logger)

	mRepository := repository.NewMealRepository(db, logger)
	mService := service.NewMealService(mRepository, logger)
	mHandler := handler.NewMealHandler(mService, logger)

	r := mux.NewRouter()
	r.HandleFunc("/workouts/{user}/{from}/{to}", wHandler.GetWorkoutsFromDateRange).Methods(http.MethodGet)
	r.HandleFunc("/workouts/{user}/{type}", wHandler.GetWorkoutsByType).Methods(http.MethodGet)
	r.HandleFunc("/workouts/{user}", wHandler.GetAllWorkouts).Methods(http.MethodGet)
	r.HandleFunc("/workouts", wHandler.CreateWorkout).Methods(http.MethodPost)

	r.HandleFunc("/meals/{id}", mHandler.GetMealById).Methods(http.MethodGet)
	r.HandleFunc("/meals/{user}/{from}/{to}", mHandler.GetMealsFromDateRange).Methods(http.MethodGet)
	r.HandleFunc("/meals", mHandler.CreateMeal).Methods(http.MethodPost)

	fmt.Println("Life Server Running on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
