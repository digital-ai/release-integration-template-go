package main

import (
	"context"
	"github.com/digital-ai/release-integration-sdk-go/api/release"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"github.com/digital-ai/release-integration-sdk-go/task/property"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd"
	"github.com/digital-ai/release-integration-template-go/task/server"
	"os"
)

var PluginVersion = os.Getenv("VERSION")
var BuildDate = os.Getenv("BUILD_DATE")

func prepareCommandFactory(input task.InputContext) (command.CommandFactory, error) {
	var httpClient *http.HttpClient

	// If there is no server as an input property, return empty http httpClient
	if _, err := property.ExtractByName(server.ApiServerNameField, input.Task.Properties); err == nil {

		// Otherwise, deserialize server and return http httpClient
		apiServer, err := server.DeserializeApiServer(input.Task.Properties)
		if err != nil {
			return nil, err
		}

		httpClient, err = apiServer.GetHttpClient()
		if err != nil {
			return nil, err
		}
	}

	ctx := task.ReleaseContext{
		Id: input.Release.Id,
		AutomatedTaskAsUser: task.AutomatedTaskAsUserContext{
			Username: input.Release.AutomatedTaskAsUser.Username,
			Password: input.Release.AutomatedTaskAsUser.Password,
		},
		Url: input.Release.Url,
	}

	releaseClient, err := release.NewReleaseApiClient(ctx)
	if err != nil {
		return nil, err
	}

	return cmd.NewCommandFactory(httpClient, releaseClient), nil

}

var commandRunner = runner.NewCommandRunner(prepareCommandFactory)

func main() {
	context.Background()
	runner.Execute(PluginVersion, BuildDate, commandRunner)
}
