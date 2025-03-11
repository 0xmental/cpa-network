package click_in_memory

import (
	"CPAPlatform/internal/domain"
	"errors"
)

var ErrClickIDNotFound = errors.New("the clickID does not exist")

func (r *Repo) GetByClickID(clickID string) (*domain.Click, error) {
	var result *domain.Click

	for _, click := range r.data {
		if click.ClickID == clickID {
			result = click
		}
	}

	if result == nil {
		return nil, ErrClickIDNotFound
	}

	return result, nil
}
