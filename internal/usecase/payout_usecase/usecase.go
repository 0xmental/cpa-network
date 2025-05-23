package payout_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
	"time"
)

type (
	UseCase struct {
		payoutRepo  repoPayout
		partnerRepo repoPartner
		timer       timer
	}

	repoPayout interface {
		GetAllPayouts(filter dto.PayoutFilter) []*domain.Payout
		UpdatePayoutStatus(payout *domain.Payout) *domain.Payout
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
