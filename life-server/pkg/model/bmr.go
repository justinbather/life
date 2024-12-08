package model

import "time"

type Bmr struct {
	Id            int       `json:"id,omitempty"`
	UserId        string    `json:"userId,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	TotalCalories int       `json:"totalCalories"`
	NumWorkouts   int       `json:"numWorkouts"`
}
