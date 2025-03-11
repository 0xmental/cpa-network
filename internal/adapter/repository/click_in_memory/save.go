package click_in_memory

import "CPAPlatform/internal/domain"

func (r *Repo) Save(click *domain.Click) *domain.Click {
	click.ID = r.NewID()
	r.data[click.ID] = click

	return click
}
