package http

func ResponseWithHTTP(payload any, errorCode *string) HTTPResponse[any] {
	var errInResponse ErrorInResponse
	if errorCode != nil {
		errInResponse = ErrorInResponse{
			Code: *errorCode,
		}

		return HTTPResponse[any]{
			Data:  payload,
			Error: &errInResponse,
		}
	}

	return HTTPResponse[any]{
		Data: payload,
	}
}
