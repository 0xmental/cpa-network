package offer_usecase

import "CPAPlatform/internal/domain"

type (
	UseCase struct {
		offerRepo repoOffer
	}

	repoOffer interface {
		GetAllOffers() []*domain.Offer
		GetOfferByID(offerID int64) (*domain.Offer, error)
		Save(offer *domain.Offer) *domain.Offer
		Update(offer *domain.Offer) *domain.Offer
	}
)

func NewUseCase(offerRepo repoOffer) *UseCase {
	return &UseCase{
		offerRepo: offerRepo,
	}
}
