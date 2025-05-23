package click_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

type GetAllClicksReq struct {
	PartnerID int64
	OfferID   int64
	ClickID   string
}

func (u *UseCase) GetAllClicks(req GetAllClicksReq) []*domain.Click {
	return u.repoClick.GetAllClicks(dto.ClickFilter{
		PartnerID: req.PartnerID,
		OfferID:   req.OfferID,
		ClickID:   req.ClickID,
	})
}
