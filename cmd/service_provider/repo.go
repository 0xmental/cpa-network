package service_provider

import (
	"CPAPlatform/internal/adapter/repository/in_memory/partner_in_memory"
	"CPAPlatform/internal/adapter/repository/postgres/partner_postgres"
)

func (s *ServiceProvider) getPartnerInMemoryRepo() *partner_in_memory.Repo {
	if s.partnerInMemoryRepo == nil {
		s.partnerInMemoryRepo = partner_in_memory.NewRepo()
	}

	return s.partnerInMemoryRepo
}

func (s *ServiceProvider) getPartnerPostgresRepo() *partner_postgres.Repo {
	if s.partnerPostgresRepo == nil {
		s.partnerPostgresRepo = partner_postgres.NewRepo(s.dbCluster)
	}

	return s.partnerPostgresRepo
}
