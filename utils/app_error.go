package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

// product

func GetProductError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Get Product Failed",
		ErrorType:    http.StatusConflict,
	}
}

func CreateProductError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Create Product Failed",
		ErrorType:    http.StatusConflict,
	}
}

// register
func EmailFoundError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Email found inside Database",
		ErrorType:    http.StatusConflict,
	}
}

func ReqBodyNotValidError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Didn't pass Validation",
		ErrorType:    http.StatusBadRequest,
	}
}

func ServerError() error {
	return AppError{
		ErrorCode:    "500",
		ErrorMessage: "Server Error",
		ErrorType:    http.StatusInternalServerError,
	}
}

// login

func PasswordCannotBeEncodeError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password cannot be encode",
		ErrorType:    http.StatusBadRequest,
	}
}

func UserNotFoundError() error {
	return AppError{
		ErrorCode:    "404",
		ErrorMessage: "User Not Found",
		ErrorType:    http.StatusInternalServerError,
	}
}

func PasswordWrongError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Password Is Wrong",
		ErrorType:    http.StatusInternalServerError,
	}
}
