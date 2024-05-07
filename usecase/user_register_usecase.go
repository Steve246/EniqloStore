package usecase

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"

	"gorm.io/gorm"
)

// import (
// 	"7Zero4/model"
// 	"7Zero4/model/dto"
// 	"7Zero4/repository"
// 	"7Zero4/utils"
// 	"time"
// )

type UserRegistrationUsecase interface {
	StaffRegister(reqRegistBody dto.RequestRegistBody) (model.StaffResponse, error)
}

type userRegistrationUsecase struct {
	staffRepo    repository.StaffRepository
	passWordRepo repository.PasswordRepository
	tokenRepo    repository.TokenRepository
}

func (p *userRegistrationUsecase) StaffRegister(reqUserData dto.RequestRegistBody) (model.StaffResponse, error) {

	// validation check request body
	errValidate := p.staffRepo.ValidateUser(reqUserData.PhoneNumber, reqUserData.Name, reqUserData.Password, "register")
	if errValidate != nil {
		return model.StaffResponse{}, errValidate
	}

	// validation check email already registered
	exist := p.staffRepo.FindByPhone(reqUserData.PhoneNumber)
	if exist {
		return model.StaffResponse{}, utils.EmailFoundError()
	}

	// Hash the password
	hashedPasswordStr, errHashed := p.passWordRepo.HashAndSavePassword(reqUserData.Password)
	if errHashed != nil {
		return model.StaffResponse{}, utils.PasswordCannotBeEncodeError()
	}

	// Get token auth
	token, tokenErr := p.tokenRepo.CreateTokenV2(reqUserData.Name, 12)
	if tokenErr != nil {
		return model.StaffResponse{}, tokenErr
	}

	generateUserId, err := utils.GenerateUserID()

	if err != nil {
		return model.StaffResponse{}, err
	}

	// insert to database
	errInsert := p.staffRepo.Register("staffdata", model.Staff{
		Model:       gorm.Model{},
		UserID:      generateUserId,
		Name:        reqUserData.Name,
		PhoneNumber: reqUserData.PhoneNumber,
		Password:    hashedPasswordStr,
	})

	if errInsert != nil {
		return model.StaffResponse{}, errInsert

	}

	responseData := model.StaffResponse{
		UserID:      generateUserId,
		Name:        reqUserData.Name,
		PhoneNumber: reqUserData.PhoneNumber,
		AccessToken: token,
	}

	return responseData, nil
}

func NewUserRegistrationUsecase(staffRepo repository.StaffRepository, passWordRepo repository.PasswordRepository, tokenRepo repository.TokenRepository) UserRegistrationUsecase {
	usecase := new(userRegistrationUsecase)
	usecase.staffRepo = staffRepo
	usecase.passWordRepo = passWordRepo
	usecase.tokenRepo = tokenRepo

	return usecase
}
