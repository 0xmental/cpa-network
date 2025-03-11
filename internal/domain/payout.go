package domain

import "time"

type PayoutStatus int8

var PendingPayoutStatus PayoutStatus = 1
var PaidPayoutStatus PayoutStatus = 2
var CanceledPayoutStatus PayoutStatus = 3

type Payout struct {
	ID           int64
	PartnerID    int64
	WithdrawInfo WithdrawInfo
	Amount       int64
	Status       PayoutStatus
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func NewPayout(partnerID int64, withdrawInfo WithdrawInfo, amount int64, time time.Time) *Payout {
	return &Payout{
		PartnerID:    partnerID,
		WithdrawInfo: withdrawInfo,
		Amount:       amount,
		Status:       PendingPayoutStatus,
		CreatedAt:    time,
		UpdateAt:     time,
	}
}
