package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func assert(response *http.Response, expected testexpected) (string, error) {
	if !assertStatusCode(response, expected) {
		result := "FAIL\nExpected [%v], Received [%v]"
		return fmt.Sprintf(result, expected.Status, response.StatusCode), nil
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return "FAIL", err
	}

	if expected.Body.IsArray() {
		var receivedBody []map[string]interface{}
		err := json.Unmarshal(responseBody, &receivedBody)

		if err != nil {
			return "FAIL", err
		}

		expectedBody, err := fromJSONToArrayMap(expected.Body.JSON)

		if err != nil {
			return "FAIL", err
		}

		var result bool
		if result = assertArrayBody(receivedBody, expectedBody); !result {
			result := "FAIL\nExpected\n%s\n\nReceived\n%s"

			expectedBytes, err := json.MarshalIndent(expectedBody, "", "  ")
			if err != nil {
				fmt.Println(err)
			}

			return fmt.Sprintf(result, string(expectedBytes), string(responseBody)), nil
		}
	} else {
		var receivedBody map[string]interface{}
		err := json.Unmarshal(responseBody, &receivedBody)

		if err != nil {
			return "FAIL", err
		}

		expectedBody, err := fromJSONToMap(expected.Body.JSON)

		if err != nil {
			return "FAIL", err
		}

		var result bool
		if result = assertSimpleBody(receivedBody, expectedBody); !result {
			result := "FAIL\nExpected\n%s\n\nReceived\n%s"

			expectedBytes, err := json.MarshalIndent(expectedBody, "", "  ")
			if err != nil {
				fmt.Println(err)
			}

			return fmt.Sprintf(result, string(expectedBytes), string(responseBody)), nil
		}
	}

	return "OK", nil
}

func assertStatusCode(response *http.Response, expected testexpected) bool {
	return response.StatusCode == expected.Status
}

func assertArrayBody(receivedBody []map[string]interface{}, expected []map[string]interface{}) bool {
	if len(receivedBody) != len(expected) {
		return false
	}

	for index, received := range receivedBody {
		if !assertSimpleBody(received, expected[index]) {
			return false
		}
	}

	return true
}

func assertSimpleBody(receivedBody map[string]interface{}, expectedBody map[string]interface{}) bool {
	return reflect.DeepEqual(receivedBody, expectedBody)
}
