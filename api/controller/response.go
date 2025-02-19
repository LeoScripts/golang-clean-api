package controller

type Response struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func NewResponseMessage(msg string) *Response {
	return &Response{
		Message: msg,
	}
}

func NewResponseMessageError(msg string) *ResponseError {
	return &ResponseError{
		Error: msg,
	}
}
