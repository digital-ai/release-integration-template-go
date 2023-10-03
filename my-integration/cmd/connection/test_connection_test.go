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
	type args struct {
		client *http.HttpClient
	}

	type Tests struct {
		name        string
		getFunc     func(ctx context.Context, httpClient *http.HttpClient) ([]byte, error)
		args        args
		returnValue *task.Result
		expectedErr error
	}

	SuccessResponse := func() *task.Result {
		resultJson, _ := json.Marshal(map[string]interface{}{"output": "", "success": true})
		result := task.NewResult().Json(task.DefaultResponseResultField, resultJson)
		return result
	}

	ErrorResponse := func() *task.Result {
		resultJson, _ := json.Marshal(map[string]interface{}{"output": "some error", "success": false})
		result := task.NewResult().Json(task.DefaultResponseResultField, resultJson)
		return result
	}

	tests := []Tests{
		{
			name: "test connection success",
			getFunc: func(ctx context.Context, httpClient *http.HttpClient) ([]byte, error) {
				return json.Marshal(nil)
			},
			args: args{
				client: &http.HttpClient{},
			},
			returnValue: SuccessResponse(),
			expectedErr: nil,
		},
		{
			name: "test connection fail",
			getFunc: func(ctx context.Context, httpClient *http.HttpClient) ([]byte, error) {
				return nil, fmt.Errorf("some error")
			},
			args: args{
				client: &http.HttpClient{},
			},
			returnValue: ErrorResponse(),
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TestConnectionRequest = tt.getFunc
			got, err := TestConnection(context.Background(), tt.args.client)
			test.AssertRequestResult(t, got, err, tt.returnValue, tt.expectedErr)
		})
	}
}
