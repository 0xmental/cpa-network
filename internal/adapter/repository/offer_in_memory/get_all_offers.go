package offer_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllOffers() []*domain.Offer {
	result := make([]*domain.Offer, 0, len(r.data))
	for _, offer := range r.data {
		if offer.IsActive {
			result = append(result, offer)
		}
	}

	return result
}
