package click_usecase

import (
	"CPAPlatform/internal/adapter/repository/click_in_memory"
	"CPAPlatform/internal/adapter/repository/offer_in_memory"
	"CPAPlatform/internal/adapter/repository/partner_in_memory"
)

type UseCase struct {
	repoClick   *click_in_memory.Repo
	repoPartner *partner_in_memory.Repo
	repoOffer   *offer_in_memory.Repo
}

func NewUseCase(clickRepo *click_in_memory.Repo, offerRepo *offer_in_memory.Repo, partnerRepo *partner_in_memory.Repo) *UseCase {
	return &UseCase{
		repoClick:   clickRepo,
		repoPartner: partnerRepo,
		repoOffer:   offerRepo,
	}
}
