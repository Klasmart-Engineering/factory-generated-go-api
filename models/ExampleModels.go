package models

import (
	"net/http"
)

type Healthcheck struct {
	Up      bool   `json:"up"`
	Details string `json:"details"`
}

func (*Healthcheck) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ExampleStruct struct {
	ExampleId   string `json:"ID"`
	ExampleName string `json:"name"`
}

type ExampleStructList struct {
	ExampleStruct []ExampleStruct `json:"exampleResponse"`
}

func (*ExampleStructList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
