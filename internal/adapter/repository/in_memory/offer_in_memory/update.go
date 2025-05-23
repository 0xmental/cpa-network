package offer_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Update(offer *domain.Offer) *domain.Offer {
	r.data[offer.ID] = offer

	return offer
}
