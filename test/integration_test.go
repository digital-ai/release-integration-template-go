package test

import (
	"github.com/digital-ai/release-integration-sdk-go/api/release"
	"github.com/digital-ai/release-integration-sdk-go/http"
	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
	"github.com/digital-ai/release-integration-sdk-go/test"
	"github.com/digital-ai/release-integration-template-go/my-integration/cmd"
	"testing"
)

var testsLabels = []string{
	"hello",
	"server-query",
	//"set-system-message", //todo after defining mock
	"connection",
}

var commandRunner = runner.NewCommandRunner(
	func(input task.InputContext) (command.CommandFactory, error) {
		var httpClient http.HttpClient
		httpClient.Client(
			test.NewMockHttpClient(
				[]test.MockResult{
					{
						Method:     "GET",
						Path:       "/products/1",
						Filename:   "fixtures/product.json",
						StatusCode: 200,
					},
					{
						Method:     "GET",
						Path:       "/",
						Filename:   "fixtures/void.json",
						StatusCode: 200,
					},
				},
			),
		)

		releaseClient, err := release.NewReleaseApiClient(input.Release) //todo define a mock release client
		if err != nil {
			return nil, err
		}

		return cmd.NewCommandFactory(&httpClient, releaseClient), nil
	},
)

func GenerateFixtures() []test.ExecutorTest {
	testMap := make(map[string]runner.Runner)
	for _, label := range testsLabels {
		testMap[label] = commandRunner
	}
	return test.CreateExecutorTestSet("testdata", testMap)
}

func TestSpec(t *testing.T) {
	test.ConveyTest(t, GenerateFixtures())
}
