package model

import (
	"time"
)

type Workout struct {
	Id             int       `json:"id"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"createdAt"`
	Duration       int       `json:"duration"`
	CaloriesBurned int       `json:"caloriesBurned"`
	Workload       int       `json:"workload"`
	Description    string    `json:"description"`
}
