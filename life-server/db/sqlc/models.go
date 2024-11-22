// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Meal struct {
	ID          int32            `json:"id"`
	Username    string           `json:"username"`
	Type        string           `json:"type"`
	Calories    int32            `json:"calories"`
	Protein     int32            `json:"protein"`
	Carbs       int32            `json:"carbs"`
	Fat         int32            `json:"fat"`
	Description *string          `json:"description"`
	Date        pgtype.Timestamp `json:"date"`
}

type Workout struct {
	ID             int32            `json:"id"`
	Username       string           `json:"username"`
	Type           string           `json:"type"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	Duration       int32            `json:"duration"`
	CaloriesBurned int32            `json:"calories_burned"`
	Workload       int32            `json:"workload"`
	Description    *string          `json:"description"`
}
