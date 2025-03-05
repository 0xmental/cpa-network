package partner_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Update(partner *domain.Partner, partnerID int64) *domain.Partner {
	r.data[partnerID] = partner

	return partner
}
