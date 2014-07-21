package gospark

import (
	"reflect"
)

const (
	APIVersion = "/v1"
	BaseUrl    = "https://api.spark.io"
)

type APIUrl struct {
	BaseUrl    string
	APIVersion string
	Endpoint   string
}

func GetCompleteEndpointUrl(a *APIUrl) string {
	return a.BaseUrl + a.APIVersion + a.Endpoint

}

func GetEndpoint(i interface{}) string {
	ref := reflect.TypeOf(i)
	field := ref.Field(0)
	return field.Tag.Get("endpoint")
}
