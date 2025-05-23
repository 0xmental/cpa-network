package offer_usecase

import (
	"CPAPlatform/internal/usecase/offer_usecase/mocks"
	"errors"
	"testing"
)

var errTest = errors.New("internal")

type useCaseMocks struct {
	repoOffer *mocks.RepoOffer
	timer     *mocks.Timer
}

func makeServiceWithMocks(t *testing.T) (s *UseCase, m useCaseMocks) {
	m = useCaseMocks{
		repoOffer: mocks.NewRepoOffer(t),
		timer:     mocks.NewTimer(t),
	}
	u := &UseCase{
		offerRepo: m.repoOffer,
		timer:     m.timer,
	}

	return u, m
}
