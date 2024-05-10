package usecase

import (
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"
	"fmt"
)

type ProductUsecase interface {
	GetProduct(data dto.ProductQueryParams) ([]dto.ResponseGetProduct, error)
	CreateProduct(data dto.RequestProduct) (dto.ProductInfo, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func (p *productUsecase) GetProduct(data dto.ProductQueryParams) ([]dto.ResponseGetProduct, error) {
	// Call the repository to get the list of products

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

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	usecase := new(productUsecase)

	usecase.productRepo = productRepo

	return usecase
}
