package example

import (
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/integration"
	"github.com/digital-ai/release-integration-template-go/integration/cmd/test"
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
			output:  task.NewResult().String(integration.DefaultResponseResultField, "System message updated"),
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

	for _, ts := range tests {
		t.Run(fmt.Sprintf("SetSystemMessage [message = %s]", ts.message), func(t *testing.T) {
			UpdateSystemMessage = ts.response
			result, err := SetSystemMessage(ts.client, ts.message)
			test.AssertRequestResult(t, result, err, ts.output, ts.err)
		})
	}
}
