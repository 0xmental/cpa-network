package conversion_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllConversions() []*domain.Conversion {
	return u.repoConversion.GetAllConversions()
}
