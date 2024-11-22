package model

type Workout struct {
	Id             int    `json:"id,omitempty"`
	User           string `json:"user,omitempty"`
	Type           string `json:"type"`
	Duration       int    `json:"duration"`
	CaloriesBurned int    `json:"caloriedBurned"`
	Workload       int    `json:"workload"`
	Description    string `json:"description"`
}
