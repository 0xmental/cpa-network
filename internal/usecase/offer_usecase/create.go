package offer_usecase

import (
	"CPAPlatform/internal/domain"
)

type CreateOfferRequest struct {
	TargetURL      string
	Name           string
	Description    string
	RedirectDomain string
	ConversionType domain.ConversionType
	Payout         map[string]int64
}

func (u *UseCase) CreateOffer(req CreateOfferRequest) *domain.Offer {
	now := u.timer.Now()
	offer := domain.NewOffer(req.TargetURL, req.Name, req.Description, req.RedirectDomain, req.ConversionType, req.Payout, now)

	return u.offerRepo.Save(offer)
}
