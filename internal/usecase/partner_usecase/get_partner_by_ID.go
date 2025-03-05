package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (u *UseCase) GetPartnerByID(partnerID int64) (*domain.Partner, error) {
	partner, err := u.repo.GetByID(partnerID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPartnerByID: %w", err)
	}

	return partner, nil
}
