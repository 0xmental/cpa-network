package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type UpdatePayoutReq struct {
	PayoutID int64
	Status   domain.PayoutStatus
}

func (u *UseCase) UpdatePayoutStatus(req UpdatePayoutReq) (*domain.Payout, error) {
	result, err := u.payoutRepo.UpdatePayoutStatus(req.PayoutID, req.Status)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}

	return result, err
}
