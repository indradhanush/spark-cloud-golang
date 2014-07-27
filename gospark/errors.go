package gospark

import (
	"fmt"
)

type ApiError struct {
	ErrorMsg string
}

func (e ApiError) Error() string {
	return e.ErrorMsg
}

type ErrorResponse struct {
	Code             int32  `json:"code, omitempty"`
	Err              string `json:"error, omitempty"`
	ErrorDescription string `json:"error_description, omitempty"`
	Info             string `json:"info, omitempty"`
}

func (e ErrorResponse) Error() string {
	errorMsg := ""

	// Just hoping that there isn't any error code of 0
	if e.Code != 0 {
		errorMsg = fmt.Sprintf("Error Code: %v Error:%v", e.Code, e.Err)
	}
	if e.ErrorDescription != "" {
		errorMsg += fmt.Sprintf("Error Description:%v",
			e.ErrorDescription)
	}
	if e.Info != "" {
		errorMsg += fmt.Sprint("Info:", e.Info)
	}

	return errorMsg
}
