package dto

import "CPAPlatform/internal/domain"

type PayoutFilter struct {
	PartnerID int64
	PayoutID  int64
	Status    domain.PayoutStatus
}
