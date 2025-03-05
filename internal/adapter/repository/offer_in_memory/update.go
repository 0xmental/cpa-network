package offer_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Update(offer *domain.Offer, offerID int64) *domain.Offer {
	r.data[offerID] = offer

	return offer
}
