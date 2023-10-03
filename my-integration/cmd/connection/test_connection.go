package connection

import (
	"context"
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

func TestConnection(ctx context.Context, httpClient *http.HttpClient) (*task.Result, error) {
	var result task.TestConnectionResult
	result.Success = true

	_, err := TestConnectionRequest(ctx, httpClient)
	if err != nil {
		result.Success = false
		result.Output = err.Error()
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		result.Success = false
		result.Output = err.Error()
	}

	return task.NewResult().Json(task.DefaultResponseResultField, resultJson), nil
}

var TestConnectionRequest = func(ctx context.Context, httpClient *http.HttpClient) ([]byte, error) {
	return httpClient.Get(ctx, "")
}
