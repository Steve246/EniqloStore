package usecase

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"
	"fmt"
	"net/url"
	"strconv"
)

type ProductUsecase interface {
	GetProduct(data dto.ProductQueryParams) ([]dto.ResponseGetProduct, error)
	CreateProduct(data dto.RequestProduct) (dto.ProductInfo, error)
	SearchProduct(params url.Values) ([]model.ProductList, error)
}

type productUsecase struct {
	productRepo  repository.ProductRepository
	categoryEnum map[string]struct{}
	priceEnum    map[string]struct{}
	inStockEnum  map[string]struct{}
}

func (p *productUsecase) GetProduct(data dto.ProductQueryParams) ([]dto.ResponseGetProduct, error) {
	// Call the repository to get the list of products
	fmt.Println("ini isi getAvailable --> ", data.IsAvailable)
	products, err := p.productRepo.GetProduct(data)
	if err != nil {
		return nil, utils.GetProductError()
	}

	// Map products to dto.ResponseGetProduct
	var responseProducts []dto.ResponseGetProduct
	for _, product := range products {
		responseProduct := dto.ResponseGetProduct{
			ID:          product.ID,
			Name:        product.Name,
			SKU:         product.SKU,
			Category:    product.Category,
			ImageURL:    product.ImageURL,
			Notes:       product.Notes,
			Price:       product.Price,
			Stock:       product.Stock,
			Location:    product.Location,
			IsAvailable: product.IsAvailable,
			CreatedAt:   product.CreatedAt,
		}
		responseProducts = append(responseProducts, responseProduct)
	}

	return responseProducts, nil
}

func (p *productUsecase) CreateProduct(data dto.RequestProduct) (dto.ProductInfo, error) {
	validation := p.productRepo.ValidateProduct(data)
	fmt.Println("ini hasil validasi --> ", validation)
	if !validation {
		return dto.ProductInfo{}, utils.ReqBodyNotValidError()
	}

	err := p.productRepo.InsertProduct(data)
	if err != nil {
		return dto.ProductInfo{}, utils.CreateProductError()
	}

	dataProduct, errProduct := p.productRepo.FindIdCreatedAtBy(data)

	if errProduct != nil {
		return dto.ProductInfo{}, utils.CreateProductError()

	}

	NewData := dto.ProductInfo{
		ID:        dataProduct.ID,
		CreatedAt: dataProduct.CreatedAt,
	}

	return NewData, nil

}

func (p *productUsecase) validateSearchProductParams(params map[string]string) map[string]string {
	result := make(map[string]string, 0)

	// default value limit and offset
	limit, _ := strconv.Atoi(params["limit"])
	if limit == 0 {
		result["limit"] = "5"
	} else {
		result["limit"] = params["limit"]
	}

	offset, _ := strconv.Atoi(params["offset"])
	if offset == 0 {
		result["offset"] = "0"
	} else {
		result["offset"] = params["offset"]
	}

	result["name"] = params["name"]
	result["sku"] = params["sku"]
	result["inStock"] = params["inStock"]
	result["price"] = params["price"]

	if category, ok := params["category"]; ok {
		if _, categoryExist := p.categoryEnum[category]; categoryExist {
			result["category"] = category
		}
	}

	return result
}

func parseURLValuesToMap(params url.Values) map[string]string {
	result := make(map[string]string, 0)

	for key, val := range params {
		result[key] = val[0]
	}
	return result
}

func (p *productUsecase) SearchProduct(params url.Values) ([]model.ProductList, error) {
	validateParams := p.validateSearchProductParams(parseURLValuesToMap(params))
	return p.productRepo.SearchProduct(validateParams)
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	usecase := new(productUsecase)
	usecase.productRepo = productRepo
	usecase.categoryEnum = map[string]struct{}{
		"Clothing":    {},
		"Accessories": {},
		"Footwear":    {},
		"Beverages":   {},
	}
	usecase.priceEnum = map[string]struct{}{
		"asc":  {},
		"desc": {},
	}
	usecase.inStockEnum = map[string]struct{}{
		"true":  {},
		"false": {},
	}

	return usecase
}
