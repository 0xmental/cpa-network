package offer_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllOffers() []*domain.Offer {
	return u.offerRepo.GetAllOffers()
}
