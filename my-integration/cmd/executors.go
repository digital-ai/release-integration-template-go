package cmd

import (
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd/example"
)

func (command *Hello) FetchResult() (*task.Result, error) {
	return example.Hello(command.YourName)
}

func (command *SetSystemMessage) FetchResult() (*task.Result, error) {
	return example.SetSystemMessage(command.releaseClient, command.Message)
}

func (command *ServerQuery) FetchResult() (*task.Result, error) {
	return example.ServerQuery(command.httpClient, command.ProductId)
}
