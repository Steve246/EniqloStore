package controller

import (
	"eniqloStore/delivery/api"
	"eniqloStore/model/dto"
	"eniqloStore/usecase"
	"eniqloStore/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucProduct usecase.ProductUsecase
	api.BaseApi
}

func (u *ProductController) createProduct(c *gin.Context) {
	var bodyRequest dto.RequestProduct
	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, createProducErr := u.ucProduct.CreateProduct(bodyRequest)

	if createProducErr != nil {
		if createProducErr == utils.CreateProductError() {

			u.Failed(c, createProducErr)

		}

		u.Failed(c, utils.ServerError())

	}

	detailMsg := "User logged successfully "

	u.Success(c, data, detailMsg, "")
}

func NewProductController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucProduct usecase.ProductUsecase) *ProductController {
	controller := ProductController{
		router:    router,
		routerDev: routerDev,

		ucProduct: ucProduct,

		BaseApi: api.BaseApi{},
	}

	router.POST("/v1/create", controller.createProduct)

	// router.POST("/v1/staff/login", controller.userLogin)

	return &controller
}
