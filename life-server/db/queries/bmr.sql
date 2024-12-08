-- name: CreateBmr :one
INSERT INTO bmr (created_at, total_calories, num_workouts, user_id) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_id, created_at, total_calories, num_workouts;

-- name: GetBmrById :one
SELECT * FROM bmr WHERE id = $1;

-- name: GetBmrFromDateRange :many
SELECT * FROM bmr WHERE user_id = $1 AND created_at BETWEEN $2 AND $3;
