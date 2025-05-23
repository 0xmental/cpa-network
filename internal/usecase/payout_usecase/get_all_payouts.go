package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

type GetAllPayoutsReq struct {
	PartnerID int64
	PayoutID  int64
	Status    domain.PayoutStatus
}

func (u *UseCase) GetAllPayouts(req GetAllPayoutsReq) []*domain.Payout {
	return u.payoutRepo.GetAllPayouts(dto.PayoutFilter{
		PartnerID: req.PartnerID,
		PayoutID:  req.PayoutID,
		Status:    req.Status,
	})
}
