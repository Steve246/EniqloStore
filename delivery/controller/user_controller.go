package controller

import (
	"eniqloStore/delivery/api"
	"eniqloStore/model/dto"
	"eniqloStore/usecase"
	"eniqloStore/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucRegist  usecase.UserRegistrationUsecase
	ucLogin   usecase.UserLoginUsecase
	api.BaseApi
}

func (u *UserController) userLogin(c *gin.Context) {
	var bodyRequest dto.RequestLoginBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, err := u.ucLogin.StaffLogin(bodyRequest)
	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "User logged successfully "
	u.Success(c, data, detailMsg, "login")
}

func (u *UserController) userRegister(c *gin.Context) {
	var bodyRequest dto.RequestRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, err := u.ucRegist.StaffRegister(bodyRequest)
	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "User register successfully "
	u.Success(c, data, detailMsg, "register")

}

func NewUserController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucRegist usecase.UserRegistrationUsecase, ucLogin usecase.UserLoginUsecase) *UserController {
	controller := UserController{
		router:    router,
		routerDev: routerDev,

		ucRegist: ucRegist,
		ucLogin:  ucLogin,

		BaseApi: api.BaseApi{},
	}

	router.POST("/v1/staff/register", controller.userRegister)

	router.POST("/v1/staff/login", controller.userLogin)

	return &controller
}
