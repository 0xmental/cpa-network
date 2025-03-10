package click_usecase

import "CPAPlatform/internal/domain"

type (
	UseCase struct {
		repoClick   clickRepo
		repoPartner partnerRepo
		repoOffer   offerRepo
	}
	clickRepo interface {
		GetAllClicksByPartner(partnerID int64) []*domain.Click
		GetAllClicksByOffer(offerID int64) []*domain.Click
		GetByClickID(clickID string) (*domain.Click, error)
		GetAllClicks() []*domain.Click
		IsUnique(clickID string) bool
		Save(click *domain.Click) *domain.Click
	}

	partnerRepo interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
	}

	offerRepo interface {
		GetOfferByID(offerID int64) (*domain.Offer, error)
	}
)

func NewUseCase(clickRepo clickRepo, offerRepo offerRepo, partnerRepo partnerRepo) *UseCase {
	return &UseCase{
		repoClick:   clickRepo,
		repoPartner: partnerRepo,
		repoOffer:   offerRepo,
	}
}
