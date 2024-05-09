package usecase

import (
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"
	"fmt"
)

type ProductUsecase interface {
	CreateProduct(data dto.RequestProduct) (dto.ProductInfo, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
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
