-- name: GetWorkoutsByType :many
SELECT * FROM workout WHERE user_id = $1 AND type = $2;

-- name: GetAllWorkouts :many
SELECT * FROM workout WHERE user_id = $1;

-- name: CreateWorkout :one
INSERT INTO workout (type, user_id, duration, calories_burned, workload, description)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, user_id, type, created_at, duration, calories_burned, workload, description;

-- name: GetWorkoutsFromDateRange :many
SELECT * FROM workout WHERE user_id = $1 AND created_at BETWEEN $2 AND $3;
