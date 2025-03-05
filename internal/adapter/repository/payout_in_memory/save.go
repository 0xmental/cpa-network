package payout_in_memory

import (
	"CPAPlatform/internal/domain"
)

func (r *Repo) Save(payout *domain.Payout) *domain.Payout {
	payout.ID = r.NewID()
	r.data[payout.ID] = payout
	
	return payout
}
