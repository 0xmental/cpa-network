package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type GetPayoutReq struct {
	PayoutID int64
}

func (u *UseCase) GetPayoutByID(req GetPayoutReq) (*domain.Payout, error) {
	payout, err := u.payoutRepo.GetPayoutByID(req.PayoutID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}

	return payout, nil
}
