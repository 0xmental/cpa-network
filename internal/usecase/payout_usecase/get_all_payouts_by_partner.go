package payout_usecase

import "CPAPlatform/internal/domain"

type GetPayoutsByPartnerReq struct {
	PartnerID int64
}

func (u *UseCase) GetAllPayoutsByPartnerID(req GetPayoutsByPartnerReq) []*domain.Payout {
	return u.payoutRepo.GetAllPayoutsByPartnerID(req.PartnerID)
}
