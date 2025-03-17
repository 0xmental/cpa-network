package partner_usecase

import (
	"CPAPlatform/internal/domain"
)

type (
	UseCase struct {
		partnerRepo repoPartner
	}
	repoPartner interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
		GetAllPartners() []*domain.Partner
		Update(partner *domain.Partner, partnerID int64) *domain.Partner
		Save(partner *domain.Partner) *domain.Partner
	}
)

func NewUseCase(partnerRepo repoPartner) *UseCase {
	return &UseCase{
		partnerRepo: partnerRepo,
	}
}
