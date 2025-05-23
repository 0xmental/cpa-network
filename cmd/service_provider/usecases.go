package service_provider

import (
	"CPAPlatform/internal/usecase/partner_usecase"
	"CPAPlatform/pkg"
)

func (s *ServiceProvider) GetPartnerUseCase() *partner_usecase.UseCase {
	if s.partnerUseCase == nil {
		s.partnerUseCase = partner_usecase.NewUseCase(s.getPartnerPostgresRepo(), pkg.NewTimer())
	}

	return s.partnerUseCase
}
