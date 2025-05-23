package click_in_memory

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

func (r *Repo) GetAllClicks(filter dto.ClickFilter) []*domain.Click {
	result := make([]*domain.Click, 0, len(r.data))

	for _, click := range r.data {
		shouldAdd := true

		if filter.PartnerID > 0 && click.PartnerID != filter.PartnerID {
			shouldAdd = false
		}

		if filter.OfferID > 0 && click.OfferID != filter.OfferID {
			shouldAdd = false
		}

		if filter.ClickID == "" && click.ClickID != filter.ClickID {
			shouldAdd = false
		}

		if shouldAdd {
			result = append(result, click)
		}
	}

	return result
}
