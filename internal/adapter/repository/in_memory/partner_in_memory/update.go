package partner_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Update(partner *domain.Partner) *domain.Partner {
	r.data[partner.ID] = partner

	return partner
}
