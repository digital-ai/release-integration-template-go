package example

import (
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"net/http"
	"testing"
)

func TestSetSystemMessage(t *testing.T) {

	tests := []struct {
		client   *openapi.APIClient
		message  string
		output   *task.Result
		response func(releaseClient *openapi.APIClient, systemMessage openapi.SystemMessageSettings) (*openapi.SystemMessageSettings, *http.Response, error)
		err      error
	}{
		{
			client:  &openapi.APIClient{},
			message: "Welcome user!",
			output:  task.NewResult(),
			response: func(releaseClient *openapi.APIClient, systemMessage openapi.SystemMessageSettings) (*openapi.SystemMessageSettings, *http.Response, error) {
				return &systemMessage, nil, nil
			},
			err: nil,
		},
		{
			client:  &openapi.APIClient{},
			message: "Welcome user!",
			output:  nil,
			response: func(releaseClient *openapi.APIClient, systemMessage openapi.SystemMessageSettings) (*openapi.SystemMessageSettings, *http.Response, error) {
				return nil, nil, errors.New("401 unauthorized")
			},
			err: errors.New("401 unauthorized"),
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("SetSystemMessage [message = %s]", testCase.message), func(t *testing.T) {
			UpdateSystemMessage = testCase.response
			result, err := SetSystemMessage(testCase.client, testCase.message)
			test.AssertRequestResult(t, result, err, testCase.output, testCase.err)
		})
	}
}
