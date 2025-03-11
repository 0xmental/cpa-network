package payout_in_memory

import "CPAPlatform/internal/domain"

type Repo struct {
	data      map[int64]*domain.Payout
	currentID int64
}

func NewRepo() *Repo {
	return &Repo{
		data:      make(map[int64]*domain.Payout),
		currentID: 1,
	}
}

func (r *Repo) NewID() int64 {
	nextID := r.currentID
	r.currentID++

	return nextID
}
