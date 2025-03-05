package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type UpdateInfo struct {
	Pass         string
	ContactInfo  domain.ContactInfo
	WithdrawInfo *domain.WithdrawInfo
	PostbackURL  *string
}

func (u *UseCase) UpdatePartnerInfo(upd UpdateInfo, partnerID int64) (*domain.Partner, error) {
	partner, err := u.repo.GetByID(partnerID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %w", err)
	}

	partner.Pass = upd.Pass
	partner.ContactInfo = upd.ContactInfo
	partner.WithdrawInfo = upd.WithdrawInfo
	partner.PostbackURL = upd.PostbackURL

	updatedInfo := u.repo.Update(partner, partnerID)

	return updatedInfo, nil
}
