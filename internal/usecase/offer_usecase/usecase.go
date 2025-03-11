package offer_usecase

import "CPAPlatform/internal/domain"

type (
	UseCase struct {
		repo repo
	}

	repo interface {
		GetAllOffers() []*domain.Offer
		GetOfferByID(offerID int64) (*domain.Offer, error)
		Save(offer *domain.Offer) *domain.Offer
		Update(offer *domain.Offer) *domain.Offer
	}
)

func NewUseCase(repo repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
