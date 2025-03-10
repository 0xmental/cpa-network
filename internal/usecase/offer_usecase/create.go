package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"time"
)

type CreateOfferRequest struct {
	TargetURL      string
	Name           string
	Description    string
	isActive       bool
	RedirectDomain string
	ConversionType domain.ConversionType
	Payout         map[string]int64
}

func (u *UseCase) CreateOffer(req CreateOfferRequest) *domain.Offer {
	now := time.Now()
	offer := domain.NewOffer(req.TargetURL, req.Name, req.Description, req.RedirectDomain, req.ConversionType, req.Payout, now)

	return u.repo.Save(offer)
}
