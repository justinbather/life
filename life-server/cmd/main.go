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
	var authCacheTTL int = 1

	db, err := setupDb(url)
	if err != nil {
		fmt.Printf("Error making connection to postgres db on app startup. Err: %s\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	logger := prettylog.New()

	userRepository := repository.NewUserRepository(db, logger)
	userService := service.NewUserService(userRepository)

	// TODO: un fuck this dependency between user and auth services
	authService := service.NewAuthService(db, userService, authCacheTTL, logger)
	userHandler := handlers.NewUserHandler(userService, authService, logger)

	healthHandler := handlers.NewHealthHandler(logger)
	workoutHandler := registerWorkoutDomain(db, logger)
	mealHandler := registerMealDomain(db, logger)

	mw := middleware.NewMiddleware(authService, userService, logger)

	global := mux.NewRouter()

	global.Use(mw.Recoverer)
	global.Use(mw.Tracer)

	// Non-protected routes
	global.HandleFunc("/auth/login", userHandler.Login).Methods(http.MethodPost)
	global.HandleFunc("/auth/signup", userHandler.Signup).Methods(http.MethodPost)

	protected := global.PathPrefix("/").Subrouter()
	protected.Use(mw.Protect)

	// TODO: Move health check to global after
	protected.HandleFunc("/health-check", healthHandler.HealthCheck).Methods(http.MethodGet)

	protected.HandleFunc("/workouts/{from}/{to}", workoutHandler.GetWorkoutsFromDateRange).Methods(http.MethodGet)
	protected.HandleFunc("/workouts/{type}", workoutHandler.GetWorkoutsByType).Methods(http.MethodGet)
	protected.HandleFunc("/workouts", workoutHandler.GetAllWorkouts).Methods(http.MethodGet)
	protected.HandleFunc("/workouts", workoutHandler.CreateWorkout).Methods(http.MethodPost)

	protected.HandleFunc("/meals/{id}", mealHandler.GetMealById).Methods(http.MethodGet)
	protected.HandleFunc("/meals/{from}/{to}", mealHandler.GetMealsFromDateRange).Methods(http.MethodGet)
	protected.HandleFunc("/meals", mealHandler.CreateMeal).Methods(http.MethodPost)

	fmt.Println("Life Server Running on 8080")
	if err := http.ListenAndServe(":8080", global); err != nil {
		panic(err)
	}
}

func setupDb(url string) (*pgx.Conn, error) {
	fmt.Println("Setting up database...")

	err := migrateDb(url)
	if err != nil {
		fmt.Printf("Error migrating database: %s\n", err)
		return nil, err
	}

	fmt.Println("Migration Complete...")
	fmt.Println("Connecting to database...")

	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Error connecting: %s\n", err)
		return nil, err
	}

	fmt.Println("Connected to database successfully...")

	return db, nil
}

func migrateDb(url string) error {
	m, err := migrate.New(
		"file://./db/migrations",
		url)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func registerUserDomain(db *pgx.Conn, authService service.AuthService, logger *prettylog.Logger) (*handlers.UserHandler, service.UserService) {

	uRepository := repository.NewUserRepository(db, logger)
	uService := service.NewUserService(uRepository)
	uHandler := handlers.NewUserHandler(uService, authService, logger)

	return uHandler, uService
}

func registerWorkoutDomain(db *pgx.Conn, logger *prettylog.Logger) *handlers.WorkoutHandler {
	wRepository := repository.NewWorkoutRepository(db, logger)
	wService := service.NewWorkoutService(wRepository, logger)
	wHandler := handlers.NewWorkoutHandler(wService, logger)

	return wHandler
}

func registerMealDomain(db *pgx.Conn, logger *prettylog.Logger) *handlers.MealHandler {

	mRepository := repository.NewMealRepository(db, logger)
	mService := service.NewMealService(mRepository, logger)
	mHandler := handlers.NewMealHandler(mService, logger)

	return mHandler
}
