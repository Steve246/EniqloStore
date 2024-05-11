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

// customer

func CreateCustomerError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "Create Customer Failed",
		ErrorType:    http.StatusConflict,
	}
}

func GetCustomerError() error {
	return AppError{
		ErrorCode:    "400",
		ErrorMessage: "No Customer Found",
		ErrorType:    http.StatusConflict,
	}
}

// product

func ProductsNotFound() error {
	return AppError{
		ErrorCode:    "404",
		ErrorMessage: "id is not found",
		ErrorType:    http.StatusNotFound,
	}
}

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
func PhoneNumberFoundError() error {
	return AppError{
		ErrorCode:    "409",
		ErrorMessage: "Phone Number found inside Database",
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
