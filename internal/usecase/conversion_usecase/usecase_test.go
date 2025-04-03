package conversion_usecase

import (
	"CPAPlatform/internal/usecase/conversion_usecase/mocks"
	"errors"
	"testing"
)

var errTest = errors.New("internal")

type useCaseMocks struct {
	repoPartner    *mocks.PartnerRepo
	repoConversion *mocks.ConversionRepo
	repoOffer      *mocks.OfferRepo
	repoClick      *mocks.ClickRepo
	timer          *mocks.Timer
}

func makeServiceWithMocks(t *testing.T) (s *UseCase, m useCaseMocks) {
	m = useCaseMocks{
		repoPartner:    mocks.NewPartnerRepo(t),
		repoConversion: mocks.NewConversionRepo(t),
		repoOffer:      mocks.NewOfferRepo(t),
		repoClick:      mocks.NewClickRepo(t),
		timer:          mocks.NewTimer(t),
	}
	u := &UseCase{
		repoPartner:    m.repoPartner,
		repoConversion: m.repoConversion,
		repoOffer:      m.repoOffer,
		repoClick:      m.repoClick,
		timer:          m.timer,
	}

	return u, m
}
