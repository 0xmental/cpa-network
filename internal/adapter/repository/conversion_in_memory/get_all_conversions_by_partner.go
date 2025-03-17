package conversion_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllConversionsByPartner(partnerID int64) []*domain.Conversion {
	var result []*domain.Conversion

	for _, conversion := range r.data {
		if conversion.PartnerID == partnerID {
			result = append(result, conversion)
		}
	}

	return result
}
