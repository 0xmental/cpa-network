package conversion_usecase

import (
	"CPAPlatform/internal/domain"
)

type GetConversionsByOfferReq struct {
	OfferID int64
}

func (u *UseCase) GetAllConversionsByOffer(req GetConversionsByOfferReq) []*domain.Conversion {
	return u.repoConversion.GetAllConversionsByOffer(req.OfferID)
}
