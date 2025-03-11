package payout_in_memory

import (
	"CPAPlatform/internal/domain"
	"errors"
)

var ErrPayoutNotFound = errors.New("the payout with this ID does not exist")

func (r *Repo) GetPayoutByID(payoutID int64) (*domain.Payout, error) {
	result, exist := r.data[payoutID]
	if !exist {
		return nil, ErrPayoutNotFound
	}

	return result, nil
}
