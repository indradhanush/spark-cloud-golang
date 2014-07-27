package gospark

import (
	"reflect"
)

const (
	APIVersion     = "/v1"
	BaseUrl        = "https://api.spark.io"
	MaxVariableLen = 12
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

func ParseToken(i interface{}) (string, error) {
	if t, ok := i.(*OAuthResponse); ok {
		return t.AccessToken, nil
	} else if t, ok := i.(*AccessToken); ok {
		return t.Token, nil
	} else {
		return "", &ApiError{"Pass either OAuthResponse, or AccessToken"}
	}
}
