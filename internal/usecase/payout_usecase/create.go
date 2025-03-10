package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
	"time"
)

type CreatePayoutRequest struct {
	PartnerID    int64
	WithdrawInfo domain.WithdrawInfo
	Amount       int64
	Status       domain.PayoutStatus
}

func (u *UseCase) CreatePayout(req CreatePayoutRequest) (*domain.Payout, error) {
	partner, err := u.partnerRepo.GetPartnerByID(req.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repo.PartnerGetByID: %w", err)
	}

	err = partner.DeductBalance(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("partner.DeductBalance: %w", err)
	}

	req.WithdrawInfo = *partner.WithdrawInfo
	now := time.Now()
	payout := domain.NewPayout(req.PartnerID, req.WithdrawInfo, req.Amount, now)

	return u.payoutRepo.Save(payout), nil
}
