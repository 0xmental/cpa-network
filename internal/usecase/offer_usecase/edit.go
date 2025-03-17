package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
	"time"
)

type UpdateOfferReq struct {
	OfferID        int64
	TargetURL      string
	Name           string
	Description    string
	RedirectDomain string
	ConversionType domain.ConversionType
	Payout         map[string]int64
}

func (u *UseCase) EditOffer(req UpdateOfferReq) (*domain.Offer, error) {
	offer, err := u.offerRepo.GetOfferByID(req.OfferID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}

	offer.TargetUrl = req.TargetURL
	offer.Name = req.Name
	offer.Description = req.Description
	offer.RedirectDomain = req.RedirectDomain
	offer.ConversionType = req.ConversionType
	offer.Payout = req.Payout
	offer.UpdatedAt = time.Now()

	return u.offerRepo.Update(offer), nil
}
