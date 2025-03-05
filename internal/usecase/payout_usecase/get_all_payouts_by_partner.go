package payout_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllPayoutsByPartnerID(partnerID int64) []*domain.Payout {
	payouts := u.payoutRepo.GetByPartnerID(partnerID)

	return payouts
}
