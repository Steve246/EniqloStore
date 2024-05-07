package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

type JsonResponseSuccess struct {
	c              *gin.Context
	httpStatusCode int
	response       ResponseSuccess
}

// Send implements AppHttpResponse.
func (j *JsonResponseSuccess) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, data interface{}, detailMsg interface{}, condition string) AppHttpResponse {

	if condition == "login" {
		httpStatusCode, resp := NewSuccessMessageLogin(data, detailMsg)
		return &JsonResponseSuccess{
			c,
			httpStatusCode,
			resp,
		}

	}

	if condition == "register" {
		httpStatusCode, resp := NewSuccessMessageRegister(data, detailMsg)
		return &JsonResponseSuccess{
			c,
			httpStatusCode,
			resp,
		}
	}

	httpStatusCode, resp := NewSuccessMessage(data, detailMsg)
	return &JsonResponseSuccess{
		c,
		httpStatusCode,
		resp,
	}

}

func NewErrorJsonResponse(c *gin.Context, err error) AppHttpResponse {
	fmt.Println("===>", err)
	httpStatusCode, resp := NewErrorMessage(err)
	return &JsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

func NewGlobalJsonResponse(c *gin.Context, httpStatusCode int, response Response) AppHttpResponse {
	return &JsonResponse{
		c,
		httpStatusCode,
		response,
	}
}
