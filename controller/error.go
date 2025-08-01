package controller

import "fmt"

func (app *TrakerCron) respondWithError(message string, err error)(int, error) {
	return fmt.Print(map[string]string{
		"status":  "invalid",
		"message": message,
		"problem": err.Error(),
	})
}