package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type UpdateOffer struct {
	TargetURL      string
	Name           string
	Description    string
	RedirectDomain string
	ConversionType domain.ConversionType
	Payout         map[string]int64
}

func (u *UseCase) EditOffer(upd UpdateOffer, offerID int64) (*domain.Offer, error) {
	offer, err := u.repo.GetByID(offerID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}
	
	offer.TargetUrl = upd.TargetURL
	offer.Name = upd.Name
	offer.Description = upd.Description
	offer.RedirectDomain = upd.RedirectDomain
	offer.ConversionType = upd.ConversionType
	offer.Payout = upd.Payout

	updatedOffer := u.repo.Update(offer, offerID)

	return updatedOffer, nil
}
