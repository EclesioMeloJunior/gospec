package marshal

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
)

//JSONMarshal struct implements BodyMarshal to json spec
type JSONMarshal struct {
}

// NewJSONMarshal returns BodyMarshal implementation
func NewJSONMarshal() BodyMarshal {
	return &JSONMarshal{}
}

func isMap(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Map
}

func parser(value interface{}) interface{} {
	switch converted := value.(type) {
	case map[string]interface{}:
		for k, v := range converted {
			converted[k] = parser(v)
		}

		return value

	case map[interface{}]interface{}:
		second := map[string]interface{}{}

		for vk, v := range converted {
			second[vk.(string)] = parser(v)
		}

		return second

	case []interface{}:
		for k, v := range converted {
			converted[k] = parser(v)
		}
	}

	return value
}

// Marshal transform in interface to a json struct
func (jm *JSONMarshal) Marshal(marshaler interface{}) (io.ReadCloser, error) {
	marshaler = parser(marshaler)
	parsedBytes, err := json.Marshal(marshaler)

	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(strings.NewReader(string(parsedBytes))), nil
}
