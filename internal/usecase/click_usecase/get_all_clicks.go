package click_usecase

import "CPAPlatform/internal/domain"

func (u *UseCase) GetAllClicks() []*domain.Click {

	return u.repoClick.GetAllClicks()
}
