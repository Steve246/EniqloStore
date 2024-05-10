package repository

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/utils"
	"errors"
	"net/url"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProduct()
	FindIdCreatedAtBy(requestData dto.RequestProduct) (model.ProductList, error)
	InsertProduct(requestData dto.RequestProduct) error
	ValidateProduct(requestData dto.RequestProduct) bool
	SearchProduct(params map[string]string) ([]model.ProductList, error)
}

type productdRepository struct {
	db *gorm.DB
}

// TODO: NAMBAIN SATU REPO UNTUK DAPETIN ID DAN CREATED_AT KAPAN

func (p *productdRepository) GetProduct() {

}

func (p *productdRepository) FindIdCreatedAtBy(requestData dto.RequestProduct) (model.ProductList, error) {

	var product model.ProductList

	p.db.Raw("SELECT * FROM ProductList WHERE name = ?", requestData.Name).Scan(&product)

	if (product == model.ProductList{}) {
		return model.ProductList{}, errors.New("Product Not Found")
	}

	return product, nil
}

func (p *productdRepository) InsertProduct(requestData dto.RequestProduct) error {
	result := p.db.Exec("INSERT INTO ProductList (name, sku, category, imageUrl, notes, price, stock, location, isAvailable) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", requestData.Name, requestData.SKU, requestData.Category, requestData.ImageURL, requestData.Notes, requestData.Price, requestData.Stock, requestData.Location, requestData.IsAvailable)
	if result.Error != nil {
		return utils.CreateProductError()
	}

	return nil
}

func (p *productdRepository) ValidateProduct(requestData dto.RequestProduct) bool {
	// Name validation
	if len(requestData.Name) < 1 || len(requestData.Name) > 30 {
		// return fmt.Errorf("name must be between 1 and 30 characters")
		return false
	}

	// SKU validation
	if len(requestData.SKU) < 1 || len(requestData.SKU) > 30 {
		// return fmt.Errorf("sku must be between 1 and 30 characters")
		return false
	}

	// Category validation
	validCategories := map[string]bool{
		"Clothing":    true,
		"Accessories": true,
		"Footwear":    true,
		"Beverages":   true,
	}
	if !validCategories[requestData.Category] {
		// return fmt.Errorf("invalid category")
		return false
	}

	// ImageURL validation
	_, err := url.ParseRequestURI(requestData.ImageURL)
	if err != nil {
		return false

	}

	// Notes validation
	if len(requestData.Notes) < 1 || len(requestData.Notes) > 200 {
		return false

	}

	// Price validation
	if requestData.Price < 1 {
		return false

	}

	// Stock validation
	if requestData.Stock < 0 || requestData.Stock > 100000 {
		return false

	}

	// Location validation
	if len(requestData.Location) < 1 || len(requestData.Location) > 200 {
		return false

	}

	return true
}

func generateLikeQuery(params string) string {
	result := "%" + params + "%"
	return result
}

func generateQuery(params map[string]string) string {
	var result string
	if category, ok := params["category"]; ok {
		result += " AND category = " + "'" + category + "'"
	}

	if sku, ok := params["sku"]; ok && params["sku"] != "" {
		result += " AND sku = " + "'" + sku + "'"
	}

	if params["inStock"] == "true" {
		result += " AND stock > 0"
	}

	if params["price"] == "asc" {
		result += " ORDER BY price ASC"
	} else if params["price"] == "desc" {
		result += " ORDER BY price DESC"
	} else {
		result += " ORDER BY created_at ASC"
	}

	result += " LIMIT " + params["limit"] + " OFFSET " + params["offset"]

	return result
}

func (p *productdRepository) SearchProduct(params map[string]string) ([]model.ProductList, error) {
	var product []model.ProductList
	additionalQuery := generateQuery(params)
	result := p.db.Raw(`SELECT id,name,sku,category,imageUrl,stock,price,location,created_at FROM productlist WHERE LOWER(name) LIKE ?`+additionalQuery, generateLikeQuery(params["name"])).Scan(&product)
	return product, result.Error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productdRepository)
	repo.db = db
	return repo
}
