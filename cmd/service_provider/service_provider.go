package service_provider

import (
	"CPAPlatform/internal/adapter/repository/in_memory/partner_in_memory"
	"CPAPlatform/internal/adapter/repository/postgres/click_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/conversion_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/offer_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/partner_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/payout_postgres"
	"CPAPlatform/internal/config"
	"CPAPlatform/internal/usecase/partner_usecase"
)

type ServiceProvider struct {
	//use cases
	partnerUseCase *partner_usecase.UseCase

	//repos
	partnerInMemoryRepo    *partner_in_memory.Repo
	partnerPostgresRepo    *partner_postgres.Repo
	clickPostgresRepo      *click_postgres.Repo
	conversionPostgresRepo *conversion_postgres.Repo
	offerPostgresRepo      *offer_postgres.Repo
	payoutPostgresRepo     *payout_postgres.Repo

	dbCluster *config.Cluster
}

func (s *ServiceProvider) ProviderService() error {

	return nil
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
