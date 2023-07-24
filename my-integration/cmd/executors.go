package cmd

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd/example"
	"time"
)

func (command *Hello) FetchResult(ctx context.Context) (*task.Result, error) {
	if command.Sleep {
		time.Sleep(60 * time.Second)
	}
	task.Comment("Hello world sleep stopped")
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
