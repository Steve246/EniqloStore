package manager

import (
	"eniqloStore/usecase"
)

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
	ProductUsecase() usecase.ProductUsecase
	LoginUsecase() usecase.UserLoginUsecase
	TokenUsecase() usecase.TokenUsecase
	RegisterUsecase() usecase.UserRegistrationUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func (u *usecaseManager) ProductUsecase() usecase.ProductUsecase {
	return usecase.NewProductUsecase(u.repoManager.ProductRepo())
}

func (u *usecaseManager) LoginUsecase() usecase.UserLoginUsecase {
	return usecase.NewUserLoginUsecase(u.repoManager.TokenRepo(), u.repoManager.StaffRepo(), u.repoManager.PasswordRepo())
}

func (u *usecaseManager) TokenUsecase() usecase.TokenUsecase {
	return usecase.NewTokenUsecase(u.repoManager.TokenRepo())
}

func (u *usecaseManager) RegisterUsecase() usecase.UserRegistrationUsecase {
	return usecase.NewUserRegistrationUsecase(u.repoManager.StaffRepo(), u.repoManager.PasswordRepo(), u.repoManager.TokenRepo())

}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
