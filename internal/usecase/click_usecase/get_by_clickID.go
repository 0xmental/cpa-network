package click_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type GetClickByClickID struct {
	ClickID string
}

func (u *UseCase) GetClickByClickID(req GetClickByClickID) (*domain.Click, error) {
	click, err := u.repoClick.GetByClickID(req.ClickID)
	if err != nil {
		return nil, fmt.Errorf("repoClick.GetByClickID: %w", err)
	}

	return click, nil
}
