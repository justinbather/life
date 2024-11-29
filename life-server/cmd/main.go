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
	"github.com/justinbather/life/life-server/pkg/http/handlers"
	"github.com/justinbather/life/life-server/pkg/http/middleware"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"

	"github.com/jackc/pgx/v5"
)

func main() {
	url := os.Getenv("DATABASE_URL")

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

	healthHandler := handlers.NewHealthHandler(logger)

	authService := service.NewAuthService()
	mw := middleware.NewMiddleware(authService, logger)

	uRepository := repository.NewUserRepository(db, logger)
	uService := service.NewUserService(uRepository)
	uHandler := handlers.NewUserHandler(uService, authService, logger)

	wRepository := repository.NewWorkoutRepository(db, logger)
	wService := service.NewWorkoutService(wRepository, logger)
	wHandler := handlers.NewWorkoutHandler(wService, logger)

	mRepository := repository.NewMealRepository(db, logger)
	mService := service.NewMealService(mRepository, logger)
	mHandler := handlers.NewMealHandler(mService, logger)

	global := mux.NewRouter()

	global.Use(mw.Recoverer)
	global.Use(mw.Tracer)

	// Non-protected routes
	global.Handle("/health-check", mw.Protect(http.HandlerFunc(healthHandler.HealthCheck))).Methods(http.MethodGet)
	global.HandleFunc("/auth/login", uHandler.Login).Methods(http.MethodPost)
	global.HandleFunc("/auth/signup", uHandler.Signup).Methods(http.MethodPost)

	protected := global.PathPrefix("/").Subrouter()
	protected.Use(mw.Protect)

	protected.HandleFunc("/workouts/{user}/{from}/{to}", wHandler.GetWorkoutsFromDateRange).Methods(http.MethodGet)
	protected.HandleFunc("/workouts/{user}/{type}", wHandler.GetWorkoutsByType).Methods(http.MethodGet)
	protected.HandleFunc("/workouts/{user}", wHandler.GetAllWorkouts).Methods(http.MethodGet)
	protected.HandleFunc("/workouts", wHandler.CreateWorkout).Methods(http.MethodPost)

	protected.HandleFunc("/meals/{id}", mHandler.GetMealById).Methods(http.MethodGet)
	protected.HandleFunc("/meals/{user}/{from}/{to}", mHandler.GetMealsFromDateRange).Methods(http.MethodGet)
	protected.HandleFunc("/meals", mHandler.CreateMeal).Methods(http.MethodPost)

	fmt.Println("Life Server Running on 8080")
	if err := http.ListenAndServe(":8080", global); err != nil {
		panic(err)
	}
}
