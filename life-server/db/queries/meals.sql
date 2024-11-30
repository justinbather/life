-- name: CreateMeal :one
INSERT INTO meal (type, user_id, calories, protein, carbs, fat, description, date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, user_id, type, calories, protein, carbs, fat, description, date;

-- name: GetMealById :one
SELECT * FROM meal WHERE id = $1;

-- name: GetMealsByType :many
SELECT * FROM meal WHERE user_id = $1 and type = $2;

-- name: GetMealsFromDateRange :many
SELECT * FROM meal WHERE user_id = $1 AND date BETWEEN $2 AND $3;
