package payout_in_memory

import (
	"CPAPlatform/internal/domain"
)

func (r *Repo) GetByPartnerID(partnerID int64) []*domain.Payout {
	var result []*domain.Payout
	allPayouts := r.GetAll()

	for _, payout := range allPayouts {
		if payout.PartnerID == partnerID {
			result = append(result, payout)
		}
	}

	return result
}
