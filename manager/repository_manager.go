package manager

type RepositoryManager interface {
}

type repositoryManager struct {
	infra Infra
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
