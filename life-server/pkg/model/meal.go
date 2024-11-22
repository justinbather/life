package model

import "time"

type Meal struct {
	Id          int       `json:"id"`
	User        string    `json:"user"`
	Type        string    `json:"type"`
	Calories    int       `json:"calories"`
	Protein     int       `json:"protein"`
	Carbs       int       `json:"carbs"`
	Fat         int       `json:"fat"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
