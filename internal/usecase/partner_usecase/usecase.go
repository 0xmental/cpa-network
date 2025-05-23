package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"time"
)

type (
	UseCase struct {
		partnerRepo repoPartner
		timer       timer
	}
	repoPartner interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
		GetAllPartners() ([]*domain.Partner, error)
		Update(partner *domain.Partner) (*domain.Partner, error)
		Save(partner *domain.Partner) (*domain.Partner, error)
	}

	timer interface {
		Now() time.Time
	}
)

func NewUseCase(partnerRepo repoPartner, timer timer) *UseCase {
	return &UseCase{
		partnerRepo: partnerRepo,
		timer:       timer,
	}
}
