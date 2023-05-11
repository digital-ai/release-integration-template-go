package cmd

import (
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/http"
)

type GetCiTags struct {
	client   *http.HttpClient
	DeployCi string `json:"deployCI"`
}

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
