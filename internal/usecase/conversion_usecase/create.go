package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"errors"
	"fmt"
)

var ErrPayoutNotDefined = errors.New("payout is not defined for the specified country")
var ErrOfferInactive = errors.New("this offer is inactive")

type CreateConversionRequest struct {
	ClickID string
}

func (u *UseCase) CreateConversion(req CreateConversionRequest) (*domain.Conversion, error) {
	clicks := u.repoClick.GetAllClicks(dto.ClickFilter{
		ClickID: req.ClickID,
	})
	
	if len(clicks) == 0 {
		return nil, fmt.Errorf("click with ID %s not found", req.ClickID)
	}
	
	click := clicks[0]

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

	now := u.timer.Now()
	conversion := domain.NewConversion(req.ClickID, payout, partner.ID, offer.ID, now)

	conversion = u.repoConversion.Save(conversion)

	return conversion, nil
}
