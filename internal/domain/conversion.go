package domain

import "time"

type Conversion struct {
	ID        int64
	ClickID   string
	Payout    int64
	CreatedAt time.Time
}

func NewConversion(clickID string, payout int64, time time.Time) *Conversion {
	return &Conversion{
		ClickID:   clickID,
		Payout:    payout,
		CreatedAt: time,
	}
}
