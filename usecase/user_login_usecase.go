package usecase

import (
	"eniqloStore/model"
	"eniqloStore/model/dto"
	"eniqloStore/repository"
	"eniqloStore/utils"
)

type UserLoginUsecase interface {
	StaffLogin(reqLoginBody dto.RequestLoginBody) (model.StaffResponse, error)
}

type userLoginUsecase struct {
	tokenRepo    repository.TokenRepository
	userRepo     repository.StaffRepository
	passWordRepo repository.PasswordRepository
}

func (u *userLoginUsecase) StaffLogin(reqLoginBody dto.RequestLoginBody) (model.StaffResponse, error) {

	errValidate := u.userRepo.ValidateUser(reqLoginBody.PhoneNumber, "", reqLoginBody.Password, "login")
	if errValidate != nil {
		return model.StaffResponse{}, errValidate
	}

	dbPass, errdbPass := u.userRepo.FindPasswordByPhoneNumber(reqLoginBody.PhoneNumber)
	if errdbPass != nil {
		return model.StaffResponse{}, utils.UserNotFoundError()
	}

	errPassword := u.passWordRepo.VerifyPassword([]byte(dbPass.Password), []byte(reqLoginBody.Password))
	if errPassword != nil {
		return model.StaffResponse{}, utils.PasswordWrongError()
	}

	// TODO: kalau udh register, dan masih ada token aktif dia gak create

	// Get token auth
	token, tokenErr := u.tokenRepo.CreateTokenV2(dbPass.UserUniqueID, 12)
	if tokenErr != nil {
		return model.StaffResponse{}, tokenErr
	}

	// Populate the success data struct
	successData := model.StaffResponse{
		UserUniqueID: dbPass.UserUniqueID,
		Name:         dbPass.Name,
		PhoneNumber:  dbPass.PhoneNumber,
		AccessToken:  token,
	}

	return successData, nil
}

func NewUserLoginUsecase(
	tokenRepo repository.TokenRepository,
	staffRepo repository.StaffRepository,
	passWordRepo repository.PasswordRepository) UserLoginUsecase {
	usecase := new(userLoginUsecase)

	usecase.tokenRepo = tokenRepo
	usecase.userRepo = staffRepo
	usecase.passWordRepo = passWordRepo
	return usecase
}
