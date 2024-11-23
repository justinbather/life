-- name: CreateMeal :exec
INSERT INTO meal (type, username, calories, protein, carbs, fat, description, date) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMealById :one
SELECT * FROM meal WHERE id = ?;

-- name: GetMealsByType :many
SELECT * FROM meal WHERE username = ? AND type = ?;

-- name: GetMealsFromDateRange :many
SELECT * FROM meal WHERE username = ? AND date BETWEEN ? AND ?;
