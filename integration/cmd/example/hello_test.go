package example

import (
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/integration/cmd/test"
	"testing"
)

func TestHello(t *testing.T) {

	tests := []struct {
		name   string
		output *task.Result
		err    error
	}{
		{
			name:   "John",
			output: task.NewResult().String("greeting", "Hello John"),
			err:    nil,
		},
		{
			name:   "",
			output: nil,
			err:    errors.New("the 'yourName' field cannot be empty"),
		},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("Hello [message = %s]", ts.name), func(t *testing.T) {
			result, err := Hello(ts.name)
			test.AssertRequestResult(t, result, err, ts.output, ts.err)
		})
	}
}
