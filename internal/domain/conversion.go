package domain

import "time"

type Conversion struct {
	ID        int64
	ClickID   string
	Payout    int64
	OfferID   int64
	PartnerID int64
	CreatedAt time.Time
}

func NewConversion(clickID string, payout, offerID, partnerID int64, time time.Time) *Conversion {
	return &Conversion{
		ClickID:   clickID,
		Payout:    payout,
		OfferID:   offerID,
		PartnerID: partnerID,
		CreatedAt: time,
	}
}
