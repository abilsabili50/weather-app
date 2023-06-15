package errs

import "net/http"

type IErrResponse interface {
	GetCode() int
	Error() string
}

type ErrResponse struct {
	ErrError   bool   `json:"error"`
	ErrMessage string `json:"message"`
	ErrStatus  string `json:"status"`
	ErrCode    int    `json:"code"`
}

func (e *ErrResponse) GetCode() int {
	return e.ErrCode
}

func (e *ErrResponse) Error() string {
	return e.ErrStatus
}

func NewNotFoundError(message string) IErrResponse {
	return &ErrResponse{
		ErrError:   true,
		ErrMessage: message,
		ErrStatus:  "NOT_FOUND",
		ErrCode:    http.StatusNotFound,
	}
}

func NewBadRequestError(message string) IErrResponse {
	return &ErrResponse{
		ErrError:   true,
		ErrMessage: message,
		ErrStatus:  "BAD_REQUEST",
		ErrCode:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) IErrResponse {
	return &ErrResponse{
		ErrError:   true,
		ErrMessage: message,
		ErrStatus:  "INTERNAL_SERVER_ERROR",
		ErrCode:    http.StatusInternalServerError,
	}
}
