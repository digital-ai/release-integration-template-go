package connection

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"testing"
)

func TestConnectionTest(t *testing.T) {

	type Tests struct {
		name        string
		client      *http.HttpClient
		getFunc     func(ctx context.Context, httpClient *http.HttpClient) error
		returnValue *task.Result
		expectedErr error
	}

	Success := func() *task.Result {
		result, _ := json.Marshal(map[string]interface{}{"output": "", "success": true})
		return task.NewResult().Json(task.DefaultResponseResultField, result)
	}

	Error := func() *task.Result {
		result, _ := json.Marshal(map[string]interface{}{"output": "some error", "success": false})
		return task.NewResult().Json(task.DefaultResponseResultField, result)
	}

	tests := []Tests{
		{
			name:   "test connection success",
			client: &http.HttpClient{},
			getFunc: func(ctx context.Context, httpClient *http.HttpClient) error {
				return nil
			},
			returnValue: Success(),
			expectedErr: nil,
		},
		{
			name:   "test connection fail",
			client: &http.HttpClient{},
			getFunc: func(ctx context.Context, httpClient *http.HttpClient) error {
				return fmt.Errorf("some error")
			},
			returnValue: Error(),
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := &ExampleConnectionTester{
				Ctx:        context.Background(),
				HttpClient: tt.client,
			}
			TestConnectionRequest = tt.getFunc
			got, err := test.TestConnection(tester)
			test.AssertRequestResult(t, got, err, tt.returnValue, tt.expectedErr)
		})
	}
}
