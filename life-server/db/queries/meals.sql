-- name: CreateMeal :one
INSERT INTO meal (type, username, calories, protein, carbs, fat, description, date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, type, username, calories, protein, carbs, fat, description, date;

-- name: GetMealById :one
SELECT * FROM meal WHERE id = $1;

-- name: GetMealsByType :many
SELECT * FROM meal WHERE username = $1 and type = $2;

-- name: GetMealsFromDateRange :many
SELECT * FROM meal WHERE username = $1 AND date BETWEEN $2 AND $3;
