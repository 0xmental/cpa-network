package partner_usecase

import "CPAPlatform/internal/adapter/repository/partner_in_memory"

type UseCase struct {
	repo *partner_in_memory.Repo
}

func NewUseCase(repo *partner_in_memory.Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
