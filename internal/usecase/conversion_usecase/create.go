package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"errors"
	"fmt"
	"time"
)

var ErrPayoutNotDefined = errors.New("payout is not defined for the specified country")
var ErrOfferInactive = errors.New("this offer is inactive")

type CreateConversionRequest struct {
	ClickID string
}

func (u *UseCase) CreateConversion(req CreateConversionRequest) (*domain.Conversion, error) {
	click, err := u.repoClick.GetByClickID(req.ClickID)
	if err != nil {
		return nil, fmt.Errorf("repoClick.GetByClickID: %w", err)
	}

	offer, err := u.repoOffer.GetOfferByID(click.OfferID)
	if err != nil {
		return nil, fmt.Errorf("repoOffer.GetByOfferID: %w", err)
	}

	if !offer.IsActive {
		return nil, ErrOfferInactive
	}

	payout, exist := offer.Payout[click.Country]
	if !exist {
		return nil, ErrPayoutNotDefined
	}

	partner, err := u.repoPartner.GetPartnerByID(click.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repoPartner.GetByID: %w", err)
	}
	partner.AddBalance(payout)

	now := time.Now()
	conversion := domain.NewConversion(req.ClickID, payout, partner.ID, offer.ID, now)

	conversion = u.repoConversion.Save(conversion)

	return conversion, nil
}
