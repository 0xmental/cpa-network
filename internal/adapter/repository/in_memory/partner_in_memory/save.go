package partner_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Save(partner *domain.Partner) (*domain.Partner, error) {
	partner.ID = r.NewID()
	r.data[partner.ID] = partner

	return partner, nil
}
