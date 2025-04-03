package conversion_usecase

import (
	"CPAPlatform/internal/domain"
)

type GetConversionsByPartnerReq struct {
	PartnerID int64
}

func (u *UseCase) GetAllConversionsByPartner(req GetConversionsByPartnerReq) []*domain.Conversion {
	return u.repoConversion.GetAllConversionsByPartner(req.PartnerID)
}
