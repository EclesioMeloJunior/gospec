package marshal_test

import (
	"io/ioutil"
	"testing"

	"github.com/eclesiomelojunior/gospec/marshal"
)

func TestJSONMarshalFactory(t *testing.T) {
	factory := marshal.FactoryMarshal("json")

	if _, ok := factory.(*marshal.JSONMarshal); !ok {
		t.Errorf("The factory does not return JSONMarsal")
	}
}

func TestNotAllowedMarshal(t *testing.T) {
	notAllowedFactory := marshal.FactoryMarshal("yaml")

	if notAllowedFactory != nil {
		t.Errorf("The factory does not return nil to not allowed marshal")
	}
}

func TestJSONMarshaler(t *testing.T) {
	toTransform := map[string]interface{}{
		"key": "value",
		"nested": map[string]string{
			"other": "value",
		},
	}

	expected := `{"key":"value","nested":{"other":"value"}}`

	factory := marshal.FactoryMarshal("json")

	if factory == nil {
		t.Errorf("The factory must return JSONMarshal but returns nil")
	}

	body, err := factory.Marshal(toTransform)

	if err != nil {
		t.Error(err)
	}

	readBody, err := ioutil.ReadAll(body)

	if err != nil {
		t.Error(err)
	}

	if string(readBody) != expected {
		t.Errorf(
			"Expected %s\n The marshal returns %s", expected, string(readBody))
	}
}
