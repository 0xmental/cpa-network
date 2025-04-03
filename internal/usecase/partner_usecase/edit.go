package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type UpdateInfoReq struct {
	PartnerID    int64
	Pass         string
	ContactInfo  domain.ContactInfo
	WithdrawInfo *domain.WithdrawInfo
	PostbackURL  *string
}

func (u *UseCase) UpdatePartnerInfo(req UpdateInfoReq) (*domain.Partner, error) {
	partner, err := u.partnerRepo.GetPartnerByID(req.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}

	partner.Pass = req.Pass
	partner.ContactInfo = req.ContactInfo
	partner.WithdrawInfo = req.WithdrawInfo
	partner.PostbackURL = req.PostbackURL
	partner.UpdatedAt = u.timer.Now()

	return u.partnerRepo.Update(partner), nil
}
