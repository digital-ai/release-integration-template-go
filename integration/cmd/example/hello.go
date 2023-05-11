package example

import (
	"errors"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// Hello Creates a greeting based on a message
func Hello(yourName string) (*task.Result, error) {
	if len(yourName) == 0 {
		return nil, errors.New("the 'yourName' field cannot be empty")
	}

	// Create greeting
	greeting := "Hello " + yourName

	// Add greeting to the task's comment section in the UI
	task.Comment(greeting)

	// Return greeting in the output of the task
	return task.NewResult().String("greeting", greeting), nil
}
