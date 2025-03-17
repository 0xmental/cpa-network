package conversion_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllConversionsByOffer(offerID int64) []*domain.Conversion {
	var result []*domain.Conversion

	for _, conversion := range r.data {
		if conversion.OfferID == offerID {
			result = append(result, conversion)
		}
	}

	return result
}
