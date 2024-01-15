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

func TestListNames(t *testing.T) {

	tests := []struct {
		output *task.Result
		err    error
	}{
		{
			output: task.NewResult().LookupResultElements(task.DefaultResponseResultField, []task.LookupResultElement{
				{Label: "World", Value: "World"},
				{Label: "John", Value: "John"},
				{Label: "Bob", Value: "Bob"},
				{Label: "Mary", Value: "Mary"},
			}),
			err: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("ListNames"), func(t *testing.T) {
			result, err := ListNames()
			test.AssertRequestResult(t, result, err, testCase.output, testCase.err)
		})
	}
}
