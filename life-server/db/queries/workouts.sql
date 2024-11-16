-- name: GetWorkoutsByType :many
SELECT * FROM workout WHERE type = $1;

-- name: CreateWorkout :one
INSERT INTO workout (type) VALUES ($1) RETURNING id, type;
