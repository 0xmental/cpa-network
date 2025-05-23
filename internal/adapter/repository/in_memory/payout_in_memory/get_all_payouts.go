package payout_in_memory

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

func (r *Repo) GetAllPayouts(filter dto.PayoutFilter) []*domain.Payout {
	result := make([]*domain.Payout, 0, len(r.data))

	for _, payout := range r.data {
		shouldAdd := true

		if filter.PartnerID > 0 && payout.PartnerID != filter.PartnerID {
			shouldAdd = false
		}

		if filter.PayoutID > 0 && payout.ID != filter.PayoutID {
			shouldAdd = false
		}

		if filter.Status > 0 && payout.Status != filter.Status {
			shouldAdd = false
		}

		if shouldAdd {
			result = append(result, payout)
		}
	}

	return result
}
