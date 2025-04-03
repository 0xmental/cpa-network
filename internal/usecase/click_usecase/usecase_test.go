package click_usecase

import (
	"CPAPlatform/internal/usecase/click_usecase/mocks"
	"errors"
	"testing"
)

var errTest = errors.New("internal")

type useCaseMocks struct {
	repoPartner *mocks.PartnerRepo
	repoOffer   *mocks.OfferRepo
	repoClick   *mocks.ClickRepo
	timer       *mocks.Timer
}

func makeServiceWithMocks(t *testing.T) (s *UseCase, m useCaseMocks) {
	m = useCaseMocks{
		repoPartner: mocks.NewPartnerRepo(t),
		repoOffer:   mocks.NewOfferRepo(t),
		repoClick:   mocks.NewClickRepo(t),
		timer:       mocks.NewTimer(t),
	}
	u := &UseCase{
		repoPartner: m.repoPartner,
		repoOffer:   m.repoOffer,
		repoClick:   m.repoClick,
		timer:       m.timer,
	}

	return u, m
}
