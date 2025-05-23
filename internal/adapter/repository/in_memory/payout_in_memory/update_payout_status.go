package payout_in_memory

import (
	"CPAPlatform/internal/domain"
)

func (r *Repo) UpdatePayoutStatus(payout *domain.Payout) *domain.Payout {
	r.data[payout.ID] = payout

	return payout
}
