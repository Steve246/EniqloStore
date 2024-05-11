package manager

import "eniqloStore/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
	ProductRepo() repository.ProductRepository
	TokenRepo() repository.TokenRepository
	StaffRepo() repository.StaffRepository
	PasswordRepo() repository.PasswordRepository
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
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
