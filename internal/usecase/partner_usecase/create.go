package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
	"time"
)

type CreatePartnerRequest struct {
	Email        string
	Pass         string
	ContactInfo  domain.ContactInfo
	WithdrawInfo *domain.WithdrawInfo
	PostbackURL  *string
	IsActive     bool
	Balance      int64
}

func (u *UseCase) CreatePartner(req CreatePartnerRequest) (*domain.Partner, error) {
	now := time.Now()
	partner, err := domain.NewPartner(req.Email, req.Pass, req.ContactInfo, req.WithdrawInfo, req.PostbackURL, req.Balance, now)
	if err != nil {
		return nil, fmt.Errorf("domain.NewPartner: %w", err)
	}

	return u.repo.Save(partner), nil
}
