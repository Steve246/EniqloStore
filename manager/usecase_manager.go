package manager

import (
	"eniqloStore/usecase"
)

type UsecaseManager interface {
	TokenUsecase() usecase.TokenUsecase

	RegisterUsecase() usecase.UserRegistrationUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
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
