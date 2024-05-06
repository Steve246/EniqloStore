package usecase

import (
	"eniqloStore/model"
	"eniqloStore/repository"
	"eniqloStore/utils"
)

// import (
// 	"7Zero4/model"
// 	"7Zero4/model/dto"
// 	"7Zero4/repository"
// 	"7Zero4/utils"
// 	"time"
// )

type UserRegistrationUsecase interface {
	RegisterUser(reqRegistBody model.Staff) (model.StaffResponse, error)
}

type userRegistrationUsecase struct {
	userRepo     repository.UserRepository
	passWordRepo repository.PasswordRepository
	tokenRepo    repository.TokenRepository
}

func (p *userRegistrationUsecase) RegisterUser(reqUserData model.Staff) (model.StaffResponse, error) {

	// validation check request body
	errValidate := p.userRepo.ValidateUser(reqUserData.PhoneNumber, reqUserData.Name, reqUserData.Password, "register")
	if errValidate != nil {
		return model.StaffResponse{}, errValidate
	}

	// validation check email already registered
	exist := p.userRepo.FindByPhone(reqUserData.PhoneNumber)
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

	// insert to database
	err := p.userRepo.Register("staffdata", model.Staff{
		Name:     reqUserData.Name,
		Password: hashedPasswordStr,
		UserID:   "3",
		// RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return model.StaffResponse{}, err
	}

	generateUserId, err := utils.GenerateUserID()

	if err != nil {
		return model.StaffResponse{}, err
	}

	responseData := model.StaffResponse{
		UserID:      generateUserId,
		Name:        reqUserData.Name,
		PhoneNumber: "+6281711818",
		AccessToken: token,
	}

	return responseData, nil
}

func NewUserRegistrationUsecase(userRepo repository.UserRepository, passWordRepo repository.PasswordRepository, tokenRepo repository.TokenRepository) UserRegistrationUsecase {
	usecase := new(userRegistrationUsecase)
	usecase.userRepo = userRepo
	usecase.passWordRepo = passWordRepo
	usecase.tokenRepo = tokenRepo

	return usecase
}
