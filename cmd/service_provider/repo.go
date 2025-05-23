package service_provider

import (
	"CPAPlatform/internal/adapter/repository/postgres/click_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/conversion_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/offer_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/partner_postgres"
	"CPAPlatform/internal/adapter/repository/postgres/payout_postgres"
	"context"
)

func (s *ServiceProvider) getClickPostgresRepo() *click_postgres.Repo {
	if s.clickPostgresRepo == nil {
		s.clickPostgresRepo = click_postgres.NewRepo(s.getDbCluster(context.Background()))
	}

	return s.clickPostgresRepo
}

func (s *ServiceProvider) getConversionPostgresRepo() *conversion_postgres.Repo {
	if s.conversionPostgresRepo == nil {
		s.conversionPostgresRepo = conversion_postgres.NewRepo(s.getDbCluster(context.Background()))
	}

	return s.conversionPostgresRepo
}

func (s *ServiceProvider) getOfferPostgresRepo() *offer_postgres.Repo {
	if s.offerPostgresRepo == nil {
		s.offerPostgresRepo = offer_postgres.NewRepo(s.getDbCluster(context.Background()))
	}

	return s.offerPostgresRepo
}

func (s *ServiceProvider) getPartnerPostgresRepo() *partner_postgres.Repo {
	if s.partnerPostgresRepo == nil {
		s.partnerPostgresRepo = partner_postgres.NewRepo(s.getDbCluster(context.Background()))
	}

	return s.partnerPostgresRepo
}

func (s *ServiceProvider) getPayoutPostgresRepo() *payout_postgres.Repo {
	if s.payoutPostgresRepo == nil {
		s.payoutPostgresRepo = payout_postgres.NewRepo(s.getDbCluster(context.Background()))
	}

	return s.payoutPostgresRepo
}
