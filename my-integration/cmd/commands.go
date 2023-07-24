package cmd

import (
	"github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/digital-ai/release-integration-sdk-go/http"
)

type Hello struct {
	YourName string `json:"yourName"`
	Sleep    bool   `json:"sleep"`
}

type SetSystemMessage struct {
	releaseClient *openapi.APIClient
	Message       string `json:"message"`
}

type ServerQuery struct {
	httpClient *http.HttpClient
	ProductId  string `json:"productId"`
}

type AbortHello struct {
	YourName string `json:"yourName"`
}
