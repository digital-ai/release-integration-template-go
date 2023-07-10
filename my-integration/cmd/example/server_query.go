package example

import (
	"encoding/json"
	"fmt"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/task"
)

// ServerQuery Fetches product details from a remote server
func ServerQuery(httpClient *http.HttpClient, productId string) (*task.Result, error) {
	response, err := GetProducts(httpClient, productId)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	taskResult := task.NewResult()
	taskResult.String("productName", result["title"].(string))
	taskResult.String("productBrand", result["brand"].(string))

	return taskResult, nil
}

var GetProducts = func(httpClient *http.HttpClient, productId string) ([]byte, error) {
	return httpClient.Get(fmt.Sprintf("products/%s", productId))
}
