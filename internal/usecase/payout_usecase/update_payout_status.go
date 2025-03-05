package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (u *UseCase) UpdatePayoutStatus(payoutID int64, status domain.PayoutStatus) (*domain.Payout, error) {
	result, err := u.UpdatePayoutStatus(payoutID, status)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}

	return result, err
}
