package gospark

import (
	"reflect"
)

const (
	Version = "v1"
	BaseUrl = "https://api.spark.io/" + Version
)

func GetEndpoint(i interface{}) string {
	ref := reflect.TypeOf(i)
	field := ref.Field(0)
	return field.Tag.Get("endpoint")
}
