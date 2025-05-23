package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type CreatePartnerRequest struct {
	Email        string
	Pass         string
	ContactInfo  domain.ContactInfo
	WithdrawInfo *domain.WithdrawInfo
	PostbackURL  *string
}

func (u *UseCase) CreatePartner(req CreatePartnerRequest) (*domain.Partner, error) {
	now := u.timer.Now()
	partner, err := domain.NewPartner(req.Email, req.Pass, req.ContactInfo, req.WithdrawInfo, req.PostbackURL, now)
	if err != nil {
		return nil, fmt.Errorf("domain.NewPartner: %w", err)
	}

	partner, err = u.partnerRepo.Save(partner)
	if err != nil {
		return nil, fmt.Errorf("u.partnerRepo.Save: %w", err)
	}
	return partner, nil
}
