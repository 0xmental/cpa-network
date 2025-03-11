package payout_usecase

import "CPAPlatform/internal/domain"

type (
	UseCase struct {
		payoutRepo  repoPayout
		partnerRepo repoPartner
	}

	repoPayout interface {
		UpdatePayoutStatus(payoutID int64, status domain.PayoutStatus) (*domain.Payout, error)
		GetPayoutByID(payoutID int64) (*domain.Payout, error)
		GetAllPayoutsByPartnerID(partnerID int64) []*domain.Payout
		GetAllPayouts() []*domain.Payout
		Save(payout *domain.Payout) *domain.Payout
	}
	
	repoPartner interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
	}
)

func NewUseCase(payoutRepo repoPayout, partnerRepo repoPartner) *UseCase {
	return &UseCase{
		payoutRepo:  payoutRepo,
		partnerRepo: partnerRepo,
	}
}
