package cmd

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd/connection"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd/example"
)

func (command *Hello) FetchResult(ctx context.Context) (*task.Result, error) {
	return example.Hello(command.YourName)
}

func (command *SetSystemMessage) FetchResult(ctx context.Context) (*task.Result, error) {
	return example.SetSystemMessage(command.releaseClient, command.Message)
}

func (command *ServerQuery) FetchResult(ctx context.Context) (*task.Result, error) {
	return example.ServerQuery(ctx, command.httpClient, command.ProductId)
}

func (command *AbortHello) FetchResult(ctx context.Context) (*task.Result, error) {
	return task.NewResult().String("aborted", "successfully"), nil
}

func (command *TestConnectionCommand) FetchResult(ctx context.Context) (*task.Result, error) {
	tester := &connection.ExampleConnectionTester{
		Ctx:        ctx,
		HttpClient: command.httpClient,
	}
	return test.TestConnection(tester)
}

func (command *HelloWithLookup) FetchResult(ctx context.Context) (*task.Result, error) {
	return example.Hello(command.YourName)
}

func (command *LookupNames) FetchResult(ctx context.Context) (*task.Result, error) {
	return example.ListNames()
}
