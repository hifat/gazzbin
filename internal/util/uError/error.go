package uError

import "net/http"

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Status  int    `json:"-"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func InternalError() ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: Message.INTERNAL_SERVER_ERROR,
		Code:    Code.INTERNAL_SERVER_ERROR,
	}
}

func RecordNotFound() ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: Message.RECORD_NOT_FOUND,
		Code:    Code.RECORD_NOT_FOUND,
	}
}
