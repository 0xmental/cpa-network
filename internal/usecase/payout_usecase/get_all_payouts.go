package payout_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllPayouts() []*domain.Payout {
	return u.payoutRepo.GetAllPayouts()
}
