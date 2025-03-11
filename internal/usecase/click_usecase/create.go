package click_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
	"time"
)

type CreateClickRequest struct {
	OfferID   int64
	PartnerID int64
	ClickID   string // Если не передан, генерируется автоматически
	UTMParams map[string]string
	IPAddress string
	UserAgent string
	Country   string
}

func (u *UseCase) CreateClick(req CreateClickRequest) (*domain.Click, error) {
	now := time.Now()

	partner, err := u.repoPartner.GetPartnerByID(req.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repoPartner.GetByID: %w", err)
	}

	offer, err := u.repoOffer.GetOfferByID(req.OfferID)
	if err != nil {
		return nil, fmt.Errorf("repoOffer.GetByID: %w", err)
	}

	clickID := req.ClickID
	if clickID == "" {
		clickID = domain.GenerateClickID(offer.ID, partner.ID, req.IPAddress, req.UserAgent)
	}

	IsUnique := u.repoClick.IsUnique(req.ClickID)

	click := domain.NewClick(
		offer.ID, partner.ID, clickID, req.Country,
		req.IPAddress, req.UserAgent, req.UTMParams, IsUnique, now,
	)

	click = u.repoClick.Save(click)

	return click, nil
}
