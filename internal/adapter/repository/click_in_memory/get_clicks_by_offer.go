package click_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetByOffer(offerID int64) []*domain.Click {
	var result []*domain.Click

	for _, click := range r.data {
		if click.OfferID == offerID {
			result = append(result, click)
		}
	}

	return result
}
