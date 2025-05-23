package partner_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllPartners() ([]*domain.Partner, error) {
	return u.partnerRepo.GetAllPartners()
}
