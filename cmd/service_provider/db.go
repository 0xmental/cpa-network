package service_provider

import (
	"CPAPlatform/internal/config"
	"context"
	"log"
)

func (s *ServiceProvider) getDbCluster(ctx context.Context) *config.Cluster {
	if s.dbCluster == nil {
		dbCluster, err := config.NewCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}

		s.dbCluster = dbCluster
	}

	return s.dbCluster
}
