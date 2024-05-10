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

	router.GET("/v1/product/customer", controller.searchProduct)

	router.GET("/v1/product", controller.getProduct)

	router.POST("/v1/product", controller.createProduct)

	// router.POST("/v1/staff/login", controller.userLogin)

	return &controller
}
