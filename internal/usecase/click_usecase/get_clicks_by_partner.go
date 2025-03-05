package click_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (u *UseCase) GetClicksByPartner(partnerID int64) ([]*domain.Click, error) {
	_, err := u.repoPartner.GetByID(partnerID)
	if err != nil {
		return nil, fmt.Errorf("repoPartner.GetByID: %w", err)
	}

	return u.repoClick.GetByPartner(partnerID), nil
}
