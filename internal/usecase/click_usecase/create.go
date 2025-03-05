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
	Useragent string
	Country   string
	IsUnique  bool
}

func (u *UseCase) CreateClick(req CreateClickRequest) (*domain.Click, error) {
	now := time.Now()

	_, err := u.repoPartner.GetByID(req.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repoPartner.GetByID: %w", err)
	}

	_, err = u.repoOffer.GetByID(req.OfferID)
	if err != nil {
		return nil, fmt.Errorf("repoOffer.GetByID: %w", err)
	}

	if req.ClickID == "" {
		req.ClickID = domain.GenerateClickID(req.OfferID, req.PartnerID, req.IPAddress, req.Useragent)
	}

	req.IsUnique = u.repoClick.IsUnique(req.ClickID)

	click := domain.NewClick(
		req.OfferID, req.PartnerID, req.ClickID, req.Country,
		req.IPAddress, req.Useragent, req.UTMParams, req.IsUnique, now,
	)

	click = u.repoClick.Save(click)

	return click, nil
}
