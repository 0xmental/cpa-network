package click_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllClicksByPartner(partnerID int64) []*domain.Click {
	var result []*domain.Click

	for _, click := range r.data {
		if click.PartnerID == partnerID {
			result = append(result, click)
		}
	}

	return result
}
