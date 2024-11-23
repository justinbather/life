-- name: GetWorkoutsByType :many
SELECT * FROM workout WHERE username = ? AND type = ?;

-- name: GetAllWorkouts :many
SELECT * FROM workout WHERE username = ?;

-- name: CreateWorkout :exec
INSERT INTO workout (type, username, duration, calories_burned, workload, description)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetWorkoutsFromDateRange :many
SELECT id, username, type, created_at, duration, calories_burned, workload, description
FROM workout
WHERE username = ? AND created_at BETWEEN ? AND ?;

