package manager

type UsecaseManager interface {
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
