package connection

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/http"
)

// ExampleConnectionTester is a struct for testing connection to server, containing needed properties.
// Extends ConnectionTester from SDK. See https://github.com/digital-ai/release-integration-sdk-go/blob/master/test/connection.go
type ExampleConnectionTester struct {
	Ctx        context.Context
	HttpClient *http.HttpClient
}

// TestConnection implements logic for testing the connection to the server.
// If there is no error, connection is successful.
func (tester *ExampleConnectionTester) TestConnection() error {
	return TestConnectionRequest(tester.Ctx, tester.HttpClient)
}

var TestConnectionRequest = func(ctx context.Context, httpClient *http.HttpClient) error {
	_, err := httpClient.Get(ctx, "")
	return err
}
