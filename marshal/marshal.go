package marshal

import (
	"io"
)

// BodyMarshal converts a apispec body
type BodyMarshal interface {
	Marshal(interface{}) (io.ReadCloser, error)
}

// FactoryMarshal returns a BodyMarshal interface implementation
func FactoryMarshal(marshaltype string) BodyMarshal {
	switch marshaltype {
	case "json":
		return NewJSONMarshal()
	default:
		return nil
	}
}
