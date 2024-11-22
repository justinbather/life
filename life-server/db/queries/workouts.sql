-- name: GetWorkoutsByType :many
SELECT * FROM workout WHERE username = $1 AND type = $2;

-- name: GetAllWorkouts :many
SELECT * FROM workout WHERE username = $1;

-- name: CreateWorkout :one
INSERT INTO workout (type, username, duration, calories_burned, workload, description)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, username, type, created_at, duration, calories_burned, workload, description;
