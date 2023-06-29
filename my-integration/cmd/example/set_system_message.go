package example

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"net/http"
)

// SetSystemMessage Sets the system message in the Release UI by invoking the API.
func SetSystemMessage(releaseClient *openapi.APIClient, message string) (*task.Result, error) {
	// Define parameter object to send through the API client
	systemMessage := openapi.SystemMessageSettings{}
	systemMessage.SetType("xlrelease.SystemMessageSettings")
	systemMessage.SetId("Configuration/settings/SystemMessageSettings")
	systemMessage.SetMessage(message)
	systemMessage.SetEnabled(true)
	systemMessage.SetAutomated(false)

	// Make the actual rest call to the designated endpoint
	_, _, err := UpdateSystemMessage(releaseClient, systemMessage)
	if err != nil {
		return nil, err
	}

	// Return message in the output of the task
	return task.NewResult(), nil
}

var UpdateSystemMessage = func(releaseClient *openapi.APIClient, systemMessage openapi.SystemMessageSettings) (*openapi.SystemMessageSettings, *http.Response, error) {
	return releaseClient.ConfigurationApi.UpdateSystemMessage(context.TODO()).SystemMessageSettings(systemMessage).Execute()
}
