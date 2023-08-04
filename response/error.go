package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(msg string) Response {
	return errors(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errors(msg, http.StatusNotFound)
}
func Unauthorized(msg string) Response {
	return errors(msg, http.StatusUnauthorized)
}
func Forbidden(msg string) Response {
	return errors(msg, http.StatusForbidden)
}
func BadRequest(msg string) Response {
	return errors(msg, http.StatusBadRequest)
}

func errors(msg string, status int) Response {
	return &ErrorResponse{
		Status:  status,
		Message: msg,
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}
func (e *ErrorResponse) StatusCode() int {
	return e.Status
}
func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}
func (e *ErrorResponse) GetData() interface{} {
	return nil
}
