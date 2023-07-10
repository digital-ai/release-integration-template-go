package cmd

import (
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/http"
)

type Hello struct {
	YourName string `json:"yourName"`
}

type SetSystemMessage struct {
	releaseClient *openapi.APIClient
	Message       string `json:"message"`
}

type ServerQuery struct {
	httpClient *http.HttpClient
	ProductId  string `json:"productId"`
}