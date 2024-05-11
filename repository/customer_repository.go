package repository

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/utils"
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	SearchCustomers(phoneNumber string, name string) ([]dto.CustomerResponse, error)

	FindByPhone(phoneNumber string) bool

	FindUserByPhoneNumber(phoneNumber string) (model.Customer, error)

	Register(tableName string, data model.Customer) error

	ValidateUser(phoneNumber string, name string) error
}

type customerRepository struct {
	db *gorm.DB
}

func (r *customerRepository) SearchCustomers(phoneNumber string, name string) ([]dto.CustomerResponse, error) {
	query := "SELECT userId, phoneNumber, name FROM customer WHERE deleted_at IS NULL"
	args := []interface{}{}

	if phoneNumber != "" {
		query += " AND phoneNumber LIKE ?"
		args = append(args, "%"+phoneNumber+"%")
	}
	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	var customers []dto.CustomerResponse
	if err := r.db.Raw(query, args...).Scan(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (u *customerRepository) FindByPhone(phoneNumber string) bool {
	result := u.db.Exec("SELECT * FROM customerData WHERE phone_number = ?", phoneNumber)
	return result.RowsAffected != 0
}

func (u *customerRepository) FindUserByPhoneNumber(phoneNumber string) (model.Customer, error) {
	var user model.Customer
	u.db.Raw("SELECT * FROM customerData WHERE phone_number = ?", phoneNumber).Scan(&user)
	if (user == model.Customer{}) {
		return model.Customer{}, errors.New("User not found")
	}
	return user, nil

}

func (u *customerRepository) ValidateUser(phoneNumber string, name string) error {

	if phoneNumber == "" {
		return utils.ReqBodyNotValidError()
	}
	// `^\+\d{1,3}-?\d{1,14}$`
	// phoneRegex := `^\+(?:[0-9] ?){6,14}[0-9]$`
	phoneRegex := `^\+\d{1,3}-?\d{1,14}$`

	match, _ := regexp.MatchString(phoneRegex, phoneNumber)

	if !match {
		return utils.ReqBodyNotValidError()
	}

	// Check if name is not null and length is between 5 and 50
	if name == "" {
		return utils.ReqBodyNotValidError()
	}
	nameLength := len(strings.TrimSpace(name))

	if nameLength < 5 || nameLength > 50 {
		return utils.ReqBodyNotValidError()
	}

	return nil

}

func (u *customerRepository) Register(tableName string, data model.Customer) error {
	result := u.db.Exec("INSERT INTO "+tableName+" (created_at, updated_at, name, user_unique_id, phone_number) VALUES (?, ?, ?, ?, ?, ?)", data.CreatedAt, data.UpdatedAt, data.Name, data.UserUniqueID, data.PhoneNumber)

	return result.Error
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
