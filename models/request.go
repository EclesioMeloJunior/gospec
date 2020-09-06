package models

type Request struct {
	Headers     []map[string]string
	QueryParams []map[string]string
	Body        map[string]interface{}
}
