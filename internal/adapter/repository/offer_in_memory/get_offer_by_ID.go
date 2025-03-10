package offer_in_memory

import (
	"CPAPlatform/internal/domain"
	"errors"
)

var ErrOfferNotFound = errors.New("the offer with this ID does not exist")

func (r *Repo) GetOfferByID(offerID int64) (*domain.Offer, error) {
	offer, exist := r.data[offerID]
	if !exist {
		return nil, ErrOfferNotFound
	}

	return offer, nil
}
