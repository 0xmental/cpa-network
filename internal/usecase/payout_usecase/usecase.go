package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"time"
)

type (
	UseCase struct {
		payoutRepo  repoPayout
		partnerRepo repoPartner
		timer       timer
	}

	repoPayout interface {
		UpdatePayoutStatus(payout *domain.Payout) *domain.Payout
		GetPayoutByID(payoutID int64) (*domain.Payout, error)
		GetAllPayoutsByPartnerID(partnerID int64) []*domain.Payout
		GetAllPayouts() []*domain.Payout
		Save(payout *domain.Payout) *domain.Payout
	}

	repoPartner interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
	}

	timer interface {
		Now() time.Time
	}
)

func NewUseCase(payoutRepo repoPayout, partnerRepo repoPartner, timer timer) *UseCase {
	return &UseCase{
		payoutRepo:  payoutRepo,
		partnerRepo: partnerRepo,
		timer:       timer,
	}
}
