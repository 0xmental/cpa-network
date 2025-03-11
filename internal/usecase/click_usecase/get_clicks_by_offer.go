package click_usecase

import (
	"CPAPlatform/internal/domain"
)

type GetClicksByOfferReq struct {
	OfferID int64
}

func (u *UseCase) GetAllClicksByOffer(req GetClicksByOfferReq) []*domain.Click {
	return u.repoClick.GetAllClicksByOffer(req.OfferID)
}
