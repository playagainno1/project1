package domain

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest       = NewError(http.StatusBadRequest, 10400, "Request parameter error, please check the submitted content")
	ErrNotFound         = NewError(http.StatusNotFound, 10404, "The requested content does not exist, please check the submitted content")
	ErrInternalError    = NewError(http.StatusInternalServerError, 10500, "Server internal error, please try again")
	ErrUnauthorized     = NewError(http.StatusUnauthorized, 10401, "The account is not logged in, please log in")
	ErrForbidden        = NewError(http.StatusForbidden, 10403, "No operation permission")
	ErrUsernameTooLong  = NewError(http.StatusBadRequest, 20001, "The username is too long")
	ErrEmailFormat      = NewError(http.StatusBadRequest, 20002, "Incorrect email format")
	ErrUserShouldVerify = NewError(http.StatusBadRequest, 20003, "Your account needs to be verified")
	ErrVerifyCode       = NewError(http.StatusBadRequest, 20004, "The verification code is incorrect")
	ErrEmailNotMatch    = NewError(http.StatusBadRequest, 20005, "The email does not match the username")
)

type Error struct {
	status  int
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(status, code int, message string) Error {
	return Error{status, code, message}
}

func (e Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e Error) Status() int {
	return e.status
}
