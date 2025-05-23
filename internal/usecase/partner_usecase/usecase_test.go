package partner_usecase

import (
	"CPAPlatform/internal/usecase/partner_usecase/mocks"
	"errors"
	"testing"
)

var errTest = errors.New("internal")

type useCaseMocks struct {
	repoPartner *mocks.RepoPartner
	timer       *mocks.Timer
}

func makeServiceWithMocks(t *testing.T) (s *UseCase, m useCaseMocks) {
	m = useCaseMocks{
		repoPartner: mocks.NewRepoPartner(t),
		timer:       mocks.NewTimer(t),
	}
	s = &UseCase{
		partnerRepo: m.repoPartner,
		timer:       m.timer,
	}

	return s, m
}
