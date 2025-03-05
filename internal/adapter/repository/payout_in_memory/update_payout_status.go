package payout_in_memory

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (r *Repo) Update(payoutID int64, status domain.PayoutStatus) (*domain.Payout, error) {
	payout, err := r.GetPayoutByID(payoutID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}

	payout.Status = status

	return payout, nil
}
