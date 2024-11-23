package model

import (
	"fmt"
	"time"
)

type Meal struct {
	Id          int       `json:"id,omitempty"`
	Type        string    `json:"type"`
	User        string    `json:"user"`
	Calories    int       `json:"calories"`
	Protein     int       `json:"protein"`
	Carbs       int       `json:"carbs"`
	Fat         int       `json:"fat"`
	Description string    `json:"description"`
	Date        time.Time `json:"date,omitempty"`
}

func (m Meal) String() string {
	return fmt.Sprintf("Meal {\nId=%d\nUser=%s\nType=%s\nCals=%d\nProtein=%d\nCarbs=%d\nFat=%d\nDesc=%s\nDate=%s\n}\n",
		m.Id, m.User, m.Type, m.Calories, m.Protein, m.Carbs, m.Fat, m.Description, m.Date)
}
