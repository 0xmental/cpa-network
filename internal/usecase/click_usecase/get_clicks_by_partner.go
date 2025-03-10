package click_usecase

import (
	"CPAPlatform/internal/domain"
)

type GetClicksByPartnerReq struct {
	PartnerID int64
}

func (u *UseCase) GetClicksByPartner(req GetClicksByPartnerReq) []*domain.Click {
	return u.repoClick.GetAllClicksByPartner(req.PartnerID)
}
