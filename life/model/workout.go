package model

import (
	"fmt"
	"time"
)

type Workout struct {
	Id             int       `json:"id,omitempty"`
	User           string    `json:"user,omitempty"`
	Type           string    `json:"type"`
	Duration       int       `json:"duration"`
	CaloriesBurned int       `json:"caloriesBurned"`
	Workload       int       `json:"workload"`
	Description    string    `json:"description"`
	Date           time.Time `json:"createdAt,omitempty"`
}

func (w Workout) String() string {
	return fmt.Sprintf("Workout {\nId=%d\nUser=%s\nType=%s\nDuration=%d\nCalsBurned=%d\nLoad=%d\nDesc=%s\nDate=%s\n}\n", w.Id, w.User, w.Type,
		w.Duration, w.CaloriesBurned, w.Workload, w.Description, w.Date)
}
