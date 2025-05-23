package click_in_memory

func (r *Repo) IsUnique(clickID string) bool {
	for _, click := range r.data {
		if click.ClickID == clickID {
			return false
		}
	}

	return true
}
