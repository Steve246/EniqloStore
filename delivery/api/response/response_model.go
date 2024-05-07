package response

import (
	"eniqloStore/utils"
	"errors"
	"net/http"
)

type Status struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

type ResponseSuccess struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessMessageRegister(data interface{}, detailMsg interface{}) (httpStatusCode int, apiResponse ResponseSuccess) {
	// status := Status{
	// 	ResponseCode:    SuccessCode,
	// 	ResponseMessage: SuccessMessage,
	// }
	httpStatusCode = http.StatusCreated
	apiResponse = ResponseSuccess{
		Message: detailMsg,
		Data:    data,
	}
	return
}

func NewSuccessMessageLogin(data interface{}, detailMsg interface{}) (httpStatusCode int, apiResponse ResponseSuccess) {
	// status := Status{
	// 	ResponseCode:    SuccessCode,
	// 	ResponseMessage: SuccessMessage,
	// }
	httpStatusCode = http.StatusOK
	apiResponse = ResponseSuccess{
		Message: detailMsg,
		Data:    data,
	}
	return
}

func NewSuccessMessage(data interface{}, detailMsg interface{}) (httpStatusCode int, apiResponse ResponseSuccess) {
	// status := Status{
	// 	ResponseCode:    SuccessCode,
	// 	ResponseMessage: SuccessMessage,
	// }
	httpStatusCode = http.StatusOK
	apiResponse = ResponseSuccess{
		Message: detailMsg,
		Data:    data,
	}
	return
}

func NewErrorMessage(err error) (httpStatusCode int, apiResponse Response) {
	var userError utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponseCode:    DefaultErrorCode,
			ResponseMessage: DefaultErrorMessage,
		}
		httpStatusCode = http.StatusInternalServerError
	}
	apiResponse = Response{
		Status: status,
		Data:   nil,
	}
	return
}
