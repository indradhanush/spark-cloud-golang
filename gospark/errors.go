package gospark

type ApiError struct {
	ErrorMsg string
}

func (e ApiError) Error() string {
	return e.ErrorMsg
}
