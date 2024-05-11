package controller

import (
	"eniqloStore/delivery/api"
	"eniqloStore/model/dto"
	"eniqloStore/usecase"
	"eniqloStore/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router     *gin.RouterGroup
	routerDev  *gin.RouterGroup
	ucCustomer usecase.CustomerUsecase
	ucRegist   usecase.UserRegistrationUsecase
	ucLogin    usecase.UserLoginUsecase
	api.BaseApi
}

func (u *UserController) getCustomer(c *gin.Context) {
	phoneNumber := c.Query("phoneNumber ")
	name := c.Query("name")

	data, err := u.ucCustomer.FindCustomer(phoneNumber, name)

	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "success"

	u.Success(c, data, detailMsg, "")
}

func (u *UserController) customerRegister(c *gin.Context) {
	var bodyRequest dto.RequestCustomerRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, err := u.ucCustomer.CreateCustomer(bodyRequest)
	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "Customer Created successfully "
	u.Success(c, data, detailMsg, "register")

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

func NewUserController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucCustomer usecase.CustomerUsecase, ucRegist usecase.UserRegistrationUsecase, ucLogin usecase.UserLoginUsecase) *UserController {
	controller := UserController{
		router:     router,
		routerDev:  routerDev,
		ucCustomer: ucCustomer,
		ucRegist:   ucRegist,

		ucLogin: ucLogin,
		BaseApi: api.BaseApi{},
	}

	router.GET("/v1/customer", controller.getCustomer)

	router.POST("/v1/customer/register", controller.customerRegister)

	router.POST("/v1/staff/register", controller.userRegister)
	router.POST("/v1/staff/login", controller.userLogin)

	return &controller
}
