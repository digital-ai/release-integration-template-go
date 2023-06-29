package example

import (
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"testing"
)

func TestHello(t *testing.T) {

	tests := []struct {
		yourName string
		output   *task.Result
		err      error
	}{
		{
			yourName: "John",
			output:   task.NewResult().String("greeting", "Hello John"),
			err:      nil,
		},
		{
			yourName: "",
			output:   nil,
			err:      errors.New("the 'yourName' field cannot be empty"),
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("Hello [message = %s]", testCase.yourName), func(t *testing.T) {
			result, err := Hello(testCase.yourName)
			test.AssertRequestResult(t, result, err, testCase.output, testCase.err)
		})
	}
}
