package conversion_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Save(conversion *domain.Conversion) *domain.Conversion {
	conversion.ID = r.NewID()
	r.data[conversion.ID] = conversion

	return conversion
}
