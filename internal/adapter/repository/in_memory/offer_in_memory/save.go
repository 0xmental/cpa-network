package offer_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Save(offer *domain.Offer) *domain.Offer {
	offer.ID = r.NewID()
	r.data[offer.ID] = offer

	return offer
}
