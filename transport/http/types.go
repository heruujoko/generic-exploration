package http

type ErrorInResponse struct {
	Code string `json:"code"`
}

type HTTPResponse[T any | []any] struct {
	Data  T                `json:"data,omitempty"`
	Error *ErrorInResponse `json:"error,omitempty"`
}
