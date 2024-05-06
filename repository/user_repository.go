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

type UserRepository interface {
	FindPasswordByPhoneNumber(phoneNumber string) (model.Staff, error)
	Register(tableName string, data model.Staff) error
	FindByPhone(phoneNumber string) bool
	ValidateUser(email string, name string, password string, user string) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) ValidateUser(email string, name string, password string, user string) error {

	if user == "register" {
		if email == "" {
			return utils.ReqBodyNotValidError()
		}

		fmt.Println("ini email --> ", email)

		// emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

		emailRegex := `(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		match, _ := regexp.MatchString(emailRegex, email)
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
		if email == "" {
			return utils.ReqBodyNotValidError()
		}

		emailRegex := `(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		match, _ := regexp.MatchString(emailRegex, email)
		if !match {
			return utils.ReqBodyNotValidError()
		}

		// Check if name is not null and length is between 5 and 50
		// if name == "" {
		// 	return utils.ErrNameNull
		// }
		// nameLength := len(strings.TrimSpace(name))
		// if nameLength < 5 || nameLength > 50 {
		// 	return utils.ErrInvalidName
		// }

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

	return nil

}

func (u *userRepository) FindPasswordByPhoneNumber(phoneNumber string) (model.Staff, error) {
	var user model.Staff
	u.db.Raw("SELECT * FROM staffdata WHERE phone_number = ?", phoneNumber).Scan(&user)
	if (user == model.Staff{}) {
		return model.Staff{}, errors.New("User not found")
	}
	return user, nil

}

func (u *userRepository) Register(tableName string, data model.Staff) error {
	result := u.db.Exec("INSERT INTO "+tableName+" (created_at, updated_at, name, password, user_unique_id) VALUES (?, ?, ?, ?, ?)", data.CreatedAt, data.UpdatedAt, data.Name, data.Password, data.UserID)
	return result.Error
}

func (u *userRepository) FindByPhone(phoneNumber string) bool {
	result := u.db.Exec("SELECT * FROM staffdata WHERE phone_number = ?", phoneNumber)
	return result.RowsAffected != 0
}

func NewUserRepository(db *gorm.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
