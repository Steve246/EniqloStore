package controller

import (
	"eniqloStore/delivery/api"
	"eniqloStore/model/dto"
	"eniqloStore/usecase"
	"eniqloStore/utils"
	"fmt"

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
	fmt.Println("berhasil masuk create Produt")
	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, createProducErr := u.ucProduct.CreateProduct(bodyRequest)

	fmt.Println("ini isi data --> ", data)

	if createProducErr != nil {
		if createProducErr == utils.CreateProductError() {

			u.Failed(c, createProducErr)

		}

		u.Failed(c, utils.ServerError())

	} else {
		detailMsg := "User logged successfully "

		u.Success(c, data, detailMsg, "")
	}

}

func (u *ProductController) searchProduct(c *gin.Context) {
	requestParams := c.Request.URL.Query()
	res, err := u.ucProduct.SearchProduct(requestParams)
	if err != nil {
		u.Failed(c, err)
		return
	}

	u.Success(c, res, "success", "")
}

func NewProductController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucProduct usecase.ProductUsecase) *ProductController {
	controller := ProductController{
		router:    router,
		routerDev: routerDev,

		ucProduct: ucProduct,

		BaseApi: api.BaseApi{},
	}

	router.POST("/v1/create", controller.createProduct)
	router.GET("/v1/product/customer", controller.searchProduct)

	return &controller
}
