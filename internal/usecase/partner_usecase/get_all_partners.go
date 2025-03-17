package partner_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllPartners() []*domain.Partner {
	return u.partnerRepo.GetAllPartners()
}
