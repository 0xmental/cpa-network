package payout_usecase

import (
	"CPAPlatform/internal/usecase/payout_usecase/mocks"
	"errors"
	"testing"
)

var errTest = errors.New("internal")

type useCaseMocks struct {
	repoPartner *mocks.RepoPartner
	repoPayout  *mocks.RepoPayout
	timer       *mocks.Timer
}

func makeServiceWithMocks(t *testing.T) (s *UseCase, m useCaseMocks) {
	m = useCaseMocks{
		repoPartner: mocks.NewRepoPartner(t),
		repoPayout:  mocks.NewRepoPayout(t),
		timer:       mocks.NewTimer(t),
	}
	u := &UseCase{
		partnerRepo: m.repoPartner,
		payoutRepo:  m.repoPayout,
		timer:       m.timer,
	}

	return u, m
}
