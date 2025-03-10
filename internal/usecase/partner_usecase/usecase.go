package partner_usecase

import (
	"CPAPlatform/internal/domain"
)

type (
	UseCase struct {
		repo repo
	}
	repo interface {
		GetPartnerByID(partnerID int64) (*domain.Partner, error)
		GetAllPartners() []*domain.Partner
		Update(partner *domain.Partner, partnerID int64) *domain.Partner
		Save(partner *domain.Partner) *domain.Partner
	}
)

func NewUseCase(repo repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
