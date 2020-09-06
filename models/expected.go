package models

type Expected struct {
	Headers    []map[string]string
	StatusCode int
	Body       map[string]string
}
