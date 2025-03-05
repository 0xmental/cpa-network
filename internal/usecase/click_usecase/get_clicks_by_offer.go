package click_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

func (u *UseCase) GetClicksByOffer(offerID int64) ([]*domain.Click, error) {
	_, err := u.repoOffer.GetByID(offerID)
	if err != nil {
		return nil, fmt.Errorf("repoOffer.GetByID: %w", err)
	}

	return u.repoClick.GetByOffer(offerID), nil
}
