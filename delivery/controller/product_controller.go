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
	router     *gin.RouterGroup
	routerDev  *gin.RouterGroup
	ucProduct  usecase.ProductUsecase
	ucCustomer usecase.CustomerUsecase
	api.BaseApi
}

func (u *ProductController) CheckoutCustomer(c *gin.Context) {
	var bodyRequest dto.CheckoutRequest

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	err := u.ucCustomer.Checkout(bodyRequest)
	if err != nil {
		u.Failed(c, err)
		return
	}

	u.Success(c, nil, "Successfully checked out product", "")
}

func (u *ProductController) UpdateProduct(c *gin.Context) {
	var bodyRequest dto.RequestProduct
	id := c.Param("id")

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	err := u.ucProduct.UpdateProduct(bodyRequest, id)

	if err != nil {
		if err == utils.CreateProductError() {

			u.Failed(c, err)

		}

		u.Failed(c, utils.ServerError())

	} else {
		detailMsg := "successfully edit product"

		u.Success(c, "", detailMsg, "")
	}
}

func (u *ProductController) delProduct(c *gin.Context) {
	id := c.Param("id")

	err := u.ucProduct.DeleteProduct(id)

	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "Succesfully Delete Product"

	u.Success(c, "", detailMsg, "")

}

func (u *ProductController) getProduct(c *gin.Context) {

	// requestParams := c.Request.URL.Query()

	id := c.Query("id")
	name := c.Query("name")
	limit := utils.StringToInt(c.DefaultQuery("limit", "5"))
	offset := utils.StringToInt(c.DefaultQuery("offset", "0"))
	isAvailable := c.Query("isavailable")
	category := c.Query("category")
	sku := c.Query("sku")
	price := c.Query("price")
	inStock := c.Query("inStock")
	createdAt := c.Query("createdAt")

	bodyRequestData := dto.ProductQueryParams{
		ID:          id,
		Name:        name,
		Limit:       limit,
		Offset:      offset,
		IsAvailable: isAvailable,
		Category:    category,
		SKU:         sku,
		Price:       price,
		InStock:     inStock,
		CreatedAt:   createdAt,
	}

	fmt.Println("ini request data first --> ", bodyRequestData)

	// TODO: add validation via params

	// if err := u.ParseRequestBody(c, &bodyRequestData); err != nil {
	// 	u.Failed(c, utils.ReqBodyNotValidError())
	// 	return
	// }

	fmt.Println("ini bodyRequest second --> ", bodyRequestData)

	data, err := u.ucProduct.GetProduct(bodyRequestData)

	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "success"

	u.Success(c, data, detailMsg, "")

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

func NewProductController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucProduct usecase.ProductUsecase, ucCustomer usecase.CustomerUsecase) *ProductController {
	controller := ProductController{
		router:    router,
		routerDev: routerDev,

		ucProduct: ucProduct,

		BaseApi: api.BaseApi{},
	}

	router.POST("/v1/product/checkout", controller.CheckoutCustomer)

	router.PUT("/v1/product/:id", controller.UpdateProduct)

	router.DELETE("/v1/product/:id", controller.delProduct)

	router.GET("/v1/product/customer", controller.searchProduct)

	router.GET("/v1/product", controller.getProduct)

	router.POST("/v1/product", controller.createProduct)

	// router.POST("/v1/staff/login", controller.userLogin)

	return &controller
}
