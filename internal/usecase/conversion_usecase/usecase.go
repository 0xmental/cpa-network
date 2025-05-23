package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"time"
)

type (
	UseCase struct {
		repoClick      clickRepo
		repoOffer      offerRepo
		repoPartner    partnerRepo
		repoConversion conversionRepo
		timer          timer
	}

	clickRepo interface {
		GetAllClicks(filter dto.ClickFilter) []*domain.Click
	}

	offerRepo interface {
		GetOfferByID(offerID int64) (*domain.Offer, error)
	}

	partnerRepo interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
	}

	conversionRepo interface {
		Save(conversion *domain.Conversion) *domain.Conversion
		GetAllConversions(filter dto.ConversionFilter) []*domain.Conversion
	}

	timer interface {
		Now() time.Time
	}
)

func NewUseCase(repoClick clickRepo, repoOffer offerRepo,
	repoPartner partnerRepo, repoConversion conversionRepo, timer timer) *UseCase {
	return &UseCase{
		repoClick:      repoClick,
		repoOffer:      repoOffer,
		repoPartner:    repoPartner,
		repoConversion: repoConversion,
		timer:          timer,
	}
}
