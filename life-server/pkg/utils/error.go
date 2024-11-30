package utils

import "fmt"

type LifeErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e LifeErr) Error() string {
	return fmt.Sprintf("Error: %s. Code: %d", e.Message, e.Code)
}
