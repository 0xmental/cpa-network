package payout_in_memory

import (
	"CPAPlatform/internal/domain"
)

func (r *Repo) GetAllPayouts() []*domain.Payout {
	result := make([]*domain.Payout, 0, len(r.data))
	for _, v := range r.data {
		result = append(result, v)
	}

	return result
}
