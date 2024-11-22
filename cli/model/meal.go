package model

type Meal struct {
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
	User        string `json:"user"`
	Calories    int    `json:"calories"`
	Protein     int    `json:"protein"`
	Carbs       int    `json:"carbs"`
	Fat         int    `json:"fat"`
	Description string `json:"description"`
}
