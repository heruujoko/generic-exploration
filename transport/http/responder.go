package http

func ResponseWithHTTP[T any | []any | *any](payload T, errorCode *string) HTTPResponse[T] {
	var errInResponse ErrorInResponse
	if errorCode != nil {
		errInResponse = ErrorInResponse{
			Code: *errorCode,
		}

		return HTTPResponse[T]{
			Data:  payload,
			Error: &errInResponse,
		}
	}

	return HTTPResponse[T]{
		Data: payload,
	}
}
