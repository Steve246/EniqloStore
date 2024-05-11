package usecase

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CustomerUsecase interface {
	validateCheckoutRequest(request dto.CheckoutRequest) error
	Checkout(request dto.CheckoutRequest) error
	FindCustomer(phoneNumber string, name string) ([]dto.CustomerResponse, error)
	CreateCustomer(data dto.RequestCustomerRegistBody) (dto.CustomerResponse, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
	productRepo  repository.ProductRepository
}

// validateCheckoutRequest validates the checkout request
func (p *customerUsecase) validateCheckoutRequest(request dto.CheckoutRequest) error {
	// Check if customer ID is empty
	if request.CustomerID == "" {
		return errors.New("customer ID is required")
	}

	// Check if product details are empty
	if len(request.ProductDetails) == 0 {
		return errors.New("at least one product detail is required")
	}

	// Check if paid amount is non-negative
	if request.Paid < 0 {
		return errors.New("paid amount cannot be negative")
	}

	// Check if change amount is non-negative
	if request.Change < 0 {
		return errors.New("change amount cannot be negative")
	}

	// Additional custom validations can be added here

	return nil
}

func (p *customerUsecase) Checkout(request dto.CheckoutRequest) error {
	// Validate request
	if err := p.validateCheckoutRequest(request); err != nil {
		return err
	}

	// Check if customer exists (not implemented in this example)

	// Check if all product IDs are valid
	for _, product := range request.ProductDetails {
		availability, err := p.productRepo.CheckAvailability(product.ProductID)
		if err != nil {
			return err
		}
		if !availability {
			return fmt.Errorf("product ID %s is not found", product.ProductID)
		}
	}

	// Validate paid amount
	totalPrice := 0
	for _, product := range request.ProductDetails {
		price, err := p.productRepo.GetProductPrice(product.ProductID)
		if err != nil {
			return err
		}
		totalPrice += price * product.Quantity
	}
	if totalPrice > request.Paid {
		return errors.New("paid amount is not enough")
	}

	// Decrease stock and perform checkout
	for _, product := range request.ProductDetails {
		err := p.productRepo.DecreaseStock(product.ProductID, product.Quantity)
		if err != nil {
			return err
		}
	}

	// Validate change
	if request.Paid-totalPrice != request.Change {
		return errors.New("incorrect change amount")
	}

	return nil
}

func (p *customerUsecase) FindCustomer(phoneNumber string, name string) ([]dto.CustomerResponse, error) {

	customers, err := p.customerRepo.SearchCustomers(phoneNumber, name)

	if err != nil {
		return nil, utils.GetCustomerError()
	}

	var responseCustomers []dto.CustomerResponse

	for _, customer := range customers {
		responseCustomer := dto.CustomerResponse{
			UserUniqueID: customer.UserUniqueID,
			Name:         customer.Name,
			PhoneNumber:  customer.PhoneNumber,
		}

		responseCustomers = append(responseCustomers, responseCustomer)
	}

	return responseCustomers, nil

}

func (p *customerUsecase) CreateCustomer(data dto.RequestCustomerRegistBody) (dto.CustomerResponse, error) {

	validation := p.customerRepo.ValidateUser(data.PhoneNumber, data.Name)

	if validation != nil {
		return dto.CustomerResponse{}, validation
	}

	// validation check email already registered
	exist := p.customerRepo.FindByPhone(data.PhoneNumber)
	if exist {
		return dto.CustomerResponse{}, utils.PhoneNumberFoundError()
	}

	generateUserId, err := utils.GenerateUserID()

	if err != nil {
		return dto.CustomerResponse{}, err
	}

	errInsert := p.customerRepo.Register("customerData", model.Customer{
		Model:        gorm.Model{},
		UserUniqueID: generateUserId,
		Name:         data.Name,
		PhoneNumber:  data.PhoneNumber,
	})

	if errInsert != nil {
		return dto.CustomerResponse{}, utils.CreateCustomerError()
	}

	NewData := dto.CustomerResponse{
		UserUniqueID: generateUserId,
		Name:         data.Name,
		PhoneNumber:  data.PhoneNumber,
	}

	return NewData, nil

}

func NewCustomerUsecase(customerRepo repository.CustomerRepository) CustomerUsecase {
	usecase := new(customerUsecase)
	usecase.customerRepo = customerRepo
	return usecase
}
