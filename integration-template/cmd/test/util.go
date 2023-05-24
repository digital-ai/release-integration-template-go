package test

import (
	"encoding/json"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"reflect"
	"testing"
)

func AssertRequestResult(t *testing.T, actual *task.Result, actualErr error, expected *task.Result, expectedErr error) {
	if actualErr != nil {
		if !reflect.DeepEqual(actualErr, expectedErr) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", actualErr, expectedErr)
		} else {
			t.Logf("Success!")
		}
	} else {
		mapResult, err := actual.Get()
		if !reflect.DeepEqual(err, expectedErr) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", err, expectedErr)
		}
		response, err := json.Marshal(mapResult)
		if !reflect.DeepEqual(err, expectedErr) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", err, expectedErr)
		}

		expectedMap, err := expected.Get()
		if err != nil {
			t.Fatalf("Error while trying to get value from expected: [%v]", err)
		}

		expectedJson, err := json.Marshal(expectedMap)
		if err != nil {
			t.Fatalf("Error while trying to marshal expected: [%v]", err)
		}

		if !reflect.DeepEqual(string(response), string(expectedJson)) {
			t.Fatalf("Actual: [%v]; Expected: [%v]", string(response), string(expectedJson))
		} else {
			t.Logf("Success!")
		}
	}
}
