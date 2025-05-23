package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"time"
)

type (
	UseCase struct {
		offerRepo repoOffer
		timer     timer
	}

	repoOffer interface {
		GetAllOffers() []*domain.Offer
		GetOfferByID(offerID int64) (*domain.Offer, error)
		Save(offer *domain.Offer) *domain.Offer
		Update(offer *domain.Offer) *domain.Offer
	}

	timer interface {
		Now() time.Time
	}
)

func NewUseCase(offerRepo repoOffer, timer timer) *UseCase {
	return &UseCase{
		offerRepo: offerRepo,
		timer:     timer,
	}
}
