-- name: CreateMeal :one
INSERT INTO meal (type, calories, protein, carbs, fat, description, date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, type, calories, protein, carbs, fat, description, date;

-- name: GetMealById :one
SELECT * FROM meal WHERE id = $1;

-- name: GetMealsByType :many
SELECT * FROM meal WHERE type = $1;

-- name: GetMealsFromDateRange :many
SELECT * FROM meal WHERE date BETWEEN $1 AND $2;
