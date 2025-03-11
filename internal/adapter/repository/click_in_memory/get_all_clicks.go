package click_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) GetAllClicks() []*domain.Click {
	result := make([]*domain.Click, 0, len(r.data))

	for _, click := range r.data {
		result = append(result, click)
	}

	return result
}
