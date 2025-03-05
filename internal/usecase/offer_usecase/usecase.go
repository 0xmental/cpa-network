package offer_usecase

import "CPAPlatform/internal/adapter/repository/offer_in_memory"

type UseCase struct {
	repo *offer_in_memory.Repo
}

func NewUseCase(repo *offer_in_memory.Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
