-- name: GetWorkoutsByType :many
SELECT * FROM workout WHERE type = $1;

-- name: CreateWorkout :one
INSERT INTO workout (type, duration, calories_burned, workload, description)
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id, type, created_at, duration, calories_burned, workload, description;
