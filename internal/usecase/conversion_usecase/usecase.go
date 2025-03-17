package conversion_usecase

import (
	"CPAPlatform/internal/domain"
)

type (
	UseCase struct {
		repoClick      clickRepo
		repoOffer      offerRepo
		repoPartner    partnerRepo
		repoConversion conversionRepo
	}

	clickRepo interface {
		GetByClickID(clickID string) (*domain.Click, error)
	}

	offerRepo interface {
		GetOfferByID(offerID int64) (*domain.Offer, error)
	}

	partnerRepo interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
	}

	conversionRepo interface {
		Save(conversion *domain.Conversion) *domain.Conversion
		GetAllConversions() []*domain.Conversion
		GetAllConversionsByOffer(offerID int64) []*domain.Conversion
		GetAllConversionsByPartner(partnerID int64) []*domain.Conversion
	}
)

func NewUseCase(repoClick clickRepo, repoOffer offerRepo,
	repoPartner partnerRepo, repoConversion conversionRepo) *UseCase {
	return &UseCase{
		repoClick:      repoClick,
		repoOffer:      repoOffer,
		repoPartner:    repoPartner,
		repoConversion: repoConversion,
	}
}
