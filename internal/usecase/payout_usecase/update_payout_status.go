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
	payout, err := u.payoutRepo.GetPayoutByID(req.PayoutID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPayoutByID: %w", err)
	}
	payout.Status = req.Status
	payout.UpdateAt = u.timer.Now()

	return u.payoutRepo.UpdatePayoutStatus(payout), nil
}
