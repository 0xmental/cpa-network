package conversion_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllConversions() []*domain.Conversion {
	result := make([]*domain.Conversion, 0, len(r.data))

	for _, conversion := range r.data {
		result = append(result, conversion)
	}

	return result
}
