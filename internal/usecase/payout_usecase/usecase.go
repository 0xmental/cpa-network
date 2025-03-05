package payout_usecase

import (
	"CPAPlatform/internal/adapter/repository/partner_in_memory"
	"CPAPlatform/internal/adapter/repository/payout_in_memory"
)

type UseCase struct {
	payoutRepo  *payout_in_memory.Repo
	partnerRepo *partner_in_memory.Repo
}

func NewUseCase(payoutRepo *payout_in_memory.Repo, partnerRepo *partner_in_memory.Repo) *UseCase {
	return &UseCase{
		payoutRepo:  payoutRepo,
		partnerRepo: partnerRepo,
	}
}
