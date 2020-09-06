package models

type Host struct {
	Description string
	Protocol    string
	Host        string
	Port        string
	Prefix      string
	Endpoints   []Endpoint
}
