package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (u *UseCase) GetPayoutByID(payoutID int64) (*domain.Payout, error) {
	payout, err := u.payoutRepo.GetPayoutByID(payoutID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}

	return payout, nil
}
