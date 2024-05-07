package manager

import "eniqloStore/repository"

type RepositoryManager interface {
	TokenRepo() repository.TokenRepository
	StaffRepo() repository.StaffRepository
	PasswordRepo() repository.PasswordRepository
}

func (r *repositoryManager) TokenRepo() repository.TokenRepository {
	return repository.NewTokenRepository(r.infra.TokenConfig(), r.infra.SqlDb())
}

func (r *repositoryManager) StaffRepo() repository.StaffRepository {
	return repository.NewStaffRepository(r.infra.SqlDb())
}

func (r *repositoryManager) PasswordRepo() repository.PasswordRepository {
	return repository.NewPasswordRepository(r.infra.SqlDb())
}

type repositoryManager struct {
	infra Infra
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
