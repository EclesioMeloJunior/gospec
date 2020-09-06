package models

type Endpoint struct {
	Path        string
	Description string
	Method      string
	Request     Request
	Expected    Expected
}
