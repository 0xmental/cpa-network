package partner_in_memory

import (
	"CPAPlatform/internal/domain"
	"errors"
)

var ErrPartnerNotFound = errors.New("the partner with this ID does not exist")

func (r *Repo) GetByID(partnerID int64) (*domain.Partner, error) {
	partner, exist := r.data[partnerID]
	if !exist {
		return nil, ErrPartnerNotFound
	}

	return partner, nil
}
