package offer_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type GetOfferReq struct {
	OfferID   int64
	PartnerID int64
}
type GetOfferResponse struct {
	ID             int64
	Name           string
	TargetURL      string
	Description    string
	ConversionType domain.ConversionType
	Payout         map[string]int64
	TrackingURL    string
}

func (u *UseCase) GetOffer(req GetOfferReq) (*GetOfferResponse, error) {
	offer, err := u.offerRepo.GetOfferByID(req.OfferID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}

	trackingURL := fmt.Sprintf("https://%s/click?offer=%d&partner=%d", offer.RedirectDomain, offer.ID, req.PartnerID)

	return &GetOfferResponse{
		ID:             offer.ID,
		Name:           offer.Name,
		TargetURL:      offer.TargetUrl,
		Description:    offer.Description,
		ConversionType: offer.ConversionType,
		Payout:         offer.Payout,
		TrackingURL:    trackingURL,
	}, nil
}
