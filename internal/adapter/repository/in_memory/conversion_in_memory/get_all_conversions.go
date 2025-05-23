package conversion_in_memory

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

func (r *Repo) GetAllConversions(filter dto.ConversionFilter) []*domain.Conversion {
	result := make([]*domain.Conversion, 0, len(r.data))

	for _, conversion := range r.data {
		shouldAdd := true
		if filter.PartnerID > 0 && conversion.PartnerID != filter.PartnerID {
			shouldAdd = false
		}

		if filter.OfferID > 0 && conversion.OfferID != filter.OfferID {
			shouldAdd = false
		}

		if shouldAdd {
			result = append(result, conversion)
		}
	}

	return result
}
