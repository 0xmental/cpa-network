package conversion_usecase

import (
	"CPAPlatform/internal/domain"
	"CPAPlatform/internal/domain/dto"
)

type GetAllConversionsReq struct {
	OfferID   int64
	PartnerID int64
}

func (u *UseCase) GetAllConversions(req GetAllConversionsReq) []*domain.Conversion {
	return u.repoConversion.GetAllConversions(dto.ConversionFilter{
		OfferID:   req.OfferID,
		PartnerID: req.PartnerID,
	})
}
