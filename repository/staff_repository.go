package repository

import (
	"eniqloStore/model"
	"eniqloStore/utils"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type StaffRepository interface {
	FindPasswordByPhoneNumber(phoneNumber string) (model.Staff, error)
	Register(tableName string, data model.Staff) error
	FindByPhone(phoneNumber string) bool
	ValidateUser(phoneNumber string, name string, password string, user string) error
}

type staffRepository struct {
	db *gorm.DB
}

func (u *staffRepository) ValidateUser(phoneNumber string, name string, password string, user string) error {

	if user == "register" {
		if phoneNumber == "" {
			return utils.ReqBodyNotValidError()
		}

		phoneRegex := `^\+(?:[0-9] ?){6,14}[0-9]$`
		match, _ := regexp.MatchString(phoneRegex, phoneNumber)

		fmt.Println("ini match --> ", match)
		if !match {
			return utils.ReqBodyNotValidError()
		}

		// Check if name is not null and length is between 5 and 50
		if name == "" {
			return utils.ReqBodyNotValidError()
		}
		nameLength := len(strings.TrimSpace(name))

		fmt.Println("ini nameLength --> ", nameLength)

		if nameLength < 5 || nameLength > 50 {
			return utils.ReqBodyNotValidError()
		}

		// Check if password is not null and length is between 5 and 15
		if password == "" {
			return utils.ReqBodyNotValidError()
		}
		passwordLength := len(password)
		if passwordLength < 5 || passwordLength > 15 {
			return utils.ReqBodyNotValidError()
		}

		return nil
	}

	if user == "login" {
		if phoneNumber == "" {
			return utils.ReqBodyNotValidError()
		}

		phoneRegex := `^\+(?:[0-9] ?){6,14}[0-9]$`
		match, _ := regexp.MatchString(phoneRegex, phoneNumber)

		if !match {
			return utils.ReqBodyNotValidError()
		}

		if password == "" {
			return utils.ReqBodyNotValidError()
		}
		passwordLength := len(password)
		if passwordLength < 5 || passwordLength > 15 {
			return utils.ReqBodyNotValidError()
		}

		return nil
	}

	return nil

}

func (u *staffRepository) FindPasswordByPhoneNumber(phoneNumber string) (model.Staff, error) {
	var user model.Staff
	u.db.Raw("SELECT * FROM staffdata WHERE phone_number = ?", phoneNumber).Scan(&user)
	if (user == model.Staff{}) {
		return model.Staff{}, errors.New("User not found")
	}
	return user, nil

}

func (u *staffRepository) Register(tableName string, data model.Staff) error {
	result := u.db.Exec("INSERT INTO "+tableName+" (created_at, updated_at, name, password, user_unique_id, phone_number) VALUES (?, ?, ?, ?, ?, ?)", data.CreatedAt, data.UpdatedAt, data.Name, data.Password, data.UserID, data.PhoneNumber)

	fmt.Println("ini result --> ", result)
	return result.Error
}

func (u *staffRepository) FindByPhone(phoneNumber string) bool {
	result := u.db.Exec("SELECT * FROM staffdata WHERE phone_number = ?", phoneNumber)
	return result.RowsAffected != 0
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	repo := new(staffRepository)
	repo.db = db
	return repo
}
