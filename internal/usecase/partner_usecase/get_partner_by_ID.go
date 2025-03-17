package partner_usecase

import (
	"CPAPlatform/internal/domain"
	"fmt"
)

type GetPartnerReq struct {
	PartnerID int64
}

func (u *UseCase) GetPartnerByID(req GetPartnerReq) (*domain.Partner, error) {
	partner, err := u.partnerRepo.GetPartnerByID(req.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetPartnerByID: %w", err)
	}

	return partner, nil
}
