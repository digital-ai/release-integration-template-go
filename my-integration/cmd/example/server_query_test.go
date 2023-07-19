package example

import (
	"context"
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"os"
	"testing"
)

func TestServerQuery(t *testing.T) {

	tests := []struct {
		client    *http.HttpClient
		productId string
		output    *task.Result
		response  func(ctx context.Context, httpClient *http.HttpClient, productId string) ([]byte, error)
		err       error
		ctx       context.Context
	}{
		{
			client:    &http.HttpClient{},
			productId: "1",
			output:    task.NewResult().String("productBrand", "Apple").String("productName", "iPhone 9"),
			response: func(ctx context.Context, httpClient *http.HttpClient, productId string) ([]byte, error) {
				return os.ReadFile("../../../test/fixtures/product.json")
			},
			err: nil,
			ctx: context.Background(),
		},
		{
			client:    &http.HttpClient{},
			productId: "non-existing",
			output:    nil,
			response: func(ctx context.Context, httpClient *http.HttpClient, productId string) ([]byte, error) {
				return nil, errors.New("not found")
			},
			err: errors.New("not found"),
			ctx: context.Background(),
		},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("ServerQuery [message = %s]", testCase.productId), func(t *testing.T) {
			GetProducts = testCase.response
			result, err := ServerQuery(testCase.ctx, testCase.client, testCase.productId)
			test.AssertRequestResult(t, result, err, testCase.output, testCase.err)
		})
	}
}
