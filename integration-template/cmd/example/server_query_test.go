package example

import (
	"errors"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/integration-template/cmd/test"
	"os"
	"testing"
)

func TestServerQuery(t *testing.T) {

	tests := []struct {
		client    *http.HttpClient
		productId string
		output    *task.Result
		response  func(httpClient *http.HttpClient, productId string) ([]byte, error)
		err       error
	}{
		{
			client:    &http.HttpClient{},
			productId: "1",
			output:    task.NewResult().String("productBrand", "Apple").String("productName", "iPhone 9"),
			response: func(httpClient *http.HttpClient, productId string) ([]byte, error) {
				return os.ReadFile("../../../test/fixtures/product.json")
			},
			err: nil,
		},
		{
			client:    &http.HttpClient{},
			productId: "non-existing",
			output:    nil,
			response: func(httpClient *http.HttpClient, productId string) ([]byte, error) {
				return nil, errors.New("not found")
			},
			err: errors.New("not found"),
		},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("ServerQuery [message = %s]", ts.productId), func(t *testing.T) {
			GetProducts = ts.response
			result, err := ServerQuery(ts.client, ts.productId)
			test.AssertRequestResult(t, result, err, ts.output, ts.err)
		})
	}
}
