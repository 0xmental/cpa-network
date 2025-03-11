package partner_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllPartners() []*domain.Partner {
	result := make([]*domain.Partner, 0, len(r.data))
	for _, v := range r.data {
		result = append(result, v)
	}

	return result
}
