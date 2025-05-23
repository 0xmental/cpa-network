package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"fmt"
)

type UpdatePayoutReq struct {
	PayoutID int64
	Status   domain.PayoutStatus
}

func (u *UseCase) UpdatePayoutStatus(req UpdatePayoutReq) (*domain.Payout, error) {

	payouts := u.payoutRepo.GetAllPayouts(dto.PayoutFilter{
		PayoutID: req.PayoutID,
	})

	if len(payouts) == 0 {
		return nil, fmt.Errorf("payout with ID %d not found", req.PayoutID)
	}
	payout := payouts[0]
	payout.Status = req.Status
	payout.UpdateAt = u.timer.Now()

	return u.payoutRepo.UpdatePayoutStatus(payout), nil
}
